package skipgraph

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

const IdentifierSize = 32
const CompareEqaul = "compare equal"
const CompareGreater = "compare greater"
const CompareLess = "compare less"

// Identifier represents a 32-byte unique identifier a Skip Graph node.
type Identifier [IdentifierSize]byte

type IdentifierList []Identifier

func (i Identifier) String() string {
	return hex.EncodeToString(i[:])
}

// Compare
func (i Identifier) Compare(other Identifier) string {
	cmp := bytes.Compare(i[:], other[:])
	switch cmp {
	case 1:
		return CompareGreater
	case -1:
		return CompareLess
	default:
		return CompareEqaul
	}
}

// ToIdentifier converts s to an Identifier
// returns error if length of s is more than Identifier's length i.e., 32 bytes
func ToIdentifier(s []byte) (Identifier, error) {
	res := Identifier{0}
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
