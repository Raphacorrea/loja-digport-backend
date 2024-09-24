package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Raphacorrea/loja-digport-backend/model"
)

func CriaUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	err := model.CriaUsuario(usuario)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
