package models

type ResponseProperty struct {
	ID        string `json:"ID"`
	Feed      uint64 `json:"Feed"`
	Published bool   `json:"Published"`
	GeoInfo   GeoInfo
	Property  PropertyDetails
}

type GeoInfo struct {
	Breadcrumbs []string `json:"Breadcrumbs"`
	City        string   `json:"City"`
	Country     string   `json:"Country"`
	CountryCode string   `json:"CountryCode"`
	Name        string   `json:"Name"`
	LocationID  string   `json:"LocationID"`
	Lat         float64  `json:"Lat"`
	Lon         float64  `json:"Lon"`
	State       string   `json:"State"`
	StateAbbr   string   `json:"StateAbbr"`
}

type PropertyDetails struct {
	Amenities    []string `json:"Amenities"`
	Name         string   `json:"name"`
	Slug         string   `json:"Slug"`
	PropertyType string   `json:"PropertyType"`
	Price        float64  `json:"Price"`
	ReviewScore  float64  `json:"ReviewScore"`
	StarRating   uint64   `json:"StarRating"`
	Counts       Counts   `json:"Counts"`
	Image        Image    `json:"Image"`
}

type Counts struct {
	Bathroom  uint64 `json:"Bathroom"`
	Bedroom   uint64 `json:"Bedroom"`
	Reviews   uint64 `json:"Reviews"`
	Occupancy uint64 `json:"Occupancy"`
}

type Image struct {
	Count  uint64   `json:"Count"`
	Images []string `json:"Images"`
}
