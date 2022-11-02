FROM golang:latest
RUN mkdir /app
ENV GO_ENV dev
ADD . /app
WORKDIR /app/cmd/orderService
RUN go build -o main .
EXPOSE 4000
CMD ["/app/cmd/orderService/main"]