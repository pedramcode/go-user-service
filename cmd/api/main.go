package main

import (
	"dovenet/user-service/internal/application"
	"dovenet/user-service/internal/infrastructure/persistent"
	"dovenet/user-service/internal/infrastructure/persistent/repository"
	"dovenet/user-service/internal/interfaces/http"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load .env file
	godotenv.Load(".env")

	// Initialize database
	db := persistent.InitializeDatabase()
	if err := persistent.SyncMigrations(db); err != nil {
		panic(err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	credRepo := repository.NewCredentialRepository(db)
	otpRepo := repository.NewOtpRepository(db)

	// Initialize services
	userService := application.NewUserService(userRepo, credRepo, otpRepo)

	// Initialize interfaces dependencies
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Initialize interfaces
	httpServer := http.NewHttpServer(logger, userService)

	// Run interfaces
	httpServer.Run()
}
