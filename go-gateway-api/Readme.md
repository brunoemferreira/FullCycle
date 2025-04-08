


```bash
# Roda a Aplicaćão e mostra no terminal
$ go run cmd/app/main.go

# Compila a aplicaćão e gera um executável para qualquer so
$ go build

# Baixa e instala qualquer biblioteca que foi declarada nos arquivos no go 
$ go mod tidy

# Executa o arquivo docker-compose.yaml
$ docker compose up -d

# 
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest 

# Executar a migration 
$ migrate -database "postgres://postgres:postgres@localhost:5432/gateway?sslmode=disable" -path migrations up


```

