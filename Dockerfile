FROM golang:1.9

RUN mkdir /go/src/sample
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/wakashiyo/compose-sample/users

COPY main.go /go/src/sample

CMD ["go", "run", "/go/src/sample/main.go"]
