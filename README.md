# Scrum poker backend
Desenvolvido no início de 2023, o Scrum Poker é uma REST API que gerencia usuários e sessões provendo uma ferramenta de pontuação para a etapa de planning do processo do SCRUM.

## Funcionalidades
- Auto cadastro de usuários
- Criação de sessões de scrum-poker
- Sistema de envio de pontuações e revelação

## Instalação
Para instalar e rodar esse projeto na sua máquina, a aplicação foi conteinerizada usando o Docker, visando prover um ambiente de desenvolvimento simplificado e homogêneo para qualquer pessoa que seja.

Para acessar o banco de dados, basta acessar o painel do pgAdmin no seu navegador digitando localhost:54321 e fazer login com o usuário e senha definidos no `docker-compose.yml`. Com o painel aberto, basta registrar um servidor informando na aba connection o host, nome do banco, usuário e senha definidos também no `docker-compose.yml` (O host pode ser obtido através do comando `sudo docker inspect <container id> | grep IPAddress`).

### Comandos:
- `$ git clone https://github.com/ropehapi/scrum-poker-backend.git`
- `$ docker-compose up`

## Endpoints
- Players
    - Index
        - Method: GET
        - Endpoint: /players
    - Show
        - Method: GET
        - Endpoint: /player/{id}
    - Store
        - Method: POST
        - Endpoint: /player
        - Body Ex:

            {
                "name":"Tainara Queiroz de Marchi",
                "email":"taiqdm@gmail.com",
                "password":"Tainara0123",
                "role":"GUEST"
            }
    - Update
        - Method: PUT
        - Endpoint: /player/{id}
        - Body Ex:

            {
                "name":"Pedro Monteiro Yoshimura",
                "email":"ropehapi@gmail.com",
                "password":"Tainara0123",
                "role":"GUEST"
            }

    - Delete
        - Method: DELETE
        - Endpoint: /player/{id}
