package api

import (
	"LearnGo/api/usuario"
	"net/http"

	"github.com/gorilla/mux"
)

func UsuarioRoutes(r *mux.Router) {

	r.HandleFunc("/usuario", usuario.Create).Methods(http.MethodPost)
	r.HandleFunc("/usuario/{id}", usuario.Get).Methods(http.MethodGet)
	r.HandleFunc("/usuario/{id}", usuario.Update).Methods(http.MethodPut)
	r.HandleFunc("/usuario/{id}", usuario.Delete).Methods(http.MethodDelete)
}
