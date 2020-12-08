FROM golang:1.15.5-alpine3.12

WORKDIR /go/src/app/

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . .

CMD go run main.go