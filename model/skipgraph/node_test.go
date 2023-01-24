package skipgraph_test

import (
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"github/yhassanzadeh13/skipgraph-go/unittest"
	"testing"
)

func TestLookUpQueryToJSON(t *testing.T) {
	// create a LookupQuery and covert it to JSON
	lq := skipgraph.LookupQuery{unittest.IdentityFixture(t), skipgraph.EntryPosition{skipgraph.LeftDirection, unittest.LevelFixture(t)}}
 	jsonStr, err := skipgraph.LookUpQueryToJSON(lq)
	 require.NoError(t, err)
	 require.NotEmpty(t, jsonStr)

	 // convert the JSON back to a LookupQuery
	 lq2, err := skipgraph.JSONToLookUPQuery(jsonStr)
	 require.NoError(t, err)
	 require.Equal(t, lq, lq2)
}
