package services

import (
	"encoding/json"
	"errors"
	"os"
	"rental-property-api/models"
	"strings"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

var InMemoryProperties []models.ResponseProperty

func LoadData() {

	jsonPath, err := config.String("dataPath")
	if err != nil || jsonPath == "" {
		jsonPath = "data/rental_properties.json"
		logs.Warning("dataPath not found in config, using default fallback: %s", jsonPath)
	}

	data, err := os.ReadFile(jsonPath)

	if err != nil {
		logs.Error("Something error to read file", err)
		return
	}

	logs.Info("File data len", len(data))

	var properties []models.SourceProperty

	err = json.Unmarshal(data, &properties)

	if err != nil {
		logs.Error("Error parsing JSON:", err)
		return
	}
	TransformData(properties)

}

func TransformData(properties []models.SourceProperty) {

	for index, property := range properties {

		var Categories []models.CategoryDetail
		var breadcrumbs []string

		if property.Categories != "" && property.Categories != "[]" {
			err := json.Unmarshal([]byte(property.Categories), &Categories)
			if err != nil {
				logs.Error("Index:%d of Can't parse categories: %s \n", index, err)
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
			ID:        property.ID,
			Feed:      property.Feed,
			Published: property.Published,
			GeoInfo: models.GeoInfo{
				Breadcrumbs: breadcrumbs,
				City:        property.City,
				Country:     property.Country,
				CountryCode: property.CountryCode,
				Name:        property.Display,
				LocationID:  property.LocationID,
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
					Count:  int64(len(property.Images)),
					Images: property.Images,
				},
			},
		}
		InMemoryProperties = append(InMemoryProperties, transformedItem)
	}

}

func FilterProperties(minPrice, maxPrice float64, minStarRating int64, minReviewScore float64, minReviews int64, published string, feed, minBedroom int64, propertyType string, amenities string, limit int64) []models.ResponseProperty {

	var filtered []models.ResponseProperty

	for _, p := range InMemoryProperties {
		var matchFound bool

		if minPrice >= 0.0 && p.Property.Price < minPrice {
			continue
		}

		if maxPrice >= 0.0 && p.Property.Price > maxPrice {
			continue
		}

		if minStarRating >= 0 && p.Property.StarRating < minStarRating {
			continue
		}
		if minReviewScore >= 0 && p.Property.ReviewScore < minReviewScore {
			continue
		}
		if minReviews >= 0 && p.Property.Counts.Reviews < minReviews {
			continue
		}

		if feed >= 0 && p.Feed != feed {
			continue
		}

		if minBedroom >= 0 && p.Property.Counts.Bedroom < minBedroom {
			continue
		}

		if propertyType != "" && p.Property.PropertyType != propertyType {
			continue
		}

		if published != "" {
			isPublished := published == "true"
			if isPublished != p.Published {
				continue
			}
		}

		if amenities != "" {
			searchAmenities := strings.Split(amenities, ",")
			for _, searchAmenity := range searchAmenities {
				for _, propertyAmenity := range p.Property.Amenities {
					if propertyAmenity == searchAmenity {
						matchFound = true
						break
					}
				}
				if matchFound == true {
					break
				}
			}

			if !matchFound {
				continue
			}
		}

		filtered = append(filtered, p)

	}

	if filtered == nil {
		filtered = make([]models.ResponseProperty, 0)
	}

	if limit > 0 && limit < int64(len(filtered)) {
		filtered = filtered[:limit]
	}
	return filtered
}

func GetPropertyByID(id string) (models.ResponseProperty, error) {
	for _, property := range InMemoryProperties {
		if property.ID == id {
			return property, nil
		}
	}

	return models.ResponseProperty{}, errors.New("Property not found")
}
