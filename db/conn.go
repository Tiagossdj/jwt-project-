package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// ConnDB inicia a conexão com o postgreSQL
func ConnDB() (*sqlx.DB, error) {
	var err error
	psqlInfo := "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"
	DB, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error establishing a database connection:%v", err)
	}
	fmt.Println("Connect to Database!")
	return DB, nil

}
