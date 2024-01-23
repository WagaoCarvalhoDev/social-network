FROM golang:latest

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
