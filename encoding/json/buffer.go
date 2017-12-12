package json

import "strconv"

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

// AppendBool  adds a boolean value to the buffer.
func (b *Buffer) AppendBool(val bool) error {
	b.ba = strconv.AppendBool(b.ba, val)
	return nil
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
