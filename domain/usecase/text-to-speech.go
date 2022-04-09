package usecase

import (
	"fmt"

	"github.com/EdlanioJ/tts/domain/entity"
	"github.com/EdlanioJ/tts/domain/gateway"
)

const ContentType = "audio/mpeg"

type TextToSpeech struct {
	client gateway.Client
}

// NewTextToSpeech returns a new TextToSpeech service.
// The client is used to make requests to the API.
func NewTextToSpeech(client gateway.Client) *TextToSpeech {
	return &TextToSpeech{
		client: client,
	}
}

// Exec executes the TextToSpeech use case.
// It receives an InputTextToSpeech struct and returns an OutputTextToSpeech struct.
func (speech *TextToSpeech) Exec(input InputTextToSpeech) (OutputTextToSpeech, error) {
	clientInput := gateway.ClientInput{
		Text: input.Text,
		Lang: input.Language,
	}
	res, err := speech.client.GetAudio(clientInput)
	if err != nil {
		return nil, errorf("error getting audio: %v", err)
	}

	audio := entity.NewAudio()

	err = audio.ReadFrom(res)
	if err != nil {
		return nil, errorf("error decoding response: %v", err)
	}

	return audio.Audio, nil
}

func errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
