package gateway

import "io"

type ClientInput struct {
	Text string
	Lang string
}

type Client interface {
	GetAudio(ClientInput) (io.ReadCloser, error)
}
