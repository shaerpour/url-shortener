FROM golang:1.22.2 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/src/app/url-shortener

FROM alpine:latest

WORKDIR /app

COPY --from=builder /usr/src/app/url-shortener /app/url-shortener

EXPOSE 8080

ENTRYPOINT ["/app/url-shortener"]
