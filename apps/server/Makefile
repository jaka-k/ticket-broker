build:
	@go build -o bin/server

run: build
	@./bin/server -port=$(PORT)

test:
	@go test -v ./...