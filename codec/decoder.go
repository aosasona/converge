package codec

import "io"

/*
Decoder defines the interface for deserializing data from a specific format.

It includes methods for correctly deserializing data for a specified format (e.g., JSON, XML) into a Go struct.
*/
type Decoder[T any] interface {
	// Decode deserializes data from the provided io.Reader into the target struct in a type-safe manner.
	Decode(reader io.Reader) (T, error)
}
