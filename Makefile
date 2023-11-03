build:
	@go build -o bin/handtalkerapi cmd/api/main.go

run: build
	@./bin/handtalkerapi

seed: 
	@go run cmd/seed/main.go