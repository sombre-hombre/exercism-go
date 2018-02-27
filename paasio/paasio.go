package paasio

import (
	"io"
	"sync"
)

// NewWriteCounter creates new WriteCounter
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

// NewReadCounter creates new ReadCounter
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

// NewReadWriteCounter creates new ReadWriteCounter
func NewReadWriteCounter(readerWriter io.ReadWriter) ReadWriteCounter {
	return &readerWriterCounter{
		WriteCounter: NewWriteCounter(readerWriter),
		ReadCounter:  NewReadCounter(readerWriter),
	}
}

type readerWriterCounter struct {
	WriteCounter
	ReadCounter
}

type counter struct {
	byteCounter int64
	callCounter int32
	mutex       sync.Mutex
}

type writeCounter struct {
	writer io.Writer
	counter
}

type readCounter struct {
	reader io.Reader
	counter
}

func (c *counter) Count(bytesCount int) {
	c.mutex.Lock()
	c.byteCounter += int64(bytesCount)
	c.callCounter++
	c.mutex.Unlock()
}

func (c *counter) GetCount() (n int64, nops int) {
	c.mutex.Lock()
	n, nops = c.byteCounter, int(c.callCounter)
	c.mutex.Unlock()
	return
}

func (c *writeCounter) Write(p []byte) (n int, err error) {
	n, err = c.writer.Write(p)
	c.counter.Count(n)
	return
}

func (c *writeCounter) WriteCount() (n int64, nops int) {
	return c.GetCount()
}

func (c *readCounter) Read(p []byte) (n int, err error) {
	n, err = c.reader.Read(p)
	c.counter.Count(n)
	return
}

func (c *readCounter) ReadCount() (n int64, nops int) {
	return c.GetCount()
}
