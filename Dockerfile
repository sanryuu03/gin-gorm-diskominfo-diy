FROM golang:1.23.2

WORKDIR /app

# Installing dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

EXPOSE 8080/tcp

CMD ["go", "run", "main.go"]