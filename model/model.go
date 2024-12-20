package models

type Message struct {
	Message string
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
