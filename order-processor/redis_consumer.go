package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	pb "pedidos-redis-grpc/proto"

	"google.golang.org/grpc"
)

type OrderMessage struct {
	OrderID string `json:"order_id"`
	Message string `json:"message"`
}

func StartRedisConsumer() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	conn, err := grpc.Dial("email-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Erro ao conectar gRPC: %v", err)
	}
	defer conn.Close()
	client := pb.NewEmailServiceClient(conn)

	for {
		result, err := rdb.BLPop(ctx, 0*time.Second, "orderCreated").Result()
		if err != nil {
			log.Printf("Erro no Redis: %v", err)
			continue
		}

		var order OrderMessage
		if err := json.Unmarshal([]byte(result[1]), &order); err != nil {
			log.Printf("Erro ao decodificar: %v", err)
			continue
		}

		log.Printf("Processando pedido: %s", order.OrderID)

		resp, err := client.SendConfirmation(ctx, &pb.EmailRequest{
			OrderId: order.OrderID,
			Body:    order.Message,
		})
		if err != nil {
			log.Printf("Erro ao chamar gRPC: %v", err)
			continue
		}

		log.Printf("Resposta do EmailService: %s", resp.Status)
	}
}
