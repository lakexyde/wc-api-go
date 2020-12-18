package url

import (
	URL "net/url"

	"github.com/lakexyde/wc-api-go/request"
)

// QueryEnricherMock ...
type QueryEnricherMock struct {
	query URL.Values
}

// GetEnrichedQuery ...
func (q *QueryEnricherMock) GetEnrichedQuery(url string, query URL.Values, req request.Request) URL.Values {
	return q.query
}
