run: build

build:
	@echo " >> building binaries"
	@go build -o update cmd/update/main.go
