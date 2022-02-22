package doer

type Doer interface {
	DoSomething(int, string) error
}

type DoDo interface {
	DoSomething(string, string) error
}
