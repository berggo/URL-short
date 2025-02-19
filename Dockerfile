FROM golang:alpine

WORKDIR /app

COPY main.go .

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]
