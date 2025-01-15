package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tiagossdj/jwt-project-/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// Teste de Middleware com requisição sem Token (sem o header para autorização!)
func TestJwtMiddleware_MissingToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/protected", nil) //requisição sem o Token!
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Cria um handler simples para passar para o middleware

	handler := JwtMiddleware(func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.Message{
			Message: "Success",
		})
	})

	err := handler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "Missing or malformed jwt")
	}

}

// Teste de Middleware com token mal formado (não segue o formato)!

func TestMiddlware_MalFormedToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "BearerMalFormedToken") // token mal formado.
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := JwtMiddleware(func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.Message{
			Message: "Success",
		})
	})

	err := handler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "Missing or malformed jwt")
	}

}

// Teste de Middleware com Token Inválido
func TestJwtMiddleware_InvalidToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidTokenHere") //token inválido
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := JwtMiddleware(func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.Message{
			Message: "Success",
		})
	})

	err := handler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid Token")

	}
}

// Teste de Middleware com Token Válido!

func TestJwtMiddleware_ValidToken(t *testing.T) {

	// Gerar um token válido (você precisa de uma chave secreta para isso)
	token := "token valido aqui!" // Simule um token válido

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := JwtMiddleware(func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.Message{
			Message: "Success",
		})
	})

	err := handler(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Success")
	}

}
