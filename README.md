# wiza.core

## Запуск

```bash
# 1. Установить migrate (один раз)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
export PATH="$PATH:$(go env GOPATH)/bin"

# 2. Скопировать конфиг
cp .env.example .env

# 3. Поднять базу и накатить миграции
make db-up
make migrate-up

# 4. Запустить
go run ./cmd/main.go
```

Сервис на `http://localhost:8080`.

## Эндпоинты

```
GET /health
GET /api/v1/clients/:iin
```
