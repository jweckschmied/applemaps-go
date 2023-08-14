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
	search_SuccessResponse             string = `{"displayMapRegion":{"southLatitude":51.03468425106257,"westLongitude":13.722250806167722,"northLatitude":51.05481251142919,"eastLongitude":13.747509084641933},"results":[{"coordinate":{"latitude":51.0489734,"longitude":13.7382522},"displayMapRegion":{"southLatitude":51.0444774235794,"westLongitude":13.731076071050252,"northLatitude":51.0534605764206,"eastLongitude":13.74536552894975},"name":"Coffeeplace","formattedAddressLines":["Altmarkt 7","01067 Dresden","Germany"],"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01067","subLocality":"Friedrichstadt","thoroughfare":"Altmarkt","subThoroughfare":"7","fullThoroughfare":"Altmarkt 7","dependentLocalities":["Altstadt Centre","InnereAltstadt"]},"country":"Germany","countryCode":"DE","poiCategory":"Cafe"},{"coordinate":{"latitude":51.0453064,"longitude":13.7359337},"displayMapRegion":{"southLatitude":51.04041895715622,"westLongitude":13.728854314450537,"northLatitude":51.049402109997416,"eastLongitude":13.74314252037848},"name":"Coffeeplace","formattedAddressLines":["Prager Straße","01069 Dresden","Germany"],"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01069","subLocality":"Friedrichstadt","thoroughfare":"Prager Straße","fullThoroughfare":"Prager Straße","areasOfInterest":["Centrum Galerie"],"dependentLocalities":["Seevorstadt-Ost/GroßerGarten/Strehlen-Nordwest","Seevorstadt-Ost/Großer Garten/Strehlen-Nordwest"]},"country":"Germany","countryCode":"DE","poiCategory":"Cafe"},{"coordinate":{"latitude":51.040260581504135,"longitude":13.732845783233643},"displayMapRegion":{"southLatitude":51.0360362235794,"westLongitude":13.72439357287398,"northLatitude":51.0450193764206,"eastLongitude":13.738680427126019},"name":"Coffeeplace","formattedAddressLines":["Wiener Platz 4","01069 Dresden","Germany"],"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01069","subLocality":"Friedrichstadt","thoroughfare":"Wiener Platz","subThoroughfare":"4","fullThoroughfare":"Wiener Platz 4","dependentLocalities":["Seevorstadt-Ost/GroßerGarten/Strehlen-Nordwest","Seevorstadt-Ost/Großer Garten/Strehlen-Nordwest"]},"country":"Germany","countryCode":"DE","poiCategory":"Cafe"}]}`
	searchAutocomplete_SuccessResponse string = `{"results":[{"completionUrl":"/v1/search?q=Coffeeplace","displayLines":["Coffeeplace","Dresden, Saxony, Germany"]},{"completionUrl":"/v1/search?q=Coffeeplace%20Prager%20Stra%C3%9Fe%2C%2001069%20Dresden%2C%20Germany&metadata=ChEKCVN0YXJidWNrcxIECAAQCRIoCiZQcmFnZXIgU3RyYcOfZSwgMDEwNjkgRHJlc2RlbiwgR2VybWFueRgBKiEKEgkAAACgzIVJQBEAAABAzHgrQBCj6qHho8e6rE0Yrk1iEwoRU3RhcmJ1Y2tzIERyZXNkZW6C8QQRU3RhcmJ1Y2tzIERyZXNkZW6I8QQA2vEEFgkAAACA9b22QBEAAACgcN1TQDICZW7q8QQA","displayLines":["Coffeeplace","Prager Straße, 01069 Dresden, Germany"],"location":{"latitude":51.04530715942383,"longitude":13.735933303833008},"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01069","subLocality":"Friedrichstadt","thoroughfare":"Prager Straße","fullThoroughfare":"Prager Straße","areasOfInterest":["Centrum Galerie"],"dependentLocalities":["Seevorstadt-Ost/GroßerGarten/Strehlen-Nordwest","Seevorstadt-Ost/Großer Garten/Strehlen-Nordwest"]}},{"completionUrl":"/v1/search?q=Coffeeplace%20Wiener%20Platz%204%2C%2001069%20Dresden%2C%20Germany&metadata=ChEKCVN0YXJidWNrcxIECAAQCRIoCiZXaWVuZXIgUGxhdHogNCwgMDEwNjkgRHJlc2RlbiwgR2VybWFueRgBKiIKEgkAAABAJ4VJQBEAAACAN3crQBCQw9zw9ZCLrZwBGK5NYhMKEVN0YXJidWNrcyBEcmVzZGVugvEEEVN0YXJidWNrcyBEcmVzZGVuiPEEANrxBBYJAAAA4Fu9tkARAAAAYI%2BCVEAyAmVu6vEEAA%3D%3D","displayLines":["Coffeeplace","Wiener Platz 4, 01069 Dresden, Germany"],"location":{"latitude":51.040260314941406,"longitude":13.732845306396484},"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01069","subLocality":"Friedrichstadt","thoroughfare":"Wiener Platz","subThoroughfare":"4","fullThoroughfare":"Wiener Platz 4","dependentLocalities":["Seevorstadt-Ost/GroßerGarten/Strehlen-Nordwest","Seevorstadt-Ost/Großer Garten/Strehlen-Nordwest"]}},{"completionUrl":"/v1/search?q=Coffeeplace%20Altmarkt%207%2C%2001067%20Dresden%2C%20Germany&metadata=ChEKCVN0YXJidWNrcxIECAAQCRIkCiJBbHRtYXJrdCA3LCAwMTA2NyBEcmVzZGVuLCBHZXJtYW55GAEqIQoSCQAAAMBEhklAEQAAAED8eStAEMTXtL3EzI%2BFGBiuTWITChFTdGFyYnVja3MgRHJlc2RlboLxBBFTdGFyYnVja3MgRHJlc2RlbojxBADa8QQWCQAAAIBlvrZAEQAAAKBwGXFAMgJlburxBAA%3D","displayLines":["Coffeeplace","Altmarkt 7, 01067 Dresden, Germany"],"location":{"latitude":51.048973083496094,"longitude":13.738252639770508},"structuredAddress":{"administrativeArea":"Saxony","locality":"Dresden","postCode":"01067","subLocality":"Friedrichstadt","thoroughfare":"Altmarkt","subThoroughfare":"7","fullThoroughfare":"Altmarkt 7","dependentLocalities":["Altstadt Centre","InnereAltstadt"]}},{"completionUrl":"/v1/search?q=Coffeeplace","displayLines":["Coffeeplace","Dresden, TN, United States"]},{"completionUrl":"/v1/search?q=Coffeeplace%20Dresden","displayLines":["Coffeeplace Dresden","Search Nearby"]}]}`
)

