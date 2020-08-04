package event

import (
	"sync"
)

type Manager struct {
	eventListenerMap map[string][]Listener
	lock             sync.Mutex
}

func New() *Manager {
	return &Manager{eventListenerMap: make(map[string][]Listener)}
}

func (manager *Manager) Add(eventName string, listener Listener) {
	defer func() {
		manager.lock.Unlock()
	}()
	manager.lock.Lock()
	if listeners, ok := manager.eventListenerMap[eventName]; ok {
		listeners := append(listeners, listener)
		manager.eventListenerMap[eventName] = listeners
	} else {
		listeners := make([]Listener, 100)
		manager.eventListenerMap[eventName] = append(listeners, listener)
	}
}

func (manager *Manager) Dispatch(event Event) {
	if listeners, ok := manager.eventListenerMap[event.Name()]; ok {
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
