package test

import "uuid"

type TodoEventID struct {
	value uuid.UUID
}

func NewTodoEventID() TodoEventID {
	return TodoEventID{uuid.New()}
}

func RestoreTodoEventID(value uuid.UUID) TodoEventID {
	return TodoEventID{value}
}

func (i TodoEventID) String() string { return i.value.String() }
