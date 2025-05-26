package usuario

import (
	"LearnGo/domain"
	usuarioFacade "LearnGo/facade"
	"LearnGo/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body := domain.Usuario{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, 400, "Erro ao decodificar usuario")
		return
	}

	if err := body.Validate(); err != nil {
		response.Error(w, 400, "Erro ao validar usuario")
		return
	}

	ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, 400, "ID de usuario inválido")
		return
	}

	body.ID = ID

	err = usuarioFacade.Get().Update(body)
	if err != nil {
		response.Error(w, 500, "Erro ao atualizar usuario")
		return
	}

	response.JSON(w, 200, body)
}
