package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/g-vinokurov/pyramidum-backend-api-router/internal/config"
	"github.com/g-vinokurov/pyramidum-backend-api-router/internal/env"
	authReg "github.com/g-vinokurov/pyramidum-backend-api-router/internal/http-server/handlers/auth/register"
	regImpl "github.com/g-vinokurov/pyramidum-backend-api-router/internal/service/auth/register"
	"github.com/gin-gonic/gin"
)

type App struct {
	srv *http.Server
}

func NewApp(log *slog.Logger, cfg *config.Config, envVars *env.Env) (*App, error) {
	const op = "app.NewApp"

	router := gin.Default()
	newService, err := regImpl.NewService(cfg.GrpcAuthServer.Address)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	router.POST("/auth/register", authReg.MakeGetHandlerFunc(log, newService))

	srv := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router.Handler(),
	}

	return &App{srv: srv}, nil
}

func (a *App) Run() error {
	const op = "app.Run"

	if err := a.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	const op = "app.Stop"

	if err := a.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
