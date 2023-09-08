# 🆒 API de Bandeira Tarifária desenvolvida com Go

API de bandeira tarifária de energia elátrica, exibe os dados históricos das bandeiras como mês, ano, bandeira e multiplicador.

A aplicação realiza um webscraping dos dados, armazena na base de dados e exibe no formato JSON.

### 📋 Requisitos

- Go

- PostgreSQL

- Docker

### 🔧 Instalação

Executar o comando abaixo na raiz do projeto para criar o container do banco de dados:

```
docker compose up
```

Criar o banco de dados "flags".

### ⚙️ Execução

Executar o comando abaixo para iniciar a aplicação:

```
go run . 
```

### ▶️ Requisição

GET

```
http://localhost:8080/
```
