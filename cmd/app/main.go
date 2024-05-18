package main

import (
	"context"
	"github.com/google/logger"
	"os"
	"os/signal"
	"pastebin"
	"syscall"
)

func main() {
	srv := pastebin.NewServer(nil, "8080")
	if err := srv.Run(); err != nil {
		logger.Errorf("error with running server: %x", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error with shutting down server: %s", err.Error())
	}
}