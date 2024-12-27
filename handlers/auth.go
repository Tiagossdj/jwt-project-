package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Tiagossdj/jwt-project-/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// chave secreta para assinar o token jwt
var jwtSecret = []byte("your secret key!")

// Login é o handler para o endPoint /auth/login
func Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Message{
			Message: "Invalid data!",
		})
	}

	// validação básica para substituir com a lógica real de autenticação depois
	if req.Email != "---" || req.Password != "---" {
		return c.JSON(http.StatusUnauthorized, model.Message{
			Message: "invalid credentials!",
		})
	}

	// geração de token JWT

	claims := &jwt.RegisteredClaims{
		Subject:   req.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 24h de duração.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "Error to generate Token!",
		})
	}
	return c.JSON(http.StatusOK, model.Message{
		Message: "Token:" + signedToken,
	})
}

// Register é o handler para o endPoint /auth/register
func Register(c echo.Context, db *sqlx.DB) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Message{
			Message: "Invalid Data",
		})
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, model.Message{
			Message: "You need to fill in all the fields!",
		})
	}

	var existingUser model.User
	err := db.Get(&existingUser, "SELECT * FROM usuario WHERE email = $1", req.Email)
	if err == nil {
		return c.JSON(http.StatusConflict, model.Message{
			Message: "Email already registered",
		})
	} else if err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "Error verifying Email.",
		})

	}

	// HASH da senha
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "error processing password",
		})
	}

	// insere o novo usuario no banco de dados com a senha hash
	_, err = db.DB.Exec(
		"INSERT INTO usuario (nome, email, password, created_at) VALUES ($1, $2, $3, $4)", req.Name, req.Email, string(hashPassword), req.Created_at,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "Error to Register user",
		})
	}

	return c.JSON(http.StatusCreated, model.Message{
		Message: "successfully registered user",
	})

}
