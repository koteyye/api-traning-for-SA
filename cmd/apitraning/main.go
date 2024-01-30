package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/koteyye/api-traning-for-SA/internal/config"
	"github.com/koteyye/api-traning-for-SA/internal/handlers"
	"github.com/koteyye/api-traning-for-SA/internal/server"
)

// @Title APITraningSA
// @Description Локальный веб-сервер для практики системных аналитиков по работе с документацией Swagger
// @Version 1.0

// @Contact.email arpetuk1@mts.ru

// @BasePath /api/traning
// @Host localhost:8081

// @Tag.name Health
// @Tag.desctiption "Состояние сервиса"

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var cfg config.Config
	if err := cfg.GetConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	logger := newLogger(&cfg)
	handlers := handlers.NewHandlers(*logger)

	runServer(ctx, &cfg, handlers, *logger)
}

func newLogger(c *config.Config) *slog.Logger {
	opts := &slog.HandlerOptions{Level: c.Level}
	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler)
}

func runServer(ctx context.Context, cfg *config.Config, handler *handlers.Handlers, logger slog.Logger) error {
	httpServer := new(server.Server)
	logger.Info("Starting server")
	go func() {
		if err := httpServer.Run(cfg.Server, handler.InitRoutes()); err != nil {
			logger.Error(err.Error())
		}
	}()

	<-ctx.Done()

	logger.Info("Shutting down server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shotdown: %w", err)
	}

	return nil
}
