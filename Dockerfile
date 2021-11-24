FROM alpine:3.13.7
RUN apk update
RUN apk add apache2
ENTRYPOINT [ "echo" ]
CMD [ "hello world v29" ]
