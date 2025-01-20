package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// teste de Login com credenciais válidas!

func TestLogin_Success(t *testing.T) {
	e := echo.New()

	// simula requisição http com email e a senha válida!
	req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(`{"email":"test@user.com", "password":"123456"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// simular mock do BD
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db:%v", err)
	}

	db := sqlx.NewDb(mockDB, "postgres")

	// simular a query esperada!
	mock.ExpectQuery("SELECT \\* FROM usuario WHERE email = \\$1").
		WithArgs("test@user.com").
		WillReturnRows(sqlmock.NewRows([]string{"email", "password"}).
			AddRow("test@user.com", "$2y$10$CKI3SWwBvoe6R44E39kgl.fP9MqaozCraXvd3d.MWLRMP5EAns8TW"))

	//Executar o handler de Login
	if assert.NoError(t, Login(c, db)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "here's your token")

	}
	assert.NoError(t, mock.ExpectationsWereMet())
}

// teste com credenciais Inválidas!

func TestLogin_InvalidCredencials(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(`{"email":"invalid@user.com", "password":"wrongpass"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db:%v", err)
	}

	db := sqlx.NewDb(mockDB, "postgres")

	mock.ExpectQuery("SELECT \\* FROM usuario WHERE email = \\$1").
		WithArgs("invalid@user.com").
		WillReturnError(sql.ErrNoRows)

	assert.NoError(t, Login(c, db))
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "User not Found!!!")
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Teste com dados Imcompletos!

func TestLogin_MissingData(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(`{"email":"", "password":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// mesmo que não use diretamente, é necessário criar o mock pra não dar panic
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db:%v", err)
	}

	db := sqlx.NewDb(mockDB, "postgres")

	if assert.NoError(t, Login(c, db)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "invalid credentials!")
	}
}

// Teste de Register com Sucesso!

func TestRegister_Success(t *testing.T) {
	e := echo.New()

	// simula requisição http com email e a senha válida!
	req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(`{"nome":"test user", "email":"test@user.com", "password":"123456"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// simula mock do Banco de Dados
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db:%v", err)
	}
	db := sqlx.NewDb(mockDB, "postgres")

	//expectativa para o SELECT do email
	mock.ExpectQuery("SELECT \\* FROM usuario WHERE email = \\$1").
		WithArgs("test@user.com").
		WillReturnError(sql.ErrNoRows) //simula que o email não está registrado no banco.

	//Caso o modelo de usuario seja registrado corretamente!
	mock.ExpectExec("INSERT INTO usuario"). // o comando de inserção de dados
						WithArgs("test user", "test@user.com", sqlmock.AnyArg()). // user, email e hash password(AnyArg() pois o hash muda a cada chamada da função!)
						WillReturnResult(sqlmock.NewResult(1, 1))                 // retorna sucesso na inserção.

	// executa o handler de register
	if assert.NoError(t, Register(c, db)) { //Chama a função Register passando o contexto e a DB simulada
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "successfully registered user")
	}

	assert.NoError(t, mock.ExpectationsWereMet())

}

// Teste de Register com Email Existente!

func TestRegister_EmailExists(t *testing.T) {
	e := echo.New()

	// simula requisição http com email e a senha válida!
	req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(`{"nome":"test user", "email":"test@user.com", "password":"123456"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// simula mock do Banco de Dados
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db:%v", err)
	}
	db := sqlx.NewDb(mockDB, "postgres")

	//expectativa para o SELECT do email
	mock.ExpectQuery("SELECT \\* FROM usuario WHERE email = \\$1").
		WithArgs("test@user.com").
		WillReturnRows(sqlmock.NewRows([]string{"email"}).
			AddRow("test@user.com"))

	// Executar handler de Register
	if assert.NoError(t, Register(c, db)) { //Chama a função Register passando o contexto e a DB simulada
		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.Contains(t, rec.Body.String(), "Email already registered")
	}

	assert.NoError(t, mock.ExpectationsWereMet()) // Verifica se todas as expectativas de mock foram atendidas
}

// Função para gerar token válidos para teste!

func GenerateValidToken() string {
	//claims do token
	claims := jwt.MapClaims{
		"name": "Test User",                            // O nome que você quiser!
		"exp":  time.Now().Add(time.Minute * 1).Unix(), //tempo de expiração do token.
	}

	// Criando o token com a chave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// A chave secreta que você usaria para assinar o token (o mesmo valor usado na aplicação)!
	secretKey := []byte("yoursecretkey!")

	//gerando o token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

// Teste de GetProfile com Token Válido!

func TestGetProfile_Success(t *testing.T) {
	e := echo.New()

	// Criando um tokenJWT de Exemplo (simulação)
	token := GenerateValidToken()

	req := httptest.NewRequest(http.MethodGet, "/protected/profile", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//Simulando a inserção do usuario no contexto

	claims := jwt.MapClaims{
		"name": "test user",
	}
	c.Set("user", claims)

	// Executando o handler!
	if assert.NoError(t, GetProfile(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Token validated successfully! Welcome test user!")
	}
}

// Teste de GetProfile com Token Inválido ou ausente!

func TestGetProfile_InvalidToken(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/protected/profile", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer invalid_token")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Executando o handler!
	if assert.NoError(t, GetProfile(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid Token claims")
	}
}

// Teste de GetProfile com Claims Inválidas!

func TestGetProfile_InvalidTokenClaimsType(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/protected/profile", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer valid_token_here")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//simula claims com o tipo errado (não é  Jwt_MapClaims)!!!
	c.Set("user", "invalid_claims_type") //tipo inválido no contexto!

	// Executando o handler!
	if assert.NoError(t, GetProfile(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid Token claims")
	}
}
