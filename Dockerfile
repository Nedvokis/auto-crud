FROM golang:1.17.2-alpine3.14

RUN mkdir /src
ADD . /src
WORKDIR /src
RUN go build -o main .
CMD ["/src/main"]
