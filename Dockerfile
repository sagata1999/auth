FROM golang:1.21-alpine AS builder

COPY . /github.com/sagata1999/auth/
WORKDIR /github.com/sagata1999/auth/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/sagata1999/auth/bin/auth_server .

CMD ["./auth_server"]