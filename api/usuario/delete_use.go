package usuario

import (
	usuarioFacade "LearnGo/facade"
	"LearnGo/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, 400, "ID inválido")
		return
	}

	err = usuarioFacade.Get().Delete(ID)
	if err != nil {
		response.Error(w, 500, err.Error())
		return
	}

	response.JSON(w, 204, nil)
}
