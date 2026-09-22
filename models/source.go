package models

type SourceProperty struct {
	ID                   string   `json:"id"`
	Feed                 uint64      `json:"feed"`
	Country              string   `json:"country"`
	CountryCode          string   `json:"country_code"`
	State                string   `json:"state"`
	StateAbbr    string   `json:"state_abbr"`
	City                 string   `json:"city"`
	Display              string   `json:"display"`
	LocationID           string   `json:"location_id"`
	PropertyName         string   `json:"property_name"`
	PropertySlug         string   `json:"property_slug"`
	PropertyTypeCategory string   `json:"property_type_category"`
	UsdPrice             float64  `json:"usd_price"`
	Occupancy            uint64    `json:"occupancy"`
	BedroomCount         uint64    `json:"bedroom_count"`
	BathroomCount        uint64    `json:"bathroom_count"`
	NumberOfReview       uint64    `json:"number_of_review"`
	ReviewScoreGeneral   float64  `json:"review_score_general"`
	StarRating           uint64    `json:"star_rating"`
	AmenityCategories    []string `json:"amenity_categories"`
	LonLat               LonLat   `json:"lonlat"`
	Categories           string   `json:"categories"`
	Published            bool     `json:"published"`
	Images               []string `json:"images"`
}

type LonLat struct {
	Coordinates []float64
}

type CategoryDetail struct {
	LocationID string   `json:"LocationID"`
	Name       string   `json:"Name"`
	Type       string   `json:"Type"`
	Slug       string   `json:"Slug"`
	Display    []string `json:"Display"`
}
