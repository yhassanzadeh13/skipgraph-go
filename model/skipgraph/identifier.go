package skipgraph

import (
	"bytes"
	"encoding/hex"
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

// compare
func (i Identifier) compare(other Identifier) string {
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
