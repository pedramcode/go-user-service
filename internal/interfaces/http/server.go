package http

import (
	"dovenet/user-service/internal/application"
	"dovenet/user-service/internal/interfaces/http/router"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpServer struct {
	engine *gin.Engine
	logger *zap.Logger
}

func NewHttpServer(logger *zap.Logger, userService *application.UserService) *HttpServer {
	server := &HttpServer{
		engine: gin.Default(),
		logger: logger,
	}

	router.NewRouter(server.engine, logger, userService)

	return server
}

func (s *HttpServer) Run() {
	s.engine.Run(fmt.Sprintf("%s:%s", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT")))
}
