package integration_tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tiagossdj/jwt-project-/handlers"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Integration(t *testing.T) {
	// Configura o banco de dados de testes
	db := SetupTestDataBase(t)
	defer CleanUpTestDataBase(db)

	// Cria dados necessários para teste
	PopulateTestData(db)

	// Configura o echo e cria o contexto para o teste
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(`{"nome":"Test User", "email":"test@example.com", "password":"123456"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Executa o handler Register
	if assert.NoError(t, handlers.Register(c, db)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "successfully registered user")
	}

	// verifica no banco se o usuário foi Registrado
	var count int
	err := db.QueryRow("select count(*) from usuario where email= $1", "test@example.com").Scan(&count)
	if err != nil {
		t.Fatalf("Error to Verify test Database:%v", err)
	}
	assert.Equal(t, 1, count)
}
