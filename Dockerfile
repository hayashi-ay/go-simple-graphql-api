FROM golang:1.15.5-alpine

# gitはalpineにないのでインストールする
RUN apk update && \
    apk add git 

WORKDIR /app

RUN go get -u github.com/cosmtrek/air

EXPOSE 8080

CMD air -c .air.toml
