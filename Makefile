include .env
export

DB_URL := $(DATABASE_URL)

.PHONY: db-up db-down migrate-up migrate-down migrate-create migrate-status

db-up:
	docker compose up -d postgres
	docker compose up --wait

db-down:
	docker compose down

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir ./migrations -seq $$name

migrate-status:
	migrate -path ./migrations -database "$(DB_URL)" version
