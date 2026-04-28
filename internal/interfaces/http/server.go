package http

import (
	"context"
	"dovenet/user-service/internal/application"
	"dovenet/user-service/internal/interfaces/http/router"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpServer struct {
	server *http.Server
	logger *zap.Logger
}

func NewHttpServer(logger *zap.Logger, userService *application.UserService) *HttpServer {
	engine := gin.Default()
	addr := fmt.Sprintf("%s:%s", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))

	server := &HttpServer{
		server: &http.Server{
			Addr:         addr,
			Handler:      engine,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		logger: logger,
	}

	router.NewRouter(engine, logger, userService)

	return server
}

func (s *HttpServer) Run() error {
	go func() {
		s.logger.Info("Starting HTTP server", zap.String("addr", s.server.Addr))
		if os.Getenv("ENV") != "production" {
			s.logger.Info(fmt.Sprintf("Swagger address: http://%s/swagger/index.html", s.server.Addr))
		}
		if err := s.server.ListenAndServe(); err != nil {
			s.logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.logger.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("Server forced to shutdown", zap.Error(err))
		return err
	}

	s.logger.Info("Server exited gracefully")
	return nil
}
