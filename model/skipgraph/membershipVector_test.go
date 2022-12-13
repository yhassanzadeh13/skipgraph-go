package skipgraph

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMembershipVectorCompare(t *testing.T) {
	// create two membershipVectors with 32 * 8 common prefix
	v1 := MembershipVector{0}
	res := v1.CommonPrefix(v1)
	require.Equal(t, 256, res)

	// create two membershipVectors with no common prefix
	v2 := MembershipVector{0}
	v2[0] = 255
	res = v1.CommonPrefix(v2)
	require.Equal(t, 0, res)

	// create two membershipVectors with non-zero common prefix
	v1[0] = 254
	res = v1.CommonPrefix(v2)
	require.Equal(t, 7, res)

}
