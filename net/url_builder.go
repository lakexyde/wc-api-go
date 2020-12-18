package net // import "github.com/lakexyde/wc-api-go/net"

import (
	"github.com/lakexyde/wc-api-go/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
