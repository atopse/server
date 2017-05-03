package test

import (
	"testing"

	"github.com/atopse/server/app"
	_ "github.com/atopse/server/routers"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	app.TestInit()
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	// r, _ := http.NewRequest("GET", "/v1/object", nil)
	// w := httptest.NewRecorder()
	// beego.BeeApp.Handlers.ServeHTTP(w, r)

	// beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	// Convey("Subject: Test Station Endpoint\n", t, func() {
	// 	Convey("Status Code Should Be 200", func() {
	// 		So(w.Code, ShouldEqual, 200)
	// 	})
	// 	Convey("The Result Should Not Be Empty", func() {
	// 		So(w.Body.Len(), ShouldBeGreaterThan, 0)
	// 	})
	// })
}
