package handler

import (
	"github.com/GolangSriLanka/go-puso/config"
	"net/http"
)

// GetHealth godoc
// @Summary Returns health of the service
// @Router /healthz [get]
func GetHealth(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, "Go Puso up and running version: "+config.GetEnv("go-puso.VERSION"))
}
