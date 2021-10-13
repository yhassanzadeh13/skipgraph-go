package mocknet

import (
	"fmt"
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/network"
)

type mockUnderlay struct {
	// it behaves the same for all the sharing skip graph nodes
	// there is only one handler per message type (but not per caller)
	messageHandlers map[messages.MessageType]network.MessageHandler
}

// SetMessageHandler determines the handler of a message based on its message type.
func (m *mockUnderlay) SetMessageHandler(msgType messages.MessageType, handler network.MessageHandler) error {
	m.messageHandlers[msgType] = handler
	return nil
}

// Send sends a message to a list of target recipients in the underlying network.
func (m *mockUnderlay) Send(message messages.Message) error {
	// check the support of the supplied message
	handler := m.messageHandlers[message.Type]
	if handler == nil {
		return fmt.Errorf("no handler for message type")
	}

	// call the installed handler
	return handler(message)
}

func (m *mockUnderlay) Start() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		//	implement something
	}()
	return ch
}
func (m *mockUnderlay) Stop() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		//	 implement the stop procedure
	}()
	return ch
}
