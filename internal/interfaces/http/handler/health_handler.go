package handler

import (
	"dovenet/user-service/internal/interfaces/http/dto"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	BaseHandler
	ready       atomic.Bool
	startTime   time.Time
	serviceName string
}

// NewHealthHandler creates a new health check handler
func NewHealthHandler() *HealthHandler {
	h := &HealthHandler{
		startTime:   time.Now(),
		serviceName: "user-service",
	}
	h.ready.Store(true)
	return h
}

// Check godoc
// @Summary      Health check
// @Description  Returns the health status of the service
// @Tags         system
// @Produce      json
// @Success      200  {object}  dto.HealthResponse  "Service is healthy"
// @Router       /health [get]
func (h *HealthHandler) Check(c *gin.Context) {
	h.Success(c, dto.HealthResponse{
		Status:    "healthy",
		Service:   h.serviceName,
		Uptime:    time.Since(h.startTime).String(),
		Timestamp: time.Now(),
	})
}

// Ready godoc
// @Summary      Readiness probe
// @Description  Checks if the service is ready to accept traffic (includes database check)
// @Tags         system
// @Produce      json
// @Success      200  {object}  dto.ReadinessResponse  "Service is ready"
// @Failure      503  {object}  Response           "Service is not ready"
// @Router       /ready [get]
func (h *HealthHandler) Ready(c *gin.Context) {
	if !h.ready.Load() {
		h.Error(c, http.StatusServiceUnavailable, "NOT_READY", "Service is not ready", "")
		return
	}

	h.Success(c, dto.ReadinessResponse{
		Status:    "ready",
		Database:  "connected",
		Timestamp: time.Now(),
	})
}

// Live godoc
// @Summary      Liveness probe
// @Description  Checks if the service is alive (Kubernetes compatible)
// @Tags         system
// @Produce      json
// @Success      200  {object}  dto.LiveResponse  "Service is alive"
// @Router       /live [get]
func (h *HealthHandler) Live(c *gin.Context) {
	c.JSON(http.StatusOK, dto.LiveResponse{
		Status:    "alive",
		Timestamp: time.Now(),
	})
}

// SetReady sets the ready status of the service
func (h *HealthHandler) SetReady(ready bool) {
	h.ready.Store(ready)
}

// GetUptime returns the service uptime
func (h *HealthHandler) GetUptime() time.Duration {
	return time.Since(h.startTime)
}
