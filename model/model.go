package model

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
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// estrutura para o Banco de Dados
type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
