FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o my-go-app .

EXPOSE 8080

CMD ["./my-go-app"]