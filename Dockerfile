FROM golang:1.16.7-alpine

WORKDIR /go/amechan

COPY src ./src
COPY go.mod .

RUN apk add --update && apk add --no-cache git && go mod tidy

WORKDIR /go/amechan/src

RUN go build -o app

CMD ["./app"]