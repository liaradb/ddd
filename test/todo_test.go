package test

import "github.com/liaradb/ddd"

const (
	TodoAggregateName = "todo"
	TodoTextChanged   = "todo_text_changed"
)

type Todo struct {
	ddd.Stream[TodoID, TodoEventID]
	complete bool
	text     Text
}

func NewTodo() *Todo {
	return &Todo{
		Stream: ddd.NewStream[TodoID, TodoEventID](NewTodoID()),
	}
}

func (t *Todo) Complete() bool { return t.complete }
func (t *Todo) Text() Text     { return t.text }

func (t *Todo) ChangeText(text Text) {
	e := TextChangedEvent{}
	t.AppendEvent(e)
	t.applyChangeText(e)
}

func (t *Todo) applyChangeText(e TextChangedEvent) {
	t.text = e.text
}

type TextChangedEvent struct {
	text Text
}

func (t TextChangedEvent) AggregateName() string { return TodoAggregateName }
func (t TextChangedEvent) Name() string          { return TodoTextChanged }

func (t TextChangedEvent) ID() TodoEventID {
	panic("unimplemented")
}
