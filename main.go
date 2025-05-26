package main

import (
	"LearnGo/api"
	"LearnGo/config"
	"LearnGo/facade"
	"LearnGo/migrations"
	"LearnGo/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dbConnection := config.SetupDB()
	defer dbConnection.Close()

	migrations.Migrate(dbConnection)

	repo := repository.UsuarioRepository{DB: dbConnection}
	facade.Init(repo)

	router := mux.NewRouter()
	api.UsuarioRoutes(router)

	log.Println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
