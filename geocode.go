package applemaps

import (
	"context"
	"net/url"
)

// Geocode returns the latitude and longitude of the specified address.
func (c *client) Geocode(ctx context.Context, query string, opts ...RequestOption) ([]Place, error) {
	values := url.Values{}
	values.Add("q", query)
	for _, opt := range opts {
		opt(values)
	}
	res, err := exec[SearchResponse](ctx, c, geocodeEndpoint, values)
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}

// ReverseGeocode returns a slice of addresses present at the specified location coordinates.
func (c *client) ReverseGeocode(ctx context.Context, location Location, opts ...RequestOption) ([]Place, error) {
	values := url.Values{}
	values.Add("loc", location.String())
	for _, opt := range opts {
		opt(values)
	}
	res, err := exec[SearchResponse](ctx, c, reverseGeocodeEndpoint, values)
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}
