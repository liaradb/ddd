package ddd

type Event[T comparable] interface {
	ID() T
	Name() string
	AggregateName() string
}

type OutboxEvent[T comparable] struct {
	Event[T]
	version Version
}

func (e *OutboxEvent[T]) Version() Version { return e.version }
