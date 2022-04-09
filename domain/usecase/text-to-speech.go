package usecase

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/EdlanioJ/tts/domain/entity"
	"github.com/EdlanioJ/tts/domain/gateway"
)

const ContentType = "audio/mpeg"

type TextToSpeech struct {
	client  gateway.HTTPClient
	baseURL string
}

// NewTextToSpeech returns a new TextToSpeech service.
// The client is used to make requests to the API.
func NewTextToSpeech(client gateway.HTTPClient) *TextToSpeech {
	return &TextToSpeech{
		client:  client,
		baseURL: "http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s",
	}
}

// Exec executes the TextToSpeech use case.
// It receives an InputTextToSpeech struct and returns an OutputTextToSpeech struct.
func (speech *TextToSpeech) Exec(input InputTextToSpeech) (OutputTextToSpeech, error) {
	fileURL := fmt.Sprintf(speech.baseURL, url.QueryEscape(input.Text), input.Language)
	res, err := speech.client.Get(fileURL)
	if err != nil {
		return nil, errorf("error making request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errorf("error invalid response: %v", http.StatusText(res.StatusCode))
	}

	audio := entity.NewAudio()

	err = audio.ReadFrom(res.Body)
	if err != nil {
		return nil, errorf("error decoding response: %v", err)
	}

	return audio.Audio, nil
}

func errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
