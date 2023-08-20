FROM golang:1.20-alpine as builder

COPY go.mod go.sum /go/src/github.com/aseldar/test_task/
WORKDIR /go/src/github.com/aseldar/test_task/
RUN go mod download
COPY . /go/src/github.com/aseldar/test_task/

# Копируем файл миграции на этапе сборки
COPY /app/db/migration /go/src/github.com/aseldar/test_task/app/db/migration  

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/test_task ./app/main.go

# Create a minimal production image
FROM alpine:latest

COPY --from=builder /go/src/github.com/aseldar/test_task/app/db/migration /app/db/migration

COPY --from=builder /go/src/github.com/aseldar/test_task/build/test_task  /usr/bin/test_task

ENV HTTP_ADDR=:8080

EXPOSE 8080

# Run the binary when the container starts
ENTRYPOINT ["/usr/bin/test_task"]
