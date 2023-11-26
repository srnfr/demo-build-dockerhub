FROM golang:1.19-alpine
##FROM alpine:3.18.3
## MaJ des listes de packages apk
RUN apk update
## Ajout de logiciels via apk
RUN apk add apache2
RUN apk add atop
RUN apk add libcap
RUN apk add bash
RUN apk add jq

#RUN mkdir /app
## Changement du dir courant dans l'image
WORKDIR /app
## Ajout des modules
COPY go.mod go.sum ./
RUN go mod download

## Copie des fichiers GO
COPY *.go ./
## Compilation du binaire go
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-world

## (essai de) nettoyage
RUN rm -f hello-word.go

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot
USER nonroot

CMD [ "/hello-world" ]
