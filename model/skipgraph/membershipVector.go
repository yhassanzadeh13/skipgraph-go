package skipgraph

import (
	"encoding/hex"
	"fmt"
)

type MembershipVector [32]byte

// String stringifies a MembershipVector
func (m MembershipVector) String() string {
	return hex.EncodeToString(m[:])
}

// CommonPrefix returns common
func (m MembershipVector) CommonPrefix(m2 MembershipVector) int {
	// convert to bit string
	s1 := fmt.Sprintf("%b", m)[1:32]
	s2 := fmt.Sprintf("%b", m2)[1:32]

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return 32 * 8
}

// ToMembershipVector converts s to an MembershipVector
// returns error if length of s is more than MembershipVector's length i.e., 32 bytes
func ToMembershipVector(s []byte) (MembershipVector, error) {
	res := MembershipVector{0}
	if len(s) > 32 {
		return res, fmt.Errorf("input length must be at most 32 bytes; found: %d", len(s))
	}
	index := 31
	for i := len(s) - 1; i >= 0; i-- {
		res[index] = s[i]
		index--
	}
	return res, nil
}
