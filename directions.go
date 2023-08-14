package applemaps

import (
	"context"
	"errors"
	"net/url"
)

type Avoid string

const Tolls Avoid = "Tolls"

func (a Avoid) String() string {
	return string(a)
}

// Directions returns directions between origin and destination.
// Both origin and destination can be specified as either an address string or coordinates in the format "Lat|Lon",
// for example origin="37.7857,-122.4011"
func (c *client) Directions(ctx context.Context, origin, destination string, opts ...RequestOption) (*DirectionsResponse, error) {
	values := url.Values{}
	values.Add("origin", origin)
	values.Add("destination", destination)
	for _, opt := range opts {
		opt(values)
	}
	return exec[DirectionsResponse](ctx, c, directionsEndpoint, values)
}

// Etas returns the estimated time of arrival (ETA) and distance between origin and destination locations.
func (c *client) Etas(ctx context.Context, origin Location, destinations []Location, opts ...RequestOption) (*EtaResponse, error) {
	if len(destinations) == 0 {
		return nil, errors.New("destinations cannot be empty")
	}
	values := url.Values{}
	values.Add("origin", origin.String())
	values.Add("destinations", queryParameterString(destinations))
	for _, opt := range opts {
		opt(values)
	}
	return exec[EtaResponse](ctx, c, etasEndpoint, values)
}
