package eventsystem

import (
	"errors"
	"testing"
	"time"
)

func TestSaveEvent(t *testing.T) {
	sut := MemoryDatastore{}

	sut.SaveEvent(&Event{
		Name: "Foo",
	})

	if len(sut.events) == 0 {
		t.Fatalf("events should be 1; got %d", len(sut.events))
	}
	if sut.events[0].ID == "" {
		t.Fatal("event should have an ID")
	}
}

func TestGetEvent(t *testing.T) {
	t.Run("newEvent", func(t *testing.T) {
		sut := MemoryDatastore{}
		sut.SaveEvent(&Event{
			Name: "Foo",
		})

		event, err := sut.GetEvent()
		if err != nil {
			t.Fatal(err)
		}
		if event == nil {
			t.Fatal("should return event; got none")
		}
	})

	t.Run("finishedEvent", func(t *testing.T) {
		sut := MemoryDatastore{}
		sut.SaveEvent(&Event{
			Name:       "Foo",
			StartedAt:  time.Now(),
			FinishedAt: time.Now(),
		})

		event, err := sut.GetEvent()
		if err != nil {
			t.Fatal(err)
		}
		if event != nil {
			t.Fatal("should not return event; got one")
		}
	})

	t.Run("failedEvent", func(t *testing.T) {
		sut := MemoryDatastore{}
		sut.SaveEvent(&Event{
			Name:      "Foo",
			StartedAt: time.Now(),
			Error:     errors.New("failed"),
		})

		event, err := sut.GetEvent()
		if err != nil {
			t.Fatal(err)
		}
		if event == nil {
			t.Fatal("should return event; got none")
		}
		if event.Error == nil {
			t.Fatal("event should have error; got none")
		}
	})
}
