package codec

import "io"

/*
Encoder defines the interface for serializing data into a specific format.

It includes methods for correctly serializing data, and retrieving the appropriate mime type to use in HTTP headers.
*/
type Encoder[T any] interface {
	// Encode serializes the given value into the provided io.Writer.
	Encode(data T, writer io.Writer) error

	// ContentType returns the mime type for the encoded data.
	ContentType() string
}
