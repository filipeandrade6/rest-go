package core

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/filipeandrade6/rest-go/pkg/database"
	"github.com/filipeandrade6/rest-go/pkg/logger"
)

var build = "develop" // TODO entender isso aqui

func Run() {
	// ==================================================================================
	// Load logger

	log, err := logger.New("rest")
	if err != nil {
		fmt.Println(err)
		return
	}

	// ==================================================================================
	// Load configuration

	cfg := "cfg" // TODO load config
	log.Infow("startup", "config", fmt.Sprintf("%+v", cfg))

	// ==================================================================================
	// App stating

	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	// ==================================================================================
	// Database support

	log.Infow("startup", "status", "initializing database support", "host", cfg) // TODO cfg.Database.Host

	dbCfg := database.Config{
		User:         cfg,
		Password:     cfg,
		Host:         cfg,
		Name:         cfg,
		MaxIDLEConns: 1,
		MaxOpenConns: 1,
		DisableTLS:   true,
	}

	db, err := database.Connect(dbCfg)
	if err != nil {
		return Tfmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Infow("shutdown", "status", "stopping database support", "host", cfg.Database.Host)
		db.Close()
	}()

	// ==================================================================================
	// Start service

	log.Infow("startup", "status", "initializing gerencia")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	go func() {
		lis, err := net.Listen(cfg.Service.Conn, fmt.Sprintf(":%d", cfg.Service.Port))
		if err != nil {
			log.Errorw("startup", "status", "could not open socket", cfg.Service.Conn, cfg.Service.Port, "ERROR", err)
		}

		log.Infow("startup", "status", "gRPC server started", cfg.Service.Address)
		serverErrors <- grpcServer.Serve(lis)
	}()

	// ==================================================================================
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

		grpcServer.GracefulStop()
	}

	return
}
