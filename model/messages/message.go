package messages

// Type determines the type of an underlay network message
type Type string

const (
	// GetEntry is a message type for requesting an entry from a node.
	GetEntry = Type("Get Entry")
	// AddEntry is a message type for adding an entry to a node.
	AddEntry = Type("Add Entry")
)

// Message is an underlay network message
type Message struct {
	// Type indicates type of the0 message
	Type Type
	// Payload denotes the content of a message
	Payload interface{}
}
