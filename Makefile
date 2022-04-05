test:
	go test -v ./...

build:
	go build -o bin/ main.exe

run:
	go run main.go

start:
	go test -v ./...
	go run main.go