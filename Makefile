.PHONY = help clean test lint

help:
	@echo make help: show help
	@echo make clean: remove unused files
	@echo make test: run unit tests
	@echo make lint: check lint

clean:
	rm -f coverage.out coverage.html

test:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...
