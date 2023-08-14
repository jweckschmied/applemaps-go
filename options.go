package applemaps

import (
	"net/url"
	"strings"
	"time"

	"golang.org/x/text/language"
)

type RequestOption func(v url.Values)

// WithExcludePoiCategories provides an option to add a slice of strings that describes the points of interest to exclude from the search results. For example, excludePoiCategories=Restaurant,Cafe.
// See Category type for a complete list of possible values.
func WithExcludePoiCategories(categories ...Category) RequestOption {
	return func(v url.Values) {
		if len(categories) == 0 {
			return
		}
		v.Add("excludePoiCategories", queryParameterString(categories))
	}
}

// WithIncludePoiCategories provides an option to add a slice of Categories that describes the points of interest to include in the search results. For example, includePoiCategories=Restaurant,Cafe.
// See Category type for a complete list of possible values.
func WithIncludePoiCategories(categories ...Category) RequestOption {
	return func(v url.Values) {
		if len(categories) == 0 {
			return
		}
		v.Add("includePoiCategories", queryParameterString(categories))
	}
}

// WithLimitToCountries provides an option to add a slice of ISO ALPHA-2 codes of the countries to limit the results to. For example, limitToCountries=US,CA limits the search to the United States and Canada.
// If you specify two or more countries, the results reflect the best available results for some or all of the countries rather than everything related to the query for those countries.
func WithLimitToCountries(countries ...string) RequestOption {
	return func(v url.Values) {
		if len(countries) == 0 {
			return
		}
		v.Add("limitToCountries", queryParameterString(countries))
	}
}

// WithResultTypeFilter provides an option to add a slice of strings that describes the kind of result types to include in the response. For example, resultTypeFilter=Poi.
// Possible Values: Poi, Address
func WithResultTypeFilter(filters ...string) RequestOption {
	return func(v url.Values) {
		if len(filters) == 0 {
			return
		}
		v.Add("resultTypeFilter", queryParameterString(filters))
	}
}

// WithLanguage provides an option to add the language the server should use when returning the response, specified using a BCP 47 language code. For example, for English use lang=en-US. Defaults to en-US.
func WithLanguage(lang language.Tag) RequestOption {
	return func(v url.Values) {
		if lang.IsRoot() {
			return
		}
		v.Add("lang", lang.String())
	}
}

// WithArrivalDate provides an option to add the date and time to arrive at the destination.
// You can specify only arrivalDate or departureDate. If you don’t specify either option, the departureDate defaults to now, which the server interprets as the current time.
func WithArrivalDate(arrival time.Time) RequestOption {
	return func(v url.Values) {
		v.Add("arrivalDate", arrival.Format(time.RFC3339))
	}
}

// WithDepartureDate provides an option to add the date and time to depart from the origin.
// You can only specify arrivalDate or departureDate. If you don’t specify either option, the departureDate defaults to now, which the server interprets as the current time.
func WithDepartureDate(departure time.Time) RequestOption {
	return func(v url.Values) {
		v.Add("departureDate", departure.Format(time.RFC3339))
	}
}

// WithRequestsAlternateRoutes provides an option for the server to return additional routes, when available. For example, requestsAlternateRoutes=true.
// Default: false
func WithRequestsAlternateRoutes() RequestOption {
	return func(v url.Values) {
		v.Add("requestsAlternateRoutes", "true")
	}
}

// WithTransportType provides an option to set the mode of transportation the server returns directions for.
// Default: Automobile
// Possible Values: Automobile, Walking
func WithTransportType(transportType string) RequestOption {
	return func(v url.Values) {
		v.Add("transportType", transportType)
	}
}

// WithSearchLocation provides an option to set a hint for the query input for origin or destination.
// If you don’t provide a searchLocation, the server uses userLocation and searchLocation as fallback hints.
func WithSearchLocation(location Location) RequestOption {
	return func(v url.Values) {
		v.Add("searchLocation", location.String())
	}
}

// WithAvoid provides an option to add a slice of the features to avoid when calculating direction routes. For example, avoid=Tolls.
// See Avoid type for a complete list of possible values.
func WithAvoid(avoid ...Avoid) RequestOption {
	return func(v url.Values) {
		if len(avoid) == 0 {
			return
		}
		v.Add("avoid", queryParameterString(avoid))
	}
}

// WithSearchRegion provides an option to set a region the app defines as a hint for the query input for origin or destination.
// If you don’t provide a searchLocation, the server uses userLocation and searchRegion as fallback hints.
func WithSearchRegion(region MapRegion) RequestOption {
	return func(v url.Values) {
		v.Add("searchRegion", region.String())
	}
}

// WithUserLocation provides an option to set the location of the user.
// If you don’t provide a searchLocation, the server uses userLocation and searchRegion as fallback hints.
func WithUserLocation(location Location) RequestOption {
	return func(v url.Values) {
		v.Add("userLocation", location.String())
	}
}

// queryParameterString formats slices of different data types to a string representation that can be used for query parameters.
func queryParameterString[P []Category | []Location | []Avoid | []string](params P) string {
	var str = make([]string, len(params))
	switch v := any(params).(type) {
	case []Category:
		for i, elem := range v {
			str[i] = elem.String()
		}
		return strings.Join(str[:], ",")
	case []Location:
		for i, elem := range v {
			str[i] = elem.String()
		}
		return strings.Join(str[:], "|")
	case []Avoid:
		for i, elem := range v {
			str[i] = elem.String()
		}
		return strings.Join(str[:], ",")
	case []string:
		return strings.Join(v[:], ",")
	default:
		return ""
	}
}
