package mocknet

import (
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnderlay(t *testing.T) {
	// construct an empty mocked underlay
	mUnderlay := NewMockUnderlay()

	//start
	mUnderlay.Start()

	// stop when the test terminates
	defer mUnderlay.Stop()

	// create a message type
	mtype := messages.MessageType("input")
	var payload interface{}
	var msg = messages.Message{Type: mtype,
		Payload: payload}
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
