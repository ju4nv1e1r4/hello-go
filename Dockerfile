FROM golang:1.19

WORKDIR /hello-go

COPY hello.go ./

CMD ["go", "run", "hello.go"]
