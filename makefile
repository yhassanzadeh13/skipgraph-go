tidy:
	go mod tidy
lint:
	golangci-lint run ./...
lint-fix:
	golangci-lint run --fix ./...
test:
	GO111MODULE=on go test ./...
build:
	go build -v ./...
	