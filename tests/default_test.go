package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "dana-tech.com/web-api/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
// func TestGet(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/user", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
// 	fmt.Println("[", w.Code, "]")

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 		Convey("Status Code Should Be 200", func() {
// 			So(w.Code, ShouldEqual, 200)
// 		})
// 		Convey("The Result Should Not Be Empty", func() {
// 			So(w.Body.Len(), ShouldBeGreaterThan, 0)
// 		})
// 	})
// }

func TestPost(t *testing.T) {

	r, _ := http.NewRequest("POST", "/user/login", nil)
	w := httptest.NewRecorder()
	var buff *bytes.Buffer
	buff.WriteString("username=222")
	w.Body = buff
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
	fmt.Println("[", w.Body, "]")

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
