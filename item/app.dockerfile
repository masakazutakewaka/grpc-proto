FROM golang:1.10.3-alpine3.8 AS builder
WORKDIR /go/src/github.com/masakazutakewaka/grpc-proto/item
COPY vendor ../vendor
COPY item ./
RUN go build -o /go/bin/app main/main.go

FROM alpine:3.8
WORKDIR /usr/bin
COPY --from=builder /go/bin .
EXPOSE 8080
CMD ["app"]
