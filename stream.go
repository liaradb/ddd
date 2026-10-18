package ddd

type Stream[T comparable, U comparable] struct {
	Aggregate[T]
	events []Event[U]
	outbox []OutboxEvent[U]
}

func NewStream[T comparable, U comparable](id T) *Stream[T, U] {
	return &Stream[T, U]{
		id: id,
	}
}

func RestoreStream[T comparable, U comparable](
	id T,
	v Version,
	es []Event[U],
) *Stream[T, U] {
	return &Stream[T, U]{
		id:      id,
		version: v,
		events:  es,
	}
}

func (s *Stream[T, U]) Events() []Event[U]       { return s.events }
func (s *Stream[T, U]) Outbox() []OutboxEvent[U] { return s.outbox }

func (s *Stream[T, U]) AppendEvent(e Event[U]) {
	s.IncrementVersion()
	s.outbox = append(s.outbox, OutboxEvent[U]{
		Event:   e,
		version: s.version,
	})
}
