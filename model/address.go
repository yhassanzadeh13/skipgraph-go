package model

import "fmt"

// Address contains network address information
type Address struct {
	hostName string
	port     string
}

// NewAddress initializes and returns an instance of Address with the supplied inputs
func NewAddress(hostname string, port string) Address {
	return Address{
		hostName: hostname,
		port:     port,
	}
}

// HostName returns the hostName
func (a Address) HostName() string {
	return a.hostName
}

// Port returns the port
func (a Address) Port() string {
	return a.port
}

// String stringifies an Address
func (a Address) String() string {
	s := fmt.Sprintf("host name: %s, port: %s", a.HostName(), a.Port())
	return s
}
