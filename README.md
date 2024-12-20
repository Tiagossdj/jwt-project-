# jwt-project-
Projeto de Microserviço de Autenticação com JWT 

  //ROTAS
POST /register
POST /login
GET /profile

  //BANCO DE DADOS
Email - (unico)
Id  - (auto-incremento)
Password - (bcrypt)

  //SEGURANÇA
 - Usar chaves secretas fortes para gerar tokens JWT.


  //FLUXO
>> Registro com e-mail e senha.
>> Login com as credenciais e recebe um token JWT.
>> Utilizar o token para acessar endpoints protegidos.

