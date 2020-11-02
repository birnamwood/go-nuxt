FROM golang:1.15

# コンテナ側作業ディレクトリ
WORKDIR /go/src/github.com/birnamwood/go-nuxt

# プロジェクトのソースをコンテナに追加
COPY ./ /go/src/github.com/birnamwood/go-nuxt

# Package追加
RUN go get github.com/google/wire/cmd/wire
#マイグレーション

# go module非対応のものを入れる
ENV GO111MODULE=off
RUN go get github.com/oxequa/realize 
ENV GO111MODULE=on