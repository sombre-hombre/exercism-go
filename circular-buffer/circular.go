package circular

import (
	"errors"
	"sync"
)

// BenchmarkOverwrite-8   	30000000	        60.0 ns/op	499723125.63 MB/s	       0 B/op	       0 allocs/op
// BenchmarkWriteRead-8   	10000000	       119 ns/op	83368035.28 MB/s	       0 B/op	       0 allocs/op

// Buffer implements circular buffer
type Buffer struct {
	data     []byte
	head     int
	tail     int
	dataSize int
	mux      sync.Mutex
}

func (b *Buffer) advance(idx *int) {
	if *idx == len(b.data)-1 {
		*idx = 0
	} else {
		*idx++
	}
}

func (b *Buffer) isFull() bool {
	return b.dataSize == len(b.data)
}

// NewBuffer creates a new circular buffer
func NewBuffer(size int) *Buffer {
	return &Buffer{
		data: make([]byte, size),
	}
}

// ReadByte reads a byte from the buffer
func (b *Buffer) ReadByte() (byte, error) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.dataSize == 0 {
		return 0, errors.New("Read from empty buffer")
	}

	result := b.data[b.tail]
	b.advance(&b.tail)
	b.dataSize--

	return result, nil
}

// WriteByte writes a byte to the buffer
func (b *Buffer) WriteByte(c byte) error {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.isFull() {
		return errors.New("Buffer is full")
	}

	b.data[b.head] = c
	b.advance(&b.head)
	b.dataSize++

	return nil
}

// Overwrite writes a byte to the buffer overwriting old data if needed
func (b *Buffer) Overwrite(c byte) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.isFull() {
		b.data[b.tail] = c
		b.advance(&b.tail)
	} else {
		b.data[b.head] = c
		b.advance(&b.head)
		b.dataSize++
	}
}

// Reset resets buffer's state
func (b *Buffer) Reset() {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.data = make([]byte, len(b.data))
	b.dataSize = 0
	b.head, b.tail = 0, 0
}
