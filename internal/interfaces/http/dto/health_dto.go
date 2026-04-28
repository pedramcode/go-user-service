package dto

import "time"

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status" example:"healthy" description:"Service health status"`
	Service   string    `json:"service" example:"user-service" description:"Service name"`
	Version   string    `json:"version" example:"1.0.0" description:"Service version"`
	Uptime    string    `json:"uptime" example:"24h30m15s" description:"Service uptime"`
	Timestamp time.Time `json:"timestamp" example:"2024-01-01T00:00:00Z" description:"Check timestamp"`
}

// ReadinessResponse represents the readiness check response
type ReadinessResponse struct {
	Status    string    `json:"status" example:"ready" description:"Ready status"`
	Database  string    `json:"database" example:"connected" description:"Database connection status"`
	Timestamp time.Time `json:"timestamp" example:"2024-01-01T00:00:00Z" description:"Check timestamp"`
}

// LiveResponse represents the liveness check response
type LiveResponse struct {
	Status    string    `json:"status" example:"alive" description:"Liveness status"`
	Timestamp time.Time `json:"timestamp" example:"2024-01-01T00:00:00Z" description:"Check timestamp"`
}
