package event

import (
	"sync"
)

type Manager struct {
	EventListenerMap map[string][]Listener
	lock             sync.Mutex
}

func (manager *Manager) Add(eventName string, listener Listener) {
	defer func() {
		manager.lock.Unlock()
	}()
	manager.lock.Lock()
	if listeners, ok := manager.EventListenerMap[eventName]; ok {
		listeners := append(listeners, listener)
		manager.EventListenerMap[eventName] = listeners
	} else {
		listeners := make([]Listener, 100)
		manager.EventListenerMap[eventName] = append(listeners, listener)
	}
}

func (manager *Manager) Dispatch(event Event) {
	if listeners, ok := manager.EventListenerMap[event.Name()]; ok {
		for _, listener := range listeners {
			if listener != nil {
				if listener.IsAsync() {
					manager.execAsync(listener, event)
				} else {
					manager.exec(listener, event)
				}
			}
		}
	}
}

func (manager *Manager) exec(listener Listener, event Event) {
	listener.Process(event)
}

func (manager *Manager) execAsync(listener Listener, event Event) {
	go listener.Process(event)
}
