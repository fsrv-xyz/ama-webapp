FROM golang:alpine@sha256:df479aa7d1298917c7de7fd56c7231ec450149f1e63d960ea96aebf9ff9240c5 AS builder
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
