FROM golang:1.14

# コンテナ側作業ディレクトリ
WORKDIR /go/src/

# プロジェクトのソースをコンテナに追加
ADD ./go .

# Package追加
RUN go get
