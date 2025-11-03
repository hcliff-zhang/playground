package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hcliff-zhang/playground/application"
	"github.com/hcliff-zhang/playground/database"
	"google.golang.org/grpc"
	"gorm.io/gorm/logger"
)

const (
	HTTPPort = ":8080"
	GRPCPort = ":9090"
)

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt retrieves an environment variable as an integer or returns a default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func main() {
	// Database configuration from environment variables
	dbConfig := database.PostgresConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "playground"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Build DSN from config
	dsn := database.BuildPostgresDSN(dbConfig)

	// Initialize database connection
	db, err := database.NewPostgres(
		dsn,
		25,            // maxOpenConns
		25,            // maxIdleConns
		5*time.Minute, // connMaxLifetime
		logger.Info,   // logLevel
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(db, &database.Patient{}, &database.Prescription{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create the service implementation with database
	service := application.NewService(db)

	// Start gRPC server in a goroutine
	go func() {
		// Create a TCP listener on the gRPC port
		listener, err := net.Listen("tcp", GRPCPort)
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port %s: %v", GRPCPort, err)
		}

		// Create a new gRPC server
		grpcServer := grpc.NewServer()

		// Register the gRPC handlers
		application.RegisterGRPCHandlers(grpcServer, service)

		log.Printf("Starting gRPC server on port %s", GRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Give the gRPC server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Register HTTP gateway
	ctx := context.Background()
	httpHandler, err := application.RegisterHTTPGateway(ctx, GRPCPort)
	if err != nil {
		log.Fatalf("Failed to register HTTP gateway: %v", err)
	}

	// Start HTTP server
	httpServer := &http.Server{
		Addr:    HTTPPort,
		Handler: httpHandler,
	}

	log.Printf("Starting HTTP gateway on port %s", HTTPPort)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
