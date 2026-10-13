.PHONY: up down test

up:
	docker compose up --build -d

down:
	docker compose down

test:
	go test ./...
