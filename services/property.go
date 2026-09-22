package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"rental-property-api/models"
	"runtime"
)

func PropertyExtract() {
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

	if err != nil {
		fmt.Println("Can't extract Amenities: ", err)
	}

	for index, property := range properties {
		fmt.Println("Index:", index, "Category: ", property.Categories, "\n\n\n")
		var Categories []models.CategoryDetail

		err = json.Unmarshal([]byte(property.Categories), &Categories)
		if err != nil {
			fmt.Printf("Index:%d of Can't parse categories: %s \n", index, err)
			return
		}

		for _, category := range Categories {
			fmt.Println(category)
		}

	}

}
