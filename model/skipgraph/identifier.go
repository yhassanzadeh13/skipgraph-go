package skipgraph

import "encoding/hex"

const IdentifierSize = 32

// Identifier represents a 32-byte unique identifier a Skip Graph node.
type Identifier [IdentifierSize]byte

type IdentifierList []Identifier

func (i Identifier) String() string {
	return hex.EncodeToString(i[:])
}
