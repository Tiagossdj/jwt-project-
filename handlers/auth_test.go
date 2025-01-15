package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
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
