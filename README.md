# pedidos-redis-grpc

Este projeto simula **3 microserviços** (`order-publisher`, `order-processor`, `email-service`) que se comunicam de formas diferentes — usando **Redis** como fila de mensagens e **gRPC** entre serviços — para processar pedidos de forma assíncrona.

---

## 📦 Visão Geral dos Serviços

| Serviço         | Responsabilidade                         | Porta  | Comunicação           |
|----------------|-------------------------------------------|--------|------------------------|
| Redis           | Fila intermediária                       | `6379` | `BLPOP` / `RPUSH`      |
| order-publisher | Recebe pedidos via HTTP e envia à fila   | `8080` | REST → Redis           |
| order-processor | Consome fila e envia mensagem via gRPC   | —      | Redis → gRPC           |
| email-service   | Recebe gRPC com dados do pedido          | `50051`| gRPC → Log/simulação   |

---

## 🚀 Subir os containers

Para iniciar todos os serviços com rebuild e exibir os logs no terminal principal:

```bash
docker compose up --build
```
Ou no modo detached

```bash
docker compose up -d
```

Simule o envio de um pedido via curl:

curl -X POST http://localhost:8080/publisher \
  -H "Content-Type: application/json" \
  -d '{"order_id": "pedido456", "message": "Pedido gerado via POST"}'


Após o order-publisher receber o pedido, ele publica o pedido numa fila com Redis que será lida pelo serviço order-processor. O order-processor se comunica com o email-service via gRPC e envia o pedido, que informa o e-mail enviado com sucesso.

Antes de iniciar confirmar se os arquivos do protobuffer foram gerados na pasta order-processor. 
Esses arquivos são compartilhados entre o order-processor e email-service para comunicação grpc.
O Makefile foi criado para tentar executar o comando para gerar o protobuffer com make proto

## Para ver os logs em tempo real

## Sugestão de teste no modo 'detached'
> docker compose up -d

## Para ver os logs em tempo real abra outro terminal e rode:

> docker compose logs -f order-processor
> docker compose logs -f email-service

## Exemplo de saida de logs:

order-publisher  | Publicando pedido ID pedido456 no Redis...
order-publisher  | Mensagem publicada com sucesso.
order-processor  | Processando pedido: pedido456
email-service    | Recebido pedido para enviar email da OrderID: pedido456, body: Pedido gerado via POST
order-processor  | Resposta do EmailService: E-mail enviado com sucesso

### 🔧 Requisitos
Docker
Docker Compose
Protoc (caso queira regenerar os arquivos .pb.go)


### ✅ Resumo da Comunicação
REST (HTTP)
Exposto por order-publisher em /publisher.

Redis (Fila)
order-publisher faz RPUSH na chave orderCreated.

order-processor consome com BLPOP.

gRPC
order-processor chama email-service.SendConfirmation() com os dados do pedido.
