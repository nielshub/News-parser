# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS build

ADD . /app
WORKDIR /app
RUN apk add --update --no-cache ca-certificates
# Run Build binary
RUN go build -v -o incrowd ./src/cmd

FROM alpine:3.5
EXPOSE 8080
CMD [ "incrowd" ]
COPY /env /usr/local/bin
COPY --from=build /app/incrowd /usr/local/bin/incrowd
COPY --from=build /etc/ssl/certs ./etc/ssl/certs
RUN chmod +x /usr/local/bin/incrowd