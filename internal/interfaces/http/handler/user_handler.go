package handler

import (
	"dovenet/user-service/internal/application"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	BaseHandler
	userService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) SayHello(c *gin.Context) {
	h.Success(c, "Hello World!")
}
