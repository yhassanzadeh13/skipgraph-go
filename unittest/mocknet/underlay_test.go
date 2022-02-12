package mocknet_test

import (
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/unittest"
	"github/yhassanzadeh13/skipgraph-go/unittest/mocknet"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUnderlay(t *testing.T) {
	// construct an empty mocked underlay
	u := mocknet.NewMockUnderlay()

	//start
	unittest.ChannelMustCloseWithinTimeout(t, u.Start(), 100 * time.Millisecond, "could not start underlay on time")

	// stop when the test terminates
	defer u.Stop()

	// create a message type
	mtype := messages.MessageType("input")
	var msg = messages.Message{Type: mtype}
	var flag bool = false
	f := func(messages.Message) error {
		flag = true
		return nil
	}
	err := u.SetMessageHandler(mtype, f)
	require.True(t, (err == nil))

	err = u.Send(msg)
	require.True(t, (err == nil))

	// the handler is called
	require.True(t, flag)
}
