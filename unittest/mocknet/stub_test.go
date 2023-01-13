package mocknet_test

import (
	"errors"
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/unittest"
	"github/yhassanzadeh13/skipgraph-go/unittest/mocknet"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestTwoUnderlays checks two mock underlays can send message to each other
func TestTwoUnderlays(t *testing.T) {
	// construct an empty mocked underlay
	stub := mocknet.NewNetworkStub()

	// create a random identifier
	id1 := unittest.IdentifierFixture(t)
	u1 := stub.NewMockUnderlay(t, id1)

	// create a random identifier
	id2 := unittest.IdentifierFixture(t)
	u2 := stub.NewMockUnderlay(t, id2)

	// make sure they are not equal
	require.NotEqual(t, id1, id2)

	// starts underlay
	unittest.ChannelsMustCloseWithinTimeout(t,
		100*time.Millisecond, "could not start underlays on time", u1.Start(), u2.Start())

	// sets message handler at u1
	received := false
	var receivedPayload interface{}
	f := func(msg messages.Message) error {
		received = true
		receivedPayload = msg.Payload
		return nil
	}
	require.NoError(t, u1.SetMessageHandler(unittest.TestMessageType, f))

	// sends message from u2 -> u1
	msg := unittest.TestMessageFixture(t)
	// TODO: refactor message as an interface
	// TODO: add test for u1 -> u2
	require.NoError(t, u2.Send(*msg, id1))

	// the handler is called
	require.True(t, received)
	require.Equal(t, msg.Payload, receivedPayload)

	// stops underlay
	unittest.ChannelsMustCloseWithinTimeout(t,
		100*time.Millisecond, "could not stop underlay on time", u1.Stop(), u2.Stop())
}

// TestStoppedMockedUnderlays checks communication between inactive mocked underlays
func TestStoppedMockedUnderlays(t *testing.T) {
	// construct an empty mocked underlay
	stub := mocknet.NewNetworkStub()

	// create a random identifier
	id1 := unittest.IdentifierFixture(t)
	u1 := stub.NewMockUnderlay(t, id1)

	// create a random identifier
	id2 := unittest.IdentifierFixture(t)
	u2 := stub.NewMockUnderlay(t, id2)

	// make sure they are not equal
	require.NotEqual(t, id1, id2)

	// starts underlays
	unittest.ChannelsMustCloseWithinTimeout(t,
		100*time.Millisecond, "could not start underlays on time", u1.Start(), u2.Start())

	// sets message handler at u1
	received := 0
	var receivedPayload interface{}
	f := func(msg messages.Message) error {
		received++
		receivedPayload = msg.Payload
		return nil
	}
	require.NoError(t, u1.SetMessageHandler(unittest.TestMessageType, f))

	// sets message handler at u2
	received2 := 0
	f2 := func(msg messages.Message) error {
		received2++
		return nil
	}
	require.NoError(t, u2.SetMessageHandler(unittest.TestMessageType, f2))

	// Test the interaction
	// sends message from u2 -> u1, it should work
	msg := unittest.TestMessageFixture(t)
	require.NoError(t, u2.Send(*msg, id1))

	// the u1's handler is called
	require.Equal(t, 1, received)
	require.Equal(t, msg.Payload, receivedPayload)

	// stops u2
	unittest.ChannelsMustCloseWithinTimeout(t,
		100*time.Millisecond, "could not stop underlay on time", u2.Stop())

	// sends message from u2 -> u1, it should not work
	msg2 := unittest.TestMessageFixture(t)
	// the send should error out as u2 is inactive
	require.True(t, errors.Is(u2.Send(*msg2, id1), mocknet.InactiveUnderlayError))

	// send a message from u1 -> u2, it should not work as u2 is not active
	msg3 := unittest.TestMessageFixture(t)
	// the send should error out as u2 is inactive
	err := u1.Send(*msg3, id2)
	require.True(t, mocknet.IsCouldNotHandleMessageError(err))

}
