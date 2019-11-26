FROM golang:1.8

WORKDIR /go/src/github.com/go-redis/redis/
RUN git clone https://github.com/go-redis/redis.git .
RUN git checkout v6

WORKDIR /go/src/app
COPY . .

RUN go get -u ./...
RUN go install -v ./...
ENTRYPOINT go run main.go
EXPOSE 8080