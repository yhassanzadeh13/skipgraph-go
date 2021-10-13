package mocknet

import (
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/network"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnderlay(t *testing.T) {
	mUnderlay := newMockUnderlay(make(map[messages.MessageType]network.MessageHandler))
	mUnderlay.Start()
	defer mUnderlay.Stop()
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
