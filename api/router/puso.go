package router

import (
	"github.com/go-chi/chi"

	"github.com/GolangSriLanka/go-puso/api/handler"
	"github.com/GolangSriLanka/go-puso/transact/puso"
)

func (o *Router) PusoRouter() chi.Router {
	r := chi.NewRouter()
	pusoHandler := handler.NewPusoHandler(puso.NewPusoRepo(o.db))

	r.Post("/", pusoHandler.CreatePuso)
	r.Get("/", pusoHandler.PusoList)
	r.Get("/{id}", pusoHandler.GetPuso)
	r.Put("/{id}", pusoHandler.UpdatePuso)
	r.Delete("/{id}", pusoHandler.DeletePuso)

	return r
}
