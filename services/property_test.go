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
			Country:              "Spain",
			CountryCode:          "ES",
			State:                "Catalonia",
			StateAbbr:            "CT",
			City:                 "Barcelona",
			Display:              "Beautiful Apartment in Barcelona",
			LocationID:           "LOC-01",
			PropertyName:         "Modern Cozy Stay",
			PropertySlug:         "modern-cozy-stay",
			PropertyTypeCategory: "Apartment",
			UsdPrice:             120.0,
			Occupancy:            4,
			BedroomCount:         2,
			BathroomCount:        1,
			NumberOfReview:       25,
			ReviewScoreGeneral:   8.5,
			StarRating:           4,
			AmenityCategories:    []string{"Internet", "Parking"},
			LonLat: models.LonLat{
				Coordinates: []float64{2.1734, 41.3851},
			},
			Categories: `[{"LocationID":"L1","Name":"Spain","Type":"Country","Slug":"spain","Display":["Spain"]},{"LocationID":"L2","Name":"Barcelona","Type":"City","Slug":"barcelona","Display":["Barcelona"]}]`,
			Published:  true,
			Images:     []string{"img1.jpg", "img2.jpg"},
		},
		{
			ID:                   "BC-1002",
			Feed:                 12,
			Country:              "Japan",
			CountryCode:          "JP",
			State:                "Tokyo",
			StateAbbr:            "TYO",
			City:                 "Tokyo",
			Display:              "Luxury Hotel in Tokyo",
			LocationID:           "LOC-02",
			PropertyName:         "Shinjuku Grand Tower",
			PropertySlug:         "shinjuku-grand-tower",
			PropertyTypeCategory: "Hotel",
			UsdPrice:             200.0,
			Occupancy:            2,
			BedroomCount:         1,
			BathroomCount:        1,
			NumberOfReview:       5,
			ReviewScoreGeneral:   6.0,
			StarRating:           3,
			AmenityCategories:    []string{"Gym"},
			LonLat: models.LonLat{
				Coordinates: []float64{139.6917, 35.6895},
			},
			Categories: `[{"LocationID":"L3","Name":"Japan","Type":"Country","Slug":"japan","Display":["Japan"]},{"LocationID":"L4","Name":"Tokyo","Type":"City","Slug":"tokyo","Display":["Tokyo"]}]`,
			Published:  false,
			Images:     []string{"hotel1.jpg"},
		},
		{
			ID:                   "BC-1003",
			Feed:                 22,
			Country:              "USA",
			CountryCode:          "US",
			State:                "Florida",
			StateAbbr:            "FL",
			City:                 "Miami",
			Display:              "Oceanfront Villa in Miami",
			LocationID:           "LOC-03",
			PropertyName:         "Miami Breeze Mansion",
			PropertySlug:         "miami-breeze-mansion",
			PropertyTypeCategory: "Villa",
			UsdPrice:             450.0,
			Occupancy:            8,
			BedroomCount:         4,
			BathroomCount:        3,
			NumberOfReview:       12,
			ReviewScoreGeneral:   9.2,
			StarRating:           5,
			AmenityCategories:    []string{"Internet", "Kitchen"},
			LonLat: models.LonLat{
				Coordinates: []float64{-80.1918, 25.7617},
			},
			Categories: `[{"LocationID":"L5","Name":"USA","Type":"Country","Slug":"usa","Display":["USA"]},{"LocationID":"L6","Name":"Miami","Type":"City","Slug":"miami","Display":["Miami"]}]`,
			Published:  true,
			Images:     []string{"villa1.jpg", "villa2.jpg", "villa3.jpg"},
		},
		{
			ID:                   "BC-1004",
			Feed:                 24,
			Country:              "Thailand",
			CountryCode:          "TH",
			State:                "Phuket",
			StateAbbr:            "HKT",
			City:                 "Phuket",
			Display:              "Tropical Resort in Phuket",
			LocationID:           "LOC-04",
			PropertyName:         "Phuket Paradise Resort",
			PropertySlug:         "phuket-paradise-resort",
			PropertyTypeCategory: "Resort",
			UsdPrice:             80.0,
			Occupancy:            2,
			BedroomCount:         1,
			BathroomCount:        1,
			NumberOfReview:       50,
			ReviewScoreGeneral:   4.5,
			StarRating:           2,
			AmenityCategories:    []string{"Pool", "Beachfront"},
			LonLat: models.LonLat{
				Coordinates: []float64{98.3923, 7.8804},
			},
			Categories: `[{"LocationID":"L7","Name":"Thailand","Type":"Country","Slug":"thailand","Display":["Thailand"]},{"LocationID":"L8","Name":"Phuket","Type":"City","Slug":"phuket","Display":["Phuket"]}]`,
			Published:  true,
			Images:     []string{"resort1.jpg", "resort2.jpg", "resort3.jpg", "resort4.jpg"},
		},
	}
}

