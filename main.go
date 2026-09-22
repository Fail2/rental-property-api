// @APIVersion 1.0.0
// @Title Rental Property REST API
// @Description This API reads property data from a JSON file, loads it into memory at startup, and exposes read-only endpoints.
// @BasePath /v1
package main

import (
	_ "rental-property-api/routers"

	"rental-property-api/services"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	services.LoadData()
	beego.Run()

}
