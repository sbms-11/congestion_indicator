# ベースイメージとしてPostgreSQLの公式イメージを使用
FROM postgres:14

# 環境変数の設定（デフォルトの設定）
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=db

# 初期化スクリプトのコピー（データベース初期化用）
COPY init.sql /docker-entrypoint-initdb.d/

# PostgreSQLのポートを公開
EXPOSE 5432