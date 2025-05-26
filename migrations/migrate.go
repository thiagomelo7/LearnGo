package migrations

import (
	"database/sql"
	"fmt"
)

func Migrate(dbConnection *sql.DB) {
	_, err := dbConnection.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)
	if err != nil {
		fmt.Println("Error creating extension pgcrypto:", err)
		return
	}

	query := `
    CREATE TABLE IF NOT EXISTS usuarios (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );
    `
	_, err = dbConnection.Exec(query)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Migration completed")
}
