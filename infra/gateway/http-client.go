package gateway

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/EdlanioJ/tts/domain/gateway"
)

//go:generate mockgen -source=http-client.go -destination=mock/http-client.go -package=mock
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

type googleTranslaterClient struct {
	client  HTTPClient
	baseURL string
}

func NewGoogleTranlaterClient(client HTTPClient) *googleTranslaterClient {
	return &googleTranslaterClient{
		client:  client,
		baseURL: "http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s",
	}
}

func (gt *googleTranslaterClient) GetAudio(input gateway.ClientInput) (io.ReadCloser, error) {
	fileURL := fmt.Sprintf(gt.baseURL, url.QueryEscape(input.Text), input.Lang)
	res, err := gt.client.Get(fileURL)
	if err != nil {
		return nil, errorf("error making request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, errorf("error invalid response: %v", http.StatusText(res.StatusCode))
	}

	return res.Body, nil
}

func errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
