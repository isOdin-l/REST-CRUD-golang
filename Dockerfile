FROM golang:1.26.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /rest-api-server ./cmd/rest-api-server/

FROM alpine:3.23.3
COPY --from=builder /rest-api-server /bin/rest-api-server

CMD ["/bin/rest-api-server"]