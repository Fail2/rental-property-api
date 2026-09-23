package services

import (
	"rental-property-api/models"
	"testing"
)

func init() {
	InMemoryProperties = []models.ResponseProperty{
		{
			ID:        "BC-1001",
			Feed:      11,
			Published: true,
			GeoInfo: models.GeoInfo{
				Breadcrumbs: []string{"Spain", "Barcelone"},
				City:        "Barcelone",
			},
			Property: models.PropertyDetails{
				Amenities:    []string{"Internet", "Parking"},
				PropertyType: "Apartment",
				Price:        120.0,
				ReviewScore:  8.5,
				StarRating:   4,
				Counts: models.Counts{
					Bedroom: 2,
					Reviews: 25,
				},
			},
		},
		{
			ID:        "BC-1002",
			Feed:      12,
			Published: false,
			GeoInfo: models.GeoInfo{
				Breadcrumbs: []string{"Japan", "Tokyo"},
				City:        "Tokyo",
			},
			Property: models.PropertyDetails{
				Amenities:    []string{"Gym"},
				PropertyType: "Hotel",
				Price:        200.0,
				ReviewScore:  6.0,
				StarRating:   3,
				Counts: models.Counts{
					Bedroom: 1,
					Reviews: 5,
				},
			},
		},
	}
}

func TestFilterProperties(t *testing.T) {
	tests := []struct {
		name          string
		minPrice      float64
		maxPrice      float64
		minStar       int64
		minScore      float64
		minReviews    int64
		published     string
		feed          int64
		minBedroom    int64
		propType      string
		amenities     string
		limit         int64
		expectedCount int
	}{
		{
			name:     "Filter - AND Case (Feed + Published)",
			minPrice: -1, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "true", feed: 11, minBedroom: -1, propType: "", amenities: "", limit: 0,
			expectedCount: 1,
		},
		{
			name:     "Filter - AND Case (Price Range match)",
			minPrice: 100, maxPrice: 150, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "", amenities: "", limit: 0,
			expectedCount: 1,
		},
		{
			name:     "Filter - Amenities OR Case (Matches at least one)",
			minPrice: -1, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "", amenities: "Internet,Pool", limit: 0,
			expectedCount: 1,
		},
		{
			name:     "Filter - Combined Case (AND filters + amenities together)",
			minPrice: 50, maxPrice: 300, minStar: 4, minScore: 8.0, minReviews: 10,
			published: "true", feed: 11, minBedroom: 2, propType: "Apartment", amenities: "Parking", limit: 0,
			expectedCount: 1,
		},
		{
			name:     "Filter - Empty result Case (No matches return empty slice)",
			minPrice: 500, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "", amenities: "", limit: 0,
			expectedCount: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := FilterProperties(
				test.minPrice, test.maxPrice, test.minStar, test.minScore, test.minReviews, test.published, test.feed, test.minBedroom, test.propType, test.amenities, test.limit,
			)

			if len(res) != test.expectedCount {
				t.Errorf("Expected count %d, got %d", test.expectedCount, len(res))
			}
		})
	}
}

func TestGetPropertyByID(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		shouldError bool
	}{
		{
			name:        "Get by ID - Found item",
			id:          "BC-1001",
			shouldError: false,
		},
		{
			name:        "Get by ID - Not found item",
			id:          "INVALID-ID",
			shouldError: true,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			_, err := GetPropertyByID(test.id)

			if (err != nil) != test.shouldError {
				t.Errorf("Unexpected error presence: %v, expected error state: %v", err, test.shouldError)
			}
		})
	}
}
