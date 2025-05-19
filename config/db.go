package config

//import (
//	"database/sql"
//	"fmt"
//	"log"
//	"os"
//
//	"github.com/joho/godotenv"
//)

//func setupDB() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}

//	dbHost := os.Getenv("DB_HOST")
//	dbPort := os.Getenv("DB_PORT")
//	dbUser := os.Getenv("DB_USER")
//	dbPassword := os.Getenv("DB_PASSWORD")
//	dbName := os.Getenv("DB_NAME")

//	connectionStr := fmt.Sprintf(
//		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

//	dbConnection, err := sql.Open("postgres", connectionStr)
//	if err != nil {
//		log.Fatal("Error connecting to the database:", err)
//	}
//}
