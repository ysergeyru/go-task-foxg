package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ysergeyru/go-task-foxg/config"
	"github.com/ysergeyru/go-task-foxg/logger"
	"github.com/ysergeyru/go-task-foxg/pg"
	"github.com/ysergeyru/go-task-foxg/server"
)

// Version is a service version
const Version = "0.1"

func main() {
	// Read config
	cfg := config.Get()
	// Init logger
	logger := logger.Get()
	// Init Postgres
	pg.DB()
	// Catch signal for graceful shutdown
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	// Create new server instance
	server := server.New(cfg)
	httpServer := &http.Server{Addr: cfg.Addr, Handler: server.HTTPHandler()}
	// Run it
	go func() {
		logger.Infof("Listening on %s", cfg.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()
	// Handle graceful shutdown
	<-stop
	logger.Info("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
	logger.Info("Server gracefully stopped")
}
