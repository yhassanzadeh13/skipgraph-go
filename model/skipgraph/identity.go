package skipgraph

import "github/yhassanzadeh13/skipgraph-go/model"

type Identity struct {
	id        Identifier       // corresponds to numerical id in traditional skip graph
	memVector MembershipVector // corresponds to name id in traditional skip graph
	addr      model.Address    // holds network address like IP
}

// getIdentifier returns the Identifier field.
func (i Identity) getIdentifier() Identifier {
	return i.id
}

// GetMembershipVector returns the MembershipVector field.
func (i Identity) GetMembershipVector() MembershipVector {
	return i.memVector
}

// GetAddress returns the Address field.
func (i Identity) GetAddress() model.Address {
	return i.addr
}

// SetId sets Identifier
func (i *Identity) SetId(id Identifier) {
	// TODO validation of the id may be needed
	i.id = id
}

// SetMemVector sets membershipVector
func (i *Identity) SetMemVector(mv MembershipVector) {
	// TODO validation of the id may be needed
	i.memVector = mv
}

// SetAddr sets address
func (i *Identity) SetAddr(addr model.Address) {
	// TODO validation of the id may be needed
	i.addr = addr
}

func NewIdentity(id Identifier, mv MembershipVector, addr model.Address) Identity {
	i := Identity{}
	i.SetMemVector(mv)
	i.SetAddr(addr)
	i.SetId(id)
	return i
}
