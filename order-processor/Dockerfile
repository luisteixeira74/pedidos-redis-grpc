FROM golang:1.23

WORKDIR /app

RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

# Copia tudo da pasta order-processor (inclusive redis_consumer.go)
COPY order-processor/ ./order-processor/

# Copia a pasta de protos para os imports funcionarem
COPY proto/ ./proto/


WORKDIR /app/order-processor

RUN go build -o order-processor


CMD ["./order-processor"]