func TestTransformData(t *testing.T) {
	source := InMemoryProperties
	res, err := TransformData(source)

	if err != nil {
		t.Fatalf("TransformData failed unexpectedly: %v", err)
	}

	if len(res) != len(source) {
		t.Fatalf("Expected %d transformed items, got %d", len(source), len(res))
	}

	if res[0].ID != source[0].ID {
		t.Errorf("Expected ID %s, got %s", source[0].ID, res[0].ID)
	}
	if res[0].Feed != source[0].Feed {
		t.Errorf("Expected Feed %d, got %d", source[0].Feed, res[0].Feed)
	}
	if res[0].Published != source[0].Published {
		t.Errorf("Expected Published %t, got %t", source[0].Published, res[0].Published)
	}

	if res[0].GeoInfo.City != source[0].City {
		t.Errorf("Expected City %s, got %s", source[0].City, res[0].GeoInfo.City)
	}
	if res[0].GeoInfo.Country != source[0].Country {
		t.Errorf("Expected Country %s, got %s", source[0].Country, res[0].GeoInfo.Country)
	}
	if res[0].GeoInfo.CountryCode != source[0].CountryCode {
		t.Errorf("Expected CountryCode %s, got %s", source[0].CountryCode, res[0].GeoInfo.CountryCode)
	}
	if res[0].GeoInfo.Name != source[0].Display {
		t.Errorf("Expected Name %s, got %s", source[0].Display, res[0].GeoInfo.Name)
	}
	if res[0].GeoInfo.LocationID != source[0].LocationID {
		t.Errorf("Expected LocationID %s, got %s", source[0].LocationID, res[0].GeoInfo.LocationID)
	}
	if res[0].GeoInfo.State != source[0].State {
		t.Errorf("Expected State %s, got %s", source[0].State, res[0].GeoInfo.State)
	}
	if res[0].GeoInfo.StateAbbr != source[0].StateAbbr {
		t.Errorf("Expected StateAbbr %s, got %s", source[0].StateAbbr, res[0].GeoInfo.StateAbbr)
	}

	if res[0].GeoInfo.Lon != source[0].LonLat.Coordinates[0] {
		t.Errorf("Expected Lon %f, got %f", source[0].LonLat.Coordinates[0], res[0].GeoInfo.Lon)
	}
	if res[0].GeoInfo.Lat != source[0].LonLat.Coordinates[1] {
		t.Errorf("Expected Lat %f, got %f", source[0].LonLat.Coordinates[1], res[0].GeoInfo.Lat)
	}

	if len(res[0].GeoInfo.Breadcrumbs) != 2 {
		t.Errorf("Expected 2 Breadcrumbs, got %d", len(res[0].GeoInfo.Breadcrumbs))
	} else {
		if res[0].GeoInfo.Breadcrumbs[0] != "Spain" || res[0].GeoInfo.Breadcrumbs[1] != "Barcelona" {
			t.Errorf("Breadcrumbs parsing mismatch: %v", res[0].GeoInfo.Breadcrumbs)
		}
	}

	if res[0].Property.Name != source[0].PropertyName {
		t.Errorf("Expected Property.Name %s, got %s", source[0].PropertyName, res[0].Property.Name)
	}
	if res[0].Property.Slug != source[0].PropertySlug {
		t.Errorf("Expected Property.Slug %s, got %s", source[0].PropertySlug, res[0].Property.Slug)
	}
	if res[0].Property.PropertyType != source[0].PropertyTypeCategory {
		t.Errorf("Expected PropertyType %s, got %s", source[0].PropertyTypeCategory, res[0].Property.PropertyType)
	}
	if res[0].Property.Price != source[0].UsdPrice {
		t.Errorf("Expected Price %f, got %f", source[0].UsdPrice, res[0].Property.Price)
	}
	if res[0].Property.ReviewScore != source[0].ReviewScoreGeneral {
		t.Errorf("Expected ReviewScore %f, got %f", source[0].ReviewScoreGeneral, res[0].Property.ReviewScore)
	}
	if res[0].Property.StarRating != source[0].StarRating {
		t.Errorf("Expected StarRating %d, got %d", source[0].StarRating, res[0].Property.StarRating)
	}

	if res[0].Property.Counts.Bedroom != source[0].BedroomCount {
		t.Errorf("Expected Bedroom count %d, got %d", source[0].BedroomCount, res[0].Property.Counts.Bedroom)
	}
	if res[0].Property.Counts.Bathroom != source[0].BathroomCount {
		t.Errorf("Expected Bathroom count %d, got %d", source[0].BathroomCount, res[0].Property.Counts.Bathroom)
	}
	if res[0].Property.Counts.Reviews != source[0].NumberOfReview {
		t.Errorf("Expected Reviews count %d, got %d", source[0].NumberOfReview, res[0].Property.Counts.Reviews)
	}
	if res[0].Property.Counts.Occupancy != source[0].Occupancy {
		t.Errorf("Expected Occupancy count %d, got %d", source[0].Occupancy, res[0].Property.Counts.Occupancy)
	}

	if res[0].Property.Image.Count != int64(len(source[0].Images)) {
		t.Errorf("Expected Image.Count %d, got %d", len(source[0].Images), res[0].Property.Image.Count)
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
