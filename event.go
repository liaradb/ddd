package ddd

type Event[T comparable, U comparable] interface {
	ID() U
	Name() string
	AggregateName() string
}

type OutboxEvent[T comparable, U comparable] struct {
	Event[T, U]
	id          U
	aggregateID T
	version     Version
}

func newOutboxEvent[T comparable, U comparable](
	e Event[T, U],
	v Version,
) *OutboxEvent[T, U] {
	return &OutboxEvent[T, U]{
		Event:   e,
		version: v,
	}
}

func (e *OutboxEvent[T, U]) AggregateID() T   { return e.aggregateID }
func (e *OutboxEvent[T, U]) Version() Version { return e.version }
