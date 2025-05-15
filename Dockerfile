FROM golang:1.24.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o excuse-service

EXPOSE 8080

CMD ["./excuse-service"]
