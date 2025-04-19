FROM golang:1.24

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -o /app/main

EXPOSE 8000

CMD ["/app/main"]