package eventsystem

type Datastore interface {
	// Save event in datastore
	SaveEvent(event *Event) error
	// Get first event that has not finished yet
	GetEvent() (*Event, error)
}
