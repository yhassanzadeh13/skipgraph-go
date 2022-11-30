# Skip Graph Middleware in Go

Skip Graph Middleware is the implementation of a SkipGraph node.
Each node is identified by a unique 32 bytes identifier.
Each node comprises two components, namely, 1) Overlay and 2) Underlay.
The overlay holds the logic for skip graph routing whereas the underlay provides network communication services between nodes.
The underlay exposes the necessary interface through which an overlay can communicate with other nodes in the network.
The overlay instructs the underlay to communicate with another node only by specifying the receiver's identifier.
Other network information such as IP address is handled by the underlay unit and is transparent to the overlay.





