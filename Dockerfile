FROM alpine
RUN apk update
RUN apk add nginx
ENTRYPOINT [ "echo" ]
CMD [ "hello world v23" ]
