package eventsystem

import "github.com/google/uuid"

type MemoryDatastore struct {
	events []*Event
}

func NewMemoryDatastore() *MemoryDatastore {
	return &MemoryDatastore{
		events: []*Event{},
	}
}

func (d *MemoryDatastore) SaveEvent(event *Event) error {
	if event.ID == "" {
		// Save new event
		event.ID = uuid.NewString()
		d.events = append(d.events, event)
	} else {
		// Update existing event
		events := []*Event{}
		for _, e := range d.events {
			if e.ID == event.ID {
				events = append(events, event)
			} else {
				events = append(events, e)
			}
		}
	}
	return nil
}

func (d *MemoryDatastore) GetEvent() (*Event, error) {
	for _, event := range d.events {
		if event.FinishedAt.IsZero() {
			return event, nil
		}
	}
	return nil, nil
}
