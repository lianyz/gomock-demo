package user

import (
	"github.com/golang/mock/gomock"
	"github.com/lianyz/gomock-demo/mocks"
	"testing"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").
		Return(nil).Times(1)

	testUser.Use()
}
