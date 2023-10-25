FROM golang:1.21-alpine AS builder

RUN #apk add --update make git curl
RUN apk add --update make

ENV CGO_ENABLED 0

WORKDIR /build

ADD go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build-api

# API Server
FROM alpine:latest as server

RUN apk --no-cache add ca-certificates

WORKDIR /opt/project

COPY --from=builder /build/bin/api .
COPY --from=builder /build/.env .

RUN chown root:root api

EXPOSE 8080

CMD ["./api"]
