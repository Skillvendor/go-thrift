
# A hello world example with Go
FROM golang:1.8

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./main.go

CMD ["./main"]