type SearchTestSuite struct {
	suite.Suite
	testServer *httptest.Server
	mapsClient Client
}

func (s *SearchTestSuite) SetupSuite() {
	mux := http.NewServeMux()
	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(accessToken_SuccessResponse))
	})
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(search_SuccessResponse))
	})
	mux.HandleFunc("/searchAutocomplete", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(searchAutocomplete_SuccessResponse))
	})
	s.testServer = httptest.NewServer(mux)
	s.mapsClient = NewAppleMaps(s.testServer.Client(), "jwt", WithCustomURL(s.testServer.URL))
}

func (s *SearchTestSuite) TestSearch_Success() {
	var expected = &SearchResponse{}
	json.Unmarshal([]byte(search_SuccessResponse), expected)
	res, err := s.mapsClient.Search(context.Background(), "test query", WithUserLocation(NewLocation(1, 1)))
	s.NoError(err)
	s.Equal(expected, res)
}

func (s *SearchTestSuite) TestSearchAutocomplete_Success() {
	var expected = &SearchAutocompleteResult{}
	json.Unmarshal([]byte(searchAutocomplete_SuccessResponse), expected)
	res, err := s.mapsClient.SearchAutocomplete(context.Background(), "test query", WithUserLocation(NewLocation(1, 1)))
	s.NoError(err)
	s.Equal(expected, res)
}

func TestSearchTestSuite(t *testing.T) {
	suite.Run(t, new(SearchTestSuite))
}
