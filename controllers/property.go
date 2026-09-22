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
// @Param   property_type      query   string  false   "Exact match category (Hotel, House, Apartment, Villa, Resort, Hostel)"
// @Param   amenities          query   string  false   "Comma-separated list of required amenities (e.g. Internet,Parking)"
// @Param   limit              query   int     false   "Maximum number of items to return"
// @Success 200 {object} map[string]interface{} "Successful transformation payload"
// @Failure 400 {object} map[string]string "Invalid parameters layout error schema"
// @router /v1/properties [get]
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

	items := services.FilterProperties(minPrice, maxPrice, minStarRating, minReviewScore, minReviews, published, feed, minBedroom, propertyType, amenities, limit)

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
