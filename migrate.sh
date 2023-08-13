#!/bin/sh
set -e

/go/bin/migrate -path=/app/migrations/ -database "postgres://username:password@postgres:5432/dbname?sslmode=disable" up

exec "$@"




