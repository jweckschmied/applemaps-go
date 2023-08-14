package applemaps

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	geocode_SuccessResponse        string = `{"results":[{"coordinate":{"latitude":51.0658585,"longitude":13.7466163},"displayMapRegion":{"southLatitude":51.0613669235794,"westLongitude":13.739468964416949,"northLatitude":51.070350076420596,"eastLongitude":13.75376363558305},"name":"Königsbrücker Straße 15","formattedAddressLines":["Königsbrücker Straße 15","01099 Dresden","Germany"],"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01099","subLocality":"Albertstadt","thoroughfare":"Königsbrücker Straße","subThoroughfare":"15","fullThoroughfare":"Königsbrücker Straße 15","dependentLocalities":["Äußere Neustadt"]},"country":"Germany","countryCode":"DE"}]}`
	reverseGeocode_SuccessResponse string = `{"results":[{"coordinate":{"latitude":51.0813007, "longitude":13.7603922}, "displayMapRegion":{"southLatitude":51.0768091235794, "westLongitude":13.753242478943124, "northLatitude":51.0857922764206, "eastLongitude":13.767541921056877}, "name":"Königsbrücker Straße 96", "formattedAddressLines":["Königsbrücker Straße 96","01099 Dresden", "Germany"], "structuredAddress":{"administrativeArea":"Saxony", "locality":"Dresden", "postCode":"01099", "subLocality":"Albertstadt", "thoroughfare":"Königsbrücker Straße","subThoroughfare":"96", "fullThoroughfare":"Königsbrücker Straße 96", "dependentLocalities":["Albertstadt"]}, "country":"Germany", "countryCode":"DE"}]}`
)

type GeocodeTestSuite struct {
	suite.Suite
	testServer *httptest.Server
	mapsClient Client
}

func (s *GeocodeTestSuite) SetupSuite() {
	mux := http.NewServeMux()
	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(accessToken_SuccessResponse))
	})
	mux.HandleFunc("/geocode", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(geocode_SuccessResponse))
	})
	mux.HandleFunc("/reverseGeocode", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(reverseGeocode_SuccessResponse))
	})
	s.testServer = httptest.NewServer(mux)
	s.mapsClient = NewAppleMaps(s.testServer.Client(), "jwt", WithCustomURL(s.testServer.URL))
}

func (s *GeocodeTestSuite) TearDownSuite() {
	s.testServer.Close()
}

func (s *GeocodeTestSuite) TestGeocode_Success() {
	var expected = &SearchResponse{}
	json.Unmarshal([]byte(geocode_SuccessResponse), expected)
	res, err := s.mapsClient.Geocode(context.Background(), "test query", WithUserLocation(NewLocation(1, 1)))
	s.NoError(err)
	s.Equal(expected.Results, res)
}

func (s *GeocodeTestSuite) TestReverseGeocode_Success() {
	var expected = &SearchResponse{}
	json.Unmarshal([]byte(reverseGeocode_SuccessResponse), expected)
	res, err := s.mapsClient.ReverseGeocode(context.Background(), NewLocation(1, 1), WithUserLocation(NewLocation(1, 1)))
	s.NoError(err)
	s.Equal(expected.Results, res)
}

func TestGeocodeTestSuite(t *testing.T) {
	suite.Run(t, new(GeocodeTestSuite))
}
