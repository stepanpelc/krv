FROM golang:1.17 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /go/src/krv/
COPY ./src /go/src/krv/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM alpine:latest
RUN apk add --no-cache tzdata
ENV TZ=Europe/Prague

COPY --from=build /go/src/krv/krv /usr/local/bin

ENTRYPOINT ["/usr/local/bin/krv"]