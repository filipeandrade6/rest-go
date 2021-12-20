package handlers

import (
	"net/http"

	"github.com/filipeandrade6/rest-go/app/handlers/camera"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type App struct {
	log    *zap.SugaredLogger
	camera *camera.Camera
}

func NewApp(log *zap.SugaredLogger, camera *camera.Camera) http.Handler {
	app := chi.NewRouter()

	return &App{
		log:    log,
		camera: camera,
	}
}
