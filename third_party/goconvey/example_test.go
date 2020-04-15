package goconvey

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	Convey("测试加法", t, func() {
		So(Add(2, 3), ShouldEqual, 5)
	})
}

func TestSubtract(t *testing.T) {
	Convey("测试减法", t, func() {
		So(Subtract(5, 3), ShouldEqual, 2)
	})
}

func TestMultiply(t *testing.T) {
	Convey("测试乘法", t, func() {
		So(Multiply(8, 9), ShouldEqual, 72)
	})

}

func TestDivide(t *testing.T) {
	Convey("测试除法", t, func() {
		Convey("除数为0", func() {
			_, err := Divide(9, 0)
			So(err, ShouldNotBeNil)
		})
		Convey("除数不为0", func() {
			res, err := Divide(9, 7)
			So(err, ShouldBeNil)
			So(res, ShouldEqual, 1)
		})
	})
}
