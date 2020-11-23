package main

import (
	"github.com/dsnet/golib/memfile"
)

type MemFile struct {
	f *memfile.File
}

func NewMemFile(b []byte) *MemFile {
	return &MemFile{f: memfile.New(b)}
}

func (mf *MemFile) Read(b []byte) (int, error) {
	return mf.f.Read(b)
}

func (mf *MemFile) Seek(offset int64, whence int) (int64, error) {
	return mf.f.Seek(offset, whence)
}

func (mf *MemFile) Close() error { return nil }
