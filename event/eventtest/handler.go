package eventtest

import (
	"context"

	"github.com/ONSdigital/dp-observation-extractor/event"
)

var _ event.Handler = (*EventHandler)(nil)

// NewEventHandler returns a new mock event handler to capture event
func NewEventHandler() *EventHandler {

	events := make([]event.DimensionsInserted, 0)

	return &EventHandler{
		Events: events,
	}
}

// EventHandler provides a mock implementation that captures events to check.
type EventHandler struct {
	Events []event.DimensionsInserted
	Error  error
}

// Handle captures the given event and stores it for later assertions
func (handler *EventHandler) Handle(ctx context.Context, event *event.DimensionsInserted) error {
	handler.Events = append(handler.Events, *event)
	return handler.Error
}
