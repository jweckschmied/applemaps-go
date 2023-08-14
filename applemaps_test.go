package applemaps

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	accessToken_UnauthorizedResponse string = `{"error":{"message":"Not Authorized","details":[]}}`
	accessToken_SuccessResponse      string = `{"accessToken":"thisis.thejwt.token","expiresInSeconds":1800}`
)

func TestGetAccessToken_Unauthorized(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(accessToken_UnauthorizedResponse))
	}))

	defer testServer.Close()

	// manually instantiating the client here to test the unexported getAccessToken() method in isolation
	mapsClient := &client{
		testServer.Client(),
		AccessToken{
			Token:      "jwt",
			Expiration: 1800,
		},
		"",
		time.Now(),
		apiBase,
	}
	_, err := mapsClient.getAccessToken()
	assert.Error(t, err)
	assert.Empty(t, mapsClient.authToken)
}

func TestGetAccessToken_NoRefresh(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer testServer.Close()

	// manually instantiating the client here to test the unexported getAccessToken() method in isolation
	nextRenewal := time.Now().Add(30 * time.Second)
	mapsClient := &client{
		testServer.Client(),
		AccessToken{
			Token:      "the.old.jwt",
			Expiration: 1800,
		},
		"authToken",
		nextRenewal,
		testServer.URL,
	}
	_, err := mapsClient.getAccessToken()
	assert.NoError(t, err)
	assert.Equal(t, "the.old.jwt", mapsClient.accessToken.Token)
	assert.Equal(t, nextRenewal, mapsClient.nextRenewal)
}

func TestGetAccessToken_Refresh(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(accessToken_SuccessResponse))
	}))
	defer testServer.Close()

	// manually instantiating the client here to test the unexported getAccessToken() method in isolation
	nextRenewal := time.Now().Add(5 * time.Second)
	mapsClient := &client{
		testServer.Client(),
		AccessToken{
			Token:      "the.old.jwt",
			Expiration: 1800,
		},
		"authToken",
		nextRenewal,
		testServer.URL,
	}

	tok, err := mapsClient.getAccessToken()
	assert.NoError(t, err)
	assert.Equal(t, "thisis.thejwt.token", tok)
	assert.Equal(t, "thisis.thejwt.token", mapsClient.accessToken.Token)
	assert.NotEqual(t, nextRenewal, mapsClient.nextRenewal)
}

func TestSetAuthToken(t *testing.T) {
	mapsClient := &client{
		http.DefaultClient,
		AccessToken{
			Token:      "jwt",
			Expiration: 1800,
		},
		"old-token",
		time.Now(),
		"url",
	}
	mapsClient.SetAuthToken("new-token")
	assert.Equal(t, "new-token", mapsClient.authToken)
}
