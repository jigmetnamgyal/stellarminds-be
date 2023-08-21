FROM golang:1.21.0-alpine3.18 AS builder
WORKDIR /src/app
COPY . .
RUN go build -o migrate migrate/migrate.go
RUN go build -o main main.go

# Run Stage
FROM golang:1.21.0-alpine3.18
WORKDIR /src/app
COPY --from=builder /src/app/main .
COPY --from=builder /src/app/migrate .
COPY .env .
COPY start.sh .

EXPOSE 3000
CMD [ "/src/app/main" ]
ENTRYPOINT ["/src/app/start.sh"]
