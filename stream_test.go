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
}

type (
	streamID uuid.UUID
	eventID  uuid.UUID
)

const (
	streamName    = "stream_event"
	aggregateName = "stream_aggregate"
)

func newStreamID() streamID { return streamID(uuid.New()) }
func newEventID() eventID   { return eventID(uuid.New()) }

type streamEvent struct {
	id eventID
}

func (s *streamEvent) ID() eventID           { return s.id }
func (s *streamEvent) Name() string          { return streamName }
func (s *streamEvent) AggregateName() string { return aggregateName }

type streamAggregate struct {
	Stream[streamID, eventID]
}

func newStreamAggregate() *streamAggregate {
	return &streamAggregate{
		Stream: NewStream[streamID, eventID](newStreamID()),
	}
}
