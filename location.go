package applemaps

import (
	"fmt"
	"strconv"
)

type (
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	MapRegion struct {
		NorthLatitude float64 `json:"northLatitude"`
		EastLongitude float64 `json:"eastLongitude"`
		SouthLatitude float64 `json:"southLatitude"`
		WestLongitude float64 `json:"westLongitude"`
	}

	Place struct {
		Country               string            `json:"country"`
		CountryCode           string            `json:"countryCode"`
		DisplayMapRegion      MapRegion         `json:"displayMapRegion"`
		FormattedAddressLines []string          `json:"formattedAddressLine"`
		Name                  string            `json:"name"`
		Coordinate            Location          `json:"coordinate"`
		StructuredAddress     StructuredAddress `json:"structuredAddress"`
	}

	StructuredAddress struct {
		AdministrativeArea     string   `json:"administrativeArea"`
		AdministrativeAreaCode string   `json:"administrativeAreaCode"`
		AreasOfInterest        []string `json:"areasOfInterest"`
		DependentLocalities    []string `json:"dependentLocalities"`
		FullThoroughfare       string   `json:"fullThoroughfare"`
		Locality               string   `json:"locality"`
		PostCode               string   `json:"postCode"`
		SubLocality            string   `json:"subLocality"`
		// SubThoroughfare The short code for the state or area.
		SubThoroughfare string `json:"subThoroughfare"`
		// Thoroughfare The state or province of the place
		Thoroughfare string `json:"thoroughfare"`
	}
)

// NewLocation creates a new coordinate object
func NewLocation(lat, lng float64) Location {
	return Location{Latitude: lat, Longitude: lng}
}

// NewRegion creates a new map region object
func NewRegion(northLat, eastLon, southLat, westLon float64) MapRegion {
	return MapRegion{
		NorthLatitude: northLat,
		EastLongitude: eastLon,
		SouthLatitude: southLat,
		WestLongitude: westLon,
	}
}

func (l Location) String() string {
	return fmt.Sprintf(
		"%s,%s",
		strconv.FormatFloat(l.Latitude, 'f', -1, 64),
		strconv.FormatFloat(l.Longitude, 'f', -1, 64),
	)
}

func (r MapRegion) String() string {
	return fmt.Sprintf("%s,%s,%s,%s",
		strconv.FormatFloat(r.NorthLatitude, 'f', -1, 64),
		strconv.FormatFloat(r.EastLongitude, 'f', -1, 64),
		strconv.FormatFloat(r.SouthLatitude, 'f', -1, 64),
		strconv.FormatFloat(r.WestLongitude, 'f', -1, 64),
	)
}
