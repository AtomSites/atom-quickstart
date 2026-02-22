package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AtomSites/atom-quickstart/internal/config"
	"github.com/AtomSites/atom-quickstart/internal/web"
)

func main() {
	port := config.GetEnvOrDefault("PORT", "3000")
	addr := "0.0.0.0:" + port
	slog.Info("starting server", "addr", addr)

	server := web.NewServer()

	go func() {
		if err := server.Start(addr); err != nil {
			slog.Error("server error", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("shutdown error", "error", err)
	}

	slog.Info("stopped")
}
