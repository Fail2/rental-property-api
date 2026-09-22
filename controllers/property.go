package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

func (c *PropertyController) ListProperties() {
	minPrice, _ := c.GetFloat("min_price")
	maxPrice, _ := c.GetFloat("max_price")
	minReviewScore, _ := c.GetFloat("min_review_score")

	minStarRating, _ := c.Getint64("min_star_rating")
	minReviews, _ := c.Getint64("min_revies")
	minBedroom, _ := c.Getint64("min_beedroom")
	feed, _ := c.Getint64("feed")

	propertyType := c.GetString("property_type")
	publishedStr := c.GetString("published")
	amenities := c.Getint64("amenities")

	limitStr := c.GetSession("limit")
	limit := 0

}
