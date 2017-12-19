package text

import (
	"encoding/json"
	"math"
	"time"

	"github.com/autopilothq/lg/encoding/buffer"
	"github.com/autopilothq/lg/encoding/types"
)

// Encoder can encode Go types to JSON
type Encoder struct {
	buf *buffer.Buffer
}

// NewEncoder returns a new Encoder
func NewEncoder() *Encoder {
	return &Encoder{
		buf: buffer.GetBuffer(),
	}
}

// String returns the encoded buffer as a string
func (e *Encoder) String() string {
	return e.buf.String()
}

// Bytes returns the raw encoded bytes
func (e *Encoder) Bytes() []byte {
	return e.buf.Bytes()
}

// AddKey appends the desired key to the buffer
func (e *Encoder) AddKey(key string) error {
	e.addSeparator()
	e.buf.AppendString(key)
	e.buf.AppendByte(':')
	return nil
}

// AddUint16 appends a uint16 to the buffer
func (e *Encoder) AddUint16(val uint16) error {
	e.addSeparator()
	e.buf.AppendUint16(val)
	return nil
}

// AddUint32 appends a uint32 to the buffer
func (e *Encoder) AddUint32(val uint32) error {
	e.addSeparator()
	e.buf.AppendUint32(val)
	return nil
}

// AddUint64 appends a uint64 to the buffer
func (e *Encoder) AddUint64(val uint64) error {
	e.addSeparator()
	e.buf.AppendUint(val)
	return nil
}

// AddInt16 appends a int16 to the buffer
func (e *Encoder) AddInt16(val int16) error {
	e.addSeparator()
	e.buf.AppendInt16(val)
	return nil
}

// AddInt32 appends a int32 to the buffer
func (e *Encoder) AddInt32(val int32) error {
	e.addSeparator()
	e.buf.AppendInt32(val)
	return nil
}

// AddInt64 appends a int64 to the buffer
func (e *Encoder) AddInt64(val int64) error {
	e.addSeparator()
	e.buf.AppendInt(val)
	return nil
}

// AddFloat32 appends a float32 to the buffer
func (e *Encoder) AddFloat32(val float32) error {
	return e.AddFloat(float64(val), 32)
}

// AddFloat64 appends a float64 to the buffer
func (e *Encoder) AddFloat64(val float64) error {
	return e.AddFloat(val, 64)
}

// AddFloat appends a float64 to the buffer
func (e *Encoder) AddFloat(val float64, bitsize int) error {
	e.addSeparator()

	switch {
	case math.IsNaN(val):
		e.buf.AppendString(`NaN`)
	case math.IsInf(val, 1):
		e.buf.AppendString(`+Inf`)
	case math.IsInf(val, -1):
		e.buf.AppendString(`-Inf`)
	default:
		e.buf.AppendFloat(val, bitsize)
	}
	return nil
}

// AddBool appends the boolean value to the buffer
func (e *Encoder) AddBool(val bool) error {
	e.addSeparator()
	e.buf.AppendBool(val)
	return nil
}

// AddByteString appends the byte string value to the buffer
func (e *Encoder) AddByteString(val string) error {
	e.addSeparator()
	e.buf.AppendString(val)
	return nil
}

// AddBytes appends the byte array value to the buffer
func (e *Encoder) AddBytes(val []byte) error {
	e.addSeparator()
	_, err := e.buf.Write(val)
	return err
}

// AddDuration appends a Duration to the buffer
func (e *Encoder) AddDuration(val time.Duration) error {
	e.addSeparator()
	e.buf.AppendInt(int64(val))
	return nil
}

// AddTime appends a Time to the buffer
func (e *Encoder) AddTime(t time.Time) error {
	e.addSeparator()
	// TODO a unix timestamp possibly isn't the most conveienent format
	e.buf.AppendInt(t.UnixNano())
	return nil
}

// AddTimestamp appends a Timestamp to the buffer
func (e *Encoder) AddTimestamp(t time.Time) error {
	e.addSeparator()
	// TODO a unix timestamp possibly isn't the most conveienent format
	e.buf.AppendInt(t.UnixNano())
	return nil
}

// AddNull appends a null value to the buffer
func (e *Encoder) AddNull() error {
	e.addSeparator()
	e.buf.AppendString("null")
	return nil
}

// AddArrayish appends an array value to the buffer
func (e *Encoder) AddArrayish(arr types.Array) error {
	e.addSeparator()
	e.buf.AppendByte('[')
	arr.MarshalArray(e)
	e.buf.AppendByte(']')
	return nil
}

// StartArray appends the necessary bytes to begin a json array
func (e *Encoder) StartArray() error {
	e.addSeparator()
	e.buf.AppendByte('[')
	return nil
}

// EndArray appends the necessary bytes to end a json array
func (e *Encoder) EndArray() error {
	e.buf.AppendByte(']')
	return nil
}

// StartObject appends the necessary bytes to begin a json object
func (e *Encoder) StartObject() error {
	e.addSeparator()
	e.buf.AppendByte('{')
	return nil
}

// EndObject appends the necessary bytes to end a json object
func (e *Encoder) EndObject() error {
	e.buf.AppendByte('}')
	return nil
}

// AddReflected encodes the value via json.Marshal and appends it to the buffer
func (e *Encoder) AddReflected(val interface{}) error {
	marshaledVal, err := json.Marshal(val)
	if err != nil {
		return err
	}

	e.addSeparator()
	_, err = e.buf.Write(marshaledVal)
	return err
}

func (e *Encoder) addSeparator() {
	last := e.buf.Len() - 1
	if last < 0 {
		return
	}

	switch e.buf.Bytes()[last] {
	case '{', '[', ':', ',', ' ':
		return

	default:
		e.buf.AppendByte(' ')
	}
}

// AddString adds a string to the encoded buffer
func (e *Encoder) AddString(s string) error {
	e.addSeparator()
	e.buf.AppendString(s)
	return nil
}
