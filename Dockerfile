FROM golang:1.8.6
ADD ./ /go/src/github.com/fionwan/todoApp
WORKDIR /go/src/github.com/fionwan/todoApp
RUN ["go", "build"]
ENTRYPOINT ["go", "run", "main.go"]
