package applemaps

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestWithExcludePoiCategories(t *testing.T) {
	type test struct {
		parameters []Category
		expected   map[string][]string
	}
	tt := map[string]test{
		"No Category":         {[]Category{}, map[string][]string{"q": {"testing"}}},
		"Single Category":     {[]Category{Bakery}, map[string][]string{"excludePoiCategories": {"Bakery"}, "q": {"testing"}}},
		"Multiple Categories": {[]Category{Bakery, Bank}, map[string][]string{"excludePoiCategories": {"Bakery,Bank"}, "q": {"testing"}}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			vals := url.Values{}
			vals.Add("q", "testing")
			WithExcludePoiCategories(tc.parameters...)(vals)
			assert.EqualValues(t, tc.expected, vals)
		})
	}
}

func TestWithIncludePoiCategories(t *testing.T) {
	type test struct {
		parameters []Category
		expected   map[string][]string
	}
	tt := map[string]test{
		"No Category":         {[]Category{}, map[string][]string{"q": {"testing"}}},
		"Single Category":     {[]Category{Bakery}, map[string][]string{"includePoiCategories": {"Bakery"}, "q": {"testing"}}},
		"Multiple Categories": {[]Category{Bakery, Bank}, map[string][]string{"includePoiCategories": {"Bakery,Bank"}, "q": {"testing"}}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			vals := url.Values{}
			vals.Add("q", "testing")
			WithIncludePoiCategories(tc.parameters...)(vals)
			assert.EqualValues(t, tc.expected, vals)
		})
	}
}

func TestWithLimitToCountries(t *testing.T) {
	type test struct {
		parameters []string
		expected   map[string][]string
	}
	tt := map[string]test{
		"No Country":         {[]string{}, map[string][]string{"q": {"testing"}}},
		"Single Country":     {[]string{"US"}, map[string][]string{"limitToCountries": {"US"}, "q": {"testing"}}},
		"Multiple Countries": {[]string{"US", "CA"}, map[string][]string{"limitToCountries": {"US,CA"}, "q": {"testing"}}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			vals := url.Values{}
			vals.Add("q", "testing")
			WithLimitToCountries(tc.parameters...)(vals)
			assert.EqualValues(t, tc.expected, vals)
		})
	}
}

func TestWithResultTypeFilter(t *testing.T) {
	type test struct {
		parameters []string
		expected   map[string][]string
	}
	tt := map[string]test{
		"No Filter":        {[]string{}, map[string][]string{"q": {"testing"}}},
		"Single Filter":    {[]string{"Poi"}, map[string][]string{"resultTypeFilter": {"Poi"}, "q": {"testing"}}},
		"Multiple Filters": {[]string{"Poi", "Address"}, map[string][]string{"resultTypeFilter": {"Poi,Address"}, "q": {"testing"}}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			vals := url.Values{}
			vals.Add("q", "testing")
			WithResultTypeFilter(tc.parameters...)(vals)
			assert.EqualValues(t, tc.expected, vals)
		})
	}
}

func TestWithLanguage(t *testing.T) {
	type test struct {
		parameter language.Tag
		expected  map[string][]string
	}
	tt := map[string]test{
		"Empty Language": {language.Tag{}, map[string][]string{"q": {"testing"}}},
		"en-US":          {language.AmericanEnglish, map[string][]string{"lang": {"en-US"}, "q": {"testing"}}},
		"de":             {language.German, map[string][]string{"lang": {"de"}, "q": {"testing"}}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			vals := url.Values{}
			vals.Add("q", "testing")
			WithLanguage(tc.parameter)(vals)
			assert.EqualValues(t, tc.expected, vals)
		})
	}
}

func TestWithArrivalDate(t *testing.T) {
	vals := url.Values{}
	vals.Add("q", "testing")
	WithArrivalDate(time.Date(2023, 04, 15, 16, 42, 0, 0, time.UTC))(vals)

	expected := url.Values{}
	expected.Add("q", "testing")
	expected.Add("arrivalDate", "2023-04-15T16:42:00Z")

	assert.Equal(t, expected, vals)
}

func TestWithDepartureDate(t *testing.T) {
	vals := url.Values{}
	vals.Add("q", "testing")
	WithDepartureDate(time.Date(2023, 04, 15, 16, 42, 0, 0, time.UTC))(vals)

	expected := url.Values{}
	expected.Add("q", "testing")
	expected.Add("departureDate", "2023-04-15T16:42:00Z")

	assert.Equal(t, expected, vals)
}

func TestWithRequestsAlternateRoutes(t *testing.T) {
	vals := url.Values{}
	vals.Add("q", "testing")
	WithRequestsAlternateRoutes()(vals)

	expected := url.Values{}
	expected.Add("q", "testing")
	expected.Add("requestsAlternateRoutes", "true")

	assert.Equal(t, expected, vals)
}

func TestWithTransportType(t *testing.T) {
	vals := url.Values{}
	vals.Add("q", "testing")
	WithTransportType("Automobile")(vals)

	expected := url.Values{}
	expected.Add("q", "testing")
	expected.Add("transportType", "Automobile")

	assert.Equal(t, expected, vals)
}

func TestWithSearchLocation(t *testing.T) {
	vals := url.Values{}
	vals.Add("q", "testing")
	WithSearchLocation(NewLocation(15.011276, 13.9955))(vals)

	expected := url.Values{}
	expected.Add("q", "testing")
	expected.Add("searchLocation", "15.011276,13.9955")

	assert.Equal(t, expected, vals)
}

func TestWithAvoid(t *testing.T) {
	type test struct {
		parameter []Avoid
		expected  map[string][]string
	}
	tt := map[string]test{
		"No Avoids":      {[]Avoid{}, map[string][]string{"q": {"testing"}}},
		"Single Avoid":   {[]Avoid{Tolls}, map[string][]string{"avoid": {"Tolls"}, "q": {"testing"}}},
		"Multiple Avoid": {[]Avoid{Tolls, "Something"}, map[string][]string{"avoid": {"Tolls,Something"}, "q": {"testing"}}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			vals := url.Values{}
			vals.Add("q", "testing")
			WithAvoid(tc.parameter...)(vals)
			assert.EqualValues(t, tc.expected, vals)
		})
	}
}

func TestWithSearchRegion(t *testing.T) {
	vals := url.Values{}
	vals.Add("q", "testing")
	WithSearchRegion(NewRegion(15.011276, 13.9955, 11.1, 17.77))(vals)

	expected := url.Values{}
	expected.Add("q", "testing")
	expected.Add("searchRegion", "15.011276,13.9955,11.1,17.77")

	assert.Equal(t, expected, vals)
}

func TestWithUserLocation(t *testing.T) {
	vals := url.Values{}
	vals.Add("q", "testing")
	WithUserLocation(NewLocation(15.011276, -13.9955))(vals)

	expected := url.Values{}
	expected.Add("q", "testing")
	expected.Add("userLocation", "15.011276,-13.9955")

	assert.Equal(t, expected, vals)
}
