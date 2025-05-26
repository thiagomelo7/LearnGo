package usuario

import (
	"LearnGo/domain"
	usuarioFacade "LearnGo/facade"
	"LearnGo/response"
	"encoding/json"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

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

	ID, err := usuarioFacade.Get().Create(body)
	if err != nil {
		response.Error(w, 500, err.Error())
		return
	}
	body.ID = int(ID)
	response.JSON(w, 200, body)
}
