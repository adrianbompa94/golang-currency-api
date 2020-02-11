FROM golang:alpine

RUN apk add --no-cache git=2.22.2-r0 \
    --repository https://alpine.global.ssl.fastly.net/alpine/v3.10/community \
    --repository https://alpine.global.ssl.fastly.net/alpine/v3.10/main

ENV GOPATH /app
ENV GOBIN=$GOPATH/bin

RUN mkdir /app 
ADD . /app/

WORKDIR /app

RUN go get .
RUN go build -o main .

RUN adduser -S -D -H -h /app appuser
USER appuser

CMD ["./main"]