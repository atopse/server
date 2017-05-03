package service

import (
	"github.com/atopse/box"
	"github.com/atopse/box/drivers"
	"github.com/atopse/comm"
	"github.com/atopse/comm/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func checkBox(b *box.Box) error {
	if b.Namespace == "" {
		return comm.ErrDataIsEmpty("namespace")
	}
	if b.Title == "" {
		return comm.ErrDataIsEmpty("title")
	}
	if len(b.Actions) == 0 {
		return comm.ErrDataIsEmpty("actions")
	}
	return nil
}

// SaveNewBox 新建Box
func SaveNewBox(b *box.Box) error {
	if err := checkBox(b); err != nil {
		return err
	}
	return db.Insert("box", b)
}

// UpdateBox 根据namespace更新Box
func UpdateBox(b *box.Box) error {
	if err := checkBox(b); err != nil {
		return err
	}
	return db.Do(func(d *mgo.Database) error {
		return d.C("box").Update(bson.M{"namespace": b.Namespace}, b)
	})
}

// FindBox 根据 namespace 查找Box
func FindBox(namespace string) (*box.Box, error) {
	if namespace == "" {
		return nil, comm.ErrDataIsEmpty("namespace")
	}
	b := &box.Box{}
	err := db.FindOne("box", bson.M{"namespace": namespace}, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// RemoteBox 移除Box
func RemoteBox(namespace string) error {
	if namespace == "" {
		return comm.ErrDataIsEmpty("namespace")
	}
	return db.Do(func(d *mgo.Database) error {
		return d.C("box").Remove(bson.M{"namespace": namespace})
	})
}
