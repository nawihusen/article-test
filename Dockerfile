FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o alpha-test app/main.go

CMD ["./alpha-test"]
