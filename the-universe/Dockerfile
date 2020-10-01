FROM golang:1.14

WORKDIR /
COPY . .

RUN go get github.com/buger/goterm

CMD go run src/main.go
