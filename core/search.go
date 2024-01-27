package core

import (
	"github/yhassanzadeh13/skipgraph-go/model/skipgraph"
)

type SearchRequest struct {
	targetId    skipgraph.Identifier
	myId        skipgraph.Identity
	lookupTable skipgraph.LookupTable
	direction   skipgraph.Direction
	level       skipgraph.Level
}

type SearchResponse struct {
	result skipgraph.Identity
}

// SearchByIdentifier implements the local part of the search protocol (by identifier, aka by numerical id) at the local
// node. It returns the result of the search.
// Arguments:
//
//	req: a SearchRequest struct containing the target identifier.
//
// Returns:
//
//	a SearchResponse struct containing the result of the search.
//	error: if any error occurs during the search.
func SearchByIdentifier(req *SearchRequest) (*SearchResponse, error) {
	// return fast if the target id is equal to my id.
	if req.targetId.Compare(req.myId.GetIdentifier()) == skipgraph.CompareEqual {
		return &SearchResponse{result: req.myId}, nil
	}

	// we set valid comparison to my id, because if we reach the bottom of the lookup table and no valid comparison is
	// found, we return the last valid comparison.
	validComparison := req.myId

	// traverse the lookup table down till a valid comparison is found
	// valid comparison means we found a neighbor that is less than my id but greater than target id, when dir is left
	// or greater than my id but less than target id, when dir is right.
	// if we reach the bottom of the lookup table and no valid comparison is found, return the last valid comparison.
	
}
