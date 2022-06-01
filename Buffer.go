package html2pdf

import (
   "io"
   "bytes"
)

type BufferReader interface{
   io.Reader
   Bytes() []byte
}

type Buffer struct {
   Buffer *bytes.Buffer
}

func (b Buffer) Write(p []byte) (n int, err error) {
   return b.Buffer.Write(p)
}

func (b Buffer) Read(p []byte) (n int, err error) {
   return b.Buffer.Read(p)
}

func (b Buffer) Bytes() []byte {
   return b.Buffer.Bytes()
}

func NewBuffer() *Buffer {
   return &Buffer{
      Buffer: &bytes.Buffer{},
   }
}
