package skipgraph

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func to32ByteArray(t *testing.T, s []byte) [32]byte {
	require.Less(t, len(s), 33)
	res := [32]byte{0}
	index := 31
	for i := len(s) - 1; i >= 0; i-- {
		res[index] = s[i]
		index--
	}
	return res
}

func TestCompare(t *testing.T) {
	s1 := []byte("12")
	s2 := []byte("22")
	b1 := to32ByteArray(t, s1)
	b2 := to32ByteArray(t, s2)
	i1 := Identifier(b1)
	i2 := Identifier(b2)

	require.Equal(t, CompareLess, i1.compare(i2))
	require.Equal(t, CompareGreater, i2.compare(i1))
	require.Equal(t, CompareEqaul, i1.compare(i1))

}
