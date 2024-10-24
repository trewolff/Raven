build:
	go generate ./...
	go build -o main server.go

run: build
	./main

watch:
	air