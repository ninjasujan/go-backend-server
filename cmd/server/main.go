package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/server/config"
	"app/server/db"

	"app/server/route"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load config
	appCfg, err := config.LoadConfig("local.yaml")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// gin.SetMode(gin.ReleaseMode)
	appEngine := gin.New()

	// Connect Postgres
	pgDb, err := db.InitPostgres(appCfg.Postgres)

	if err != nil {
		log.Println("[postgres database connection failed]")
		db.Cleanup(pgDb)
		os.Exit(1)
	}

	// Run Migration script
	err = db.RunMigration(appCfg.Postgres)
	if err != nil {
		log.Println("Migration script failed", err)
		os.Exit(1)
	}

	// Register Routes

	route.RegisterRoutes(pgDb, appEngine)

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", appCfg.Server.Host, appCfg.Server.Port),
		Handler:           appEngine,
		ReadHeaderTimeout: time.Minute,
		ReadTimeout:       time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
		MaxHeaderBytes:    1 << 20, // 1 MB
	}

	go func() {
		log.Printf("[server starting on %s]", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("[server failed to start: %v]", err)
		}
	}()

	log.Println("[server started successfully]")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	log.Println("[shutdown signal received]")
	log.Println("[shutting down server gracefully...]")

	// Create shutdown context with 30 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("[graceful shutdown failed: %v]", err)
		log.Println("[forcing server close]")

		// Force close if graceful shutdown fails
		if closeErr := server.Close(); closeErr != nil {
			log.Printf("[force close failed: %v]", closeErr)
		}
	} else {
		log.Println("[server shutdown completed gracefully]")
	}
}
