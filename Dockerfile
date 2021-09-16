FROM golang:1.17-alpine3.14 as builder
RUN apk add make gcc libc-dev pkgconf

WORKDIR /go/src/github.com/checkr/idempotent-requests
ADD . .
RUN make build

FROM alpine:3.14
RUN apk update && apk add --no-cache libc6-compat ca-certificates
COPY --from=builder /go/src/github.com/checkr/idempotent-requests/bin/idempotent-requests-server /app/idempotent-requests-server

EXPOSE 8080

CMD /app/idempotent-requests-server
