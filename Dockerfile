FROM golang:1.21.0-alpine3.18 AS builder
WORKDIR /app
COPY . .
ENV DATABASE_URL="host=postgres user=jtn password=jtn11 dbname=jtn port=5432 sslmode=disable"
RUN go run migrate/migrate.go
RUN go build -o main main.go

# Run Stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 3000
CMD [ "/app/main" ]