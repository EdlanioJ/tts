package main

import (
	"os"
	"strconv"

	"github.com/EdlanioJ/tts/domain/usecase"
	"github.com/EdlanioJ/tts/infra/gateway"
	"github.com/EdlanioJ/tts/infra/grpc"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	p := os.Getenv("GRPC_PORT")
	grpcPort, err := strconv.Atoi(p)
	if err != nil {
		panic(err)
	}

	httpClient := gateway.NewHTTPClient()
	ttsUsecase := usecase.NewTextToSpeech(httpClient)

	grpcServer := grpc.NewGRPCServer(grpcPort, ttsUsecase)
	grpcServer.Serve()
}
