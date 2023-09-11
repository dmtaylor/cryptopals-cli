package distribution

type Distribution interface {
	Add(a any, n int) error
	Ordering() []any
	Clear()
}
