package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks github.com/lianyz/gomock-demo/doer Doer,DoDo

type Doer interface {
	DoSomething(int, string) error
}

type DoDo interface {
	DoSomething(string, string) error
}
