FROM golang:1.11

WORKDIR /app

ADD . /app

EXPOSE 8000

CMD go run main.go