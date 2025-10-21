package serde

/*
Deserializer defines the interface for deserializing data from a specific format.

It includes methods for correctly deserializing data for a specified format (e.g., JSON, XML) into a Go struct.
*/
type Deserializer[Target any] interface {
	// Deserialize deserializes the given byte slice into the provided target struct in a type-safe manner.
	// The target struct must be of the type specified by the generic parameter Target.
	// If the data cannot be deserialized into the target struct, an error is returned.
	//
	// Example usage:
	//
	//	var myStruct MyStruct
	//	err := deserializer.Deserialize(data, &myStruct)
	//	if err != nil {
	//	    // handle error
	//	}
	//
	// Note: The caller is responsible for ensuring that the target struct is of the correct type.
	Deserialize(data []byte) (*Target, error)
}
