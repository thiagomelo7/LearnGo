package usuario

import (
	usuarioFacade "LearnGo/facade"
	"LearnGo/response"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	IDStr := mux.Vars(r)["id"]
	if err := validation.Validate(IDStr, is.Int); err != nil {
		response.Error(w, 400, "ID de usuario inválido")
		return
	}

	id, err := strconv.Atoi(IDStr)
	if err != nil {
		response.Error(w, 400, "ID de usuario inválido")
		return
	}

	result, err := usuarioFacade.Get().Get(id)
	if err != nil {
		response.Error(w, 404, err.Error())
		return
	}

	response.JSON(w, 200, result)
}
