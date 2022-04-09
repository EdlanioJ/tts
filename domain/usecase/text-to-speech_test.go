package usecase

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
		prepare func(*mock.MockClient)
	}{
		{
			name: "error getting audio",
			prepare: func(c *mock.MockClient) {
				c.EXPECT().GetAudio(gomock.Any()).Return(nil, fmt.Errorf("error making request")).Times(1)
			},
			hasErr: true,
		},
		{
			name: "error encoding response",
			prepare: func(c *mock.MockClient) {
				res := ioutil.NopCloser(bytes.NewReader([]byte("audio")))

				c.EXPECT().GetAudio(gomock.Any()).Return(res, nil).Times(1)
			},
			hasErr: true,
		},
		{
			name: "success",
			prepare: func(c *mock.MockClient) {
				res := ioutil.NopCloser(bytes.NewReader(audioBytes))

				c.EXPECT().GetAudio(gomock.Any()).Return(res, nil).Times(1)
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

			client := mock.NewMockClient(ctrl)
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
