# Nome do binário gerado
BINARY_NAME=myservice

# Diretórios de testes
UNIT_TEST_DIR=./handlers
INTEGRATION_TEST_DIR=./integration_tests

# Comandos gerais
build: ## Compila o binário principal
	go build -o $(BINARY_NAME) main.go

run: build ## Compila e executa o projeto
	./$(BINARY_NAME)

clean: ## Remove o binário gerado
	rm -f $(BINARY_NAME)

# Testes
test-unit: ## Roda os testes unitários
	go test -v -short $(UNIT_TEST_DIR)

test-integration: ## Roda os testes de integração
	go test -v -tags=integration $(INTEGRATION_TEST_DIR)

test-e2e: ## Roda os testes E2E
	go test -v $(INTEGRATION_TEST_DIR)

test-all: ## Roda todos os testes (unitários, integração e E2E)
	go test -v ./...

# Documentação
help: ## Mostra todos os comandos disponíveis
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
