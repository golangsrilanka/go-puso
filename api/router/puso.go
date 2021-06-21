package router

import (
	"github.com/go-chi/chi"

	"github.com/GolangSriLanka/go-puso/api/handler"
)

func PusoRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreatePuso)
	r.Get("/", handler.PusoList)
	r.Get("/{id}", handler.GetPuso)
	r.Put("/{id}", handler.UpdatePuso)
	r.Delete("/{id}", handler.DeletePuso)

	return r
}
