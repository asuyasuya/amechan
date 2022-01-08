FROM golang:1.16.7-alpine

WORKDIR /go/amechan

COPY src ./src
COPY go.mod .

RUN echo "replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0" >> go.mod
RUN apk add --no-cache git \
  && go get github.com/oxequa/realize && go mod tidy

WORKDIR /go/amechan/src

RUN go build -o app

CMD ["go", "run", "main.go"]

EXPOSE 80