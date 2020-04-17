package gomock

import (
	"errors"
	. "github.com/golang/mock/gomock"
	_ "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestObjDemo(t *testing.T) {
	Convey("test obj create", t, func() {
		// 1.控制器创建
		ctrl := NewController(t)
		// 2.延迟关闭控制器
		defer ctrl.Finish()
		// 3.创建接口对象
		mockRepo := NewMockRepository(ctrl)
		// 4.对方法进行打桩
		mockRepo.EXPECT().Create(Any(), Any()).Return(errors.New("error"))
		// 5.使用接口
		err := mockRepo.Create("err", nil)
		// 6.判断
		So(err, ShouldNotBeNil)
	})

	Convey("test obj retrieve", t, func() {
		// 1.控制器创建
		ctrl := NewController(t)
		// 2.延迟关闭控制器
		defer ctrl.Finish()
		// 3.创建接口对象
		mockRepo := NewMockRepository(ctrl)
		// 4.对方法进行打桩
		mockRepo.EXPECT().Retrieve(Any()).Return([]byte{1, 2, 3}, nil)
		// 5.使用接口
		b, err := mockRepo.Retrieve("")
		// 6.判断
		So(err, ShouldBeNil)
		So(b, ShouldResemble, []byte{1, 2, 3})
	})
}
