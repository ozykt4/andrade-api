FROM golang:1.22.0-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o /andrade cmd/main.go


EXPOSE 8000
EXPOSE 5000

CMD ["./andrade"]
