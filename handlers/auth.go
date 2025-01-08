package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Tiagossdj/jwt-project-/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// chave secreta para assinar o token jwt
var JwtSecret = []byte("yoursecretkey!")

// Login é o handler para o endPoint /auth/login
func Login(c echo.Context, db *sqlx.DB) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Message{
			Message: "Invalid data!",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusUnauthorized, model.Message{
			Message: "invalid credentials!",
		})
	}

	// verificar existencia de dados no banco
	var user model.User
	err := db.Get(&user, "SELECT * FROM usuario WHERE email = $1", req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, model.Message{
				Message: "User not Found!!!",
			})
		}
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "Error while verifying user!",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Message{
			Message: "Invalid Credencials",
		})
	}

	// geração de token JWT
	claims := jwt.MapClaims{
		"name": user.Name,                             // o nome do usuario logado!
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // A data de expiração
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "Error to generate Token!",
		})
	}

	// return token gerado
	return c.JSON(http.StatusOK, model.Message{
		Message: fmt.Sprintf("Hi %s, here's your token: %s", user.Name, signedToken),
	})

}

// Register é o handler para o endPoint /auth/register
func Register(c echo.Context, db *sqlx.DB) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		// log para saber que Erro aconteceu.
		log.Printf("Error to bind data: %+v", err)
		return c.JSON(http.StatusBadRequest, model.Message{
			Message: "Invalid Data",
		})
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, model.Message{
			Message: "You need to fill in all the fields!",
		})
	}

	var userWithEmail model.User
	err := db.Get(&userWithEmail, "SELECT * FROM usuario WHERE email = $1", req.Email)
	if err == nil {
		return c.JSON(http.StatusConflict, model.Message{
			Message: "Email already registered",
		})
	} else if err == sql.ErrNoRows {
		log.Printf("No user found with email: %s. Proceeding with registration.", req.Email)
	} else {
		log.Printf("Error verifying email: %v", err) // Log detalhado
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "Error verifying Email.",
		})
	}

	//HASH da senha
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "error processing password",
		})
	}
	// insere o novo usuario no banco de dados com a senha hash
	_, err = db.DB.Exec(
		"INSERT INTO usuario (nome, email, password) VALUES ($1, $2, $3)", req.Name, req.Email, string(hashPassword),
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

// Profile é o handler para o endPoint /protected/profile, para utiliza-lo, é necessario um token!
func GetProfile(c echo.Context) error {
	user := c.Get("user") // Obtendo o usuário do contexto
	if user == nil {
		return c.JSON(http.StatusUnauthorized, model.Message{
			Message: "Invalid Token claims", // Se o usuário não estiver no contexto, retorna erro
		})
	}

	// Verificando se as claims são do tipo jwt.MapClaims
	claims, ok := user.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, model.Message{
			Message: "Invalid Token claims", // Se não for jwt.MapClaims, retorna erro
		})
	}

	name, ok := claims["name"].(string) // "name" é o nome do usuario registrado.
	if !ok {
		name = "Usuário"
	}

	// Retornar a resposta com o email extraído do token
	return c.JSON(http.StatusOK, model.Message{
		Message: fmt.Sprintf("Token validado com sucesso! Bem vindo %s!", name),
	})

}
