FROM golang:1.16 as build

WORKDIR /app

COPY . ./

RUN go build -o app /app/cmd/name


FROM debian:stretch-slim

WORKDIR /app

COPY --from=build /app/app /app

ENTRYPOINT ["./app"]