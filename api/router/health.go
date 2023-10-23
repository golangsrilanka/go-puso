package router

import (
	"github.com/go-chi/chi"

	"github.com/golangsrilanka/go-puso/api/handler"
)

func HealthRoute() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handler.GetHealth)

	return r
}
