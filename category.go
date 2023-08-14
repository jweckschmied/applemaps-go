package applemaps

type Category string

func (c Category) String() string {
	return string(c)
}

const (
	// Airport an airport
	Airport Category = "Airport"
	// AirportGate a specific gate at an airport
	AirportGate Category = "AirportGate"
	// AirportTerminal a specific named terminal at an airport
	AirportTerminal Category = "AirportTerminal"
	// AmusementPark An amusement parka
	AmusementPark Category = "AmusementPark"
	// ATM an Automated Teller Machine
	ATM Category = "ATM"
	// Aquarium an Aquarium
	Aquarium Category = "Aquarium"
	// Bakery a bakery
	Bakery Category = "Bakery"
	// Bank a bank
	Bank Category = "Bank"
	// Beach a beach
	Beach Category = "Beach"
	// Brewery a brewery
	Brewery Category = "Brewery"
	// Cafe a Cafe
	Cafe Category = "Cafe"
	// Campground a campground
	Campground Category = "Campground"
	// CarRental a Car Rental Location
	CarRental Category = "CarRental"
	// EVCharger an Electric Vehicle (EV) Charger
	EVCharger Category = "EVCharger"
	// FireStation a fire station
	FireStation Category = "FireStation"
	// FitnessCenter a fitness center
	FitnessCenter Category = "FitnessCenter"
	// FoodMarket a food market
	FoodMarket Category = "FoodMarket"
	// GasStation a gas station
	GasStation Category = "GasStation"
	// Hospital a hospital
	Hospital Category = "Hospital"
	// Hotel a hotel
	Hotel Category = "Hotel"
	// Laundry a laundry
	Laundry Category = "Laundry"
	// Library a library
	Library Category = "Library"
	// Marina a marina
	Marina Category = "Marina"
	// MovieTheater a movie theater
	MovieTheater Category = "MovieTheater"
	// Museum a museum
	Museum Category = "Museum"
	// NationalPark a national park
	NationalPark Category = "NationalPark"
	// Nightlife a nightlife venue
	Nightlife Category = "Nightlife"
	// Park a park
	Park Category = "Park"
	// Parking a parking location for an automobile
	Parking Category = "Parking"
	// Pharmacy a pharmacy
	Pharmacy Category = "Pharmacy"
	// Playground a playground
	Playground Category = "Playground"
	// Police a police station
	Police Category = "Police"
	// PostOffice a post office
	PostOffice Category = "PostOffice"
	// PublicTransport a public transportation station
	PublicTransport Category = "PublicTransport"
	// ReligiousSite a religious site
	ReligiousSite Category = "ReligiousSite"
	// Restaurant a restaurant
	Restaurant Category = "Restaurant"
	// Restroom a restroom
	Restroom Category = "Restroom"
	// School a school
	School Category = "School"
	// Stadium a stadium
	Stadium Category = "Stadium"
	// Store a store
	Store Category = "Store"
	// Theater a theater
	Theater Category = "Theater"
	// University a university
	University Category = "University"
	// Winery a winery
	Winery Category = "Winery"
	// Zoo a zoo
	Zoo Category = "Zoo"
)
