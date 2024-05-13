package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/g-vinokurov/pyramidum-backend-api-router/internal/app"
	"github.com/g-vinokurov/pyramidum-backend-api-router/internal/config"
	"github.com/g-vinokurov/pyramidum-backend-api-router/internal/env"
)

func main() {
	log := setupLogger()

	envVars := env.MustLoadEnv()

	cfg := config.MustLoadConfig(envVars.ConfigPath)

	application, err := app.NewApp(log, cfg, envVars)
	if err != nil {
		log.Error("failed to init application", err)
		os.Exit(1)
	}

	log.Info("starting application", slog.String("address", cfg.HTTPServer.Address))

	go func() {
		if err := application.Run(); err != nil {
			log.Error("server stopped with error")
		}
	}()

	log.Info("application started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTPServer.ShutdownTimeout)
	defer cancel()

	if err := application.Stop(ctx); err != nil {
		log.Error("failed to stop application")

		return
	}

	log.Info("application stopped")
}

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
}
