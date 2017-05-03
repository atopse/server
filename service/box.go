package service

import (
	"github.com/atopse/box"
	"github.com/atopse/box/drivers"
	"github.com/atopse/comm/db"
)

// GetDrivers 获取所有已注册的Drivers
func GetDrivers() ([]drivers.DriverInterface, error) {
	return drivers.GetDrivers(), nil
}

// GetBoxs 获取所有Box
func GetBoxs() ([]box.Box, error) {
	var boxs []box.Box
	err := db.FindAll("box", nil, &boxs)
	return boxs, err
}
