FROM alpine:latest

WORKDIR /app

COPY . .

RUN apk add wget

RUN wget https://piston-data.mojang.com/v1/objects/886945bfb2b978778c3a0288fd7fab09d315b25f/server.jar -O server.jar

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

RUN apk add --no-cache ca-certificates && update-ca-certificates

RUN apk update

RUN apk add openjdk21-jre

RUN apk add go

# --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community

CMD ["go", "run", "main.go"]