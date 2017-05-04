package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/atopse/server/controllers:BoxController"] = append(beego.GlobalControllerRouter["github.com/atopse/server/controllers:BoxController"],
		beego.ControllerComments{
			Method: "GetBoxs",
			Router: `/box/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/atopse/server/controllers:DriverController"] = append(beego.GlobalControllerRouter["github.com/atopse/server/controllers:DriverController"],
		beego.ControllerComments{
			Method: "GetDrivers",
			Router: `/driver/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
