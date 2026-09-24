package controllers

import (
	"rental-property-api/services"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

// ListProperties
// @Title List Properties
// @Description Get a list of transformed rental properties with optional AND/OR query filters and limit caps.
// @Param   min_price          query   float   false   "Minimum property price"
// @Param   max_price          query   float   false   "Maximum property price"
// @Param   min_star_rating    query   int     false   "Minimum star rating"
// @Param   min_review_score   query   float   false   "Minimum review score"
// @Param   min_reviews        query   int     false   "Minimum reviews"
// @Param   published          query   bool    false   "Published or not"
// @Param   feed               query   int     false   "Exact match feed (11, 12, 22, 24)"
// @Param   min_bedroom        query   int     false   "Minimum bedroom"
// @Param   property_type      query   string  false   "Exact match category (Hotel, House, Apartment, Villa, Resort, Hostel)"
// @Param   amenities          query   string  false   "Comma-separated list of required amenities (e.g. Internet,Parking)"
// @Param   limit              query   int     false   "Maximum number of items to return"
// @Success 200 {object} models.PropertyListResponseWrapper "Successful transformation payload"
// @Failure 400 {object} map[string]string "Invalid parameters layout error schema"
// @Failure 500 {object} map[string]string "Internal server error"
// @router /properties/ [get]
func (c *PropertyController) ListProperties() {
	var err error
	var minPrice, maxPrice, minReviewScore float64
	var minStarRating, minReviews, feed, minBedroom, limit int64
	var published, propertyType, limitStr, amenities string

	queryParams := c.Ctx.Input.Context.Request.URL.Query()

	if _, exists := queryParams["min_price"]; !exists {
		minPrice = -1.0
	} else {
		if c.GetString("min_price") == "" {
			c.sendBadRequest("Invalid min_price parameter")
			return
		}
		minPrice, err = c.GetFloat("min_price")
		if err != nil || minPrice < 0 {
			c.sendBadRequest("Invalid min_price parameter")
			return
		}
	}

	if _, exists := queryParams["max_price"]; !exists {
		maxPrice = -1.0
	} else {
		if c.GetString("max_price") == "" {
			c.sendBadRequest("Invalid max_price parameter")
			return
		}
		maxPrice, err = c.GetFloat("max_price")
		if err != nil || maxPrice < 0 {
			c.sendBadRequest("Invalid max_price parameter")
			return
		}
	}

	if _, exists := queryParams["min_review_score"]; !exists {
		minReviewScore = -1.0
	} else {
		if c.GetString("min_review_score") == "" {
			c.sendBadRequest("Invalid min_review_score parameter")
			return
		}
		minReviewScore, err = c.GetFloat("min_review_score")
		if err != nil || minReviewScore < 0 {
			c.sendBadRequest("Invalid min_review_score parameter")
			return
		}
	}

	if _, exists := queryParams["min_star_rating"]; !exists {
		minStarRating = -1
	} else {
		if c.GetString("min_star_rating") == "" {
			c.sendBadRequest("Invalid min_star_rating parameter")
			return
		}
		minStarRating, err = c.GetInt64("min_star_rating")
		if err != nil || minStarRating < 0 {
			c.sendBadRequest("Invalid min_star_rating parameter")
			return
		}
	}

	if _, exists := queryParams["min_reviews"]; !exists {
		minReviews = -1
	} else {
		if c.GetString("min_reviews") == "" {
			c.sendBadRequest("Invalid min_reviews parameter")
			return
		}
		minReviews, err = c.GetInt64("min_reviews")
		if err != nil || minReviews < 0 {
			c.sendBadRequest("Invalid min_reviews parameter")
			return
		}
	}

	if _, exists := queryParams["feed"]; !exists {
		feed = -1
	} else {
		if c.GetString("feed") == "" {
			c.sendBadRequest("Invalid feed parameter")
			return
		}
		feed, err = c.GetInt64("feed")
		if err != nil || feed < 0 {
			c.sendBadRequest("Invalid feed parameter")
			return
		}
		if feed != 11 && feed != 12 && feed != 22 && feed != 24 {
			c.sendBadRequest("Invalid feed parmeter")
			return
		}
	}

	if _, exists := queryParams["min_bedroom"]; !exists {
		minBedroom = -1
	} else {
		if c.GetString("min_bedroom") == "" {
			c.sendBadRequest("Invalid min_bedroom parameter")
			return
		}
		minBedroom, err = c.GetInt64("min_bedroom")
		if err != nil || minBedroom < 0 {
			c.sendBadRequest("Invalid min_bedroom parameter")
			return
		}
	}

	if _, exists := queryParams["property_type"]; exists && c.GetString("property_type") == "" {
		c.sendBadRequest("Invalid property type parameter")
		return
	}
	propertyType = c.GetString("property_type")
	if propertyType != "" && propertyType != "Hotel" && propertyType != "House" &&
		propertyType != "Apartment" && propertyType != "Villa" && propertyType != "Resort" && propertyType != "Hostel" {
		c.sendBadRequest("Invalid property type parameter")
		return
	}

	if _, exists := queryParams["published"]; exists && c.GetString("published") == "" {
		c.sendBadRequest("Invalid published parameter")
		return
	}
	published = c.GetString("published")

	if _, exists := queryParams["amenities"]; exists && c.GetString("amenities") == "" {
		c.sendBadRequest("Invalid amenities parameter")
		return
	}
	amenities = c.GetString("amenities")

	if _, exists := queryParams["limit"]; !exists {
		limit = -1
	} else {
		limitStr = c.GetString("limit")
		if limitStr == "" {
			c.sendBadRequest("Invalid limit parameter")
			return
		}
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			c.sendBadRequest("Invalid limit parameter")
			return
		}
		limit = int64(parsedLimit)
	}

	items, err := services.FilterProperties(minPrice, maxPrice, minStarRating, minReviewScore, minReviews, published, feed, minBedroom, propertyType, amenities, limit)

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"Error": "An unexpected internal server error occurred"}
		c.ServeJSON()
		return
	}
	response := map[string]interface{}{
		"Result": map[string]interface{}{
			"Count": len(items),
			"Items": items,
		},
	}

	c.Data["json"] = response
	c.ServeJSON()
}

func (c *PropertyController) sendBadRequest(message string) {
	c.Ctx.Output.SetStatus(400)
	c.Data["json"] = map[string]string{"Error": message}
	c.ServeJSON()
}

// GetPropertyByID
// @Title Get Property By ID
// @Description Retrieve a single fully transformed rental property object directly using its unique ID.
// @Param   id     path    string  true        "The unique Property ID (e.g. BC-1000001)"
// @Success 200 {object} models.ResponseProperty "Single transformed property object"
// @Failure 400 {object} map[string]string "Bad Request error schema"
// @Failure 404 {object} map[string]string "Property not found error schema"
// @Failure 500 {object} map[string]string "Internal server error"
// @router /properties/:id [get]
func (c *PropertyController) GetPropertyByID() {
	id := c.Ctx.Input.Param(":id")

	property, err := services.GetPropertyByID(id)

	if err != nil {
		if err.Error() == "Property not found" {
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = map[string]string{"Error": err.Error()}
			c.ServeJSON()
			return
		}

		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"Error": "An unexpected internal server error occurred"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = property
	c.ServeJSON()

}
