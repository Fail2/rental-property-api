package routers

import (
	controllers "rental-property-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSInclude(&controllers.PropertyController{}),
	)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")
}
