# Build stage
FROM golang:1.22.4 as builder
WORKDIR /app
COPY . .
RUN go mod init backend
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/lib/pq
RUN go build -o main .

# Run stage
FROM ubuntu:20.04
# タイムゾーンを設定しないとdocker-compose時に停止してしまう
ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN mkdir -p /var/lib/postgresql/data
RUN chown -R postgres:postgres /var/lib/postgresql
WORKDIR /var/lib/postgresql
COPY --from=builder /app/main .
RUN apt-get update && apt-get install -y postgresql postgresql-contrib
USER postgres

# Setup the database
RUN /etc/init.d/postgresql start && \
    psql --command "CREATE USER dbuser WITH SUPERUSER PASSWORD 'P#ssw0rd';" && \
    createdb -O dbuser db

USER root
EXPOSE 8000
CMD service postgresql start && ./main