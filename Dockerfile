FROM golang:1.18-alpine

WORKDIR /

COPY . .

RUN go mod download

ENV DB_PASS="admocode"

CMD ["mysql", "-u", "root", "-p", "$DB_PASS"]

CMD ["CREATE DATABASE IF NOT EXISTS", "bctec", ";"]

CMD ["QUIT", ";"]

CMD ["cd", "./cmd/api/migrations"]

RUN curl -fsSL \
        https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
        sh

CMD ["goose", "up"]

CMD ["cd", "/"]

RUN go build -o bctec.go ./cmd/api/main.go

EXPOSE 8080

CMD ["./bctec.go"]
