package services

import (
	"rental-property-api/models"
	"testing"
)

func init() {
	InMemoryProperties = []models.SourceProperty{
		{
			ID:                   "BC-1001",
			Feed:                 11,
			Published:            true,
			Country:              "Spain",
			City:                 "Barcelone",
			AmenityCategories:    []string{"Internet", "Parking"},
			PropertyTypeCategory: "Apartment",
			UsdPrice:             120.0,
			ReviewScoreGeneral:   8.5,
			StarRating:           4,
			BedroomCount:         2,
			NumberOfReview:       25,
		},
		{
			ID:                   "BC-1002",
			Feed:                 12,
			Published:            false,
			Country:              "Japan",
			City:                 "Tokyo",
			AmenityCategories:    []string{"Gym"},
			PropertyTypeCategory: "Hotel",
			UsdPrice:             200.0,
			ReviewScoreGeneral:   6.0,
			StarRating:           3,
			BedroomCount:         1,
			NumberOfReview:       5,
		},
		{
			ID:                   "BC-1003",
			Feed:                 22,
			Published:            true,
			Country:              "USA",
			City:                 "Miami",
			AmenityCategories:    []string{"Internet", "Kitchen"},
			PropertyTypeCategory: "Villa",
			UsdPrice:             450.0,
			ReviewScoreGeneral:   9.2,
			StarRating:           5,
			BedroomCount:         4,
			NumberOfReview:       12,
		},
		{
			ID:                   "BC-1004",
			Feed:                 24,
			Published:            true,
			Country:              "Thailand",
			City:                 "Phuket",
			AmenityCategories:    []string{"Pool", "Beachfront"},
			PropertyTypeCategory: "Resort",
			UsdPrice:             80.0,
			ReviewScoreGeneral:   4.5,
			StarRating:           2,
			BedroomCount:         1,
			NumberOfReview:       50,
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
			name:     "Filter - AND Case (Feed 11 + Published true)",
			minPrice: -1, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "true", feed: 11, minBedroom: -1, propType: "", amenities: "", limit: -1,
			expectedCount: 1,
		},
		{
			name:     "Filter - Price Range Edge Case (Budget match)",
			minPrice: 50.0, maxPrice: 150.0, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "", amenities: "", limit: -1,
			expectedCount: 2,
		},
		{
			name:     "Filter - Amenities OR Case (Internet OR Pool)",
			minPrice: -1, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "", amenities: "Internet,Pool", limit: -1,
			expectedCount: 3,
		},
		{
			name:     "Filter - Strict Star and Review Score Constraints",
			minPrice: -1, maxPrice: -1, minStar: 4, minScore: 8.0, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "", amenities: "", limit: -1,
			expectedCount: 2,
		},
		{
			name:     "Filter - Bedroom Count Constraint (Minimum 2)",
			minPrice: -1, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: 2, propType: "", amenities: "", limit: -1,
			expectedCount: 2,
		},
		{
			name:     "Filter - Property Type Exact Case-Sensitive Match",
			minPrice: -1, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "Hotel", amenities: "", limit: -1,
			expectedCount: 1,
		},
		{
			name:     "Filter - Combined Case (Complex AND + Amenities OR)",
			minPrice: 100.0, maxPrice: 500.0, minStar: 4, minScore: 8.0, minReviews: 10,
			published: "true", feed: 22, minBedroom: 3, propType: "Villa", amenities: "Kitchen,Internet", limit: -1,
			expectedCount: 1,
		},
		{
			name:     "Filter - Limit Validation Cap",
			minPrice: -1, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "", amenities: "", limit: 2,
			expectedCount: 2,
		},
		{
			name:     "Filter - Empty result Case (Zero Matches)",
			minPrice: 999.0, maxPrice: -1, minStar: -1, minScore: -1, minReviews: -1,
			published: "", feed: -1, minBedroom: -1, propType: "", amenities: "", limit: -1,
			expectedCount: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, _ := FilterProperties(
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
