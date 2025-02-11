package hooks

import "sync"

type HookHandler func(event interface{})

// Dispatcher for managing event hooks and triggering them
var hooks = make(map[string][]HookHandler)
var mu sync.Mutex

// AddHook registers a new hook for a given event name
func AddHook(eventName string, handler HookHandler) {
	mu.Lock()
	defer mu.Unlock()
	hooks[eventName] = append(hooks[eventName], handler)
}

// Trigger triggers an event and calls all registered hooks for that event
func Trigger(event Event) {
	mu.Lock()
	defer mu.Unlock()

	if handlers, found := hooks[event.GetName()]; found {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
