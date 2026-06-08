package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go-backend/config"
	"go-backend/internal/handler"
	"go-backend/internal/middleware"
)

type Handlers struct {
	User    *handler.UserHandler
	Product *handler.ProductHandler
}

func Setup(cfg *config.Config, logger *zap.Logger, h Handlers) *gin.Engine {

	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Global middleware
	r.Use(middleware.ZapLogger(logger))
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimiter(cfg.Rate.RPS, cfg.Rate.Burst))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// =========================
	// API v1
	// =========================
	v1 := r.Group("/api/v1")

	// Public auth routes
	auth := v1.Group("/auth")
	{
		auth.POST("/register", h.User.Register)
		auth.POST("/login", h.User.Login)
	}

	// Protected routes
	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// User routes
		users := protected.Group("/users")
		{
			users.GET("/:id", h.User.GetUser)
			users.PATCH("/:id", h.User.UpdateUser)
		}

		// Product routes (authenticated users)
		products := protected.Group("/products")
		{
			products.GET("/", h.Product.GetProducts)
			products.GET("/:id", h.Product.GetProduct)
		}

		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware("admin"))
		{
			admin.GET("/users", h.User.GetUsers)
			admin.DELETE("/users/:id", h.User.DeleteUser)

			adminProducts := admin.Group("/products")
			{
				adminProducts.POST("/", h.Product.CreateProduct)
				adminProducts.PUT("/:id", h.Product.UpdateProduct)
				adminProducts.DELETE("/:id", h.Product.DeleteProduct)
			}
		}
	}

	return r
}