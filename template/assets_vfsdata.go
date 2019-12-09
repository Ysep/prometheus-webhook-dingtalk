// Code generated by vfsgen; DO NOT EDIT.

// +build builtinassets

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC),
		},
		"/templates/default.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "default.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC),
			uncompressedSize: 2243,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x55\x4b\x6e\xdb\x30\x10\xdd\xe7\x14\x03\x79\x63\x11\xb0\xdc\x75\x80\xb6\x08\x9a\x26\x1b\xa3\x28\x1c\xa4\x9b\x26\x08\x18\x6b\xec\x30\xa1\x48\x95\x1c\x39\x09\x64\xed\xbb\xef\x01\x7a\xc5\x1e\xa1\x20\x29\x2b\xa2\x3f\xfd\xa5\x5d\x74\x27\xcf\xe7\x0d\xdf\xbc\x99\x71\x5d\x43\x8e\x73\xa1\x10\x92\xab\x2b\x5b\x5d\xdf\xe2\x8c\x12\x68\x9a\x8f\x75\x0d\xd9\x19\x71\xaa\x2c\xac\x80\xf4\x79\x59\xa2\x81\xa6\xa9\x6b\x10\x73\xc0\x4f\x9d\x33\x99\x0b\x23\xd4\xc2\xe5\x1c\xba\x9c\x23\x89\x86\x6c\x76\xe2\xad\xb0\x02\x89\x2a\xa4\xa1\xca\xa1\x69\x2e\xc1\x05\x9d\x1a\x5d\x95\x13\x7e\x8d\xd2\x66\x67\xda\x10\xe6\xef\xb9\x30\x36\xfb\xc0\x65\x85\xae\xe0\xad\x16\x0a\x12\x70\xa8\x10\x4a\x2e\x08\x86\x0e\x2b\x7b\xa3\x8b\x42\xab\x90\x9c\xb6\xb6\x1e\x5e\x0a\x4d\x33\xac\x6b\xb8\x17\x74\x13\x07\x67\x53\x2c\xf4\x12\xe3\xea\xef\x78\x81\x36\x3c\x70\x67\xf5\xee\xe1\x69\xf7\xd5\x7d\x1c\x44\xcd\xe3\x8e\x78\xc1\x15\x5f\xa0\x39\x9f\x4e\xda\xe4\xec\xed\x03\xa1\x51\x5c\x9e\x4f\x27\xd0\x34\xe3\xc1\xd8\xc7\xd9\xd7\x06\x67\x28\x96\x68\x5e\xba\xa0\x69\xfb\x23\x42\x8f\xe1\x09\x1f\x28\xd4\xb8\x92\xc2\x52\x0b\x6f\xb8\x5a\x20\x64\x2e\x9c\xb1\x40\x89\xb1\x83\x27\xc7\x76\x8f\xa1\x69\x5e\xc1\xc8\xab\xe0\xb8\x3b\xd9\xa0\x23\x0f\x2b\x28\xb8\xb9\xcb\xf5\xbd\x82\x15\xdc\x50\x21\x5b\x9a\xed\x93\x18\x3b\x52\x4a\x13\x27\xa1\x55\x5c\xa8\x67\xff\x8b\xd5\xce\x74\x65\x66\x78\xc8\x18\xf8\x79\x3c\x45\x85\x86\x93\x36\xa1\x99\x97\xc3\x1d\xc6\xf4\x60\x87\x50\xfd\x56\xe6\x38\xe7\x95\xa4\xec\x17\x5a\x3a\x18\x0c\x06\x70\xe1\x4b\xb7\x8d\xb4\xb8\x44\x23\xe8\x11\x56\x50\xb5\x1b\x71\x11\x46\xba\xdf\x00\x5b\x15\x05\x37\x8f\xbe\x32\x63\xc7\x68\x67\x46\x94\xce\xe5\x98\x6c\x06\xe7\x4f\xee\x36\xe1\xd4\xf0\xf2\xc6\x93\xfe\xf6\xf5\xcb\xe7\x3d\x2c\x3d\x30\x71\x21\xed\xe1\xcf\x15\x0f\x2b\xc4\x55\x0e\x43\x85\x30\xf4\x5a\xa4\x90\xac\xd9\x24\xe9\x86\x3d\xbc\x3f\x49\x9f\x23\x5f\xb4\x26\x3b\x04\x19\x33\x38\x0e\x52\x00\x1b\xc7\xcb\xb4\x96\x88\x04\x49\x6c\x65\x21\x2c\x4a\xc9\x29\xbe\x53\xd9\xbe\x75\x5c\x23\xcc\xb4\x22\x54\x5e\xda\x9e\x98\xff\xe0\xae\x5d\x5c\x02\x63\x0e\x5c\xa8\x1c\x1f\xa2\x1b\x03\x89\x9f\x31\xc5\x0b\xcf\xc5\xeb\xd9\x67\xb3\x75\x38\x1c\xab\x34\x88\xda\x3f\x7c\xd1\x0b\x52\x78\x01\xa3\xb0\x91\xde\x0c\xc1\x1c\xb2\x9e\xd0\x7f\x30\xec\x1b\x8c\x7c\xff\x46\xbd\x56\xee\x28\x3d\x45\xab\xe5\x12\xf3\xad\xe2\x6b\xc7\x1f\x94\x5f\xa7\x6e\x3d\x60\x14\xcf\xca\x04\x17\x7c\xf6\xb8\x35\x2a\xd2\x9b\x9f\x31\x29\x2d\xc0\xff\x3e\x28\xfd\xb8\xdf\x93\x3a\xb4\xf7\x44\x4b\xa9\xef\x9d\x57\xf9\x7f\xc4\xb9\x36\x30\xd3\x45\xc9\x49\x5c\x0b\xe9\x2e\xde\xd6\x92\x0a\xb5\xc8\xa4\x50\x77\x7b\x9a\xbf\xb1\xc4\xfb\x57\xb5\xc3\xe9\x69\xb0\x13\xa9\xf3\xc7\x58\xdf\x03\x00\x00\xff\xff\xe7\xf4\xeb\x1d\xc3\x08\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/default.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}