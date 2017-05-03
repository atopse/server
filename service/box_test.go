package service

import (
	"testing"

	"gopkg.in/mgo.v2"

	"github.com/atopse/box"
	"github.com/atopse/box/drivers"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInserBox(t *testing.T) {
	Convey("Box操作", t, func() {
		namespace := "boxOps.test.box.atopse"

		err := RemoteBox(namespace)
		if err != mgo.ErrNotFound {
			So(err, ShouldBeNil)
		}

		b, err := box.New("获取Go版本信息", namespace, "获取目标环境下的GO版本信息")
		So(err, ShouldBeNil)
		b.Actions = []box.ActionOption{
			{Driver: "exec.driver.atopse", Action: "execute",
				Input: drivers.Values{"command": "go version"}, OutputVar: "goVersion"},
			{Driver: "string.driver.atopse", Action: "format",
				Input: drivers.Values{
					"fmtway": "html.template",
					"format": "{{.Name}}:{{.goVersion}}",
					"data":   map[string]interface{}{"Name": "GoInfo"}},
			},
		}
		err = SaveNewBox(b)
		So(err, ShouldBeNil)

		Convey("更新Box", func() {
			b.Title = "[update]" + b.Title
			err = UpdateBox(b)
			So(err, ShouldBeNil)

			newBox, err := FindBox(b.Namespace)
			So(err, ShouldBeNil)
			So(b.Title, ShouldEqual, newBox.Title)
			So(b.Namespace, ShouldEqual, newBox.Namespace)
			So(b.Description, ShouldEqual, newBox.Description)

			newBox.Description = "[update]" + newBox.Description
			err = UpdateBox(newBox)
			So(err, ShouldBeNil)

			newBox2, err := FindBox(newBox.Namespace)
			So(err, ShouldBeNil)
			So(newBox.Title, ShouldEqual, newBox2.Title)

		})

	})

}
