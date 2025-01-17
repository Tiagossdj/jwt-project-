package integration_tests

import (
	"database/sql"
	"log"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// func para conectar-se ao banco de dados principal para criar o banco de teste
func SetupTestDataBase(t *testing.T) *sqlx.DB {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=postgres sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect test Database:%v", err)
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS jwtdb_test")
	if err != nil {
		t.Fatalf("Failed to drop existing test database: %v", err)
	}
	//cria o banco de dados
	_, err = db.Exec("create database jwtdb_test")
	if err != nil && !strings.Contains(err.Error(), "Already exists") {
		t.Fatalf("Failed to create database:%v", err)
	}

	//conecta-se ao banco de dados teste
	testDb, err := sqlx.Open("postgres", "user=postgres password=1234 dbname=jwtdb_test sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect test Database:%v", err)
	}

	// cria tabelas e outras estruturas para o banco de dados teste.
	createTables(testDb)

	return testDb
}

func createTables(db *sqlx.DB) {
	_, err := db.Exec(`
    CREATE TABLE usuario (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(100), 
        email VARCHAR(255) UNIQUE,
        password VARCHAR(255) 
    )
	`)
	if err != nil {
		log.Fatalf("Error to create table:%v", err)
	}
}

// func para limpeza do banco de dados após os testes
func CleanUpTestDataBase(db *sqlx.DB) {
	db.Close()
	//conectar ao banco de dados principal para excluir o banco de testes
	conn, err := sqlx.Open("postgres", "user=postgres password=1234 sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database for cleanup: %v", err)
	}
	defer conn.Close()

	_, err = conn.Exec("drop database jwtdb_test")
	if err != nil {
		log.Fatalf("Failed to drop test database:%v", err)
	}
}

// func para criar exemplos no banco de dados para testes
func PopulateTestData(db *sqlx.DB) {
	_, err := db.Exec(`insert into usuario (nome, email, password) values ('Test user', 'test@example.com', 'HashedPassword')`)

	if err != nil {
		log.Fatalf("Failed to Populate Database:%v", err)
	}
}

func CreateUserTest(db *sqlx.DB) {

	// Gerar uma senha hash para o usuario
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password:%v", err)
	}

	_, err = db.Exec(`insert into usuario (nome, email, password) values ('Test user', 'TestUser@example.com', '$1')`, hashedPassword)

	if err != nil {
		log.Fatalf("Failed to Populate Database:%v", err)
	}

}
