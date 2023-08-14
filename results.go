package applemaps

type SearchResponse struct {
	DisplayMapRegion MapRegion `json:"displayMapRegion"`
	Results          []Place   `json:"results"`
}

type SearchAutocompleteResult struct {
	Results []AutocompleteResult `json:"results"`
}

type AutocompleteResult struct {
	CompletionUrl     string            `json:"completionUrl"`
	DisplayLines      []string          `json:"displayLines"`
	Location          Location          `json:"location"`
	StructuredAddress StructuredAddress `json:"structuredAddress"`
}

type DirectionsResponse struct {
	Destination Place        `json:"destination"`
	Origin      Place        `json:"origin"`
	Routes      []Route      `json:"routes"`
	StepPaths   [][]Location `json:"stepPaths"`
	Steps       []Step       `json:"steps"`
}

type Route struct {
	DistanceMeters  int    `json:"distanceMeters"`
	DurationSeconds int    `json:"durationSeconds"`
	HasTolls        bool   `json:"hasTolls"`
	Name            string `json:"name"`
	StepIndexes     []int  `json:"stepIndexes"`
	TransportType   string `json:"transportType"`
}

type EtaResponse struct {
	Etas []Eta `json:"etas"`
}

type Eta struct {
	Destination               Location `json:"destination"`
	DistanceMeters            int      `json:"distanceMeters"`
	ExpectedTravelTimeSeconds int      `json:"expectedTravelTimeSeconds"`
	StaticTravelTimeSeconds   int      `json:"staticTravelTimeSeconds"`
	TransportType             string   `json:"transportType"`
}

type Step struct {
	DistanceMeters  int    `json:"distanceMeters"`
	DurationSeconds int    `json:"durationSeconds"`
	Instructions    string `json:"instructions"`
	StepPathIndex   int    `json:"stepPathIndex"`
	TransportType   string `json:"transportType"`
}
