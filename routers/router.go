package routers

import (
	controllers "rental-property-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Include(&controllers.PropertyController{})

	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/properties",
			&controllers.PropertyController{}, "get:ListProperties"),
	)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")
}
