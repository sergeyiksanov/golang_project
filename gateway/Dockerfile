FROM golang AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download
# RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

WORKDIR /app/cmd/
RUN go build -o gateway

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cmd/gateway .

# COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/cmd/gateway .
# COPY --from=builder /app/db/migrations /root/db/migrations

EXPOSE 8000

CMD ["./gateway"]

# CMD ["goose", "-dir", "/root/db/migrations", "postgres", "'postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable'", "up", "&&", "./auth_service"]
