FROM golang:latest

WORKDIR /go/src/offline
COPY . /go/src/offline

# Install go dependencies
RUN rm -rf /go/src/offline/vendor && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure -v -update && \
    dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux go build -o qcms server/*.go

EXPOSE 8010

CMD ["./qcms"]