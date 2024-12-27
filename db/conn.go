package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ConnDB inicia a conexão com o postgreSQL
func ConnDB() (*sqlx.DB, error) {
	var err error
	psqlInfo := "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("Error establishing a database connection: %w", err)
	}
	fmt.Println("Connect to Database!")
	return db, nil

}
