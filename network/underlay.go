package network

import "github/yhassanzadeh13/skipgraph-go/model/skipgraph"

type Underlay interface {
	// Start starts the networking layer.
	Start() <-chan interface{}

	// Stop stops the networking layer.
	Stop() <-chan interface{}

	// SetMessageHandler determines the handler of a message based on its message type.
	SetMessageHandler(interface{}, func(interface{}) error) error

	// Send sends a message to a list of target recipients in the underlying network.
	Send(interface{}, skipgraph.IdentifierList) error
}
