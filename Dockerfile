FROM alpine:3.13.6
RUN apk update
RUN apk add nginx
ENTRYPOINT [ "echo" ]
CMD [ "hello world v26" ]
