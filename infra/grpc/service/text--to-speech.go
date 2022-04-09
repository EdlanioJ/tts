package service

import (
	"context"

	"github.com/EdlanioJ/tts/domain/usecase"
	"github.com/EdlanioJ/tts/infra/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TextToSpeech struct {
	ttsUsecase *usecase.TextToSpeech
	pb.UnimplementedTextToSpeechServer
}

func NewTextToSpeech(ttsUsecase *usecase.TextToSpeech) *TextToSpeech {
	return &TextToSpeech{
		ttsUsecase: ttsUsecase,
	}
}

func (tts *TextToSpeech) Say(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	input := usecase.InputTextToSpeech{
		Language: req.Language,
		Text:     req.Text,
	}

	output, err := tts.ttsUsecase.Exec(input)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Response{
		Audio: output,
	}, nil
}
