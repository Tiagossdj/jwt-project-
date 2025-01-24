//go:build e2e

package integration_tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tiagossdj/jwt-project-/handlers"
	"github.com/labstack/echo/v4"
)

func TestGetProfile_E2E(t *testing.T) {
	// Banco de dados Teste
	db := SetupTestDataBase(t)
	defer CleanUpTestDataBase(db)

	// Configura Servidor echo e Contexto para teste
	e := echo.New()

	// Teste 1: Registro e Login com token válido!

	t.Run("GetProfile with valid token", func(t *testing.T) {
		// Registro de Usuário

		req := map[string]string{
			"nome":     "Test User",
			"email":    "TestUser@example.com",
			"password": "password123",
		}
		reqBody, err := json.Marshal(req)
		if err != nil {
			log.Fatalf("Error to Marshal map's information: %v", err)
		}
		reqReader := bytes.NewReader(reqBody)

		// Realiza o Registro
		registerReq := httptest.NewRequest(http.MethodPost, "/auth/register", reqReader)
		registerReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(registerReq, rec)
		if err := handlers.Register(c, db); err != nil {
			t.Fatalf("Error during registration: %v", err)
		}

		// Realiza Login para obter Token
		loginReq := map[string]string{
			"email":    "TestUser@example.com",
			"password": "password123",
		}
		loginReqBody, err := json.Marshal(loginReq)
		if err != nil {
			log.Fatalf("Error to Marshal login request: %v", err)
		}
		loginReqReader := bytes.NewReader(loginReqBody)

		loginRequest := httptest.NewRequest(http.MethodPost, "/auth/login", loginReqReader)
		loginRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recLogin := httptest.NewRecorder()
		cLogin := e.NewContext(loginRequest, recLogin)

		if err := handlers.Login(cLogin, db); err != nil {
			t.Fatalf("Error to Login: %v", err)
		}

		// Pegar o Token do Login
		var response map[string]interface{}
		if err := json.Unmarshal(recLogin.Body.Bytes(), &response); err != nil {
			t.Fatalf("Error to parsing login response: %v", err)
			token, ok := response["Message"].(string)
			if !ok {
				t.Fatalf("Token not found in login response")
			}

			// Teste 2: Obter Perfil com o Token!

			getProfileReq := httptest.NewRequest(http.MethodGet, "/protected/profile", nil)
			getProfileReq.Header.Set("Authorization", "Bearer "+token)
			recProfile := httptest.NewRecorder()
			cProfile := e.NewContext(getProfileReq, recProfile)

			if err := handlers.GetProfile(cProfile); err != nil {
				t.Fatalf("Error during profile fetch: %v", err)
			}

			// Verificar se a resposta foi ok

			if status := recProfile.Code; status != http.StatusOK {
				t.Errorf("Expected status %v,  but got %v", http.StatusOK, status)
			}

			if !strings.Contains(recProfile.Body.String(), "Bem-Vindo Test User!") {
				t.Errorf("Expected welcome message with user name, got: %v", recProfile.Body.String())
			}
		}

	})

}
