//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tiagossdj/jwt-project-/handlers"
	"github.com/labstack/echo/v4"
)

func TestLogin_Integration(t *testing.T) {
	db := SetupTestDataBase(t)
	defer CleanUpTestDataBase(db)

	// Cria dados necessários para teste
	CreateUserTest(db)

	// Configura o servidor para os testes
	e := echo.New()

	// teste 1: Login com Credenciais Corretas
	t.Run("Login com credenciais corretas", func(t *testing.T) {
		req := map[string]string{
			"email":    "TestUser@example.com",
			"password": "password123",
		}
		reqBody, err := json.Marshal(req)
		if err != nil {
			log.Fatalf("Error to Marshal map's information:%v", err)
		}
		reqReader := bytes.NewReader(reqBody)

		// Fazer a requisição para o endpoint /auth/login
		request := httptest.NewRequest(http.MethodPost, "/auth/login", reqReader)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(request, rec)

		// Chamar o handler Login

		err = handlers.Login(c, db)
		if err != nil {
			t.Errorf("Error to Login:%v", err)
		}

		// Verificar o status da resposta
		if status := rec.Code; status != http.StatusOK {

			t.Errorf("Esperando status %v, mas recebeu status %v", http.StatusOK, status)
		}

		// Verificar se a resposta contém o token
		var response map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			t.Fatalf("Erro ao deserializar resposta:%v", err)
		}
		if _, ok := response["Message"]; !ok {
			t.Errorf("Esperado campo 'Message' na resposta, mas não foi encontrado")
		}
	})

	//Teste 2: Login com credenciais Inválidas!
	t.Run("Login com credenciais inválidas", func(t *testing.T) {
		req := map[string]string{
			"email":    "InvalidUser@example.com",
			"password": "wrongpassword123",
		}
		reqBody, err := json.Marshal(req)
		if err != nil {
			log.Fatalf("Error to Marshal map's information:%v", err)
		}
		reqReader := bytes.NewReader(reqBody)

		// Fazer a requisição para o endpoint /auth/login
		request := httptest.NewRequest(http.MethodPost, "/auth/login", reqReader)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(request, rec)

		// Chamar o handler Login

		err = handlers.Login(c, db)
		if err != nil {
			t.Errorf("Error to Login:%v", err)
		}

		// Verificar o status da resposta
		if status := rec.Code; status != http.StatusUnauthorized {

			t.Errorf("Esperando status %v, mas recebeu status %v", http.StatusUnauthorized, status)
		}
	})

	// Teste 3: Login com dados faltando!
	t.Run("Login com dados faltando", func(t *testing.T) {
		req := map[string]string{
			"email":    "",
			"password": "",
		}
		reqBody, err := json.Marshal(req)
		if err != nil {
			log.Fatalf("Error to Marshal map's information:%v", err)
		}
		reqReader := bytes.NewReader(reqBody)

		// Fazer a requisição para o endpoint /auth/login
		request := httptest.NewRequest(http.MethodPost, "/auth/login", reqReader)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(request, rec)

		// Chamar o handler Login

		err = handlers.Login(c, db)
		if err != nil {
			t.Errorf("Error to Login:%v", err)
		}

		// Verificar o status da resposta
		if status := rec.Code; status != http.StatusUnauthorized {

			t.Errorf("Esperando status %v, mas recebeu status %v", http.StatusUnauthorized, status)
		}
	})

}
