package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/EdlanioJ/tts/domain/usecase"
	"github.com/EdlanioJ/tts/infra/grpc/pb"
	"github.com/EdlanioJ/tts/infra/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	port    int
	usecase *usecase.TextToSpeech
}

func NewGRPCServer(port int, usecase *usecase.TextToSpeech) *grpcServer {
	return &grpcServer{port, usecase}
}

func (s *grpcServer) Serve() {
	server := grpc.NewServer()

	reflection.Register(server)
	ttsService := service.NewTextToSpeech(s.usecase)

	pb.RegisterTextToSpeechServer(server, ttsService)
	address := fmt.Sprintf("0.0.0.0:%d", s.port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("gRPC server started at port %d", s.port)
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
