FROM golang:1.23.0-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download && go mod verify

RUN go build -o bin


FROM alpine

WORKDIR /app

COPY --from=builder /build/.env.docker ./.env
COPY --from=builder /build/db/migrations ./db/migrations
COPY --from=builder /build/bin ./bin

ENTRYPOINT ["./bin"]
