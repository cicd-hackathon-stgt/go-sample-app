FROM golang:1.10-alpine AS builder

RUN apk add --no-cache ruby ruby-dev

ADD . /app

WORKDIR /app

RUN if go test -v; then echo "Test execution succesful"; else exit 1; fi \
 &&  CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app .

FROM scratch

COPY --from=builder /app/app . 
CMD ["./app"] 