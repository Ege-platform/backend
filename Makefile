build:
	CGO_ENABLED=0 go build -o main ./cmd/server/main.go

run: build
	./main serve --http=0.0.0.0:9999