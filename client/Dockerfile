FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o client_service .

CMD ["./client_service"]