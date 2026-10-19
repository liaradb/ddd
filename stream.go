package ddd

type Stream[T comparable, U comparable] struct {
	Aggregate[T]
	events []*OutboxEvent[T, U]
	outbox []*OutboxEvent[T, U]
}

func NewStream[T comparable, U comparable](id T) Stream[T, U] {
	return Stream[T, U]{
		id: id,
	}
}

func RestoreStream[T comparable, U comparable](
	id T,
	v Version,
	es []*OutboxEvent[T, U],
) Stream[T, U] {
	return Stream[T, U]{
		id:      id,
		version: v,
		events:  es,
	}
}

func (s *Stream[T, U]) Events() []*OutboxEvent[T, U] { return s.events }
func (s *Stream[T, U]) Outbox() []*OutboxEvent[T, U] { return s.outbox }

func (s *Stream[T, U]) AppendEvent(e Event[T, U]) {
	s.IncrementVersion()
	s.outbox = append(s.outbox, newOutboxEvent[T](e, s.version))
}
