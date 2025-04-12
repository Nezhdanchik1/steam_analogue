package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
			return
		}

		roleStr, ok := role.(string)
		if !ok || roleStr != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient rights"})
			return
		}

		c.Next()
	}
}
