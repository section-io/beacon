FROM golang:1.11-alpine

ENV CGO_ENABLED=0

RUN apk add --no-cache \
    git

RUN go get -v github.com/kisielk/errcheck/

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

RUN gofmt -e -s -d . 2>&1 | tee /gofmt.out && test ! -s /gofmt.out

RUN go tool vet .

RUN go test -v ./...

# BEGIN hack fix for tools that don't support Go Modules yet, e.g. errcheck.
#  errcheck issue: https://github.com/kisielk/errcheck/issues/155
RUN go mod vendor && \
  mkdir -p /tmp/hack/ && \
  ln -s "${PWD}/vendor" /tmp/hack/src
ENV GOPATH="${GOPATH}:/tmp/hack"
# END

RUN errcheck ./...
