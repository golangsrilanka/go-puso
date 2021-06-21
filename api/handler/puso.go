package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/GolangSriLanka/go-puso/transact/puso"
)

// CreatePuso
// @Summary Create a new puso
// @Tags Puso
// @Accept json
// @Produce json
// @Param data body	puso.Puso	true	"data"
// @Success 200 {string} string	"successfully puso created"
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router	/puso	[post]
func CreatePuso(w http.ResponseWriter, r *http.Request) {
	t := puso.Puso{}

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.Save(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "successfully puso created")
}

// PusoList
// @Summary Get puso list
// @Tags Puso
// @Accept json
// @Produce json
// @Success 200 {object} []puso.Puso
// @Failure 404 {string} string
// @Router	/puso	[get]
func PusoList(w http.ResponseWriter, r *http.Request) {
	t := puso.Puso{}
	list, err := t.GetList()

	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, list)
}

// GetPuso
// @Summary Get puso
// @Tags Puso
// @Accept json
// @Produce json
// @Param   id	path	string	true	"ID"
// @Success 200 {object} puso.Puso
// @Failure 404 {string} string
// @Router /puso/{id} [get]
func GetPuso(w http.ResponseWriter, r *http.Request) {
	t := puso.Puso{}
	ID := chi.URLParam(r, "id")
	o, err := t.Get(ID)

	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, o)
}

// DeletePuso
// @Summary Delete puso
// @Tags Puso
// @Accept json
// @Produce json
// @Param   id	path	string	true	"ID"
// @Success 200 {nil}	nil
// @Failure 404 {string}	string
// @Router /puso/{id} [delete]
func DeletePuso(w http.ResponseWriter, r *http.Request) {
	t := puso.Puso{}
	ID := chi.URLParam(r, "id")

	if err := t.Delete(ID); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}

// UpdatePuso
// @Summary Update puso
// @Tags Puso
// @Description Update puso
// @Accept  json
// @Produce  json
// @Param   id	path	string	true	"ID"
// @Success 200 {string} string	"successfully updated"
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /puso/{id} [put]
func UpdatePuso(w http.ResponseWriter, r *http.Request) {
	t := puso.Puso{}
	ID := chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.Update(ID); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "successfully updated")
}
