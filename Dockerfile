FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go mod download

CMD ["go", "run", "./cmd/server/main.go"]
