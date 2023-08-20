#!/bin/sh
set -e

/go/bin/migrate -path=C:/Users/asildar.magomedov/test_task/app/db/migration/ -database "postgres://PG_USER:PG_PASS@localhost:5432/PG_DATABASE?sslmode=disable" -verbose up

exec "$@"


migrate -path=C:/Users/asildar.magomedov/test_task/app/db/migration/ -database "postgres://PG_USER:PG_PASS@localhost:5432/PG_DATABASE?sslmode=disable" -verbose up


migrate create -ext sql -dir app/db/migration -seq create_users_table


