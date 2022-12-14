package skipgraph

import (
	"encoding/hex"
	"fmt"
)

type MembershipVector [32]byte

// String returns hex encoding of a MembershipVector
func (m MembershipVector) String() string {
	return hex.EncodeToString(m[:])
}

// ToBinaryString returns binary representation of a MembershipVector
func (m MembershipVector) ToBinaryString() string {
	var s string
	for i := 0; i < len(m); i++ {
		s = s + ToBinaryString(m[i])
	}
	return s
}

// ToBinaryString returns binary representation of a byte value
func ToBinaryString(b byte) string {
	var s string
	for j := 0; j < 8; j++ {
		// m[i] is an 8 bit value i.e., x0 x1 x2 ... x7
		v := b >> (7 - j)   // v = 0 ... x0 x1 ... xj-1 xj (cuts the last 7-j bits to shift the jth bit to the least significant bit)
		b := v & 0b00000001 // get the value of the least significant bit (which is xj)
		if b == 1 {
			s = s + "1"
		} else {
			s = s + "0"
		}
	}
	return s
}

// CommonPrefix returns the longest common bit prefix of the supplied MembershipVectors
func (m MembershipVector) CommonPrefix(m2 MembershipVector) int {
	// convert to bit string
	s1 := m.ToBinaryString()
	s2 := m2.ToBinaryString()

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return 32 * 8
}

// ToMembershipVector converts a byte slice to a MembershipVector
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
