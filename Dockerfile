FROM golang:latest

RUN mkdir /go/src/sample
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/wakashiyo/compose-sample/users
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/go-xorm/xorm

COPY main.go /go/src/sample

CMD ["go", "run", "/go/src/sample/main.go"]
