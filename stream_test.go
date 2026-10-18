package ddd

import (
	"testing"
	"uuid"
)

func TestStream(t *testing.T) {
	t.Parallel()

	s := NewStream[streamID, eventID](newStreamID())
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

func newStreamID() streamID { return streamID(uuid.New()) }
func newEventID() eventID   { return eventID(uuid.New()) }

type streamEvent struct {
	id eventID
}

func (s *streamEvent) ID() eventID           { return s.id }
func (s *streamEvent) Name() string          { return "stream_event" }
func (s *streamEvent) AggregateName() string { return "stream_aggregate" }
