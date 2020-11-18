FROM alpine
RUN apk update
ENTRYPOINT [ "echo" ]
CMD [ "hello world v20" ]
