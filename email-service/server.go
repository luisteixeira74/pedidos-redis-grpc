package main

import (
	"context"
	"log"

	pb "pedidos-redis-grpc/proto"
)

type serverImpl struct {
	pb.UnimplementedEmailServiceServer
}

func (s *serverImpl) SendConfirmation(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	log.Printf("Recebido pedido para enviar email da OrderID: %s, body: %s", req.OrderId, req.Body)

	// Aqui poderia chamar um serviço real de email, mas vamos simular só com log.

	return &pb.EmailResponse{
		Status: "E-mail enviado com sucesso",
	}, nil
}
