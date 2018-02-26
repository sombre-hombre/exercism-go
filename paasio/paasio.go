package paasio

import (
	"io"
	"sync"
)

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{wrtr: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{rdr: reader}
}

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

type writeCounter struct {
	wrtr         io.Writer
	bytesWritten int64
	writeCalled  int32
	writeMux     sync.Mutex
}

func (c *writeCounter) Write(p []byte) (n int, err error) {
	n, err = c.wrtr.Write(p)
	c.writeMux.Lock()
	c.bytesWritten += int64(n)
	c.writeCalled++
	c.writeMux.Unlock()

	return
}

func (c *writeCounter) WriteCount() (n int64, nops int) {
	c.writeMux.Lock()
	defer c.writeMux.Unlock()
	return c.bytesWritten, int(c.writeCalled)
}

type readCounter struct {
	rdr        io.Reader
	bytesRead  int64
	readCalled int32
	readMux    sync.Mutex
}

func (c *readCounter) Read(p []byte) (n int, err error) {
	n, err = c.rdr.Read(p)
	c.readMux.Lock()
	c.bytesRead += int64(n)
	c.readCalled++
	c.readMux.Unlock()
	return
}

func (c *readCounter) ReadCount() (n int64, nops int) {
	c.readMux.Lock()
	defer c.readMux.Unlock()
	return c.bytesRead, int(c.readCalled)
}
