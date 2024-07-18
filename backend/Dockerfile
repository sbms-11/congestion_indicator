# ベースイメージとしてUbuntu 20.04を使用
FROM ubuntu:20.04

# メタデータの設定
LABEL maintainer="fufufu991122@gmail.com"

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

# 作業ディレクトリを設定
WORKDIR /app
COPY . .

# Goモジュールを初期化
RUN go mod init backend

# GinとGORMのインストール
RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres
RUN go mod tidy
RUN go build -o main .

# PostgreSQLのインストール
RUN apt-get update && \
    apt-get install -y \
    postgresql \
    postgresql-contrib \
    && rm -rf /var/lib/apt/lists/*



# PostgreSQLのデータディレクトリをボリュームに設定
VOLUME /var/lib/postgresql/data

# ポートのエクスポート（PostgreSQLのデフォルトポート5432）
EXPOSE 5432

# PostgreSQLの初期化用スクリプトのコピー
COPY init.sql /docker-entrypoint-initdb.d/

# アプリケーションのソースコードをコピー
COPY . /app

# 作業ディレクトリを設定
WORKDIR /app

# コンテナ起動時にPostgreSQLを自動起動し、アプリケーションを実行
CMD service postgresql start && ./main