package main

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	sut := NewEventSystem(NewMemoryDatastore(), false)
	sut.Register(EventHandler{
		Event: "foo",
		Handler: func(payload interface{}) error {
			return nil
		},
	})

	if len(sut.handlers) == 0 {
		t.Fatalf("handlers should be 1; got %d", len(sut.handlers))
	}
}

func TestPublish(t *testing.T) {
	t.Run("publish", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		sut := NewEventSystem(NewMemoryDatastore(), false)

		sut.Register(EventHandler{
			Event: "foo",
			Handler: func(payload interface{}) error {
				if payload.(string) != "bar" {
					t.Fatalf("payload should be bar; got %s", payload.(string))
				}
				wg.Done()
				return nil
			},
		})

		err := sut.Publish("foo", "bar")
		if err != nil {
			t.Fatalf("publish should succeed; got %s", err.Error())
		}

		wg.Wait()
	})

	t.Run("publishHandlerFails", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		sut := NewEventSystem(NewMemoryDatastore(), false)

		sut.Register(EventHandler{
			Event: "foo",
			Handler: func(payload interface{}) error {
				wg.Done()
				return errors.New("fail")
			},
		})

		err := sut.Publish("foo", "bar")
		if err != nil {
			t.Fatalf("publish should succeed; got %s", err.Error())
		}

		wg.Wait()

		event, _ := sut.datastore.GetEvent()
		if event.Error.Error() != "fail" {
			t.Fatalf("event error should be fail; got %s", event.Error.Error())
		}
	})
}

func TestRestart(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	sut := NewEventSystem(&MemoryDatastore{
		events: []*Event{
			{
				Name:      "foo",
				StartedAt: time.Now(),
				Error:     errors.New("fail"),
			},
		},
	}, false)
	sut.Register(EventHandler{
		Event: "foo",
		Handler: func(payload interface{}) error {
			wg.Done()
			return nil
		},
	})

	event, _ := sut.datastore.GetEvent()
	if event == nil {
		t.Fatal("event should exist; got none")
	}

	sut.Restart()

	wg.Wait()

	event, _ = sut.datastore.GetEvent()
	if event != nil {
		t.Fatal("event should not exist; got one")
	}
}
