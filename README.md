## 📌 jwt-project - Microserviço de Autenticação com JWT 📌

Este é um projeto de **Microserviço de Autenticação** utilizando **JWT** (JSON Web Token) para autenticação de usuários em uma API.

### 🔧 **Rotas da Aplicação**

- **POST /register**  
  Realiza o cadastro de um novo usuário, gerando um token JWT após a criação do perfil.

- **POST /login**  
  Realiza o login do usuário, retornando um token JWT caso as credenciais sejam válidas.

- **GET /profile**  
  Retorna as informações do perfil do usuário autenticado, exigindo um token JWT válido no cabeçalho de autorização.

### 💾 **Banco de Dados**

O banco de dados utilizado é o PostgreSQL. A tabela de usuários contém os seguintes campos:

- **Email** (único)  
- **ID** (auto-incremento)  
- **Password** (armazenada de forma segura usando bcrypt)

### 🔐 **Segurança**

- Utiliza **JWT (JSON Web Token)** para autenticação e **bcrypt** para criptografar senhas.
- É recomendado o uso de **chaves secretas fortes** para a geração e verificação de tokens JWT, garantindo a segurança da aplicação.
- **Tokens expiram** após um tempo determinado, garantindo que sessões não sejam mantidas indefinidamente.

### Tecnologias Utilizadas

