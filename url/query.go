package url // import "github.com/lakexyde/wc-api-go/url"

import (
	"net/url"

	"github.com/lakexyde/wc-api-go/request"
)

// QueryEnricher uses package auth to enrich existing query parameters with Authentication Based ones
type QueryEnricher interface {
	GetEnrichedQuery(url string, query url.Values, req request.Request) url.Values
}
