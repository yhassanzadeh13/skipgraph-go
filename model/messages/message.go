package messages

// MessageType determines the type of an underlay network message
type MessageType string

// Message is an underlay network message
type Message struct {
	// Type indicates the type of a message
	Type MessageType
	// Payload denotes the content of a message
	Payload interface{}
}
