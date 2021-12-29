package todo

type ToDo interface {
	Initialise() error
	Create(text string, isDone bool)
}
