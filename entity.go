package ddd

type Entity[T comparable] struct {
	id T
}

func NewEntity[T comparable](id T) Entity[T] {
	return Entity[T]{id: id}
}

func (e *Entity[T]) ID() T { return e.id }
