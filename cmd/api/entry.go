package main

import (
	"dovenet/user-service/internal/application"
	"dovenet/user-service/internal/infrastructure/persistent"
	"dovenet/user-service/internal/infrastructure/persistent/repository"

	"github.com/joho/godotenv"
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

	// Initialize services
	userService := application.NewUserService(userRepo, credRepo)

	_ = userService
}
