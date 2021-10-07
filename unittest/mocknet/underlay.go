package mocknet

import "github/yhassanzadeh13/skipgraph-go/model/skipgraph"

type mockUnderlay struct {
	// it behaves the same for all the sharing skip graph nodes
	// there is only one handler per message type (but not per caller)
	MessageHandlers map[interface{}]func(interface{}) error

}

// SetMessageHandler determines the handler of a message based on its message type.
func (m mockUnderlay) SetMessageHandler(node skipgraph.Identifier, message interface{}, handler func(interface{}) error) error {
	// TODO should distinguish between distinct identifiers
	m.MessageHandlers[message] = handler
	return nil
}

// Send sends a message to a list of target recipients in the underlying network.
func (m mockUnderlay) Send(interface{}, skipgraph.IdentifierList) error {
	return nil

}