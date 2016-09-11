from golang:latest

WORKDIR /go/src/github.com/bsedg/tasker
ADD . /go/src/github.com/bsedg/tasker

RUN echo "{'date': '`date`', 'build': '`git rev-parse HEAD`'}" > /version

RUN go get ./...
RUN go test ./...
RUN go install ./cmd/taskservice

CMD ["/go/bin/taskservice"]
