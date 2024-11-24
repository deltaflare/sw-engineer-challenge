FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOARCH=arm64 GOOS=linux go build -o main .

CMD ["./main"]