package mocknet_test

import (
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/unittest/mocknet"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnderlay(t *testing.T) {
	// construct an empty mocked underlay
	mUnderlay := mocknet.NewMockUnderlay()

	//start
	mUnderlay.Start()

	// stop when the test terminates
	defer mUnderlay.Stop()

	// create a message type
	mtype := messages.MessageType("input")
	var msg = messages.Message{Type: mtype}
	var flag bool = false
	f := func(messages.Message) error {
		flag = true
		return nil
	}
	err := mUnderlay.SetMessageHandler(mtype, f)
	require.True(t, (err == nil))

	err = mUnderlay.Send(msg)
	require.True(t, (err == nil))

	// the handler is called
	require.True(t, flag)
}
