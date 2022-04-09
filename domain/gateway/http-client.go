package gateway

import "net/http"

// HTTPClient is the interface that wraps the Do method.
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}
