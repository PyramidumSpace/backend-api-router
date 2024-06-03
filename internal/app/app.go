package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/g-vinokurov/pyramidum-backend-api-router/internal/config"
	"github.com/g-vinokurov/pyramidum-backend-api-router/internal/env"
	authLog "github.com/g-vinokurov/pyramidum-backend-api-router/internal/http-server/handlers/auth/login"
	authReg "github.com/g-vinokurov/pyramidum-backend-api-router/internal/http-server/handlers/auth/register"
	taskGetByUID "github.com/g-vinokurov/pyramidum-backend-api-router/internal/http-server/handlers/tasks/get"
	taskGet "github.com/g-vinokurov/pyramidum-backend-api-router/internal/http-server/handlers/tasks/getp"
	taskPost "github.com/g-vinokurov/pyramidum-backend-api-router/internal/http-server/handlers/tasks/post"
	taskPut "github.com/g-vinokurov/pyramidum-backend-api-router/internal/http-server/handlers/tasks/put"
	logImpl "github.com/g-vinokurov/pyramidum-backend-api-router/internal/service/auth/login"
	regImpl "github.com/g-vinokurov/pyramidum-backend-api-router/internal/service/auth/register"
	taskGetByUidImpl "github.com/g-vinokurov/pyramidum-backend-api-router/internal/service/tasks/get"
	taskGetImpl "github.com/g-vinokurov/pyramidum-backend-api-router/internal/service/tasks/getp"
	taskPostImpl "github.com/g-vinokurov/pyramidum-backend-api-router/internal/service/tasks/post"
	taskPutImpl "github.com/g-vinokurov/pyramidum-backend-api-router/internal/service/tasks/put"
	"github.com/gin-gonic/gin"
)

type App struct {
	srv *http.Server
}

func NewApp(log *slog.Logger, cfg *config.Config, envVars *env.Env) (*App, error) {
	const op = "app.NewApp"

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	newRegService, err := regImpl.NewService(cfg.GRPCAuthServer.Address)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	newLogService, err := logImpl.NewService(cfg.GRPCAuthServer.Address)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	newTaskPostService, err := taskPostImpl.NewService(cfg.GRPCTaskServer.Address)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	newTaskPutService, err := taskPutImpl.NewService(cfg.GRPCTaskServer.Address)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	newTaskGetByUidService, err := taskGetByUidImpl.NewService(cfg.GRPCTaskServer.Address)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	newTaskGet, err := taskGetImpl.NewService(cfg.GRPCTaskServer.Address)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	router.POST("/api/auth/register", authReg.MakeGetHandlerFunc(log, newRegService))
	router.POST("/api/auth/login", authLog.MakeGetHandlerFunc(log, newLogService))
	router.POST("/api/tasks", taskPost.MakeGetHandlerFunc(log, newTaskPostService))
	router.PUT("/api/tasks", taskPut.MakeGetHandlerFunc(log, newTaskPutService))
	router.GET("/api/tasks", taskGetByUID.MakeGetHandlerFunc(log, newTaskGetByUidService))
	router.GET("/api/tasks/:task_id", taskGet.MakeGetHandlerFunc(log, newTaskGet))

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
