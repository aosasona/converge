package serde

/*
Serializer defines the interface for serializing data into a specific format.

It includes methods for correctly serializing data, and retrieving the appropriate mime type to use in HTTP headers.
*/
type Serializer interface {
	// Serialize serializes the given value into a byte slice.
	Serialize(v interface{}) ([]byte, error)

	// ContentType returns the mime type for the serialized data.
	ContentType() string
}
