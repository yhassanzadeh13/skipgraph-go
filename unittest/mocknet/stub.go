package mocknet

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"sync"
	"testing"
)

// NetworkStub acts as a router to connect a set of MockUnderlay
// it needs to be locked using its l field before being accessed
type NetworkStub struct {
	l         sync.Mutex
	underlays map[skipgraph.Identifier]*MockUnderlay
}

// NewNetworkStub creates an empty NetworkStub
func NewNetworkStub() *NetworkStub {
	return &NetworkStub{underlays: make(map[skipgraph.Identifier]*MockUnderlay)}
}

// NewMockUnderlay creates and returns a mock underlay connected to this network stub for a non-existing Identifier.
func (n *NetworkStub) NewMockUnderlay(t *testing.T, id skipgraph.Identifier) *MockUnderlay {
	n.l.Lock()
	defer n.l.Unlock()

	_, exists := n.underlays[id]
	require.False(t, exists, "attempting to create mock underlay for already existing identifier")

	u := NewMockUnderlay(n)
	n.underlays[id] = u

	return u
}

// routeMessageTo imitates routing the message in the underlying network to the target identifier's mock underlay.
func (n *NetworkStub) routeMessageTo(msg messages.Message, target skipgraph.Identifier) error {
	n.l.Lock()
	defer n.l.Unlock()

	u, exists := n.underlays[target]
	if !exists {
		return &NoMockUnderlayError{identifier: target}
	}

	h, exists := u.messageHandlers[msg.Type]
	if !exists {
		return NoHandlerExistError{msgType: msg.Type}
	}

	err := h(msg)
	if err != nil {
		return CouldNotHandleMessageError{msg: msg, id: target}
	}

	return nil
}

/*
========================== error types ========================================
*/

// NoMockUnderlayError indicates absence of a mock underlay for a node.
type NoMockUnderlayError struct {
	identifier skipgraph.Identifier
}

// Error implements the Error interface for NoMockUnderlayError.
func (e NoMockUnderlayError) Error() string {
	return fmt.Sprintf("no mock underlay exists for %x", e.identifier)
}

// isCouldNotHandleMessageError checks whether err is of the type CouldNotHandleMessageError.
func isCouldNotHandleMessageError(err error) bool {
	return errors.As(err, &CouldNotHandleMessageError{})
}

// ======================

// NoHandlerExistError indicates absence of a handler for a message type.
type NoHandlerExistError struct {
	msgType messages.Type
}

// Error implements the Error interface for NoHandlerExistError.
func (e NoHandlerExistError) Error() string {
	return fmt.Sprintf("no handler exists for message type %v", e.msgType)
}

// IsNoHandlerExistError checks whether err is of type NoHandlerExistError.
func IsNoHandlerExistError(err error) bool {
	return errors.As(err, &NoHandlerExistError{})
}

// ======================

// CouldNotHandleMessageError indicates that a message could not be handled by a message handler.
type CouldNotHandleMessageError struct {
	msg messages.Message
	id  skipgraph.Identifier
}

// Error implements the Error interface for CouldNotHandleMessageError.
func (e CouldNotHandleMessageError) Error() string {
	return fmt.Sprintf("the message handler of mock underlay of node %v could not handle message %v", e.id, e.msg)
}

// IsCouldNotHandleMessageError  checks whether err is of type CouldNotHandleMessageError.
func IsCouldNotHandleMessageError(err error) bool {
	return errors.As(err, &CouldNotHandleMessageError{})
}
