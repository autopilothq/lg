package types

// Array is a JSON marshalable array
type Array interface {
	MarshalArray(enc Encoder) error
}
