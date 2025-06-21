
# Esse projeto simula 3 microservices (order-publisher, order-processor, email-service) que se comunicam de formas diferentes (redis pub/sub e gRPC) para receber uma Ordem de Pedido.

# O serviço order-publisher aguarda uma chamada REST com um pedido.

## subir os containers (mantendo o terminal para ver logs)
> docker compose up --build

## Para simular um pedido (POST)

curl -X POST http://localhost:8080/publisher \
  -H "Content-Type: application/json" \
  -d '{"order_id": "pedido456", "message": "Pedido gerado via POST"}'

## Após o order-publisher receber o pedido, ele publica o pedido numa fila com Redis que será lida pelo serviço order-processor. O order-processor se comunica com o email-service via gRPC e envia o pedido, que informa o e-mail enviado com sucesso.

## Antes de iniciar confirmar se os arquivos do protobuffer foram gerados na pasta order-processor. 
## Esses arquivos são compartilhados entre o order-processor e email-service para comunicação grpc.
## O Makefile foi criado para tentar executar o comando para gerar o protobuffer com make proto

### Para ver os logs em tempo real

## Sugestão de teste no modo 'detached'
> docker compose up -d

## Para ver os logs em tempo real abra outro terminal e rode:

> docker compose logs -f order-processor
> docker compose logs -f email-service

## Exemplo de saida de logs:

order-publisher  | 2025/06/21 14:33:19 Publicando pedido ID pedido456 no Redis...
order-publisher  | 2025/06/21 14:33:19 Mensagem publicada com sucesso.
order-processor  | 2025/06/21 14:33:19 Processando pedido: pedido456
email-service    | 2025/06/21 14:33:19 Recebido pedido para enviar email da OrderID: pedido456, body: Pedido gerado via POST
order-processor  | 2025/06/21 14:33:19 Resposta do EmailService: E-mail enviado com sucesso

### Containers envolvidos:

| Serviço         | Responsabilidade                 | Porta | Comunicação          |
| --------------- | -------------------------------- | ----- | -------------------- |
| Redis           | Fila intermediária               | 6379  | BLPOP / RPUSH        |
| order-publisher | Publica um pedido na fila        | 8080  | REST → Redis         |
| order-processor | Consome da fila e envia via gRPC | N/A   | Redis → gRPC         |
| email-service   | Recebe gRPC com dados do pedido  | 50051 | gRPC → log/simulação |


### Comunicação Entre Serviços
REST (HTTP)

Exposto pelo order-publisher em /publisher

Serve para simular um pedido sendo feito.

Redis

Usado como fila de mensagens

order-publisher faz RPUSH na key orderCreated

order-processor faz BLPOP na mesma key para consumo assíncrono

gRPC

order-processor chama email-service.SendConfirmation com os dados do pedido



# pedidos-redis-grpc
