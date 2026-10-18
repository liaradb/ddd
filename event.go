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

func newOutboxEvent[T comparable](e Event[T], v Version) *OutboxEvent[T] {
	return &OutboxEvent[T]{
		Event:   e,
		version: v,
	}
}

func (e *OutboxEvent[T]) Version() Version { return e.version }
