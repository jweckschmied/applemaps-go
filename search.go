package applemaps

import (
	"context"
	"net/url"
)

// Search performs a search to find places that match specific criteria.
func (c *client) Search(ctx context.Context, query string, opts ...RequestOption) (*SearchResponse, error) {
	values := url.Values{}
	values.Add("q", query)
	for _, opt := range opts {
		opt(values)
	}
	return exec[SearchResponse](ctx, c, searchEndpoint, values)
}

// SearchAutocomplete performs a request to find results for places that you can use to autocomplete searches.
func (c *client) SearchAutocomplete(ctx context.Context, query string, opts ...RequestOption) (*SearchAutocompleteResult, error) {
	values := url.Values{}
	values.Add("q", query)
	for _, opt := range opts {
		opt(values)
	}
	return exec[SearchAutocompleteResult](ctx, c, searchAutocompleteEndpoint, values)
}
