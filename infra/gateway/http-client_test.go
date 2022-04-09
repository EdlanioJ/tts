package gateway

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/EdlanioJ/tts/domain/gateway"
	"github.com/EdlanioJ/tts/infra/gateway/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHTTPClient(t *testing.T) {
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

				res := &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       body,
				}
				c.EXPECT().Get(gomock.Any()).Return(res, nil).Times(1)
			},
			hasErr: true,
		},
		{
			name: "success",
			prepare: func(c *mock.MockHTTPClient) {
				body := ioutil.NopCloser(bytes.NewReader([]byte("audio")))

				res := &http.Response{
					StatusCode: http.StatusOK,
					Body:       body,
				}
				c.EXPECT().Get(gomock.Any()).Return(res, nil).Times(1)
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

			httpClient := NewGoogleTranlaterClient(client)

			input := gateway.ClientInput{
				Text: "test",
				Lang: "en",
			}
			_, err := httpClient.GetAudio(input)

			if tc.hasErr {
				is.Error(err)
			} else {
				is.NoError(err)
			}
		})
	}
}
