package client // import "github.com/lakexyde/wc-api-go/client"

import (
	"net/http"

	"github.com/lakexyde/wc-api-go/request"
)

// Sender interface
type Sender interface {
	Send(req request.Request) (resp *http.Response, err error)
}
