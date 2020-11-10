FROM golang:1.15-alpine
WORKDIR /var/app
COPY start.server.sh /bin/start.server.sh

RUN apk add --update make

ENTRYPOINT [ "sh","/bin/start.server.sh" ]
