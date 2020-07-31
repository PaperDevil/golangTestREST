.PHONY: build

DB_LOGIN_URL="postgres://postgres:admin@localhost/testrest?sslmode=disable"

runserver:
	make build
	./apiserver.exe
build:
	go build ./cmd/apiserver
test:
	go test -v -race -timeout 30s ./...
makemigrations:
	migrate create -ext sql -dir migrations create_users
migrate:
	migrate -path ./migrations -database $(DB_LOGIN_URL) up
drop:
	migrate -path ./migrations -database $(DB_LOGIN_URL) down
.DEFAULT_GOAL := build