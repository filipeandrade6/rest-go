package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/filipeandrade6/rest-go/app/handlers/camera"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/treastech/logger"
	"go.uber.org/zap"
)

func NewRouter(log *zap.SugaredLogger, camera *camera.Camera) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(logger.Logger(log.Desugar()))                  // ! utilizar?
	r.Use(render.SetContentType(render.ContentTypeJSON)) // ! utilizar?
	r.Use(middleware.Timeout(60 * time.Second))

	// Protected routes
	r.Group(func(r chi.Router) {
		// TODO userStore := userstore.New(log?, db?)
		// TODO r.Mount("/users", usersresource.New(userStore, tokenAuth <- passado como dependencia de NewRouter))
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hi!"))
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(b))
		})
	})

	return r
}
