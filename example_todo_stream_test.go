package ddd

import (
	"testing"
	"uuid"
)

func TestStream(t *testing.T) {
	t.Parallel()

	s := newStreamAggregate()
	if v := s.Version(); v != 0 {
		t.Errorf("incorrect version: %v, expected: %v", v, 0)
	}

	s.AppendEvent(&streamEvent{id: newEventID()})
	if v := s.Version(); v != 1 {
		t.Errorf("incorrect version: %v, expected: %v", v, 1)
	}

	s.Outbox()
}

type (
	todoID  uuid.UUID
	eventID uuid.UUID
	text    string
)

const (
	streamName    = "stream_event"
	aggregateName = "stream_aggregate"
)

func newStreamID() todoID { return todoID(uuid.New()) }
func newEventID() eventID { return eventID(uuid.New()) }

type streamEvent struct {
	id eventID
}

type createdEvent struct {
	streamEvent
	text text
}

type completedEvent struct {
	streamEvent
}

type revertedEvent struct {
	streamEvent
}

type textChangedEvent struct {
	streamEvent
	text text
}

func (s *streamEvent) ID() eventID           { return s.id }
func (s *streamEvent) Name() string          { return streamName }
func (s *streamEvent) AggregateName() string { return aggregateName }

type streamAggregate struct {
	Stream[todoID, eventID]
}

func newStreamAggregate() *streamAggregate {
	return &streamAggregate{
		Stream: NewStream[todoID, eventID](newStreamID()),
	}
}
