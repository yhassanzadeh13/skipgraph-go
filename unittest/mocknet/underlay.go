package mocknet

import (
	"fmt"
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/network"
)

type MockUnderlay struct {
	// there is only one handler per message type (but not per caller)
	messageHandlers map[messages.MessageType]network.MessageHandler
}

// NewMockUnderlay initializes an empty MockUnderlay and returns a pointer to it
func NewMockUnderlay() *MockUnderlay {

	return &MockUnderlay{messageHandlers: make(map[messages.MessageType]network.MessageHandler)}
}

// SetMessageHandler determines the handler of a message based on its message type.
func (m *MockUnderlay) SetMessageHandler(msgType messages.MessageType, handler network.MessageHandler) error {
	// check whether a handler exists for the supplied message type
	if m.messageHandlers[msgType] != nil {
		return fmt.Errorf("a handler exists for the attemoted message type")
	}
	m.messageHandlers[msgType] = handler
	return nil
}

// Send sends a message to a list of target recipients in the underlying network.
func (m *MockUnderlay) Send(message messages.Message) error {
	// check the support of the supplied message
	handler, ok := m.messageHandlers[message.Type]
	if !ok {
		return fmt.Errorf("no handler for message type: %s", message.Type)
	}

	// call the installed handler
	err := handler(message)
	if err != nil {
		return fmt.Errorf("could not run the message handler: %w", err)
	}
	return nil
}

func (m *MockUnderlay) Start() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		// implement the start procedure
	}()
	return ch
}
func (m *MockUnderlay) Stop() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		// implement the stop procedure
	}()
	return ch
}
