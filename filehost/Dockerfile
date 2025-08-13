FROM golang:1.23.11-alpine

WORKDIR /app

COPY . .

RUN go mod download && \
go build -o main .

EXPOSE 80

CMD ["./main"]
