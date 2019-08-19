package client

import (
	"github.com/ducnt114/wc-api-go/request"
	"net/http"
)

// SenderMock imitates sending requests and receiving responses
type SenderMock struct {
	response http.Response
}

// Send ...
func (r *SenderMock) Send(req request.Request) (resp *http.Response, err error) {
	return &r.response, nil
}
