package main

import (
	"log"
	"net"

	pb "pedidos-redis-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao iniciar listener: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterEmailServiceServer(server, &serverImpl{})

	// Habilita reflexão para facilitar testes (ex: grpcurl)
	reflection.Register(server)

	log.Println("EmailService rodando na porta 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Erro ao rodar servidor: %v", err)
	}
}
