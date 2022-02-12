package skipgraph

const IdentifierSize = 32

// Identifier represents a 32-byte unique identifier a Skip Graph node.
type Identifier [IdentifierSize]byte

type IdentifierList []Identifier
