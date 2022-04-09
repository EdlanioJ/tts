package usecase

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/EdlanioJ/tts/domain/gateway/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTextToSpeech(t *testing.T) {
	audioBytes, err := ioutil.ReadFile("../../testdata/audio.mp3")
	assert.NoError(t, err)

	testCases := []struct {
		name    string
		hasErr  bool
		prepare func(*mock.MockHTTPClient)
	}{
		{
			name: "error making request",
			prepare: func(c *mock.MockHTTPClient) {
				c.EXPECT().Get(gomock.Any()).Return(nil, fmt.Errorf("error making request")).Times(1)
			},
			hasErr: true,
		},
		{
			name: "error invalid response",
			prepare: func(c *mock.MockHTTPClient) {
				body := ioutil.NopCloser(bytes.NewReader([]byte("audio")))

				response := &http.Response{
					Body:       body,
					StatusCode: http.StatusBadRequest,
				}
				c.EXPECT().Get(gomock.Any()).Return(response, nil).Times(1)
			},
			hasErr: true,
		},
		{
			name: "error encoding response",
			prepare: func(c *mock.MockHTTPClient) {
				body := ioutil.NopCloser(bytes.NewReader([]byte("audio")))
				response := &http.Response{
					Body:       body,
					StatusCode: http.StatusOK,
					Header:     http.Header{"Content-Type": []string{ContentType}},
				}

				c.EXPECT().Get(gomock.Any()).Return(response, nil).Times(1)
			},
			hasErr: true,
		},
		{
			name: "success",
			prepare: func(c *mock.MockHTTPClient) {
				body := ioutil.NopCloser(bytes.NewReader(audioBytes))
				response := &http.Response{
					Body:       body,
					StatusCode: http.StatusOK,
					Header:     http.Header{"Content-Type": []string{ContentType}},
				}

				c.EXPECT().Get(gomock.Any()).Return(response, nil).Times(1)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			client := mock.NewMockHTTPClient(ctrl)
			tc.prepare(client)

			ttsService := NewTextToSpeech(client)

			input := InputTextToSpeech{
				Text:     "hello",
				Language: "en",
			}
			output, err := ttsService.Exec(input)
			if tc.hasErr {
				is.Error(err)
				is.Equal(output, OutputTextToSpeech(nil))
			} else {
				is.NoError(err)
				is.Equal(output, OutputTextToSpeech(audioBytes))
			}
		})
	}
}
