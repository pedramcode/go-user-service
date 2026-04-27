package http

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		engine: gin.Default(),
	}
}

func (s *HttpServer) Run() {
	s.engine.Run(fmt.Sprintf("%s:%s", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT")))
}
