FROM alpine
RUN apk update
RUN apk install nginx
ENTRYPOINT [ "echo" ]
CMD [ "hello world v23" ]
