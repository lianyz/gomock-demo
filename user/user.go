package user

import "github.com/lianyz/gomock-demo/doer"

type User struct {
	Doer doer.Doer
}

func (user *User) Use() error {
	return user.Doer.DoSomething(123, "Hello GoMock")
}
