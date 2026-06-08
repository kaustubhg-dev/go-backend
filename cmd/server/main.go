package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-backend/config"
	"go-backend/internal/database"
	"go-backend/internal/handler"
	"go-backend/internal/models"
	"go-backend/internal/repository"
	"go-backend/internal/router"
	"go-backend/internal/service"
	"go-backend/pkg/logger"
)

func main() {


	// 1. Load config
	fmt.Println("Loading configuration...")
	cfg := config.Load()

	// 2. Init logger
	log := logger.New(cfg.App.Env)
	defer log.Sync()

	// 3. Connect databases
	pgDB := database.ConnectPostgres(cfg)
	fmt.Printf("Connected to PostgreSQL at %s:%s\n", cfg.Postgres.Host, cfg.Postgres.Port)
	// mysqlDB := database.ConnectMySQL(cfg)
	// mongoDB := database.ConnectMongo(cfg)
	// redisDB := database.ConnectRedis(cfg)

	// 4. Auto-migrate
	if err := pgDB.AutoMigrate(
		&models.User{},
		&models.Product{},
	); err != nil {
		log.Sugar().Fatalf("Migration failed: %v", err)
	}

	// 5. Dependency Injection
	userRepo := repository.NewUserRepository(pgDB)
	userSvc := service.NewUserService(userRepo, cfg)
	userH := handler.NewUserHandler(userSvc)

	// productRepo := repository.NewProductRepository(pgDB)
	// productSvc := service.NewProductService(productRepo, cfg)
	// productH := handler.NewProductHandler(productSvc)

	handlers := router.Handlers{
		User: userH,
		// Product: productH,
	}

	// 6. Setup router
	r := router.Setup(cfg, log, handlers)

	// 7. HTTP server config
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.App.Port),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 8. Start server
	go func() {
		log.Sugar().Infof(
			"Server started on port %s [env: %s]",
			cfg.App.Port,
			cfg.App.Env,
		)

		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Sugar().Fatalf("Server error: %v", err)
		}
	}()

	// 9. Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Sugar().Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Sugar().Errorf("Server forced shutdown: %v", err)
	}

	log.Sugar().Info("Server exited cleanly")
}