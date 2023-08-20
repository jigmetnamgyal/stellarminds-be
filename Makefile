start-dev:
	go run main.go

migrate:
	go run migrate/migrate.go

postgres-instance:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=jtn -e POSTGRES_PASSWORD=jtn11 -d postgres:15-alpine