package buffer

import (
	"strconv"
	"time"
)

const bufferSize = 1024

// Buffer is similar to the standard bytes.Buffer, except that it's leaner
// because it drops some extra protection around escaping html that we don't
// need.
//
// Cribs heavily from https://golang.org/src/encoding/json/encode.go
type Buffer struct {
	ba   []byte
	pool *Pool
}

// NewBuffer returns a new empty buffer with bufferSize capacity
func NewBuffer() *Buffer {
	return &Buffer{
		ba: make([]byte, 0, bufferSize),
	}
}

// Bytes returns the raw bytes that the buffer includes
func (b *Buffer) Bytes() []byte {
	return b.ba
}

// String returns the string form of the buffer
func (b *Buffer) String() string {
	return string(b.ba)
}

// Len returns the number of bytes in the buffer
func (b *Buffer) Len() int {
	return len(b.ba)
}

// AppendByte adds a single byte to the buffer
func (b *Buffer) AppendByte(bt byte) {
	b.ba = append(b.ba, bt)
}

// WriteByte is identical to AppendByte. This implements the io.ByteWriter
// interface.
func (b *Buffer) WriteByte(bt byte) error {
	b.ba = append(b.ba, bt)
	return nil
}

// AppendString adds a string to the buffer
func (b *Buffer) AppendString(s string) {
	b.ba = append(b.ba, s...)
}

// AppendInt appends an int64 to the buffer
func (b *Buffer) AppendInt(i int64) {
	b.ba = strconv.AppendInt(b.ba, i, 10)
}

// AppendPaddedInt appends a zero-padded, int64 to the buffer.
func (b *Buffer) AppendPaddedInt(i int64, minWidth int) {
	if minWidth == 0 {
		panic("The minWidth of padded ints must be between 1 and 4")
	}

	if minWidth == 2 {
		if i < 10 {
			b.ba = append(b.ba, '0')
		}
	} else if minWidth == 3 {
		switch {
		case i < 10:
			b.ba = append(b.ba, '0', '0')
		case i < 100:
			b.ba = append(b.ba, '0')
		}
	} else if minWidth == 4 {
		switch {
		case i < 10:
			b.ba = append(b.ba, '0', '0', '0')
		case i < 100:
			b.ba = append(b.ba, '0', '0')
		case i < 1000:
			b.ba = append(b.ba, '0')
		}
	}

	b.ba = strconv.AppendInt(b.ba, i, 10)
}

// AppendInt16 appends an int16 to the buffer
func (b *Buffer) AppendInt16(i int16) {
	b.AppendInt(int64(i))
}

// AppendInt32 appends an int32 to the buffer
func (b *Buffer) AppendInt32(i int32) {
	b.AppendInt(int64(i))
}

// AppendUint appends an uint64 to the buffer
func (b *Buffer) AppendUint(i uint64) {
	b.ba = strconv.AppendUint(b.ba, i, 10)
}

// AppendUint16 appends an uint16 to the buffer
func (b *Buffer) AppendUint16(i uint16) {
	b.AppendUint(uint64(i))
}

// AppendUint32 appends an uint32 to the buffer
func (b *Buffer) AppendUint32(i uint32) {
	b.AppendUint(uint64(i))
}

// AppendFloat appends a float.
func (b *Buffer) AppendFloat(f float64, bitSize int) {
	b.ba = strconv.AppendFloat(b.ba, f, 'f', -1, bitSize)
}

// AppendBool adds a boolean value to the buffer.
func (b *Buffer) AppendBool(val bool) {
	b.ba = strconv.AppendBool(b.ba, val)
}

// AppendDuration appends a Duration to the buffer
func (b *Buffer) AppendDuration(val time.Duration) {
	b.AppendInt(int64(val))
}

// AppendTime appends a Time to the buffer.
// 	This is hardcoded to use the format: 2006-01-02T15:04:05.000
//
func (b *Buffer) AppendTime(t time.Time) {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	b.AppendPaddedInt(int64(year), 4)
	b.AppendByte('-')
	b.AppendPaddedInt(int64(month), 2)
	b.AppendByte('-')
	b.AppendPaddedInt(int64(day), 2)
	b.AppendByte('T')
	b.AppendPaddedInt(int64(hour), 2)
	b.AppendByte(':')
	b.AppendPaddedInt(int64(min), 2)
	b.AppendByte(':')
	b.AppendPaddedInt(int64(sec), 2)
	b.AppendByte('.')
	b.AppendPaddedInt(int64(t.Nanosecond()/int(time.Millisecond)), 3)
}

// AppendTimestamp appends a Timestamp to the buffer
func (b *Buffer) AppendTimestamp(t time.Time) {
	b.AppendInt(t.UnixNano())
}

// Write appends raw bytes to the buffer.
func (b *Buffer) Write(bytes []byte) (int, error) {
	b.ba = append(b.ba, bytes...)
	return len(bytes), nil
}

// Reset resets the buffer to be empty, but it retains the underlying storage
// for use by future writes.
func (b *Buffer) Reset() {
	b.ba = b.ba[:0]
}

// Release releases this buffer back into the pool
func (b *Buffer) Release() {
	b.pool.Release(b)
}
