package types

import (
	"time"
)

// Bools is a collection of bools
type Bools []bool

// MarshalArray converts the collection to json
func (b Bools) MarshalArray(encoder Encoder) error {
	for _, v := range b {
		if err := encoder.AddBool(v); err != nil {
			return err
		}
	}

	return nil
}

// Bytes is a collection of bytes
type Bytes []byte

// MarshalArray converts the collection to json
func (b Bytes) MarshalArray(encoder Encoder) error {
	return encoder.AddBytes(b)
}

// Uint16s is a collection of uint16
type Uint16s []uint16

// MarshalArray converts the collection to json
func (u Uint16s) MarshalArray(encoder Encoder) error {
	for _, v := range u {
		if err := encoder.AddUint16(v); err != nil {
			return err
		}
	}

	return nil
}

// Uint32s is a collection of uint32
type Uint32s []uint32

// MarshalArray converts the collection to json
func (u Uint32s) MarshalArray(encoder Encoder) error {
	for _, v := range u {
		if err := encoder.AddUint32(v); err != nil {
			return err
		}
	}

	return nil
}

// Uint64s is a collection of uint64
type Uint64s []uint64

// MarshalArray converts the collection to json
func (u Uint64s) MarshalArray(encoder Encoder) error {
	for _, v := range u {
		if err := encoder.AddUint64(v); err != nil {
			return err
		}
	}

	return nil
}

// Int16s is a collection of int16
type Int16s []int16

// MarshalArray converts the collection to json
func (u Int16s) MarshalArray(encoder Encoder) error {
	for _, v := range u {
		if err := encoder.AddInt16(v); err != nil {
			return err
		}
	}

	return nil
}

// Int32s is a collection of int32
type Int32s []int32

// MarshalArray converts the collection to json
func (i Int32s) MarshalArray(encoder Encoder) error {
	for _, v := range i {
		if err := encoder.AddInt32(v); err != nil {
			return err
		}
	}

	return nil
}

// Int64s is a collection of int64
type Int64s []int64

// MarshalArray converts the collection to json
func (i Int64s) MarshalArray(encoder Encoder) error {
	for _, v := range i {
		if err := encoder.AddInt64(v); err != nil {
			return err
		}
	}

	return nil
}

// Strings is a collection of string
type Strings []string

// MarshalArray converts the collection to json
func (s Strings) MarshalArray(encoder Encoder) error {
	for _, v := range s {
		if err := encoder.AddString(v); err != nil {
			return err
		}
	}

	return nil
}

// Float32s is a collection of float32
type Float32s []float32

// MarshalArray converts the collection to json
func (f Float32s) MarshalArray(encoder Encoder) error {
	for _, v := range f {
		if err := encoder.AddFloat32(v); err != nil {
			return err
		}
	}

	return nil
}

// Float64s is a collection of float64
type Float64s []float64

// MarshalArray converts the collection to json
func (f Float64s) MarshalArray(encoder Encoder) error {
	for _, v := range f {
		if err := encoder.AddFloat64(v); err != nil {
			return err
		}
	}

	return nil
}

// Durations is a collection of time.Duration
type Durations []time.Duration

// MarshalArray converts the collection to json
func (f Durations) MarshalArray(encoder Encoder) error {
	for _, v := range f {
		if err := encoder.AddDuration(v); err != nil {
			return err
		}
	}

	return nil
}

// Times is a collection of time.Time
type Times []time.Time

// MarshalArray converts the collection to json
func (f Times) MarshalArray(encoder Encoder) error {
	for _, v := range f {
		if err := encoder.AddTime(v); err != nil {
			return err
		}
	}

	return nil
}
