package skipgraph_test

import (
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model"
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

func TestLookupTable_GetEntry(t *testing.T) {
	// create an identity
	id, err := skipgraph.StringToIdentifier("123")
	require.NoError(t, err)
	memVec, err := skipgraph.StringToMembershipVector("456")
	require.NoError(t, err)
	addr := model.NewAddress("localhost", "1234")
	identity := skipgraph.NewIdentity(id, memVec, addr)

	// create another identity
	id1, err := skipgraph.StringToIdentifier("111")
	require.NoError(t, err)
	memVec1, err := skipgraph.StringToMembershipVector("222")
	require.NoError(t, err)
	addr1 := model.NewAddress("localhost1", "5678")
	identity1 := skipgraph.NewIdentity(id1, memVec1, addr1)

	// create an empty lookup table
	lt := skipgraph.LookupTable{}

	// add the identity as a left neighbor into the lookup table
	err = lt.AddEntry(skipgraph.LeftDirection, 12, identity)
	require.NoError(t, err)

	// check that the inserted identity is retrievable
	retIdentity, err := lt.GetEntry(skipgraph.LeftDirection, 12)
	require.Equal(t, id, retIdentity.GetIdentifier())
	require.Equal(t, memVec, retIdentity.GetMembershipVector())
	require.Equal(t, addr, retIdentity.GetAddress())

	// add the identity as a left neighbor into the lookup table
	err = lt.AddEntry(skipgraph.RightDirection, 0, identity1)
	require.NoError(t, err)

	// check that the inserted identity is retrievable
	retIdentity1, err := lt.GetEntry(skipgraph.RightDirection, 0)
	require.Equal(t, id1, retIdentity1.GetIdentifier())
	require.Equal(t, memVec1, retIdentity1.GetMembershipVector())
	require.Equal(t, addr1, retIdentity1.GetAddress())

}
