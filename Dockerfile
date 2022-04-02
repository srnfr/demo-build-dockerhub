FROM alpine:3.13.9
RUN apk update
RUN apk add apache2
RUN apk add atop
RUN apk add libcap
RUN apk add bash
ENTRYPOINT [ "echo" ]
CMD [ "hello world v40" ]
