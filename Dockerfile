# source: https://docs.docker.com/guides/golang/build-images/

FROM golang:1.26 AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/les-discord-webhook


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/les-discord-webhook .

CMD ["./les-discord-webhook"]