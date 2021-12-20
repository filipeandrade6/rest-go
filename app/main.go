package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/filipeandrade6/rest-go/app/handlers"
	"github.com/filipeandrade6/rest-go/app/handlers/camera"
	"github.com/filipeandrade6/rest-go/app/infile"
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

	db := infile.New()

	// TODO load cores
	cameraCore := camera.NewCamera(log, db)

	app := handlers.NewApp(log, cameraCore)

	http.ListenAndServe(":9001", app)
}
