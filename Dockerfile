FROM alpine:3.13.7
RUN apk update
RUN apk add apache2
RUN apk add atop
ENTRYPOINT [ "echo" ]
CMD [ "hello world v32" ]
