package controllers

import (
	"github.com/atopse/server/service"
)

// BoxController 驱动器
type BoxController struct {
	baseController
}

// GetBoxs 魔盒列表
// @title 魔盒列表
// @Description 获取所有的魔盒Box
// @Success 200 {object}
// @router /box/list [get]
func (c *BoxController) GetBoxs() {
	boxs, err := service.GetBoxs()
	if err != nil {
		c.JSON(err)
		return
	}
	c.JSON(boxs)
}
