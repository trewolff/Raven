build:
	go generate ./...
	go build -o main server.go

run: build
	./main

watch:
	air

migrate-up:
	migrate -path ./migrations -database "$(MIGRATION_DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(MIGRATION_DB_URL)" down