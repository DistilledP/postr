package test

import (
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMockFileInfo(t *testing.T) {
	type args struct {
		name    string
		size    int64
		mode    fs.FileMode
		modTime time.Time
		isDir   bool
		sys     interface{}
	}
	testCases := []struct {
		name     string
		args     args
		expected MockFileInfo
	}{
		{
			name: "Constructs correctly",
			args: args{
				name: "test Name",
				size: 22,
			},
			expected: MockFileInfo{name: "test Name", size: 22},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewMockFileInfo(tc.args.name, tc.args.size, tc.args.mode, tc.args.modTime, tc.args.isDir, tc.args.sys)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestMockFileInfo_Name(t *testing.T) {
	type fields struct {
		name    string
		size    int64
		mode    fs.FileMode
		modTime time.Time
		isDir   bool
		sys     interface{}
	}
	testCases := []struct {
		name     string
		fields   fields
		expected string
	}{
		{
			name: "Name()",
			fields: fields{
				name: "Test Name",
			},
			expected: "Test Name",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MockFileInfo{
				name:    tc.fields.name,
				size:    tc.fields.size,
				mode:    tc.fields.mode,
				modTime: tc.fields.modTime,
				isDir:   tc.fields.isDir,
				sys:     tc.fields.sys,
			}

			assert.Equal(t, tc.expected, actual.Name())
		})
	}
}

func TestMockFileInfo_Size(t *testing.T) {
	type fields struct {
		name    string
		size    int64
		mode    fs.FileMode
		modTime time.Time
		isDir   bool
		sys     interface{}
	}
	testCases := []struct {
		name     string
		fields   fields
		expected int64
	}{
		{
			name: "Size()",
			fields: fields{
				size: 22,
			},
			expected: 22,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MockFileInfo{
				name:    tc.fields.name,
				size:    tc.fields.size,
				mode:    tc.fields.mode,
				modTime: tc.fields.modTime,
				isDir:   tc.fields.isDir,
				sys:     tc.fields.sys,
			}

			assert.Equal(t, tc.expected, actual.Size())
		})
	}
}

func TestMockFileInfo_Mode(t *testing.T) {
	type fields struct {
		name    string
		size    int64
		mode    fs.FileMode
		modTime time.Time
		isDir   bool
		sys     interface{}
	}
	testCases := []struct {
		name     string
		fields   fields
		expected fs.FileMode
	}{
		{
			name: "Mode()",
			fields: fields{
				mode: 0700,
			},
			expected: 0700,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MockFileInfo{
				name:    tc.fields.name,
				size:    tc.fields.size,
				mode:    tc.fields.mode,
				modTime: tc.fields.modTime,
				isDir:   tc.fields.isDir,
				sys:     tc.fields.sys,
			}

			assert.Equal(t, tc.expected, actual.Mode())
		})
	}
}

func TestMockFileInfo_ModTime(t *testing.T) {
	type fields struct {
		name    string
		size    int64
		mode    fs.FileMode
		modTime time.Time
		isDir   bool
		sys     interface{}
	}
	testCases := []struct {
		name     string
		fields   fields
		expected time.Time
	}{
		{
			name: "ModTime()",
			fields: fields{
				modTime: time.Date(2050, 10, 2, 1, 39, 23, 23, &time.Location{}),
			},
			expected: time.Date(2050, 10, 2, 1, 39, 23, 23, &time.Location{}),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MockFileInfo{
				name:    tc.fields.name,
				size:    tc.fields.size,
				mode:    tc.fields.mode,
				modTime: tc.fields.modTime,
				isDir:   tc.fields.isDir,
				sys:     tc.fields.sys,
			}

			assert.Equal(t, tc.expected, actual.ModTime())
		})
	}
}

func TestMockFileInfo_IsDir(t *testing.T) {
	type fields struct {
		name    string
		size    int64
		mode    fs.FileMode
		modTime time.Time
		isDir   bool
		sys     interface{}
	}
	testCases := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "IsDir() - true",
			fields: fields{
				isDir: true,
			},
			expected: true,
		},
		{
			name: "IsDir() - false",
			fields: fields{
				isDir: false,
			},
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MockFileInfo{
				name:    tc.fields.name,
				size:    tc.fields.size,
				mode:    tc.fields.mode,
				modTime: tc.fields.modTime,
				isDir:   tc.fields.isDir,
				sys:     tc.fields.sys,
			}

			assert.Equal(t, tc.expected, actual.IsDir())
		})
	}
}

func TestMockFileInfo_Sys(t *testing.T) {
	type fields struct {
		name    string
		size    int64
		mode    fs.FileMode
		modTime time.Time
		isDir   bool
		sys     interface{}
	}
	testCases := []struct {
		name     string
		fields   fields
		expected interface{}
	}{
		{
			name: "Sys()",
			fields: fields{
				modTime: time.Date(2050, 10, 2, 1, 39, 23, 23, &time.Location{}),
				sys:     map[string]int{"dd": 2},
			},
			expected: map[string]int{"dd": 2},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MockFileInfo{
				name:    tc.fields.name,
				size:    tc.fields.size,
				mode:    tc.fields.mode,
				modTime: tc.fields.modTime,
				isDir:   tc.fields.isDir,
				sys:     tc.fields.sys,
			}

			assert.Equal(t, tc.expected, actual.Sys())
		})
	}
}
