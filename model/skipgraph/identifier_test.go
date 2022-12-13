package skipgraph

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestCompare checks the correctness of the Identifier comparison function
func TestCompare(t *testing.T) {
	s1 := []byte("12")
	s2 := []byte("22")
	i1, _ := ToIdentifier(s1)
	i2, _ := ToIdentifier(s2)

	require.Equal(t, CompareLess, i1.compare(i2))
	require.Equal(t, CompareGreater, i2.compare(i1))
	require.Equal(t, CompareEqaul, i1.compare(i1))
}
