package multi_return

type MultiReturn interface {
	MultiReturn() (int, error)
}

