package client

import (
	"net/http"

	"github.com/lakexyde/wc-api-go/request"
)

// SenderMock imitates sending requests and receiving responses
type SenderMock struct {
	response http.Response
}

// Send ...
func (r *SenderMock) Send(req request.Request) (resp *http.Response, err error) {
	return &r.response, nil
}
