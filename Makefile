build:
	go build cmd/server/main.go

run: build
	./main serve