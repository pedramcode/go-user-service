package router

import (
	"dovenet/user-service/internal/application"
	"dovenet/user-service/internal/interfaces/http/handler"
	"dovenet/user-service/internal/interfaces/http/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	_ "dovenet/user-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	engine         *gin.Engine
	userHandler    *handler.UserHandler
	authMiddleware *middleware.AuthMiddleware
	healthHandler  *handler.HealthHandler
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
		healthHandler:  handler.NewHealthHandler(),
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
	system := r.engine.Group("")
	{
		system.GET("/health", r.healthHandler.Check)
		system.GET("/ready", r.healthHandler.Ready)
		system.GET("/live", r.healthHandler.Live)
	}

	if os.Getenv("ENV") != "production" {
		r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	public := r.engine.Group("/api/v1")
	{
		user := public.Group("/user")
		{
			user.GET("/hello", r.userHandler.SayHello)
		}
	}
}
