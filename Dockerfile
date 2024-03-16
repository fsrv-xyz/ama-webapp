FROM golang:alpine@sha256:6f179eca0d49ec57ed6d64067d3d2c8c77fb4ca134b687f31cf1666e467cd1a9 AS builder
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
