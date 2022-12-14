package skipgraph_test

import (
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"testing"
)

// TestCompare checks the correctness of the Identifier comparison function
func TestCompare(t *testing.T) {
	s1 := []byte("12")
	s2 := []byte("22")
	i1, err := skipgraph.ToIdentifier(s1)
	require.NoError(t, err)
	i2, err := skipgraph.ToIdentifier(s2)
	require.NoError(t, err)

	require.Equal(t, skipgraph.CompareLess, i1.Compare(i2))
	require.Equal(t, skipgraph.CompareGreater, i2.Compare(i1))
	require.Equal(t, skipgraph.CompareEqaul, i1.Compare(i1))
}
