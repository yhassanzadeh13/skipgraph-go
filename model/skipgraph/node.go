package skipgraph

import (
	"encoding/json"
	"fmt"
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/network"
)

type node struct {
	// Identity is the identity of the node.
	Identity Identity
	// LookupTable is the lookup table of the node.
	LookupTable LookupTable
	// Underlay is the underlay network of the node.
	Underlay network.Underlay
}

// JSONtoIdentity converts a JSON string to an Identity.
func JSONtoIdentity(data string) (Identity, error) {
	// convert to byte slice
	b := []byte(data)
	// unmarshal byte slice to identity
	var i Identity
	err := json.Unmarshal(b, &i)
	if err != nil {
		return Identity{}, fmt.Errorf("could not parse to Identity %v", err)
	}
	return i, nil
}
func (n *node) addLookUpEntry(message messages.Message) error {
	// parse message payload to a lookup entry.
	id, error := JSONtoIdentity(message.Payload)

	// add the entry to the lookup table.

	return nil
}
func (n *node) Start() {
	// set up node's underlay.
	n.Underlay.SetMessageHandler(messages.AddEntry)

}
