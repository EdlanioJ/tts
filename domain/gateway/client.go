package gateway

import "io"

type ClientInput struct {
	Text string
	Lang string
}

type ClientOutput io.Reader

type Client interface {
	GetAudio(ClientInput) (ClientOutput, error)
}
