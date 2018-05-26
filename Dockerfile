FROM golang:1.8

WORKDIR /go/src/app
COPY src/github.com/thommil/animals-go-ws .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]