package model

import (
	"time"
)

// estrutura para Mensagens
type Message struct {
	Message string
}

// estrutura para o Login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// estrutura para o registro
type RegisterRequest struct {
	Name       string    `json:"nome"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
}

// Estrutura para o Banco de Dados
type User struct {
	Id         uint      `json:"id" db:"id"`                 // db:"id" mapeia para a coluna 'id' no banco
	Name       string    `json:"nome" db:"nome"`             // db:"nome" mapeia para a coluna 'nome' no banco
	Email      string    `json:"email" db:"email"`           // db:"email" mapeia para a coluna 'email' no banco
	Password   string    `json:"password" db:"password"`     // db:"password" mapeia para a coluna 'password' no banco
	Created_at time.Time `json:"created_at" db:"created_at"` // db:"created_at" mapeia para a coluna 'created_at' no banco
}
