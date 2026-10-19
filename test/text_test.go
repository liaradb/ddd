package test

type Text struct {
	value string
}

func NewText(value string) Text {
	return Text{value}
}

func (t Text) String() string { return t.value }
