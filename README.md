# GoApiPersonalidades
Projeto onde crio uma API para fazer um crud de personalidades


## ✔️ Tecnologias utilizadas

- GO
- Docker
- Postgres
- gorilla/mux
- gorm


## 📚 Detalhes do projeto

Projeto criado usando o mux como router para facilitar a necessidade de buscar informações na URL e utilizando o gorm como ORM.

para executar o projeto:
- na raiz do projeto execute "docker-compose up -d" para criar o banco de dados.
- na raiz do projeto execute "go run main.go" para rodar a aplicação.

rotas disponiveis:
- "/" : pagina inicial do projeto atualmente com um "hello, world"
- "/api/personalidades" GET : busca todas personalidades no DB
- "/api/personalidades/{Id}" GET : busca a personalidade por Id
- "/api/personalidades" POST : Insere uma nova personalidade no DB (body ex: {"nome":"Test","historia":"test"}
- "/api/personalidades/{Id}" DELETE : deleta uma personalidade correspondente ao ID informado
- "/api/personalidades/{Id}" PUT : edita uma personlidade baseada no ID da rota e nas informaçoes do body (body ex: {"nome":"Test","historia":"test"}

Por hora ainda nao adicionei swagger nem dockerfile, estes serão os proximos passos.
