package skipgraph

import "fmt"

// maxLookupTableSize indicates the upper bound for the number of levels in a SkipGraph LookupTable.
const maxLookupTableLevel = IdentifierSize * 8

type Direction string

const (
	RightDirection = Direction("right")
	LeftDirection  = Direction("left")
)

// LookupTable corresponds to a SkipGraph node's lookup table.
type LookupTable struct {
	rightNeighbors [maxLookupTableLevel]Identity
	leftNeighbors  [maxLookupTableLevel]Identity
}

// AddEntry inserts the supplied Identity in the lth level of lookup table either as the left or right neighbor depending on the dir.
// lev runs from 0...maxLookupTableLevel-1.
func (l *LookupTable) AddEntry(dir Direction, lev int64, ident Identity) error {
	// validate the level value
	if lev >= maxLookupTableLevel {
		return fmt.Errorf("position is larger than the max lookup table entry number: %d", l)
	}

	switch dir {
	case RightDirection:
		l.rightNeighbors[lev] = ident
	case LeftDirection:
		l.leftNeighbors[lev] = ident
	default:
		return fmt.Errorf("invalid direction: %s", dir)
	}

	return nil
}

// GetEntry returns the lth left/right neighbor in the lookup table depending on the dir.
// lev runs from 0...maxLookupTableLevel-1.
func (l LookupTable) GetEntry(dir Direction, lev int64) (Identity, error) {
	res := Identity{}

	// validate the level value
	if lev >= maxLookupTableLevel {
		return res, fmt.Errorf("supplied level is larger than the max number of levels: %d", lev)
	}
	switch dir {
	case RightDirection:
		res = l.rightNeighbors[lev]
	case LeftDirection:
		res = l.leftNeighbors[lev]
	default:
		return res, fmt.Errorf("invalid direction: %s", dir)
	}
	return res, nil
}
