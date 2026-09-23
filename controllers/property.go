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
// @Param   feed               query   int     false   "Exact match feed"
// @Param   min_bedroom        query   int     false   "Minimum bedroom"
// @Param   property_type      query   string  false   "Exact match category (Hotel, House, Apartment, Villa, Resort, Hostel)"
// @Param   amenities          query   string  false   "Comma-separated list of required amenities (e.g. Internet,Parking)"
// @Param   limit              query   int     false   "Maximum number of items to return"
// @Success 200 {object} map[string]interface{} "Successful transformation payload"
// @Failure 400 {object} map[string]string "Invalid parameters layout error schema"
// @router /properties/ [get]
func (c *PropertyController) ListProperties() {
	var err error
	var minPrice, maxPrice, minReviewScore float64
	var minStarRating, minReviews, feed, minBedroom, limit int64
	var published, propertyType, limitStr, amenities string

	if c.GetString("min_price") == "" {
		minPrice = -1.0
	} else {
		minPrice, err = c.GetFloat("min_price")
		if err != nil {
			c.sendBadRequest("Invalid min_price parameter")
			return
		}
	}

	if c.GetString("max_price") == "" {
		maxPrice = -1.0
	} else {
		maxPrice, err = c.GetFloat("max_price")
		if err != nil {
			c.sendBadRequest("Invalid max_price parameter")
			return
		}
	}

	if c.GetString("min_review_score") == "" {
		minReviewScore = -1.0
	} else {
		minReviewScore, err = c.GetFloat("min_review_score")
		if err != nil {
			c.sendBadRequest("Invalid min_review_score parameter")
			return
		}
	}

	if c.GetString("min_star_rating") == "" {
		minStarRating = -1
	} else {
		minStarRating, err = c.GetInt64("min_star_rating")
		if err != nil {
			c.sendBadRequest("Invalid min_star_rating parameter")
			return
		}
	}

	if c.GetString("min_reviews") == "" {
		minReviews = -1
	} else {
		minReviews, err = c.GetInt64("min_reviews")
		if err != nil {
			c.sendBadRequest("Invalid min_reviews parameter")
			return
		}
	}

	if c.GetString("feed") == "" {
		feed = -1
	} else {
		feed, err = c.GetInt64("feed")
		if err != nil {
			c.sendBadRequest("Invalid feed parameter")
			return
		}
		if feed != 11 && feed != 12 && feed != 22 && feed != 24 {
			c.sendBadRequest("Invalid feed parmeter")
			return
		}
	}

	if c.GetString("min_bedroom") == "" {
		minBedroom = -1
	} else {
		minBedroom, err = c.GetInt64("min_bedroom")
		if err != nil {
			c.sendBadRequest("Invalid min_bedroom parameter")
			return
		}
	}

	propertyType = c.GetString("property_type")
	published = c.GetString("published")
	amenities = c.GetString("amenities")

	if propertyType != "" && propertyType != "Hotel" && propertyType != "House" &&
		propertyType != "Apartment" && propertyType != "Villa" && propertyType != "Resort" && propertyType != "Hostel" {
		c.sendBadRequest("Invalid property type parameter")
		return
	}

	limitStr = c.GetString("limit")
	limit = -1

	if limitStr != "" {
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
// @router /properties/:id [get]
func (c *PropertyController) GetPropertyByID() {
	id := c.Ctx.Input.Param(":id")

	property, err := services.GetPropertyByID(id)

	if err != nil {
		if err.Error() == "Property not found" {
			c.Ctx.Output.SetStatus(404)
		}

		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"Error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = property
	c.ServeJSON()

}
