# Go Event System

A simple event system that publishes events to handlers. The events are saved in a datastore so that they can get replayed if needed. In case of an event error, the system will stop until the error is resolved.

## Getting Started

```go
datastore := NewMemoryDatastore()
enableLog := true
eventSystem := NewEventSystem(datastore, enableLog)

// register event handler
eventSystem.Register(EventHandler{
    Event: "foo",
    Handler: func(payload interface{}) error {
        // do something with the payload
        return nil
    },
})

// publish event with payload
eventSystem.Publish("foo", "bar")
```

## Datastore

You can write your own datastore or use the existing one that will hold the data in memory

```go
type Datastore interface {
	// Save event in datastore
	SaveEvent(event *Event) error
	// Get first event that has not finished yet
	GetEvent() (*Event, error)
}
```

## Error Handling

In case one of the handler fails the system will stop. After you have resolved the issue, restart the system with the following function.

```go
eventSystem.Restart()
```
