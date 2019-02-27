FROM golang:1.9

RUN mkdir /go/src/sample
RUN go get -u github.com/gin-gonic/gin

COPY main.go /go/src/sample

CMD ["go", "run", "/go/src/sample/main.go"]
