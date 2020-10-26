FROM golang:1.15

# コンテナ側作業ディレクトリ
WORKDIR /go/src/github.com/birnamwood/go-nuxt

# プロジェクトのソースをコンテナに追加
ADD ./ .

# Package追加
RUN go get
# realizeはgo module非対応
ENV GO111MODULE=off
RUN go get github.com/oxequa/realize 
ENV GO111MODULE=on