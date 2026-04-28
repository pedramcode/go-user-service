package router

import (
	"dovenet/user-service/internal/application"
	"dovenet/user-service/internal/interfaces/http/handler"
	"dovenet/user-service/internal/interfaces/http/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	engine         *gin.Engine
	userHandler    *handler.UserHandler
	authMiddleware *middleware.AuthMiddleware
	logger         *zap.Logger
}

func NewRouter(
	engine *gin.Engine,
	logger *zap.Logger,
	userService *application.UserService,
) *Router {
	authMiddleware := middleware.NewAuthMiddleware(os.Getenv("SECRET"))

	router := &Router{
		engine:         engine,
		logger:         logger,
		userHandler:    handler.NewUserHandler(userService),
		authMiddleware: authMiddleware,
	}

	router.setupMiddlewares()
	router.setupRouter()

	return router
}

func (r *Router) setupMiddlewares() {
	r.engine.Use(middleware.Logging(r.logger))
	r.engine.Use(middleware.Recovery(r.logger))
	r.engine.Use(middleware.CORS())
}

func (r *Router) setupRouter() {
	public := r.engine.Group("/api/v1")
	{
		user := public.Group("/user")
		{
			user.GET("/hello", r.userHandler.SayHello)
		}
	}
}
