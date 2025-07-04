FROM golang:1.23-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deeptech-app .

FROM alpine:3.19

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/deeptech-app .
RUN chmod +x deeptech-app

CMD ["./deeptech-app"]
