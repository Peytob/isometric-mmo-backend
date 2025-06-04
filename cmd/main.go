package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"isonetric-mmo-backend/initialization"
	log "log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := initialization.Config()
	if err != nil {
		panic(err)
	}

	if err := initialization.Logging(config); err != nil {
		panic(err)
	}

	log.Info("configuration loaded. Initializing application")

	// Simplest gracefully shutdown realisation

	rootCtx, rootCtxStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer rootCtxStop()

	server := initialization.Server(config.Server, http.NewServeMux())

	errorGroup, gCtx := errgroup.WithContext(rootCtx)

	errorGroup.Go(func() error {
		log.Info("starting http server listening", "port", config.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error while listening on http server", "err", err)
			return err
		}
		return nil
	})

	errorGroup.Go(func() error {
		<-gCtx.Done()
		log.Info("graceful shutdown on http server started")
		return server.Shutdown(context.Background())
	})

	<-rootCtx.Done()
	if err := errorGroup.Wait(); err != nil {
		log.Error("error while waiting for graceful shutdown", "err", err)
	}

	log.Info("shutdown complete")
}
