FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build cmd/main.go


EXPOSE 6969

CMD ["./main"]
