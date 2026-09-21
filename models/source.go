package models

type SourceProperty struct{
	ID string `json:"id"`
	Feed int `json:"feed"`
	Country string `json:"country"`
	CountryCode string `json:"country_code"`
	State string `json:"state"`
	StateAbbreviation string `json:"state_abbr"`
	City string `json:"city"`
	Display string `json:"display"`
	LocationId string `json:"location_id"`
	PropertyName string `json:"property_name"`
	PropertySlug string `json:"property_slug"`
	PropertyTypeCategory string `json:"property_type_category"`
	UsdPrice float64 `json:"usd_price"`
	Occupancy int64 `json:"occupancy"`
	BedroomCount int64 `json:"bedroom_count"`
	BathroomCount int64 `json:"bathroom_count"`
	NumberOfReview int64 `json:"number_of_review"`
	ReviewScoreGeneral float64 `json:"review_score_general"`
	StarRating int64 `json:"star_rating"`
	AmenityCategories []string `json:"amenity_categories"`
	LonLat LonLat `json:"lonlat"`
}

type LonLat struct{
	Coordinates []float64
}




