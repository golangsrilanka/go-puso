package server

import (
	"go.uber.org/fx"

	"github.com/GolangSriLanka/go-puso/api/handler"
	"github.com/GolangSriLanka/go-puso/api/router"
	"github.com/GolangSriLanka/go-puso/database"
	"github.com/GolangSriLanka/go-puso/transact/puso"
)

var Init = fx.Options(
	database.Init,
	router.Init,
	handler.Init,
	puso.Init,
)
