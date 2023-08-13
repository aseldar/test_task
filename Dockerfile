FROM golang:latest

WORKDIR /app

RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY migrate.sh /app/migrate.sh
RUN chmod +x /app/migrate.sh

COPY . .

RUN go build -o main ./app


CMD ["/app/migrate.sh", "./main"]
