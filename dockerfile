FROM golang:1.22-alpine
WORKDIR /app
# Preparamos las dependencias
RUN go mod init core-cold || true
RUN go get github.com/go-redis/redis/v8
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
