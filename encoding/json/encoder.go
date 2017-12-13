package json

import (
	"encoding/json"
	"math"
	"time"
	"unicode/utf8"
)

// Encoder can encode Go types to JSON
type Encoder struct {
	buf *Buffer
}

// NewEncoder returns a new Encoder
func NewEncoder() *Encoder {
	return &Encoder{
		buf: GetBuffer(),
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
	e.buf.AppendByte('"')
	e.buf.AppendString(key)
	e.buf.AppendByte('"')
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
		e.buf.AppendString(`"NaN"`)
	case math.IsInf(val, 1):
		e.buf.AppendString(`"+Inf"`)
	case math.IsInf(val, -1):
		e.buf.AppendString(`"-Inf"`)
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

// AddNull appends a null value to the buffer
func (e *Encoder) AddNull() error {
	e.addSeparator()
	e.buf.AppendString("null")
	return nil
}

// AddArrayish appends an array value to the buffer
func (e *Encoder) AddArrayish(arr Array) error {
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
		e.buf.AppendByte(',')
	}
}

// AddString adds a string to the encoded buffer
func (e *Encoder) AddString(s string) error {
	e.addSeparator()
	e.buf.AppendByte('"')
	start := 0

	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if safeSet[b] {
				i++
				continue
			}

			if start < i {
				e.buf.AppendString(s[start:i])
			}

			switch b {
			case '\\', '"':
				e.buf.AppendByte('\\')
				e.buf.AppendByte(b)

			case '\n':
				e.buf.AppendByte('\\')
				e.buf.AppendByte('n')

			case '\r':
				e.buf.AppendByte('\\')
				e.buf.AppendByte('r')

			case '\t':
				e.buf.AppendByte('\\')
				e.buf.AppendByte('t')

			default:

				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				e.buf.AppendString(`\u00`)
				e.buf.AppendByte(hex[b>>4])
				e.buf.AppendByte(hex[b&0xF])
			}

			i++
			start = i
			continue
		}

		c, size := utf8.DecodeRuneInString(s[i:])

		if c == utf8.RuneError && size == 1 {
			if start < i {
				e.buf.AppendString(s[start:i])
			}

			e.buf.AppendString(`\ufffd`)
			i += size
			start = i
			continue
		}

		i += size
	}

	if start < len(s) {
		e.buf.AppendString(s[start:])
	}

	e.buf.AppendByte('"')
	return nil
}
