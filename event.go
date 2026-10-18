package ddd

type Event[T comparable] interface {
	ID() T
	Name() string
	AggregateName() string
}
