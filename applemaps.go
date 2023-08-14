package applemaps

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	apiBase                    = "https://maps-api.apple.com/v1"
	tokenEndpoint              = "token"
	geocodeEndpoint            = "geocode"
	reverseGeocodeEndpoint     = "reverseGeocode"
	searchEndpoint             = "search"
	searchAutocompleteEndpoint = "searchAutocomplete"
	etasEndpoint               = "etas"
	directionsEndpoint         = "directions"
	defaultOffset              = 10 * time.Second
)

type (
	AccessToken struct {
		Token      string `json:"accessToken"`
		Expiration int    `json:"expiresInSeconds"`
	}

	Client interface {
		Geocode(ctx context.Context, query string, opts ...RequestOption) ([]Place, error)
		ReverseGeocode(ctx context.Context, location Location, opts ...RequestOption) ([]Place, error)

		Search(ctx context.Context, query string, opts ...RequestOption) (*SearchResponse, error)
		SearchAutocomplete(ctx context.Context, query string, opts ...RequestOption) (*SearchAutocompleteResult, error)

		Directions(ctx context.Context, origin, destination string, opts ...RequestOption) (*DirectionsResponse, error)
		Etas(ctx context.Context, origin Location, destinations []Location, opts ...RequestOption) (*EtaResponse, error)

		SetAuthToken(authToken string)
	}

	client struct {
		client      *http.Client
		accessToken AccessToken
		authToken   string
		nextRenewal time.Time
		baseURL     string
	}
)

// NewAppleMaps returns a new Apple Maps Server API Client
// given a http client and a JWT Auth Token for the API.
// If you need to specify a custom API URL, use WithCustomURL() as an option.
func NewAppleMaps(httpClient *http.Client, authToken string, options ...ClientOption) Client {
	mapsClient := &client{
		client:      httpClient,
		baseURL:     apiBase,
		authToken:   authToken,
		nextRenewal: time.Now(),
	}
	for _, o := range options {
		o(mapsClient)
	}
	return mapsClient
}

type ClientOption func(c *client)

// WithCustomURL returns a functional ClientOption used to set a custom base URL
// when creating a new Apple Maps API Client using NewAppleMaps().
func WithCustomURL(baseURL string) ClientOption {
	return func(c *client) {
		c.baseURL = baseURL
	}
}

// getAccessToken either returns the current access token, or requests a new one if needed.
// A new token will be requested and returned only if the current token expires within the next 10 seconds or is already expired,
// or if no access token exists yet.
func (c *client) getAccessToken() (string, error) {
	if time.Now().After(c.nextRenewal.Add(-defaultOffset)) {
		reader, err := c.doRequest(context.Background(), c.authToken, tokenEndpoint, nil)
		if err != nil {
			return "", err
		}
		var accessToken AccessToken
		if err := json.NewDecoder(reader).Decode(&accessToken); err != nil {
			return "", err
		}
		c.nextRenewal = time.Now().Add(time.Duration(accessToken.Expiration) * time.Second)
		c.accessToken = accessToken
	}
	return c.accessToken.Token, nil
}

// doAuthenticatedRequest wraps doRequest() with a call to retrieve the currently valid access token to perform the request.
func (c *client) doAuthenticatedRequest(ctx context.Context, endpoint string, params url.Values) (io.Reader, error) {
	token, err := c.getAccessToken()
	if err != nil {
		return nil, err
	}
	return c.doRequest(ctx, token, endpoint, params)
}

// doRequest performs the http request, given the access token, api endpoint and query parameters.
func (c *client) doRequest(ctx context.Context, auth string, endpoint string, params url.Values) (io.Reader, error) {
	path, err := url.JoinPath(c.baseURL, endpoint)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+auth)
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	switch res.StatusCode {
	case http.StatusOK:
		return res.Body, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("unauthorized")
	case http.StatusBadRequest:
		return nil, fmt.Errorf("bad request: %s", unmarshalErrorResponse(res.Body).Error.Message)
	case http.StatusTooManyRequests:
		return nil, fmt.Errorf("API rate limit reached: %s", unmarshalErrorResponse(res.Body).Error.Message)
	default:
		errRes := unmarshalErrorResponse(res.Body)
		return res.Body, fmt.Errorf("API Error %d, Message: %s, Details: %s", res.StatusCode, errRes.Error.Message, errRes.Error.Details)
	}
}

// SetAuthToken sets a new JWT for the Apple Maps Client.
// You can use this method to set a new token when the old one is about to expire.
func (c *client) SetAuthToken(authToken string) {
	c.authToken = authToken
	c.nextRenewal = time.Now()
}

// exec is a generic wrapper around doAuthenticatedRequest(), that decodes the returned data into the specified
// type T, and returns a pointer to T.
func exec[T any](ctx context.Context, c *client, endpoint string, values url.Values) (*T, error) {
	var res = new(T)
	reader, err := c.doAuthenticatedRequest(ctx, endpoint, values)
	if err != nil {
		return res, err
	}

	if err := json.NewDecoder(reader).Decode(res); err != nil {
		return res, err
	}

	return res, nil
}
