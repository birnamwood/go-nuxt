FROM golang:1.15

# コンテナ側作業ディレクトリ
WORKDIR /go/src/github.com/birnamwood/go-nuxt

# プロジェクトのソースをコンテナに追加
ADD ./go .

# Package追加
RUN go get
ENV GO111MODULE=off
RUN go get github.com/oxequa/realize 
ENV GO111MODULE=on