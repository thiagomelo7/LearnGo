package api

import (
	"LearnGo/domain"
	"LearnGo/facade"
	"encoding/json"
	"fmt"
	"net/http"
)

var usuarioFacade facade.UsuarioFacade

func PostUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body := domain.Usuario{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}
	if err := body.Validate(); err != nil {
		http.Error(w, "Erro de validação: "+err.Error(), http.StatusBadRequest)
		return
	}
	ID, err := usuarioFacade.InserirUsuario(r.Context(), body)
	if err != nil {
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("ID: %d", ID)))
}
