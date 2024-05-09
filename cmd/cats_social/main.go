package main

import (
	"context"

	"github.com/syamsulhudauul/cats-social-service/internal/app/config"
	"github.com/syamsulhudauul/cats-social-service/internal/app/delivery"
	"github.com/syamsulhudauul/cats-social-service/internal/app/repository"
	"github.com/syamsulhudauul/cats-social-service/internal/app/service"
	"github.com/syamsulhudauul/cats-social-service/internal/middleware"
	"github.com/syamsulhudauul/cats-social-service/internal/pkg/log"
	"github.com/syamsulhudauul/cats-social-service/pkg/version"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := context.Background()
	// Load the configuration file
	cfg, err := config.Load(ctx)
	if err != nil {
		panic(err)
	}
	// init logger
	logger, err := log.New(zapcore.DebugLevel, version.ServiceID, version.Version)
	if err != nil {
		panic(err)
	}
	res, err := initResources(ctx, cfg, logger)
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepo(res.DB)
	s := service.NewService(cfg, logger, repo)
	h := delivery.New(s)

	secretKey := cfg.JWTSecret
	authMiddleware := middleware.AuthMiddleware(secretKey)

	initRouter(h, authMiddleware)
}
