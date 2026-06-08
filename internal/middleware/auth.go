package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-backend/config"
	"go-backend/internal/utils"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "missing token")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid auth format")
			c.Abort()
			return
		}

		token := parts[1]

		claims, err := utils.ValidateToken(token, cfg.App.Secret)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		roleVal, exists := c.Get("role")
		if !exists {
			utils.ErrorResponse(c, http.StatusForbidden, "role missing")
			c.Abort()
			return
		}

		role := roleVal.(string)

		for _, r := range roles {
			if r == role {
				c.Next()
				return
			}
		}

		utils.ErrorResponse(c, http.StatusForbidden, "forbidden")
		c.Abort()
	}
}