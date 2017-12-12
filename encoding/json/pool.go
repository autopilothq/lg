package json

import (
	"sync"
)

// Pool is represents a pool of our Buffer type.
// Cribs heavily from https://golang.org/src/encoding/json/encode.go
type Pool struct {
	pool sync.Pool
}

var (
	pool      *Pool
	GetBuffer func() *Buffer
)

func init() {
	pool = NewPool()

	// GetBuffer retrieves a Buffer from the Pool
	GetBuffer = pool.Get
}

func NewPool() *Pool {
	return &Pool{
		pool: sync.Pool{
			New: func() interface{} {
				// The Pool's New function should generally only return pointer
				// types, since a pointer can be put into the return interface
				// value without an allocation:
				return NewBuffer()
			},
		},
	}
}

func (p *Pool) Get() *Buffer {
	buf := p.pool.Get().(*Buffer)
	buf.Reset()
	buf.pool = p
	return buf
}

func (p *Pool) Release(buf *Buffer) {
	p.pool.Put(buf)
}
