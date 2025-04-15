FROM golang:1.20.3-alpine AS builder

COPY . /github.com/armanbektassov/go_chat/grpc/source/
WORKDIR /github.com/armanbektassov/go_chat/grpc/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/armanbektassov/go_chat/grpc/source/bin/chat_server .

CMD ["./chat_server"]