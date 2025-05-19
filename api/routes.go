package api

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Registro dos endpoints
	mux.HandleFunc("/usuarios", PostUsuarioHandler)

	return mux
}
