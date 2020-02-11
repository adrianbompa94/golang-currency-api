FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]