FROM golang:1.10-alpine

ENV CGO_ENABLED=0

RUN apk add --no-cache \
    git

RUN go get -v \
    github.com/stretchr/testify \
    github.com/kisielk/errcheck

WORKDIR /go/src/github.com/section-io/beacon
COPY *.go ./

RUN gofmt -e -s -d . 2>&1 | tee /gofmt.out && test ! -s /gofmt.out

RUN go tool vet .

RUN errcheck ./...

RUN go test -v ./...
