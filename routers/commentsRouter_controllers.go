package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/atopse/server/controllers:DriverController"] = append(beego.GlobalControllerRouter["github.com/atopse/server/controllers:DriverController"],
		beego.ControllerComments{
			Method: "GetDrivers",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
