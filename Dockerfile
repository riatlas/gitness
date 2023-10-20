# docker build --rm -f docker/Dockerfile -t drone/drone .
FROM golang:1.16-alpine as builder
WORKDIR /app
ADD . /app
RUN apk add --update-cache gcc musl-dev
RUN go build -tags "nolimit" ./cmd/drone-server/

FROM alpine:3.11 as alpine
RUN apk add -U --no-cache ca-certificates tzdata

FROM alpine:3.11
EXPOSE 80 443
VOLUME /data

RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

ENV GODEBUG netdns=go
ENV XDG_CACHE_HOME /data
ENV DRONE_DATABASE_DRIVER sqlite3
ENV DRONE_DATABASE_DATASOURCE /data/database.sqlite
ENV DRONE_RUNNER_OS=linux
ENV DRONE_RUNNER_ARCH=arm64
ENV DRONE_SERVER_PORT=:80
ENV DRONE_SERVER_HOST=localhost
ENV DRONE_DATADOG_ENABLED=true
ENV DRONE_DATADOG_ENDPOINT=https://stats.drone.ci/api/v1/series

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=alpine /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/drone-server /bin/drone-server
ENTRYPOINT ["/bin/drone-server"]