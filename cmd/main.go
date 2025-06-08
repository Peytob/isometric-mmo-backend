package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	ini "isonetric-mmo-backend/init"
	log "log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var err error

	config, err := ini.Config()
	if err != nil {
		panic(err)
	}

	if err := ini.Logging(config); err != nil {
		panic(err)
	}

	log.Info("initializing database")

	// todo ini.SqlDatabase(config.Database.Sql)

	log.Info("initializing application")

	_, err = ini.Application(config)
	if err != nil {
		log.Error("application structure cant be initialized", "err", err.Error())
		panic(err)
	}

	httpServer, err := ini.HttpServer(config.Server, http.NewServeMux())
	if err != nil {
		log.Error("http server cant be initialized", "err", err.Error())
		panic(err)
	}

	// Simplest gracefully shutdown realisation

	log.Info("starting application services")

	rootCtx, rootCtxStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer rootCtxStop()

	errorGroup, gCtx := errgroup.WithContext(rootCtx)

	errorGroup.Go(func() error {
		log.Info("starting http server listening", "port", config.Server.Port)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error while listening on http httpServer", "err", err.Error())
			return err
		}
		return nil
	})

	errorGroup.Go(func() error {
		<-gCtx.Done()
		log.Info("graceful shutdown on http server started")
		return httpServer.Shutdown(context.Background())
	})

	<-rootCtx.Done()
	if err := errorGroup.Wait(); err != nil {
		log.Error("error while waiting for graceful shutdown", "err", err.Error())
	}

	log.Info("shutdown complete")
}
