FROM golang:1.16.3-alpine3.13 as builder

WORKDIR /app

COPY . .

RUN go run main.go

RUN apk add wget

RUN wget https://piston-data.mojang.com/v1/objects/450698d1863ab5180c25d7c804ef0fe6369dd1ba/server.jar -O server.jar

FROM openjdk:21-jdk

WORKDIR /app

COPY --from=builder /app .

CMD ["java", "-jar", "server.jar", "nogui"]