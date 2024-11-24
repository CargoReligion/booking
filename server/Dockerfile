FROM golang:1.22-alpine
WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o bin/main ./main.go

EXPOSE 8080

CMD ["./bin/main"]