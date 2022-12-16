package skipgraph_test

import (
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"testing"
)

func TestLookupTable_AddEntry(t *testing.T) {
	// create an empty lookup table
	lt := skipgraph.LookupTable{}

	// create an empty identity
	identity := skipgraph.Identity{}

	// add the identity in a valid position
	err := lt.AddEntry(skipgraph.LeftDirection, 0, identity)
	require.NoError(t, err)

	// add the identity in a valid position
	err = lt.AddEntry(skipgraph.LeftDirection, skipgraph.MaxLookupTableLevel-1, identity)
	require.NoError(t, err)

	// add an entry with invalid level
	err = lt.AddEntry(skipgraph.LeftDirection, skipgraph.MaxLookupTableLevel, identity)
	require.Error(t, err)

	// add an entry with wrong direction
	err = lt.AddEntry(skipgraph.Direction("no where"), skipgraph.MaxLookupTableLevel, identity)
	require.Error(t, err)
}
