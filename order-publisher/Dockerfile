FROM golang:1.23

# Define diretório raiz
WORKDIR /app

# Copia go.mod da raiz do projeto
COPY ../go.mod ../go.sum ./

RUN go mod download

# Copia somente a pasta do serviço
COPY order-publisher/ ./order-publisher/

# Define o diretório correto onde estão os arquivos Go
WORKDIR /app/order-publisher

# Compila o binário
RUN go build -o order-publisher .

EXPOSE 8080

CMD ["./order-publisher"]
