FROM golang:1.23.2-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY src /app/src
RUN go build -o main src/cmd/main.go
ENTRYPOINT ["./main"]
