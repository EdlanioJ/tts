package service

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/EdlanioJ/tts/domain/gateway/mock"
	"github.com/EdlanioJ/tts/domain/usecase"
	"github.com/EdlanioJ/tts/infra/grpc/pb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type TextToSpeechSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	httpClient *mock.MockHTTPClient
	req        *pb.Request
}

func (s *TextToSpeechSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.httpClient = mock.NewMockHTTPClient(s.ctrl)
	s.req = &pb.Request{
		Language: "en",
		Text:     "Hello World",
	}
}

func (s *TextToSpeechSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *TextToSpeechSuite) TestFailure() {
	s.httpClient.EXPECT().Get(gomock.Any()).Return(nil, fmt.Errorf("error"))

	ttsUsecase := usecase.NewTextToSpeech(s.httpClient)
	addr := startTextToSpeechServer(s.T(), ttsUsecase)
	ttsClient := newTextToSpeechClient(addr, s.T())
	res, err := ttsClient.Say(context.TODO(), s.req)

	assert.Nil(s.T(), res)
	assert.Error(s.T(), err)
}

func (s *TextToSpeechSuite) TestSuccess() {
	audioBytes, err := ioutil.ReadFile("../../../testdata/audio.mp3")
	assert.NoError(s.T(), err)

	body := ioutil.NopCloser(bytes.NewReader(audioBytes))
	response := &http.Response{
		Body:       body,
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{usecase.ContentType}},
	}

	s.httpClient.EXPECT().Get(gomock.Any()).Return(response, nil)

	ttsUsecase := usecase.NewTextToSpeech(s.httpClient)
	addr := startTextToSpeechServer(s.T(), ttsUsecase)
	ttsClient := newTextToSpeechClient(addr, s.T())
	res, err := ttsClient.Say(context.TODO(), s.req)

	assert.NotNil(s.T(), res)
	assert.NoError(s.T(), err)
}

func TestTextToSpeechSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TextToSpeechSuite))
}

func startTextToSpeechServer(t *testing.T, ttsUsecase *usecase.TextToSpeech) string {
	ttsService := NewTextToSpeech(ttsUsecase)
	ttsServer := grpc.NewServer()
	pb.RegisterTextToSpeechServer(ttsServer, ttsService)

	listener, err := net.Listen("tcp", ":0")
	assert.NoError(t, err)

	// nolint
	go ttsServer.Serve(listener)
	return listener.Addr().String()
}

func newTextToSpeechClient(addr string, t *testing.T) pb.TextToSpeechClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	assert.NoError(t, err)

	return pb.NewTextToSpeechClient(conn)
}
