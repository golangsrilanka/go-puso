package handler

import (
	"net/http"

	"github.com/GolangSriLanka/go-puso/config"
)

// GetHealth godoc
// @Summary Returns health of the service
// @Router /healthz [get]
func GetHealth(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, "Go Puso up and running version: "+config.GetEnv("go-puso.VERSION"))
}
