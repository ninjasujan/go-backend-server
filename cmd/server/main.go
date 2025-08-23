package main

import (
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
	appCfg := config.LoadConfig("local.yaml")

	// gin.SetMode(gin.ReleaseMode)
	appEngine := gin.New()

	// Connect Postgres
	db.InitPostgres(&appCfg.Postgres)

	// Run Migration script
	db.RunMigration(&appCfg.Postgres)

	// Register Routes

	route.RegisterRoutes(appEngine)

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
		if err := server.ListenAndServe(); err != nil {
			log.Println("[server failed to start]")
		}
	}()

	log.Println("[server started successfully]")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	fmt.Println("[shutting down server]")

	if err := server.Close(); err != nil {
		log.Println("[server failed to close]")
	} else {
		log.Println("[server closed successfully]")
	}
}
