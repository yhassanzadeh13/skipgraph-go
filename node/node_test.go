package node_test

import (
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"github/yhassanzadeh13/skipgraph-go/node"
	"github/yhassanzadeh13/skipgraph-go/unittest"
	"testing"
)

func TestLookUpQueryToJSON(t *testing.T) {
	// create a LookupQuery and covert it to JSON
	lq := node.LookupQuery{unittest.IdentityFixture(t), node.EntryPosition{skipgraph.LeftDirection, unittest.LevelFixture(t)}}
 	jsonStr, err := node.LookUpQueryToJSON(lq)
	 require.NoError(t, err)
	 require.NotEmpty(t, jsonStr)

	 // convert the JSON back to a LookupQuery
	 lq2, err := node.JSONToLookUPQuery(jsonStr)
	 require.NoError(t, err)
	 require.Equal(t, lq, lq2)
}
