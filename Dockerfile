FROM alpine:latest

WORKDIR /app

COPY . .

RUN apk add wget

RUN wget https://launchermeta.mojang.com/mc/game/version_manifest.json -O version_manifest.json

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

RUN apk add --no-cache ca-certificates && update-ca-certificates

RUN apk update

RUN apk add openjdk21-jre

RUN apk add go

# --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community

CMD ["go", "run", "main.go"]