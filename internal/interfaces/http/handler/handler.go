package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *MetaInfo   `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type MetaInfo struct {
	Page  int   `json:"page,omitempty"`
	Limit int   `json:"limit,omitempty"`
	Total int64 `json:"total,omitempty"`
}

type BaseHandler struct{}

func (h *BaseHandler) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func (h *BaseHandler) Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func (h *BaseHandler) Error(c *gin.Context, status int, code, message, details string) {
	c.AbortWithStatusJSON(status, Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

func (h *BaseHandler) BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		h.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error(), "")
		return false
	}
	return true
}

// GetUserIDFromContext extracts user ID from authenticated context
func (h *BaseHandler) GetUserIDFromContext(c *gin.Context) (int32, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		h.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "User not authenticated", "")
		return 0, false
	}

	id, ok := userID.(int32)
	if !ok {
		h.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Invalid user ID format", "")
		return 0, false
	}

	return id, true
}
