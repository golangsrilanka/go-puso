package server

import (
	"go.uber.org/fx"

	"github.com/golangsrilanka/go-puso/api/handler"
	"github.com/golangsrilanka/go-puso/api/router"
	"github.com/golangsrilanka/go-puso/database"
	"github.com/golangsrilanka/go-puso/transact/puso"
)

var Init = fx.Options(
	database.Init,
	router.Init,
	handler.Init,
	puso.Init,
)
