FROM golang:1.17-alpine3.14
RUN apk add make gcc libc-dev pkgconf

WORKDIR /go/src/github.com/checkr/idempotent-requests
ADD . .
WORKDIR ./test
CMD make test