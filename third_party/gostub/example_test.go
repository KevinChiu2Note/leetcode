package gostub

import (
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAll(t *testing.T) {
	Convey("测试打桩", t, func() {
		Convey("全局变量", func() {
			stubs := gostub.Stub(&globalVar, 150)
			defer stubs.Reset()
			So(globalVar, ShouldEqual, 150)
		})
		Convey("方法", func() {
			stubs := gostub.StubFunc(&funcGetVar, 400)
			defer stubs.Reset()
			So(funcGetVar(), ShouldEqual, 400)
		})
		Convey("过程", func() {
			stubs := gostub.StubFunc(&funcDoMain)
			defer stubs.Reset()
			funcDoMain(5) // 函数中的过程被忽略了
			So(globalVar, ShouldEqual, 100)
		})
		Convey("复合打桩", func() {
			stubs := gostub.Stub(&globalVar, 150)
			defer stubs.Reset()
			stubs.StubFunc(&funcDoMain)
			funcDoMain(0)
			So(globalVar, ShouldEqual, 150)
		})
		Convey("无法重构的函数", func() {
			stubs := gostub.StubFunc(&osHostName, "NotKevin", nil)
			defer stubs.Reset()
			name, err := osHostName()
			So(err, ShouldBeNil)
			So(name, ShouldEqual, "NotKevin")
		})
	})
}
