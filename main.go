package main

import (
	_ "github.com/atopse/server/routers"

	// load Drivers
	_ "github.com/atopse/box/drivers/exec"
	_ "github.com/atopse/box/drivers/string"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
