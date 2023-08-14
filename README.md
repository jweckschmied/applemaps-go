# A Go client library for the Apple Maps Server API

![header_img](https://github.com/jweckschmied/applemaps-go/assets/33632397/1acf73bd-5053-404b-89f8-3c31b0822515)

To learn about the Apple Maps Server API and all its methods, check the official documentation here: 
https://developer.apple.com/documentation/applemapsserverapi

## Installation
```go 
go get github.com/jweckschmied/applemaps-go
```

## Getting Started
An Apple Developer account is required to use the Apple Maps Server API. If you have a Developer Account, you can
learn more about how to get a Maps Identifier, Private Key and Auth Token in the Apple MapKit JS Documentation: 
https://developer.apple.com/documentation/mapkitjs/creating_a_maps_identifier_and_a_private_key

The only information you need to provide to use this package is the JWT Auth Token.
### Generating the JWT
Since the JWT will expire at some point, it needs to be regenerated on a regular basis.
You can do this however you like, for example using one of the available golang JWT packages.

Details on the JWT creation can be found here: https://developer.apple.com/documentation/mapkitjs/creating_and_using_tokens_with_mapkit_js

Use the `SetAuthToken()` method to set a new token for an already existing client.

## Usage Example
```go
import (
    "context"
    "net/http"
	
    "github.com/jweckschmied/applemaps-go"
)

func main() {
    ctx := context.Background()
    httpClient := http.DefaultClient
    client := applemaps.NewAppleMaps(httpClient, "<your-auth-token>")
    result, err := client.Search(
        ctx,
        "Tour Eiffel",
        applemaps.WithLanguage(language.French),
        applemaps.WithResultTypeFilter("Poi"),
        applemaps.WithUserLocation(applemaps.NewLocation(48.858093, 2.294694)),
    )
}
```

## Request Options
The following options are available for the different methods provided by the API. 
Please check the docs for details on which parameters can be used for each API method.
```
WithExcludePoiCategories
WithIncludePoiCategories
WithLimitToCountries
WithResultTypeFilter
WithLanguage
WithArrivalDate
WithDepartureDate
WithRequestsAlternateRoutes
WithTransportType
WithSearchLocation
WithAvoid
WithSearchRegion
WithUserLocation
```

For example, if you wanted to set the `userLocation`, along with `includePoiCategories=Bank,Bakery`, the function call
should look like this:
```go
client.Search(
    "Search Query",
    applemaps.WithUserLocation(applemaps.NewLocation(51.08097475066115, 13.76077443357895)),
    applemaps.WithIncludePoiCategories(applemaps.Bank, applemaps.Bakery),
)
```
