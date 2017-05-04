package controllers

import "github.com/atopse/box/drivers"

// DriverController 驱动器
type DriverController struct {
	baseController
}

// GetDrivers 驱动列表
// @title 驱动列表
// @Description 获取所有的驱动信息
// @Success 200 {object}
// @router /driver/list [get]
func (c *DriverController) GetDrivers() {
	drivers := drivers.GetDrivers()

	result := []map[string]interface{}{}
	for _, d := range drivers {
		m := make(map[string]interface{}, 4)
		m["title"] = d.Title()
		m["namespace"] = d.Namespace()
		m["description"] = d.Description()
		m["options"] = d.Options()
		m["actions"] = d.Actions()
		result = append(result, m)
		//    m[""]=
	}

	c.JSON(result)
}
