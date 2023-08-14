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
	directions_SuccessResponse string = `{"origin":{"coordinate":{"latitude":51.0453064,"longitude":13.7359337},"displayMapRegion":{"southLatitude":51.0408148235794,"westLongitude":13.728789535983225,"northLatitude":51.0497979764206,"eastLongitude":13.743077864016776},"name":"Prager Straße 15","formattedAddressLines":["Prager Straße 15","01069 Dresden","Germany"],"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01069","subLocality":"Friedrichstadt","thoroughfare":"Prager Straße","subThoroughfare":"15","fullThoroughfare":"Prager Straße 15","areasOfInterest":["Centrum Galerie"],"dependentLocalities":["Seevorstadt-Ost/GroßerGarten/Strehlen-Nordwest","Seevorstadt-Ost/Großer Garten/Strehlen-Nordwest"]},"country":"Germany","countryCode":"DE"},"destination":{"coordinate":{"latitude":51.0414609,"longitude":13.7340304},"displayMapRegion":{"southLatitude":51.0369693235794,"westLongitude":13.726886828999827,"northLatitude":51.045952476420595,"eastLongitude":13.741173971000173},"name":"Prager Straße 1","formattedAddressLines":["Prager Straße 1","01069 Dresden","Germany"],"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01069","subLocality":"Friedrichstadt","thoroughfare":"Prager Straße","subThoroughfare":"1","fullThoroughfare":"Prager Straße 1","dependentLocalities":["Seevorstadt-Ost/GroßerGarten/Strehlen-Nordwest","Seevorstadt-Ost/Großer Garten/Strehlen-Nordwest"]},"country":"Germany","countryCode":"DE"},"routes":[{"name":"Prager Straße","distanceMeters":317,"durationSeconds":149,"transportType":"Automobile","stepIndexes":[0,1,2,3,4,5],"hasTolls":false}],"steps":[{"stepPathIndex":0,"distanceMeters":0,"durationSeconds":0},{"stepPathIndex":1,"distanceMeters":93,"durationSeconds":64,"instructions":"Turn right onto St Petersburger Straße"},{"stepPathIndex":2,"distanceMeters":77,"durationSeconds":25,"instructions":"Turn right"},{"stepPathIndex":3,"distanceMeters":78,"durationSeconds":8,"instructions":"Prepare to park your car near Prager Straße"},{"stepPathIndex":4,"distanceMeters":0,"durationSeconds":0,"instructions":"Take a left onto Prager Straße","transportType":"WALKING"},{"stepPathIndex":5,"distanceMeters":68,"durationSeconds":52,"instructions":"The destination is on your right","transportType":"WALKING"}],"stepPaths":[[{"latitude":51.042699,"longitude":13.735173}],[{"latitude":51.042699,"longitude":13.735173},{"latitude":51.042634,"longitude":13.735153},{"latitude":51.042588,"longitude":13.735206},{"latitude":51.042558,"longitude":13.735321},{"latitude":51.042557,"longitude":13.735483},{"latitude":51.042449,"longitude":13.735946},{"latitude":51.042429,"longitude":13.736012},{"latitude":51.042367,"longitude":13.736183},{"latitude":51.042341,"longitude":13.736249}],[{"latitude":51.042341,"longitude":13.736249},{"latitude":51.04223,"longitude":13.736121},{"latitude":51.041902,"longitude":13.735773},{"latitude":51.041879,"longitude":13.73575},{"latitude":51.041842,"longitude":13.735712},{"latitude":51.041765,"longitude":13.735636}],[{"latitude":51.041765,"longitude":13.735636},{"latitude":51.041796,"longitude":13.735554},{"latitude":51.041921,"longitude":13.735061},{"latitude":51.042003,"longitude":13.734589},{"latitude":51.042004,"longitude":13.734587}],[{"latitude":51.042004,"longitude":13.734587}],[{"latitude":51.042004,"longitude":13.734587},{"latitude":51.04141,"longitude":13.734339}]]}`
	etas_SuccessResponse       string = `{"etas":[{"destination":{"latitude":51.0453064,"longitude":13.7459337},"transportType":"Automobile","distanceMeters":1607,"expectedTravelTimeSeconds":409,"staticTravelTimeSeconds":333}]}`
)

type DirectionsTestSuite struct {
	suite.Suite
	testServer *httptest.Server
	mapsClient Client
}

func (s *DirectionsTestSuite) SetupSuite() {
	mux := http.NewServeMux()
	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(accessToken_SuccessResponse))
	})
	mux.HandleFunc("/directions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(directions_SuccessResponse))
	})
	mux.HandleFunc("/etas", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(etas_SuccessResponse))
	})
	s.testServer = httptest.NewServer(mux)
	s.mapsClient = NewAppleMaps(s.testServer.Client(), "jwt", WithCustomURL(s.testServer.URL))
}

func (s *DirectionsTestSuite) TestDirections_Success() {
	var expected = &DirectionsResponse{}
	json.Unmarshal([]byte(directions_SuccessResponse), expected)
	res, err := s.mapsClient.Directions(context.Background(), "Origin", "Destination", WithUserLocation(NewLocation(1, 1)))
	s.NoError(err)
	s.Equal(expected, res)
}

func (s *DirectionsTestSuite) TestEtas_Success() {
	var expected = &EtaResponse{}
	json.Unmarshal([]byte(etas_SuccessResponse), expected)
	res, err := s.mapsClient.Etas(context.Background(), NewLocation(1, 1), []Location{NewLocation(1, 1)}, WithUserLocation(NewLocation(1, 1)))
	s.NoError(err)
	s.Equal(expected, res)
}

func TestDirectionsTestSuite(t *testing.T) {
	suite.Run(t, new(DirectionsTestSuite))
}
