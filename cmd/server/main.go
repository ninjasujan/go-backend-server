package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/server/common/config"
	"app/server/common/constant"
	"app/server/common/logger"
	"app/server/common/middleware"
	"app/server/db"
	"app/server/route"

	"app/server/common/kafka/producer"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize logger first
	if err := logger.InitLogger(logger.Config{
		Level:  "debug",
		Format: "console",
	}); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	// Load config
	appConfig, err := config.LoadConfig("local.yaml")
	if err != nil {
		logger.Error().Err(err).Msg("Failed to load config")
		os.Exit(1)
	}

	// Connect Postgres
	pgDb, err := db.InitPostgres(appConfig.Postgres)
	if err != nil {
		logger.Error().Err(err).Msg("Postgres database connection failed")
		db.Cleanup(pgDb)
		os.Exit(1)
	}

	// Run Migration script
	if err := db.RunMigration(appConfig.Postgres); err != nil {
		logger.Error().Err(err).Msg("Migration script failed")
		os.Exit(1)
	}

	// Setup Gin
	gin.SetMode(appConfig.App.Mode)
	appEngine := gin.New()

	// Initialize common middleware
	appEngine.Use(middleware.RequestLogger())

	// Kafka Producer Initialization
	kafkaProducer, err := producer.NewKafkaProducer(appConfig.Kafka.Brokers, constant.KafkaClientID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to initialize Kafka producer")
		os.Exit(1)
	}

	// Register Routes
	route.RegisterRoutes(pgDb, appEngine, kafkaProducer)

	// Configure server
	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", appConfig.Server.Host, appConfig.Server.Port),
		Handler:           appEngine,
		ReadHeaderTimeout: time.Minute,
		ReadTimeout:       time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
		MaxHeaderBytes:    1 << 20, // 1 MB
	}

	// Start server in goroutine
	go func() {
		logger.ServerStartup(server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error().Err(err).Msg("Server failed to start")
		}
	}()

	logger.Info().Msg("Server started successfully")

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logger.Info().Msg("Shutdown signal received")
	logger.ServerShutdown()

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("Graceful shutdown failed")
		logger.Info().Msg("Forcing server close")

		if closeErr := server.Close(); closeErr != nil {
			logger.Error().Err(closeErr).Msg("Force close failed")
		}
	} else {
		logger.Info().Msg("Server shutdown completed gracefully")
	}
}
