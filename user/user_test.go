package user

import (
	"github.com/golang/mock/gomock"
	"github.com/lianyz/gomock-demo/match"
	"github.com/lianyz/gomock-demo/mocks"
	"testing"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	call1 := mockDoer.EXPECT().
		DoSomething(123, "Hello GoMock").
		Return(nil).
		Times(1)
	testUser.Use()

	// 使用参数匹配
	call2 := mockDoer.EXPECT().
		DoSomething(gomock.Any(), gomock.Eq("Hello GoMock")).
		Return(nil).
		Times(1)
	testUser.Use()

	call3 := mockDoer.EXPECT().
		DoSomething(123, "Hello GoMock").
		Return(nil).
		Times(1)
	testUser.Use()

	// 自定义matcher
	call4 := mockDoer.EXPECT().
		DoSomething(123, match.OfType("string")).
		Return(nil).
		Times(1)
	testUser.Use()

	// 断言调用顺序
	call2.After(call1)
	call3.After(call2)
	call4.After(call3)

	// 指定mock行为
	mockDoer.EXPECT().
		DoSomething(gomock.Any(), gomock.Any()).
		Return(nil).
		Do(func(x int, y string) {
			if x > len(y) {
				t.Fatalf("x larger then lenth of y")
			}
		})
	testUser.Use()
}
