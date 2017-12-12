package json

import (
	"github.com/pkg/errors"
)

// Marshalable is something that can be marshaled to bytes.
type Marshalable interface {
	Marshal() (data []byte, err error)
}

// Unmarshalable is something that can be unmarshaled from bytes
type Unmarshalable interface {
	Unmarshal(data []byte) error
}

// Marshaler is something that can be marshaled to and from bytes.
type Marshaler interface {
	Marshalable
	Unmarshalable
}

const (
	// These constants are the unique identifiers of the various
	// types of value that can be encoded by KV.
	//
	// Note: order is important. Changing it will cause existing
	// data to be undecodable. Also, False and True must always
	// be the 0 and 1 types respectively.

	// TypeFalse indicates that the stored type is a false value
	TypeFalse byte = iota
	// TypeTrue indicates that the stored type is a true value
	TypeTrue

	// TypeMarshalable indicates that the stored type is a value that
	// implements the Marhaller (Marhalable and Unmarshable) interface.
	TypeMarshalable

	// TypeNumerical indicates that the stored type is a numerical literal value
	TypeNumerical

	// TypeBytes indicates that the stored type is a byte array ([]byte) literal
	TypeBytes

	// TypeString indicates that the stored type is a string literal
	TypeString

	// TypeBools indicates that the stored type is a []bool literal
	TypeBools

	// TypeNumericalArray indicates that the stored type is a numerical array
	TypeNumericalArray
)

var (
	// ErrNoNilForYou indicates that the caller attempted to write nil to a value,
	// which we don't allow. Instead of writing nil the key should be deleted.
	ErrNoNilForYou = errors.New(
		"nil cannot be written to a key, delete the key instead")

	// ErrUnencodableValue indicates that the caller attemped to write a value
	// that KV didn't know how to marshal. If you do need to encode an unexpected
	// type you should implement the Marshalable interface for it
	ErrUnencodableValue = errors.New("Did not know how to marshal the given " +
		"value, you may wish to implement the Marshable interface for it")

	// ErrUndecodableValue indicates that the caller attemped to read a value
	// that KV didn't know how to unmarshal. If you do need to decode an
	// unexpected type you should implement the Marshalable interface for it
	ErrUndecodableValue = errors.New("Did not know how to unmarshal the given " +
		"value, you may wish to implement the Unmarshalable interface for it")
)
