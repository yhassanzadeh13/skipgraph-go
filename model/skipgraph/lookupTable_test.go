package skipgraph_test

import (
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"github/yhassanzadeh13/skipgraph-go/unittest"
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
	err = lt.AddEntry(skipgraph.Direction("no where"), 0, identity)
	require.Error(t, err)
}

func TestLookupTable_GetEntry(t *testing.T) {
	// create an identity
	identity := unittest.IdentityFixture(t)
	identity1 := unittest.IdentityFixture(t)
	require.NotEqual(t, identity1, identity)

	// declare an empty lookup table
	var lt skipgraph.LookupTable

	// add the identity as a left neighbor into the lookup table
	err := lt.AddEntry(skipgraph.LeftDirection, 0, identity)
	require.NoError(t, err)

	// add the identity as a right neighbor into the lookup table
	err = lt.AddEntry(skipgraph.RightDirection, 0, identity1)
	require.NoError(t, err)

	// check that the inserted identity is retrievable
	retIdentity, err := lt.GetEntry(skipgraph.LeftDirection, 0)
	require.Equal(t, identity, retIdentity)

	// check that the inserted identity is retrievable
	retIdentity1, err := lt.GetEntry(skipgraph.RightDirection, 0)
	require.Equal(t, identity1, retIdentity1)

	// access a wrong level
	_, err = lt.GetEntry(skipgraph.RightDirection, skipgraph.MaxLookupTableLevel)
	require.Error(t, err)

	// access a wrong direction
	_, err = lt.GetEntry(skipgraph.Direction("no where"), 0)
	require.Error(t, err)

}
