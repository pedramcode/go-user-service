package main

import (
	"dovenet/user-service/internal/application"
	"dovenet/user-service/internal/infrastructure/persistent"
	"dovenet/user-service/internal/infrastructure/persistent/repository"
	"dovenet/user-service/internal/interfaces/http"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	_ "dovenet/user-service/docs"
)

// @title           User Service API
// @version         1.0
// @description     User management service with authentication and authorization
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html

// @host           localhost:8080
// @schemes        http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the JWT token.

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
