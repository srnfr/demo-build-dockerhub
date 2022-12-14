FROM alpine:3.14.8
RUN apk update
RUN apk add apache2
RUN apk add atop
RUN apk add libcap
RUN apk add bash
RUN apk add jq
ENTRYPOINT [ "echo" ]
CMD [ "hello world v50" ]
