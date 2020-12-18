package net

import (
	"net/http"

	"github.com/lakexyde/wc-api-go/auth"
)

// RequestEnricherMock ...
type RequestEnricherMock struct {
}

// EnrichRequest ...
func (a *RequestEnricherMock) EnrichRequest(r *http.Request, URL string) {
	auth := auth.Authenticator{}
	if auth.IsSsl(URL) {
		r.SetBasicAuth("key", "secret")
	}
}
