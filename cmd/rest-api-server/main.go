package main

import (
	"context"

	"isOdin/RestApi/configs"
	"isOdin/RestApi/internal/database/postgresql"
	"isOdin/RestApi/internal/handler"
	"isOdin/RestApi/internal/middleware"
	"isOdin/RestApi/internal/repository"
	"isOdin/RestApi/internal/server"
	"isOdin/RestApi/internal/service"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

// @title Todo App API
// @version 1.0
// @description REST API Server
// @tag.name auth
// @tag.name lists
// @tag.name items
// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	router := echo.New()

	// Config
	var cfg configs.Config
	if err := env.Parse(&cfg); err != nil {
		router.Logger.Error("Error whie initialize config:, ", err.Error())
		return
	}

	// Database
	DB, err := postgresql.NewPostgresDB(&cfg)
	if err != nil {
		router.Logger.Error("failed to initialize db: %s", err.Error())
	}
	defer DB.Close()

	repository := repository.NewRepository(DB)                       //----- Repository -----
	service := service.NewService(&cfg.InternalConfig, repository)   // ----- Service -----
	validate := validator.New(validator.WithRequiredStructEnabled()) // ----- Validator -----
	middleware := middleware.NewMiddleware(&cfg.InternalConfig)      // ----- Middleware -----
	handler := handler.NewHandler(validate, service)                 // ----- Handler -----
	server.NewRouter(router, middleware, handler)                    // ----- Routing -----

	// Server start
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := server.RunServer(router, &ctx, ":8000"); err != nil {
		router.Logger.Error("error while running server %s", err.Error())
	}
}
