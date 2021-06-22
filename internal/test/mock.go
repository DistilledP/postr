package test

import (
	"io/fs"
	"time"
)

type MockFileInfo struct {
	name    string
	size    int64
	mode    fs.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}
}

func NewMockFileInfo(
	name string,
	size int64,
	mode fs.FileMode,
	modTime time.Time,
	isDir bool,
	sys interface{},
) MockFileInfo {
	return MockFileInfo{
		name,
		size,
		mode,
		modTime,
		isDir,
		sys,
	}
}

func (m MockFileInfo) Name() string {
	return m.name
}

func (m MockFileInfo) Size() int64 {
	return m.size
}

func (m MockFileInfo) Mode() fs.FileMode {
	return m.mode
}

func (m MockFileInfo) ModTime() time.Time {
	return m.modTime
}

func (m MockFileInfo) IsDir() bool {
	return m.isDir
}

func (m MockFileInfo) Sys() interface{} {
	return m.sys
}
