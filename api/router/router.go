package router

import (
	"gorm.io/gorm"

	"github.com/go-chi/chi"
)

type RouterRepo interface {
	Route() chi.Router
}

type Router struct {
	db *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		db: db,
	}
}

func (o *Router) Route() chi.Router {
	r := chi.NewRouter()

	r.Mount("/puso", o.PusoRouter())

	return r
}
