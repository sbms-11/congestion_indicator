# ベースイメージとしてUbuntu 20.04を使用
FROM ubuntu:20.04

# 環境変数の設定（非対話モードでのapt-get実行）
ENV DEBIAN_FRONTEND=noninteractive

# 必要なパッケージのインストール
RUN apt-get update && \
    apt-get install -y \
    wget \
    gnupg \
    software-properties-common \
    build-essential \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Go言語のインストール
RUN wget https://golang.org/dl/go1.22.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz && \
    rm go1.22.5.linux-amd64.tar.gz

# Goのパスを設定
ENV PATH="/usr/local/go/bin:${PATH}"

# ユーザーの作成
RUN useradd -m -s /bin/bash api && echo "api:api" | chpasswd && adduser api sudo

# 作業ディレクトリの作成と設定
RUN mkdir -p /app
WORKDIR /app

# ソースコードのコピー
COPY . .
RUN chown -R api:api /app

# GinとGORMのインストール
USER api
RUN go mod init github.com/sbms-11/congestion_indicator
RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres
RUN go mod tidy
RUN go build -o main .

# ポートのエクスポート（アプリケーションのデフォルトポート）
EXPOSE 8080

CMD "/app/main"