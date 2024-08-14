package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("Handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventdispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, newHandler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, handler := range ed.handlers[eventName] {
			if handler == newHandler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], newHandler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(newEventName string, comingHandler EventHandlerInterface) bool {
	eventHandlers := ed.handlers[newEventName]
	if eventHandlers == nil {
		return false
	}
	for _, handler := range eventHandlers {
		if handler == comingHandler {
			return true
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	return nil
}