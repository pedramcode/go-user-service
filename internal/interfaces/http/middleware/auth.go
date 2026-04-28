// internal/interfaces/http/middleware/auth.go
package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{jwtSecret: jwtSecret}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Extract user ID from claims
		userID, ok := (*claims)["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Set("user_id", int32(userID))
		c.Set("is_superuser", (*claims)["is_superuser"])

		c.Next()
	}
}

func (m *AuthMiddleware) RequireSuperuser() gin.HandlerFunc {
	return func(c *gin.Context) {
		isSuperuser, exists := c.Get("is_superuser")
		if !exists || !isSuperuser.(bool) {
			c.AbortWithStatusJSON(403, gin.H{"error": "Superuser access required"})
			return
		}
		c.Next()
	}
}
