FROM golang:latest as builder

LABEL org.opencontainers.image.source https://github.com/s1ntaxe770r/gopher-quotes


RUN mkdir /app

WORKDIR /app

COPY go.mod .

RUN go mod download 

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM alpine 

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server 


EXPOSE 8000

ENTRYPOINT ["/app/server"]
