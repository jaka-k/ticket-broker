build:
	@go build -o bin/server ./cmd

run: build
	@./bin/server -port=$(PORT)

test:
	@go test -v ./...
