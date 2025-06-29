package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/htooanttko/microservices/services/products/internal/config"
	"github.com/htooanttko/microservices/services/products/internal/handlers"
	"github.com/htooanttko/microservices/services/products/internal/repositories"
	"github.com/htooanttko/microservices/services/products/internal/services"
	"github.com/htooanttko/microservices/shared/pkg/db"
	"github.com/htooanttko/microservices/shared/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		logger.Error.Fatalf("Failed to load config: %v", err)
	}
	psql, err := db.InitPostgres(db.PostgresConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.Name,
		SSLMode:  cfg.Database.SSLMode,
	})
	if err != nil {
		logger.Error.Fatalf("Failed to connect to database: %v", err)
	}
	defer psql.Close()

	productRepo := repositories.NewProductRepository(psql)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Group(func(r chi.Router) {
		r.Get("/products", productHandler.GetAll)
		r.Get("/products/{id}", productHandler.GetByID)
		r.Post("/products", productHandler.Create)
	})

	v1Router.Get("/healthz", handlers.GetHealthz)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ErrorLog:     logger.Error,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	go func() {
		logger.Info.Printf("Starting server on port: %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	sig := <-c // [BLOCK] waiting for the signal
	logger.Info.Println("Got signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error.Fatalf("Server forced to shutdown: %v", err)
	}
}
