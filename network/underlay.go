package network

import (
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
)

// Underlay represents the underlying network for which skip graph node is interacting with.
type Underlay interface {
	// Start starts the networking layer.
	Start() <-chan interface{}

	// Stop stops the networking layer.
	Stop() <-chan interface{}

	// SetMessageHandler determines the handler of a message based on its message type.
	SetMessageHandler(messages.MessageType, MessageHandler) error

	// Send sends a message to a list of target recipients in the underlying network.
	Send(interface{}, skipgraph.IdentifierList) error
}

type MessageHandler func(messages.Message) error
