package mocknet

import (
	"errors"
)

type mockUnderlay struct {
	// it behaves the same for all the sharing skip graph nodes
	// there is only one handler per message type (but not per caller)
	messageHandlers map[interface{}]func(interface{}) error

}

// newMockUnderlay is the constructor for MockUnderlay object
func newMockUnderlay(messageHandlers map[interface{}]func(interface{}) error) *mockUnderlay {
	return &mockUnderlay{messageHandlers: messageHandlers}
}


// SetMessageHandler determines the handler of a message based on its message type.
func (m mockUnderlay) SetMessageHandler(message interface{}, handler func(interface{}) error) error {
	m.messageHandlers[message] = handler
	return nil
}

// Send sends a message to a list of target recipients in the underlying network.
func (m mockUnderlay) Send(message interface{}) error {
	// check the support of the supplied message
	handler := m.messageHandlers[message]
	if handler == nil{
		return errors.New("no handler for the supplied message")
	}

	// call the installed handler
	handler(message)

	// no error occurred
	return nil
}