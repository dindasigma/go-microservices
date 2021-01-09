FROM golang:latest

LABEL maintainer="Dinda <dindasigma@gmail.com>"

WORKDIR /app/services/user

COPY ./ /app/services/user

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

#RUN go get -u github.com/swaggo/swag/cmd/swag

ENTRYPOINT CompileDaemon -exclude-dir=.git -exclude-dir=docs --build="make build-dev" --command=./runserver