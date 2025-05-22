package main

import (
	"LearnGo/config"
	"log"
	"net/http"
)

func main() {

	dbConnection := config.SetupDB()

	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
