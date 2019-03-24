FROM golang:latest

# RUN mkdir /go/src/sample
# RUN go get -u github.com/gin-gonic/gin
# RUN go get -u github.com/wakashiyo/compose-sample/users
# RUN go get -u github.com/go-sql-driver/mysql
# RUN go get -u github.com/go-xorm/xorm

# COPY main.go /go/src/sample

# CMD ["go", "run", "/go/src/sample/main.go"]

#################################################################################################################

## 存在しない場合は、ディレクトリ作成される
WORKDIR /go/src/github.com/wakashiyo/compose-sample

# WORKDIRからの相対パスの位置に配置する
# mysql, nginx, users, main.go...などが全て　/go/src/github.com/wakashiyo/compose-sample 以下にコピーされる
COPY . .

# 使用している外部ライブラリは書かないと動かすことができなかった
# 以前の場合だと、別パッケージとしてgo getしていたcompose-sample/userが全てコピーしていることで存在しているので、書く必要がなくなった
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/go-xorm/xorm

# go install で /go/bin 配下にバイナリがビルドされる
RUN go install github.com/wakashiyo/compose-sample

# image自体に最初からPATHに/go/binのパスが設定されているため、バイナリファイルを書くだけで動かすことができる
ENTRYPOINT [ "compose-sample" ]