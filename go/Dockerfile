#go version
FROM golang:1.20-alpine3.17

#update & install (vimやbashを入れとくと便利）
RUN apk add --update && apk add git && apk add vim && apk add bash

# 作業ディレクトリ作成
RUN mkdir /app

# 作業ディレクトリ設定
WORKDIR /app

CMD ["go", "run", "main.go"]
