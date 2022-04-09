package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPClient(t *testing.T) {
	client := NewHTTPClient()
	assert.NotNil(t, client)
}
