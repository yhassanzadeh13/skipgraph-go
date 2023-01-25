package node

import (
	"encoding/json"
	"fmt"
	"github/yhassanzadeh13/skipgraph-go/model/messages"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"github/yhassanzadeh13/skipgraph-go/network"
)

type node struct {
	// Identity is the identity of the node.
	Identity skipgraph.Identity
	// LookupTable is the lookup table of the node.
	LookupTable skipgraph.LookupTable
	// Underlay is the underlay network of the node.
	Underlay network.Underlay
}

// LookupQuery holds a node Identity and its metadata reflecting its location in a lookup table.
type LookupQuery struct {
	skipgraph.Identity
	EntryPosition
}
type EntryPosition struct {
	skipgraph.Direction
	skipgraph.Level
}

// JSONToLookUPQuery parses a JSON string to a LookupQuery struct.
// returns error if the string is not a valid JSON.
func JSONToLookUPQuery(data string) (LookupQuery, error) {
	// convert to byte slice
	b := []byte(data)
	// unmarshal byte slice to identity
	var i LookupQuery
	err := json.Unmarshal(b, &i)
	if err != nil {
		return LookupQuery{}, fmt.Errorf("could not parse to Entry Query %v", err)
	}
	return i, nil
}

// LookUpQueryToJSON converts a lookup query to a JSON string.
func LookUpQueryToJSON(i LookupQuery) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", fmt.Errorf("could not encode query %v", err)
	}
	s := string(b)
	return s, nil
}

// addLookUpEntry extracts the lookup query embedded in the message payload and adds the entry to the lookup table.
func (n *node) addLookUpEntry(message messages.Message) error {
	// parse message payload to string.
	qJsonStr, ok := message.Payload.(string)
	if !ok {
		return fmt.Errorf("message payload is not a string")
	}
	// parse the message payload as a lookup query.
	query, err := JSONToLookUPQuery(qJsonStr)
	if err != nil {
		return fmt.Errorf("could not parse payload to identity %s", qJsonStr)
	}
	// add the entry to the lookup table.
	n.LookupTable.AddEntry(query.Direction, query.Level, query.Identity)

	return nil
}

// Start activates the node.
func (n *node) Start() {
	// set up node's underlay.
	// add a handler for add lookup entry `AddEntry` message type.
	n.Underlay.SetMessageHandler(messages.AddEntry, n.addLookUpEntry)
	n.Underlay.Start()
}

// Stop deactivates the node.
func (n *node) Stop() {
	// stop the node's underlay.
	n.Underlay.Stop()
}
