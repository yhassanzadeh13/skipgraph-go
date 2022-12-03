package messages

// Type determines the type of an underlay network message
type Type string

// Message is an underlay network message
type Message struct {
	// Type indicates type of the0 message
	Type Type
	// Payload denotes the content of a message
	Payload interface{}
}
