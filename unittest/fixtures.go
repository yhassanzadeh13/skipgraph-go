package unittest

import (
	"crypto/rand"
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"testing"
)

const TestMessageType = messages.Type("test-message")

func TestMessageFixture(t *testing.T) *messages.Message{
	return &messages.Message{
		Type:   TestMessageType,
		Payload: RandomBytesFixture(t, 100),
	}
}

func IdentifierFixture(t *testing.T) skipgraph.Identifier {
	var id skipgraph.Identifier
	bytes := RandomBytesFixture(t, skipgraph.IdentifierSize)

	for i := 0; i < skipgraph.IdentifierSize; i++{
		id[i] = bytes[i]
	}

	return id
}

func RandomBytesFixture(t *testing.T, size int) []byte {
	bytes := make([]byte, size)
	n, err := rand.Read(bytes[:])

	require.Equal(t, size, n)
	require.NoError(t, err)
	require.Len(t, bytes, size)

	return bytes
}
