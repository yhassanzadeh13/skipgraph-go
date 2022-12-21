package skipgraph

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

const IdentifierSize = 32
const CompareEqual = "compare-equal"
const CompareGreater = "compare-greater"
const CompareLess = "compare-less"

// Identifier represents a 32-byte unique identifier a Skip Graph node.
type Identifier [IdentifierSize]byte

// IdentifierList is a slice of Identifier
type IdentifierList []Identifier

// String converts Identifier to its hex representation.
func (i Identifier) String() string {
	return hex.EncodeToString(i[:])
}

// Compare compares two Identifiers and returns 0 if equal, 1 if other > i and -1 if other < i.
func (i Identifier) Compare(other Identifier) string {
	cmp := bytes.Compare(i[:], other[:])
	switch cmp {
	case 1:
		return CompareGreater
	case -1:
		return CompareLess
	default:
		return CompareEqual
	}
}

// ToIdentifier converts a byte slice b to an Identifier.
// returns error if the length of b is more than Identifier's length i.e., 32 bytes.
func ToIdentifier(b []byte) (Identifier, error) {
	res := Identifier{0}
	if len(b) > 32 {
		return res, fmt.Errorf("input length must be at most 32 bytes; found: %d", len(b))
	}
	index := 31
	for i := len(b) - 1; i >= 0; i-- {
		res[index] = b[i]
		index--
	}
	return res, nil
}

// StringToIdentifier converts a string to an Identifier.
// returns error if the byte length of the string s is more than Identifier's length i.e., 32 bytes.
func StringToIdentifier(s string) (Identifier, error) {
	b := []byte(s)
	res := Identifier{0}
	if len(b) > 32 {
		return res, fmt.Errorf("input length must be at most 32 bytes; found: %d", len(b))
	}
	index := 31
	for i := len(b) - 1; i >= 0; i-- {
		res[index] = b[i]
		index--
	}
	return res, nil
}
