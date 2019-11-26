FROM golang:1.8

WORKDIR /go/src/kgs
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
ENTRYPOINT go run main.go
EXPOSE 50052