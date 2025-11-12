FROM golang:1.23-alpine AS builder


WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p /out && go build -v -x -o /out/server ./cmd

FROM alpine:3.19

WORKDIR /src

RUN apk --no-cache upgrade

COPY --from=builder /out/server .

EXPOSE 8080

ENTRYPOINT ["/src/server"]


