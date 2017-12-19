package encoding

import (
	"time"

	"github.com/autopilothq/lg/encoding/json"
)

type UintEncoder interface {
	// AddUint16 appends a uint16 to the buffer
	AddUint16(val uint16) error
	// AddUint32 appends a uint32 to the buffer
	AddUint32(val uint32) error
	// AddUint64 appends a uint64 to the buffer
	AddUint64(val uint64) error
}

type IntEncoder interface {
	// AddInt16 appends a int16 to the buffer
	AddInt16(val int16) error
	// AddInt32 appends a int32 to the buffer
	AddInt32(val int32) error
	// AddInt64 appends a int64 to the buffer
	AddInt64(val int64) error
}

type FloatEncoder interface {
	// AddFloat32 appends a float32 to the buffer
	AddFloat32(val float32) error
	// AddFloat64 appends a float64 to the buffer
	AddFloat64(val float64) error
	// AddFloat appends a float64 to the buffer
	AddFloat(val float64, bitsize int) error
}

type BoolEncoder interface {
	// AddBool appends the boolean value to the buffer
	AddBool(val bool) error
}

type ByteEncoder interface {
	// AddByteString appends the byte string value to the buffer
	AddByteString(val string) error
	// AddBytes appends the byte array value to the buffer
	AddBytes(val []byte) error
}

type TimeEncoder interface {
	// AddDuration appends a Duration to the buffer
	AddDuration(val time.Duration) error
	// AddTime appends a Time to the buffer
	AddTime(t time.Time) error
}

type ArrayEncoder interface {
	// AddArrayish appends an array value to the buffer
	AddArrayish(arr json.Array) error

	// StartArray appends the necessary bytes to begin a json array
	StartArray() error

	// EndArray appends the necessary bytes to end a json array
	EndArray() error

	// StartObject appends the necessary bytes to begin a json object
	StartObject() error

	// EndObject appends the necessary bytes to end a json object
	EndObject() error
}

type NullEncoder interface {
	// AddNull appends a null value to the buffer
	AddNull() error
}

type ReflectionEncoder interface {
	// AddReflected encodes the value via json.Marshal and appends it to the buffer
	AddReflected(val interface{}) error
}

type StringEncoder interface {
	// AddString adds a string to the encoded buffer
	AddString(s string) error
}

type Encoder interface {
	// String returns the encoded buffer as a string
	String() string
	// Bytes returns the raw encoded bytes
	Bytes() []byte
	// AddKey appends the desired key to the buffer
	AddKey(key string) error

	UintEncoder
	IntEncoder
	FloatEncoder
	BoolEncoder
	ByteEncoder
	TimeEncoder
	ArrayEncoder
	NullEncoder
	ReflectionEncoder
	StringEncoder
}
