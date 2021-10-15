FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

ENV PORT 5001

EXPOSE 5001

CMD ["./main"]
