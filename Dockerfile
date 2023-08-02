FROM golang:alpine@sha256:7efb78dac256c450d194e556e96f80936528335033a26d703ec8146cec8c2090 AS builder
WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -trimpath -o /build/app
RUN apk add -U --no-cache ca-certificates

FROM scratch
EXPOSE 8080
EXPOSE 8081
COPY --from=builder /build/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app"]
