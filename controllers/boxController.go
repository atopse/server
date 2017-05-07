package controllers

import (
	"github.com/atopse/comm"
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

// GetBoxByNamespace 获取Box详细信息通过
// @Success 200 {object}
// @router /box/detail/:namespace [get]
func (c *BoxController) GetBoxByNamespace() {
	namespace := c.Ctx.Input.Param(":namespace")
	if namespace == "" {
		c.JSON(comm.ErrDataIsEmpty())
		return
	}
	box, err := service.FindBox(namespace)
	if err != nil {
		c.JSON(err)
		return
	}
	c.JSON(box)
}
