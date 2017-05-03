package app

import (
	"github.com/atopse/comm"
	"github.com/atopse/comm/config"
	"github.com/atopse/comm/log"

	_ "github.com/atopse/comm/db"
	_ "github.com/atopse/server/routers"

	// load Drivers
	_ "github.com/atopse/box/drivers/exec"
	_ "github.com/atopse/box/drivers/string"

	"github.com/astaxie/beego"
)

func before() {
	beego.AddAPPStartHook(func() error {
		return comm.RunAPPInitHook()
	})
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

// Run 运行APP
func Run() {
	before()
	beego.Run()
	log.Info("ATOPSE run end")
}

// TestInit 测试Init
func TestInit() {
	before()
	beego.InitBeegoBeforeTest(config.ConfigPath)
	log.Debug("test init completed")
}
