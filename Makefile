build:
	docker-compose build

up:
	docker-compose up -d

test:
	docker-compose exec api go test -v ./...

push:
	docker-compose push