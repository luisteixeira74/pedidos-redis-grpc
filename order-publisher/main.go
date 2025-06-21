package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type OrderMessage struct {
	OrderID string `json:"order_id"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/publisher", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var order OrderMessage
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "Erro ao decodificar body: "+err.Error(), http.StatusBadRequest)
			return
		}

		err := publishOrder(order)
		if err != nil {
			http.Error(w, "Erro ao publicar: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Mensagem publicada com sucesso"))
	})

	log.Println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
	}
}

func publishOrder(order OrderMessage) error {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	log.Printf("Publicando pedido ID %s no Redis...\n", order.OrderID)

	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = rdb.RPush(context.Background(), "orderCreated", data).Err()
	if err != nil {
		return err
	}

	log.Println("Mensagem publicada com sucesso.")
	return nil
}
