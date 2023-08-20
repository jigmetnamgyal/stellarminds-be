RELEASE_TAG = latest

start-dev:
	go run main.go

migrate-db:
	go run migrate/migrate.go

postgres-instance:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=jtn -e POSTGRES_PASSWORD=jtn11 -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql -U jtn

build-go-image:
	docker build -t stellarminds:$(RELEASE_TAG) .

server-docker:
	docker run --name stellarminds-be -p 3000:3000 stellarminds:$(RELEASE_TAG)