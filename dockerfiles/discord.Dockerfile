FROM golang:1.22.3 AS builder

WORKDIR /app

COPY . .

WORKDIR /app

RUN apt-get update && apt-get install -y curl g++-x86-64-linux-gnu libc6-dev-amd64-cross && rm -rf /var/lib/apt/lists/*

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV CC=x86_64-linux-gnu-gcc
ENV CGO_CFLAGS="-O -D__BLST_PORTABLE__"
ENV CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"

RUN go build -o npsm-gogo -ldflags="-w -s" ./cmd/main.go

FROM ubuntu@sha256:5d070ad5f7fe63623cbb99b4fc0fd997f5591303d4b03ccce50f403957d0ddc4

WORKDIR /app

COPY --from=builder /app/npsm-gogo /usr/bin
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /

ENV ZONEINFO=/zoneinfo.zip

ARG DISCORD_TOKEN
ARG DEFAULT_CHANNEL
ARG CHANWOONG_BLOG
ARG YANA_BLOG

ENV DISCORD_TOKEN=${DISCORD_TOKEN}
ENV DEFAULT_CHANNEL=${DEFAULT_CHANNEL}
ENV CHANWOONG_BLOG=${CHANWOONG_BLOG}
ENV YANA_BLOG=${YANA_BLOG}

CMD ["npsm-gogo"]
