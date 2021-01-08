FROM golang:latest

LABEL maintainer="Dinda <dindasigma@gmail.com>"

WORKDIR /app/services/messaging

COPY ./ /app/services/messaging

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -exclude-dir=.git -exclude-dir=docs --build="go build -v commands/runserver.go" --command=./runserver