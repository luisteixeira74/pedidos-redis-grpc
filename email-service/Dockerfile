FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY email-service/ ./email-service/
COPY proto/ ./proto/

WORKDIR /app/email-service
RUN go build -o email-service .

CMD ["./email-service"]