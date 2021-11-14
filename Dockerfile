FROM golang:1.17.3-alpine AS build

WORKDIR /app

ENV CGO_ENABLED=0

COPY go.* .
RUN go mod download

COPY . .

RUN go build -o bin/goweb .

FROM alpine:3.12
EXPOSE 8080
COPY --from=build app/bin/goweb .
CMD ["/goweb"]