FROM alpine:latest

WORKDIR /app

COPY . .

RUN apk add wget

RUN wget https://piston-data.mojang.com/v1/objects/8f3112a1049751cc472ec13e397eade5336ca7ae/server.jar -O server.jar

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

RUN apk add --no-cache ca-certificates && update-ca-certificates

RUN apk update

RUN apk add openjdk21-jre

RUN apk add go

# --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community

CMD ["go", "run", "main.go"]