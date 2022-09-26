# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS build

ADD . /app
WORKDIR /app
# Run Build binary
RUN go build -v -o incrowd ./src/cmd

FROM alpine:3.4
EXPOSE 8080
CMD [ "incrowd" ]
COPY /env /usr/local/bin
COPY --from=build /app/incrowd /usr/local/bin/incrowd
RUN chmod +x /usr/local/bin/incrowd