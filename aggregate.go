package ddd

type Aggregate[T comparable] struct {
	Entity[T]
	version Version
}

func NewAggregate[T comparable](id T) *Aggregate[T] {
	return &Aggregate[T]{
		id: id,
	}
}

func RestoreAggregate[T comparable](id T, v Version) *Aggregate[T] {
	return &Aggregate[T]{
		id:      id,
		version: v,
	}
}

func (a *Aggregate[T]) Version() Version { return a.version }

func (a *Aggregate[T]) IncrementVersion() {
	a.version = a.version.Increment()
}
