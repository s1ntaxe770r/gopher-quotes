FROM golang:latest as builder

WORKDIR /app
COPY ./go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o api

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
