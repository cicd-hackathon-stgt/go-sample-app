FROM golang:1.10-alpine AS builder

ADD . /app

WORKDIR /app

RUN if go test -v; then echo "Test execution successful"; else exit 1; fi \
 &&  CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app .

FROM scratch

COPY --from=builder /app/app . 
CMD ["./app"] 