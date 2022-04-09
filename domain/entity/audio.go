package entity

import (
	"fmt"
	"io"
	"net/http"
)

const (
	ContentTypeAudio  = "audio/mpeg"
	ContentTypeStream = "application/octet-stream"
)

func NewAudio() *Audio {
	return &Audio{}
}

type Audio struct {
	Audio []byte
}

func (a *Audio) ReadFrom(r io.Reader) error {
	b, _ := io.ReadAll(r)
	a.Audio = b

	err := a.isValid()
	if err != nil {
		return fmt.Errorf("error validate: %v", err)
	}

	return nil
}

func (a *Audio) isValid() error {
	mineType := http.DetectContentType(a.Audio)

	if mineType != ContentTypeAudio && mineType != ContentTypeStream {
		return fmt.Errorf("error invalid content type: %v", mineType)
	}
	return nil
}
