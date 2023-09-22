FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
EXPOSE 8088
RUN go build -o main .

CMD ["./main"]