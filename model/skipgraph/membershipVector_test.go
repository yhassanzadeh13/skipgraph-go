package skipgraph_test

import (
	"github.com/stretchr/testify/require"
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
	"testing"
)

// TestMembershipVectorCompare tests CommonPrefix method.
func TestMembershipVectorCompare(t *testing.T) {
	// create two membershipVectors with 32 * 8 common prefix
	v1 := skipgraph.MembershipVector{0}
	res := v1.CommonPrefix(v1)
	require.Equal(t, 256, res)

	// create two membershipVectors with no common prefix
	v2 := skipgraph.MembershipVector{0}
	v2[0] = 255
	res = v1.CommonPrefix(v2)
	require.Equal(t, 0, res)

	// create two membershipVectors with non-zero common prefix
	v1[0] = 253
	res = v1.CommonPrefix(v2)
	require.Equal(t, 6, res)
}

// TestToBinaryString tests correctness of ToBinaryString.
func TestToBinaryString(t *testing.T) {
	v1 := byte(1) // 00000001
	s1 := skipgraph.ToBinaryString(v1)
	require.Equal(t, "00000001", s1)

	v2 := byte(2) // 00000010
	s2 := skipgraph.ToBinaryString(v2)
	require.Equal(t, "00000010", s2)

	v3 := byte(128) // 10000000
	s3 := skipgraph.ToBinaryString(v3)
	require.Equal(t, "10000000", s3)

	v4 := byte(65) // 01000001
	s4 := skipgraph.ToBinaryString(v4)
	require.Equal(t, "01000001", s4)
}
