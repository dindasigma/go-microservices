FROM golang:latest

LABEL maintainer="Dinda <dindasigma@gmail.com>"

WORKDIR /app/services/ticket

COPY ./ /app/services/ticket

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -exclude-dir=.git -exclude-dir=docs --build="go build -v commands/runserver.go" --command=./runserver