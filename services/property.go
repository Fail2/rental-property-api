package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"rental-property-api/models"
	"runtime"
)

var InMemoryProperties []models.ResponseProperty

func LoadAndTransformData() {
	// First find the path of the file

	_, filename, _, _ := runtime.Caller(0)

	// Service direcotory root path
	serviceDir := filepath.Dir(filename)

	jsonPath := filepath.Join(serviceDir, "..", "data", "rental_properties.json")

	// Read the file
	data, err := os.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("Something error to read file", err)
		return
	}

	fmt.Println("File data len", len(data))

	var properties []models.SourceProperty

	err = json.Unmarshal(data, &properties)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for index, property := range properties {

		var Categories []models.CategoryDetail
		var breadcrumbs []string

		if property.Categories != "" && property.Categories != "[]" {
			err = json.Unmarshal([]byte(property.Categories), &Categories)
			if err != nil {
				fmt.Printf("Index:%d of Can't parse categories: %s \n", index, err)
				return
			} else {
				for _, category := range Categories {
					breadcrumbs = append(breadcrumbs, category.Name)
				}
			}
		}

		var lat, lon float64

		if len(property.LonLat.Coordinates) >= 2 {
			lon = property.LonLat.Coordinates[0]
			lat = property.LonLat.Coordinates[1]
		}
		transformedItem := models.ResponseProperty{
			ID:        property.Id,
			Feed:      property.Feed,
			Published: property.Published,
			GeoInfo: models.GeoInfo{
				Breadcrumbs: breadcrumbs,
				City:        property.City,
				Country:     property.Country,
				CountryCode: property.CountryCode,
				Name:        property.Display,
				LocationID:  property.LocationId,
				Lat:         lat,
				Lon:         lon,
				State:       property.State,
				StateAbbr:   property.StateAbbr,
			},
			Property: models.PropertyDetails{
				Amenities:    property.AmenityCategories,
				Name:         property.PropertyName,
				Slug:         property.PropertySlug,
				PropertyType: property.PropertyTypeCategory,
				Price:        property.UsdPrice,
				ReviewScore:  property.ReviewScoreGeneral,
				StarRating:   property.StarRating,
				Counts: models.Counts{
					Bathroom:  property.BathroomCount,
					Bedroom:   property.BedroomCount,
					Reviews:   property.NumberOfReview,
					Occupancy: property.Occupancy,
				},
				Image: models.Image{
					Count:  uint64(len(property.Images)),
					Images: property.Images,
				},
			},
		}
		InMemoryProperties = append(InMemoryProperties, transformedItem)
	}

	fmt.Printf("Slice of response properties: %+v\n", InMemoryProperties)
}
