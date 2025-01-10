// in-memory event bus using channels, to allow communcation between modules in the modular monolith architecture
package event

import (
	"fmt"
	"sync"
)

type Bus struct {
	subscribers map[string][]func(Event)
	lock        sync.RWMutex
}

func NewEventBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]func(Event)),
	}
}

func (e *Bus) Subscribe(eventName string, callback func(Event)) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.subscribers[eventName] = append(e.subscribers[eventName], callback)
}

func (e *Bus) Publish(event Event) {
	fmt.Printf("Publishing an event: %s, with payload: %s\n", event.Name, event.Payload)
	e.lock.RLock()
	defer e.lock.RUnlock()
	if callbacks, ok := e.subscribers[event.Name]; ok {
		for _, callback := range callbacks {
			go callback(event)
		}
	}
}
