package distribution

type Distribution[T any] interface {
	Add(a T, n int) error
	Ordering() []T
	Clear()
}
