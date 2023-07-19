FROM golang:1.20.5-alpine
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd/server/main.go

CMD ["./main", "serve", "--http=0.0.0.0:9999"]