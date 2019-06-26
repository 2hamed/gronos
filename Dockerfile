FROM golang:1.12 AS build

COPY . /go/src/github.com/2hamed/goronos

RUN go install github.com/2hamed/goronos

FROM alpine

COPY --from=build /go/bin/goronos /usr/bin

COPY ./commands/one.yaml /root/

CMD ["goronos", "/root"]