* [Golang](https://github.com/golang/go)
* [Docker](https://www.docker.com/)
* [postgreSQL](https://www.postgresql.org/)

## Dependências e Versões Necessárias

* **Go** - Versão: 1.22.6 
* **Docker** - Versão: 27.3.1 
* **PostgreSQL** - Versão: 12.21 
* **Make** - Versão: 4.4.1 



[![](https://mermaid.ink/img/pako:eNp1kt9u2jAUxl_F8jVUCQHa5qJS-dOVapMQpas0wsVZfACrwQfZTreO8DB7lr7YHMdBq7Sdq5zj7_vp84mPPCeBPOWbgn7kO9CWLSeZYq5uV0-mfP-tJa1Zt3tTTdWrBDYBQYYJZAvcSmM1VWy0up3PQo963bhH3jPWztJiKjZejUDlVNs9J2jHXrtAS1oBeyxzNMapJ4FrDqQMpmcOq6mCgjmE_VfEz7SVqmJTz_FN8Ey9-isUUoCjoUCVS5CmYnde-_C8DMq7D9GW9IIO-Om8GZcux-_YHPw3ULDde_Zc00YWGMT3fycJulmrs5hbFGxBpW31sw955qgdq2IPK78lY6G-djNlgs4bW_MO36PegxTuVx9rVMbtDveY8dR9CtxAWdiMZ-rkpFBaenxTOU-tLrHDNZXbXduUBwEWJxK2Gvbt8ADqG5FrN1AY16OQLuGX5mn5F-Y1PD3ynzzt9qLh8CIe9JM4cTWIk0GHv7l5Eg0vorqSXr8f966vL08d_suD46g5iaJe0o-uksHw9AcY7ODe?type=png)](https://mermaid.live/edit#pako:eNp1kt9u2jAUxl_F8jVUCQHa5qJS-dOVapMQpas0wsVZfACrwQfZTreO8DB7lr7YHMdBq7Sdq5zj7_vp84mPPCeBPOWbgn7kO9CWLSeZYq5uV0-mfP-tJa1Zt3tTTdWrBDYBQYYJZAvcSmM1VWy0up3PQo963bhH3jPWztJiKjZejUDlVNs9J2jHXrtAS1oBeyxzNMapJ4FrDqQMpmcOq6mCgjmE_VfEz7SVqmJTz_FN8Ey9-isUUoCjoUCVS5CmYnde-_C8DMq7D9GW9IIO-Om8GZcux-_YHPw3ULDde_Zc00YWGMT3fycJulmrs5hbFGxBpW31sw955qgdq2IPK78lY6G-djNlgs4bW_MO36PegxTuVx9rVMbtDveY8dR9CtxAWdiMZ-rkpFBaenxTOU-tLrHDNZXbXduUBwEWJxK2Gvbt8ADqG5FrN1AY16OQLuGX5mn5F-Y1PD3ynzzt9qLh8CIe9JM4cTWIk0GHv7l5Eg0vorqSXr8f966vL08d_suD46g5iaJe0o-uksHw9AcY7ODe)



## Como rodar o projeto ✅

Para rodar a aplicação, siga os passos abaixo:

1. **Clonar o repositório**
   
 Primeiro, clone o repositório para sua máquina local:
   
 ```
  git clone https://github.com/Tiagossdj/jwt-project-.git
 ```

Em seguida, entre no diretório do projeto:
 ```
  cd jwt-project- 
 ```


  
2. **Instalar as dependências**

  Certifique-se de que você tem o Go e o Docker instalados. Se necessário, instale as dependências executando:
  
  ```
  go mod tidy
  ```
3. **Construir o binário da aplicação (opcional)**

Se você preferir construir um binário em vez de rodar o código diretamente com o go run, você pode usar o comando abaixo para compilar o código:

```
make build
```

4. **Subir o banco de dados (se estiver usando Docker)**

  Se você estiver usando Docker para o PostgreSQL, rode o seguinte comando para subir o container do banco:
   ```
  docker-compose up -d
   ```
  Isso iniciará o banco de dados em segundo plano.

5. **Rodar a aplicação**

  Para rodar a aplicação, use o seguinte comando:
   ```
  go run main.go
   ```


  A aplicação deve começar a rodar. Você pode confirmar que está tudo funcionando corretamente verificando a seguinte mensagem no terminal:
  ```
   http server started on [::]:8888
```
## Documentação da API   

A documentação da API está disponível através do Swagger, que pode ser acessada em:

- **URL da documentação:**

```
http://localhost:8888/swagger/index.html
```

### Como Usar

- **Inicie a aplicação**
para que a documentação do Swagger esteja acessível:  

```
go run main.go
```

## Utilizando o Postman Para API

Você pode baixar o postman ou utiliza-lo online para acessar as funcionalidades da api neste link:

- [postman.com](https://www.postman.com)

#### **POST auth/register**

- **Descrição:** Registra um novo usuário.
  
- **Corpo da Requisição:**
- No Postman, selecione o método POST e insira a URL da rota:
 ```
 http://localhost:8888/auth/register
 ```
    
- Vá até a aba "Body" e selecione a opção "raw" e escolha o formato JSON.
- Cole o seguinte corpo da requisição:
    
```
    {
      "nome": "Nome do Usuário",
      "email": "email@example.com",
      "password": "senha_secreta"
    }
```
- Clique em **Send** para verificar a resposta.
  

#### **POST auth/login**

- **Descrição:** Realiza o login e retorna um token JWT.
  
- **Corpo da Requisição:**
- Selecione o método POST e insira a URL da rota:
    
 ```
 http://localhost:8888/auth/login
 ```
    
- Vá até a aba "Body" e selecione a opção "raw" e escolha o formato JSON.
- Cole o seguinte corpo da requisição:
    
```
    {
      "email": "email@example.com",
      "password": "senha_secreta"
    }
```
- Clique em **Send** para verificar a resposta.

#### **GET protected/profile**

- **Descrição:** Obtém o perfil do usuário com o token JWT.
  
- **Autorização:**
- No Postman, após realizar o login, você irá obter um token JWT.
- Em seguida, crie uma nova requisição do tipo GET e insira a URL da rota
    
 ```
 http://localhost:8888/protected/profile
 ```
    
- Vá até a aba "Headers" e adicione a chave Authorization com o valor Bearer <seu-token-jwt-aqui>, ou direto na aba "Authorization" e selecionar `Bearer Token` e adicione o valor.
  
- Clique em **Send** para verificar a resposta.

### 3. **Confirmando se a API está funcionando**

Ao executar as requisições no Postman, você verá as respostas retornadas pela API, que devem corresponder ao esperado para cada operação. 

Caso tudo esteja configurado corretamente, você verá as respostas com os respectivos status `(200 OK, 401 Unauthorized, etc.).`


## Como rodar os testes

Para rodar os testes, você pode usar o Makefile para facilitar o processo.

### Testes Unitários

Para rodar os testes unitários, execute o seguinte comando:
 ```
  make test_unit
 ```

Isso executará todos os testes unitários na pasta `handlers` ou onde você configurou seus testes unitários.

### Testes de Integração

Para rodar os testes de integração, execute o seguinte comando:
 ```
  make test_integration
 ```


Isso executará todos os testes de integração que você escreveu para testar o fluxo da aplicação.

### Testes E2E

Para rodar os testes **E2E (End-to-End)**, execute:
 ```
  make test_e2e
 ```

Isso testará o fluxo completo da aplicação, simulando a interação do usuário com a API.

Após rodar os testes, você verá no terminal se todos os testes passaram ou se algum falhou. Se tudo estiver correto, você verá a mensagem:
 ```
  PASS
 ```
## ⏭️ Próximos Passos


- **Função de Logout na API**  
  
  A adição de um endpoint para permitir que o usuário se deslogue, invalidando o token JWT e garantindo que o usuário não possa mais acessar as rotas protegidas sem se autenticar novamente.

- **Função de Renovação de Token (Refresh Token)**
  
  Adicionar suporte a refresh tokens para permitir que os usuários renovem seus tokens JWT sem precisar se autenticar novamente. O sistema de refresh token pode ser usado para aumentar a segurança e melhorar a experiência do usuário.

- **Autenticação via OAuth**
  
  Implementar a autenticação social utilizando OAuth para permitir login através de provedores como Google, Facebook, GitHub, etc.
