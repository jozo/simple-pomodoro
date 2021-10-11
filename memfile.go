package main

import (
	"github.com/dsnet/golib/memfile"
)

// MemFile represents a file in memory
// Used for audio in notification
type MemFile struct {
	f *memfile.File
}

// MemFile constructor
func NewMemFile(b []byte) *MemFile {
	return &MemFile{f: memfile.New(b)}
}

// Read bytes
func (mf *MemFile) Read(b []byte) (int, error) {
	return mf.f.Read(b)
}

// Seek in the file
func (mf *MemFile) Seek(offset int64, whence int) (int64, error) {
	return mf.f.Seek(offset, whence)
}

// Close function - does nothing in reality
// Needed to follow interface
func (mf *MemFile) Close() error { return nil }
