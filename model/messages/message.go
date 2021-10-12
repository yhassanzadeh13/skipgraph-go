package messages

type MessageType string
type Message struct {
	Type    MessageType
	Payload interface{}
}
