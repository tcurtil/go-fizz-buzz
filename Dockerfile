#Compile stage
FROM golang:1.12.6-alpine3.10 AS build-env
ENV CGO_ENABLED 0
RUN apk add --no-cache git
ADD . /go/src/github.com/tcurtil/go-fizz-buzz

# Install revel framework
RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
#build revel app
RUN revel build github.com/tcurtil/go-fizz-buzz /tmp/fizzbuzz -m prod

# Final stage
FROM alpine:3.10
EXPOSE 9000
WORKDIR /
COPY --from=build-env /tmp/fizzbuzz /
ENTRYPOINT /run.sh