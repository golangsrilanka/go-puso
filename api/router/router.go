package router

import (
	"github.com/go-chi/chi"
	"go.uber.org/fx"

	"github.com/golangsrilanka/go-puso/api/handler"
)

type RouterRepo interface {
	Route() chi.Router
}

type Config struct {
	fx.In

	PusoHandler *handler.PusoHandler
}

type Router struct {
	pusoHandler *handler.PusoHandler
}

func NewRouter(config Config) *Router {
	return &Router{
		pusoHandler: config.PusoHandler,
	}
}

func (o *Router) Route() chi.Router {
	r := chi.NewRouter()

	r.Mount("/puso", o.PusoRouter())

	return r
}
