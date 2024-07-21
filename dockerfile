FROM golang

WORKDIR /app

COPY . .

COPY cmd/api/.env ./

RUN go mod download

RUN go build -o main ./cmd/api

CMD ["./main"]