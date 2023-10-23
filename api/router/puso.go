package router

import (
	"github.com/go-chi/chi"
)

func (o *Router) PusoRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", o.pusoHandler.CreatePuso)
	r.Get("/", o.pusoHandler.PusoList)
	r.Get("/{id}", o.pusoHandler.GetPuso)
	r.Put("/{id}", o.pusoHandler.UpdatePuso)
	r.Delete("/{id}", o.pusoHandler.DeletePuso)

	return r
}
