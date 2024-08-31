FROM golang:1.22.3

WORKDIR /app

COPY . .
COPY .env /app/.env

RUN go mod tidy
RUN go build -o /main-service ./main.go

EXPOSE 8080

CMD ["/main-service"]