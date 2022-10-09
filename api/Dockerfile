FROM golang:1.16-alpine
# 設定値関連の定義
ARG WORK_DIR=/go/src/api/
ARG MODULE_NAME=api
RUN mkdir -p $WORK_DIR
WORKDIR $WORK_DIR
# RUN echo `go env`
# ARG GOMOD=$PWD/go.mod
COPY go.mod ./
RUN go mod download

# 環境構築
# RUN go mod init $MODULE_NAME

COPY . $WORK_DIR

# Build a db migration tool sql-migration
# ARG SQL_MIGRATE_SRC_PATH=`find $GOPATH/pkg | grep sql-migrate | grep main.go | sed -e "s/main.go//g"`
# WORKDIR $SQL_MIGRATE_SRC_PATH
# RUN go build -o $GOPATH/bin .
# WORKDIR $WORK_DIR

# Install a migration tool Golang-migration and Build it
# ARG MIGRATE_DIR=/usr/bin/migrate
# RUN apk update \
#   && apk add --no-cache git alpine-sdk gnupg vim
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz \
#   && mv migrate.linux-amd64 $MIGRATE_DIR
# ENV PATH $PATH:$MIGRATE_DIR

# Porting
EXPOSE 8080
