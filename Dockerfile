FROM golang:1.23.2-alpine
WORKDIR /app
COPY go.mod src ./
RUN go mod tidy