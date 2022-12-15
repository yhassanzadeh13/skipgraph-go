package skipgraph

import "github/yhassanzadeh13/skipgraph-go/model"

type Identity struct {
	id        Identifier       // corresponds to numerical id in traditional skip graph
	memVector MembershipVector // corresponds to name id in traditional skip graph
	addr      model.Address    // holds network address like IP
}

// getIdentifier returns the Identifier field.
func (id Identity) getIdentifier() Identifier {
	return id.id
}

// GetMembershipVector returns the MembershipVector field.
func (id Identity) GetMembershipVector() MembershipVector {
	return id.memVector
}

// GetAddress returns the Address field.
func (id Identity) GetAddress() model.Address {
	return id.addr
}
