package test

import "uuid"

type TodoID struct {
	value uuid.UUID
}

func NewTodoID() TodoID {
	return TodoID{uuid.New()}
}

func RestoreTodoID(value uuid.UUID) TodoID {
	return TodoID{value}
}

func (i TodoID) String() string { return i.value.String() }
