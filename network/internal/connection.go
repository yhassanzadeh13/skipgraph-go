package internal

// Connection represents a connection to a remote peer.
type Connection interface {
	// RemoteAddr returns the remote address of the connection.
	// It returns an empty string if the connection is closed, otherwise it returns the remote address,
	// which is the address of the peer that the connection is connected to.
	RemoteAddr() string

	// Send sends a message to the remote peer, returning an error if the message could not be sent.
	// Send is a blocking operation, and it will block until the message is sent.
	// It returns an error if the message could not be sent.
	// It returns io.EOF if the connection is closed.
	Send([]byte) error

	// Next returns the next message received from the remote peer.
	// It is a blocking operation, and it will block until a message is received.
	// It returns io.EOF if the connection is closed.
	// It returns an error if the message could not be read.
	Next() ([]byte, error)

	// Close gracefully closes the connection. Blocking until the connection is closed.
	Close() error
}
