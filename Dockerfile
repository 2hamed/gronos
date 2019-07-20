FROM golang:1.12 AS build

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN go test ./scheduler

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goronos .

FROM alpine

COPY --from=build /app/goronos /usr/bin

COPY ./commands/one.yaml /root/

CMD ["goronos", "/root"]