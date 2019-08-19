package net // import "github.com/ducnt114/wc-api-go/net"

import (
	"github.com/ducnt114/wc-api-go/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
