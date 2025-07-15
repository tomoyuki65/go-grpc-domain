FROM golang:1.24-alpine3.21

WORKDIR /go/src

COPY ./src .

# go.modがあれば依存関係をインストール
RUN if [ -f ./go.mod ]; then \
  go install; \
fi

# Buf CLIをインストール
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# 開発用のライブラリをインストール
RUN go install github.com/air-verse/air@latest
RUN go install go.uber.org/mock/mockgen@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

EXPOSE 8080