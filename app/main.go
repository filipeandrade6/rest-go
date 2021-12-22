package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/filipeandrade6/rest-go/app/handlers"
	"github.com/filipeandrade6/rest-go/pkg/database/inmemory"
	"github.com/filipeandrade6/rest-go/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	log, err := logger.New("rest-chi")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorw("startup", "error", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// TODO Parse config

	db := inmemory.New()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// TODO load handler
	api := handlers.NewAPI(log, db)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// * Start the service listening for api requests.
	go func() {
		log.Infow("startup", "status", "api router started", "host", "API ADDR AQUI!!!!") // TODO get addr from config
		serverErrors <- http.ListenAndServe(":9001", app)
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

		// TODO graceful stop the application
	}

	return nil
}
