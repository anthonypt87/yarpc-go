// Code generated by go-bindata.
// sources:
// internal/templatedata/base.tmpl
// internal/templatedata/bindata.go
// internal/templatedata/client.tmpl
// internal/templatedata/client_impl.tmpl
// internal/templatedata/client_stream.tmpl
// internal/templatedata/fx.tmpl
// internal/templatedata/parameters.tmpl
// internal/templatedata/server.tmpl
// internal/templatedata/server_impl.tmpl
// internal/templatedata/server_stream.tmpl
// DO NOT EDIT!

package templatedata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _internalTemplatedataBaseTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcc\x31\x4f\xc3\x30\x10\x05\xe0\xdd\xbf\xe2\x11\x75\x00\x89\x24\x7b\x25\x26\x0a\xa8\x4b\x5b\x89\xee\xe8\x70\xae\xc1\x22\x8e\xad\xb3\xa9\xa8\x2c\xff\x77\x44\x93\x02\x12\x4d\x37\xfb\x7d\xf7\x5e\x5d\xe3\xde\x35\x8c\x96\x7b\x16\x8a\xdc\xe0\xf5\x00\x2f\x2e\x3a\x5d\xb6\xdc\x97\x07\x12\xaf\xcb\xd6\xa9\xba\x46\x70\x1f\xa2\x79\x8e\x94\xaa\x47\xd3\x71\xb5\x22\xcb\x39\x7f\xcb\x62\x8d\xd5\x7a\x8b\x87\xc5\x72\x7b\xa5\x94\x27\xfd\x4e\x2d\xff\xdc\x6d\x86\x7f\xf5\xe4\xc6\x57\xce\x4a\xa5\x64\x76\x18\xfc\x99\x65\x6f\x34\x07\x94\x39\x2b\xc0\x58\xef\x24\xe2\x5a\x01\x40\x4a\x42\x7d\xcb\x98\x0d\xe9\x86\xe2\xdb\x2d\x66\xd4\x19\x0a\x98\xdf\xa1\x5a\x1e\xe3\x53\x75\x68\x0c\x9c\x33\x8a\x94\xfe\xf4\x72\x2e\xc6\x49\xee\x9b\xb1\x70\xa3\x7e\x7f\x2a\xa5\xc8\xd6\x77\x14\x19\x85\xee\x0c\xf7\xb1\x40\x75\xa4\xff\xf2\x62\xac\xef\x2e\x70\x88\xc2\x64\xcf\x1d\x04\x96\x3d\xcb\xb4\x4c\x2e\x8f\x3c\xbd\xbc\xfb\x3c\x97\x7a\x12\xb2\x1c\x59\xc2\x49\xbf\x02\x00\x00\xff\xff\xb7\xa9\x37\x42\xf4\x01\x00\x00")

func internalTemplatedataBaseTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataBaseTmpl,
		"internal/templatedata/base.tmpl",
	)
}

func internalTemplatedataBaseTmpl() (*asset, error) {
	bytes, err := internalTemplatedataBaseTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/base.tmpl", size: 500, mode: os.FileMode(420), modTime: time.Unix(1538673775, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func internalTemplatedataBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataBindataGo,
		"internal/templatedata/bindata.go",
	)
}

func internalTemplatedataBindataGo() (*asset, error) {
	bytes, err := internalTemplatedataBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1538678044, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x52\xcd\x6e\xdb\x3c\x10\xbc\xeb\x29\x06\xc6\xf7\xb5\xb6\xe1\xd2\xf7\x00\x3d\x14\x06\x52\xf4\x50\x37\x4d\x7a\xe9\x91\xa1\x57\x32\x11\x99\x64\x48\xaa\x6e\xb0\xe0\xbb\x17\x12\x15\xfd\xd8\x05\x5a\xa0\xba\x88\xdc\xc5\xce\xcc\x0e\x87\xf9\x40\xa5\x36\x84\x85\xaa\x35\x99\xb8\xc0\xbb\x94\x0a\xe6\xb3\x8e\x47\x88\x5b\x5d\x53\x77\xfd\xaf\xb2\xee\xa9\xc2\xcd\x7b\x88\x3b\xa9\x9e\x64\x45\xe2\xa3\xed\x4f\x29\x15\x05\xb3\x97\xa6\x22\x88\x07\xf2\x3f\xb4\xa2\xd0\xc1\x14\x00\xf3\x76\x8d\x5d\x07\x0d\x6d\x22\xf9\x52\x2a\xc2\x7a\x9b\xbb\xdb\x2d\x98\x45\x6e\xa7\x04\x1d\x10\x8f\xd4\x96\xf6\xf2\x44\x29\x21\x64\xb4\xb7\x01\xea\x02\x42\x14\x40\x7c\x71\x34\x9f\x1f\x08\xb8\x00\x5a\xf2\x5e\xd5\x67\x8a\x47\x7b\xc8\xa2\x80\xdc\xd2\x25\xac\x47\x3f\xfc\x10\x3d\xc9\x93\x36\x55\xde\x80\xfc\x58\x18\x67\x30\x2a\x5b\x0e\x25\x40\x59\x13\xe9\x67\x14\xbb\xfc\xdf\x4c\x5a\x1d\x8b\xb1\xf1\x9a\x66\x8a\xda\x7e\x6b\xe6\xca\x7e\x6b\x17\x12\xf7\xf4\xdc\x50\x88\xc8\x9e\xa7\x34\x07\x24\x73\xb8\x18\x16\x42\xbc\x48\xef\x94\xd8\xc9\xba\xfe\xe2\xa2\xb6\x66\x1c\x59\x61\xc9\x2c\x32\xef\xab\x4f\x1b\x90\xf7\xd6\xaf\x06\x2b\xa8\x0e\xf4\x4f\x7b\xfe\xa5\xfa\x3f\x29\x9d\xc1\x04\x67\x4d\xa0\x11\xe7\x4a\xf4\xc4\x88\xe9\xed\x2a\x76\xca\x9a\x10\x7d\xa3\x5a\xba\x69\xf2\xf6\x74\x9e\x86\xe7\xb1\xd1\xf5\x21\x40\xc2\xd0\x19\xdf\x3f\xdc\xdf\xed\x5e\x43\x57\x5a\xff\xfb\x5c\xb6\x21\x2c\x1b\xa3\x2e\xa0\x96\x0a\xfd\x9a\x5d\x61\x03\xeb\x62\x18\x96\x77\xde\x46\xfb\xd8\x94\x7d\x37\xdb\xb0\x9a\xe5\x38\xa7\xd7\x53\x6c\xbc\xc1\x9b\xa1\xf3\xe9\xe4\xea\x94\x38\x74\x8f\x79\x83\x39\xda\x9e\xce\xd3\x57\x5e\xaa\x0d\x98\x9d\xd7\x26\x96\x58\xfc\xff\xbc\x80\xb8\xfd\xba\x6f\x6d\x6c\xc5\x08\x21\x56\xbd\x57\xa3\x75\xe3\x71\xac\xfd\x0a\x00\x00\xff\xff\xd6\x5e\x65\x4c\x1f\x04\x00\x00")

func internalTemplatedataClientTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataClientTmpl,
		"internal/templatedata/client.tmpl",
	)
}

func internalTemplatedataClientTmpl() (*asset, error) {
	bytes, err := internalTemplatedataClientTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/client.tmpl", size: 1055, mode: os.FileMode(420), modTime: time.Unix(1538677589, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataClient_implTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x55\xcb\x8e\xdb\x20\x14\xdd\xfb\x2b\xce\x44\xd3\x2a\x8e\x3c\xcc\x3e\x55\x56\xd1\xb4\xea\xa2\x0f\x35\xdd\x57\x2e\xb9\xf6\xa0\xd8\x40\x80\x74\x26\x42\xfc\x7b\x05\x38\xcf\x79\xec\xa2\x51\xd5\x55\xe0\x86\x73\xcf\x83\x6b\xdb\xfb\x25\x35\x42\x12\x46\xbc\x13\x24\xdd\x2f\xd1\xeb\x6e\x84\x9b\x10\x0a\xef\x1f\x84\xbb\x07\xfb\x28\x3a\x4a\xdb\xeb\x56\xe9\x55\x8b\xe9\x0c\xec\x7b\xcd\x57\x75\x4b\xec\x93\x1a\x56\x21\x14\x85\xf7\xa6\x96\x2d\x81\x2d\xc8\xfc\x11\x9c\x6c\x6a\x03\x78\x7f\x9d\x9b\x7f\xee\x75\x97\xe0\xf3\xfd\x36\xe2\xe2\x89\xdb\x09\x72\x11\x51\x00\xf5\x24\x5d\xed\x84\x92\x98\xdc\xe6\x23\x6e\xab\xe9\xb4\x53\x08\xb0\xce\x6c\xb8\x83\x2f\x00\xc4\x0d\xd5\x3d\xb6\xb5\xd1\x5c\x1b\xe5\xd4\xef\x4d\xc3\x16\xa9\x98\x5b\x17\xc0\xc0\x36\xe8\xfc\x42\xee\x5e\x2d\x6d\x12\x19\xcb\xa2\x41\x2d\x97\x3b\x79\x19\x2a\x64\x9b\x0d\x91\x39\x14\x6e\x06\x08\xd0\x6c\x24\xc7\x98\x63\x72\x26\xad\x84\xf7\xec\x6b\xdd\x53\x08\x63\xee\x1e\xc1\x95\x74\xf4\xe8\xd8\x3c\xff\x56\x50\xda\x59\x30\xc6\x92\x5c\x36\xaf\xbb\xee\x9b\x8e\x86\x4b\x8c\xbd\x3f\x91\x1d\x42\x05\x32\x46\x99\x72\x30\x9a\xcc\xa6\x5a\x0c\x93\xb3\x6c\x3c\xf5\xc8\xb8\xc8\x58\xc1\x7b\x6d\x84\x74\x0d\x46\xef\xd6\x23\x0c\x62\x32\x31\x63\xac\xdc\xb7\x12\x4d\x6a\x75\x35\x83\x14\xdd\x11\x05\x60\xc8\x6d\x8c\x8c\xe5\xc4\xb6\xff\x27\x14\x67\x27\xde\x9f\x49\xce\x19\xf8\x2c\x6c\x0a\x1b\xaa\xd8\xa4\x38\x46\x7b\x4f\x9d\xa5\x48\xfe\x24\xee\xff\x2b\xdd\x0b\x06\x7b\x81\xb1\x35\xb4\x8e\xa0\x56\xfd\x8c\xcf\x23\xfb\x41\xeb\x0d\x59\x87\xfc\x6a\xd8\x05\xf0\x4f\x27\x3f\xc0\xa6\x33\x58\xb6\x20\xb9\x1c\x1b\x5a\x97\x1f\xde\xf0\x11\x79\xe3\x5b\x3b\xc1\x59\xad\xa4\xa5\x23\xe0\x93\x0b\xec\x6d\xfb\xfc\x15\xbe\x7a\x79\x86\xd6\x55\x32\x41\x0f\x3b\x92\x0b\xbc\xab\x6c\x05\xb5\x8a\xc2\x7a\xdb\xb2\x57\x8d\x9d\x30\x5e\xa9\xd5\x8b\x54\xa7\x5f\x9b\x79\x6d\xdd\x5d\x4c\x24\xce\xfa\x5d\xaf\xdd\xf6\xd8\x8d\x21\x5b\xbe\x3c\x20\x49\xde\x33\x23\x20\x97\xfb\xcf\xe8\x6e\x7d\x58\x1d\x96\x87\xda\xdf\x00\x00\x00\xff\xff\x67\x3f\x2b\x54\xd3\x07\x00\x00")

func internalTemplatedataClient_implTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataClient_implTmpl,
		"internal/templatedata/client_impl.tmpl",
	)
}

func internalTemplatedataClient_implTmpl() (*asset, error) {
	bytes, err := internalTemplatedataClient_implTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/client_impl.tmpl", size: 2003, mode: os.FileMode(420), modTime: time.Unix(1538677853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataClient_streamTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x4f\x6f\xdb\x20\x14\xbf\xf7\x53\xbc\x56\x3d\xd8\x51\x47\xee\x9b\x7a\x98\xaa\x6e\xda\x61\x7f\xd4\xec\x3e\x79\xf8\xd9\x41\xb1\xc1\x05\x9c\x2e\x42\xfe\xee\x13\x18\x88\x9d\xd8\x59\x1a\x6d\xb7\x9d\x42\x1e\xbc\xdf\x9f\x97\x1f\xc4\x98\x1c\x0b\xc6\x11\x6e\x68\xc5\x90\xeb\x1f\x4a\x4b\xcc\xea\x1b\x78\xd3\x75\x57\xc6\xbc\x30\xbd\x06\xf2\x81\x55\xe8\xbe\xde\x96\xa2\xd9\x94\xf0\xf6\x1e\xc8\xb7\x8c\x6e\xb2\x12\xc9\x47\xe1\x57\x5d\x77\x75\x65\x8c\xcc\x78\x89\x40\x56\x28\xb7\x8c\xa2\x72\x30\x00\xc6\xdc\xf6\xf0\xae\xf5\xc1\x2d\xed\x79\xbb\xb3\x5c\x40\x5f\x80\x9e\x1a\x18\xd7\x28\x8b\xcc\x76\x2f\x96\xe1\x94\xc7\xfd\x8c\x7a\x2d\xf2\x00\x6b\x37\x58\x01\x42\x06\xcc\x95\x43\x60\xbc\xec\x15\xa0\x8c\x05\x7f\x7e\xb9\x04\x63\x48\x5f\x0d\x32\x80\x29\xc8\x3c\xb9\x6d\x8d\xfc\xd0\x2a\xcc\x81\x71\xd0\x6b\xdc\x5b\xb0\x0d\xe1\x04\x71\xa0\x7a\xd7\xe0\x14\x6c\xc4\x31\xee\x18\xc0\x83\xe0\x1a\x7f\xe9\x24\x05\xda\xaf\x88\xaf\xec\xbd\x1c\x19\x09\x4e\x01\x56\xc8\xf3\x64\x61\x4c\x29\xbe\x5b\x42\xf2\x84\xcf\x2d\x2a\x0d\xfd\x8f\xd2\x75\x77\x40\x08\xd9\x65\xb2\xa1\x5e\xc9\xd7\x46\x33\xc1\x53\x40\x29\x85\xf4\x14\xc8\xf3\xf1\xf0\x0e\x07\x35\x20\x7c\x42\xba\x4d\x66\x30\xc7\x42\x54\x23\xb8\xc2\x81\x12\x47\x99\x06\xdb\x95\x50\xe8\xd4\x5f\x22\x30\xe3\xf9\xf1\x54\x12\x2e\xf4\x91\xf4\x74\xa0\xdd\x71\xbe\xe7\xf9\xdf\xf1\x30\x96\xd5\x1d\xd5\xf6\xeb\x99\x48\xd7\x4d\x85\x35\x72\x9d\x59\xe6\x7f\x90\xeb\xa9\x08\x7e\xaa\x9b\xaa\xeb\xac\x84\x96\xea\x98\x41\xaf\x68\xe1\x06\xd2\x48\xa1\xc5\xcf\xb6\x18\xf5\x79\x8f\xee\xa3\x68\x39\x85\x84\xc2\x62\x12\x3b\x9d\x4f\x74\xe4\x93\xa8\x5b\xc9\x81\x92\x9e\x98\xc4\x8e\x21\xcd\x64\xf4\xbd\xb5\x3f\x49\x70\xc1\x92\xf8\x0c\x27\xaf\x86\x68\xb4\x3a\x79\x3f\x66\x05\x07\xfc\x1e\x83\x10\x92\x4e\x86\x60\xfe\x46\x9d\xe9\xc3\x25\xf5\x94\xcc\xb3\xe2\x1a\x6d\xd4\xaa\x74\x35\xfb\xe2\x46\x2f\x4f\x48\x91\x6d\x31\x31\x86\x7c\xc1\x97\x00\x12\xe6\x13\xbd\x01\xb0\xc2\xf5\x5e\xdf\x03\x67\x55\xc4\x8c\xc3\xe1\xac\x72\xe0\xbe\xde\xc5\xd1\xa9\x3b\x10\x1b\x4b\x59\xab\x92\x9c\x14\x3c\x60\xba\x16\x9b\x19\x8a\x71\x48\x1f\x32\xa5\x1f\xad\x4b\xab\xff\xb1\x6e\xf4\x6e\xe8\xa0\x56\x65\x7a\xa4\xc7\x41\x39\x59\x9c\x55\xaf\x0a\x76\x7c\xb3\x2e\x4f\x8e\xc3\x48\xce\x4c\xcd\x2b\x9e\xb9\x33\x03\x35\x7a\x02\xcf\x0f\xd6\xe1\xcd\x39\xc8\x95\x4f\xc6\x30\x55\x63\x9f\xef\x2e\x49\xce\xff\xb4\x8e\xd3\x1a\xd4\xce\xe4\x65\xf2\x9f\x67\x6a\xb9\xaf\xfd\x0e\x00\x00\xff\xff\x2a\x90\x4d\x32\xe9\x09\x00\x00")

func internalTemplatedataClient_streamTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataClient_streamTmpl,
		"internal/templatedata/client_stream.tmpl",
	)
}

func internalTemplatedataClient_streamTmpl() (*asset, error) {
	bytes, err := internalTemplatedataClient_streamTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/client_stream.tmpl", size: 2537, mode: os.FileMode(420), modTime: time.Unix(1538678006, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataFxTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x41\x6f\xdb\x38\x10\x85\xef\xfa\x15\x0f\xc6\x1e\xe4\x20\xa1\x81\x3d\x1a\x58\x2c\xb0\x5b\xb8\xcd\x25\x36\x9a\xde\x8a\xa2\x61\xe4\x91\x4c\x44\x26\x59\x92\xb2\x15\x10\xfa\xef\x05\x45\xc9\x8a\x1c\xa5\x75\xe3\x83\x4d\x8f\x86\x6f\x3e\xcd\xf0\x49\xde\x6f\x29\x17\x92\x30\xcb\xeb\x19\x6e\x9a\x26\xf1\xfe\x28\xdc\x0e\x6c\x25\x4a\x6a\xff\xfe\x55\x28\xfd\x54\x60\xf9\x0f\xd8\x86\x67\x4f\xbc\x20\xf6\x51\x75\xab\xa6\x49\x12\xef\x0d\x97\x05\x81\xdd\x93\x39\x88\x8c\x6c\x2b\x93\x00\xde\x2f\xae\xb0\xaa\x91\x95\x82\xa4\xb3\xb8\x5a\xc4\xf8\x62\x01\xef\xd9\xaa\xfe\xbf\x8d\x37\xcd\x86\x1b\xbe\xb7\x88\x20\x16\x6e\x47\xd0\x21\x44\x8e\x8c\x8d\xf9\x86\x7e\x54\xc2\xd0\x16\x4e\x41\x1b\x75\x10\x5b\x02\x0f\x2a\xbd\x06\x84\x74\x0a\x5c\xc6\xf4\x55\x0d\xae\x75\x29\x32\xee\x84\x92\x2c\x01\xdc\xb3\xa6\xc9\xaa\xd6\x99\x2a\x73\xf0\x09\x00\xe4\x35\xbb\x95\x49\xbb\x8c\x69\x78\xe6\x46\x67\x5d\x95\x04\x98\xe4\xff\x4c\xb6\x2a\x5d\xcf\x65\xc7\x60\x31\xbf\xa3\xbb\x08\xac\x93\x3b\x07\x5b\x57\x6e\x44\x36\xae\xd1\x83\xdd\xd1\x71\x24\xf6\x0e\xaa\x6b\x54\x56\xc8\xa2\x9d\x43\x21\x0e\xd4\xf5\x54\xf2\x3d\x21\x57\x06\x46\x55\x4e\xc8\x82\xb5\xe1\x78\x2d\xf0\x6d\x62\xa1\xb4\x8b\x84\xf1\xc7\x93\xd3\x34\xec\x9c\x2a\x9d\xd9\x78\x56\x6e\x82\xea\x6c\x7e\x7d\xda\xc4\x18\xeb\xd6\xf3\xf8\xfb\x65\xfd\x61\x9d\xee\x49\xda\x6c\x37\x5f\xe2\x93\x3a\xe2\x28\xca\x12\x6e\x27\x2c\x8e\xca\x3c\x41\x48\x1c\xfe\xfe\x37\x01\xf2\x4a\x66\xaf\xee\x3f\xfd\x1e\x1a\x29\x64\x71\x0d\xa5\x9d\x0d\xfa\xed\x48\xb5\x51\x4e\x3d\x56\x79\xd7\x91\xb5\x0e\x77\x3e\x0f\x1d\x21\x93\xf3\x8c\x7c\xd3\x75\xde\x90\xab\x8c\x6c\xc5\x53\x3d\x75\x82\xe6\x93\xd3\x8b\x9b\x4f\xdb\x27\x52\xfa\x8c\x7e\xa0\xcb\xd7\xec\xba\xa3\x8b\xec\x8c\xb1\xb6\x51\xe1\xd3\x24\xfd\xf7\xc8\x69\xa1\xab\x64\x5e\x3b\xed\xbe\x8d\xbf\xed\x34\xfe\x6b\xa7\x85\x3c\xef\x59\xaf\x12\xe2\x19\x6d\x2b\x43\xf6\xf2\xd3\x7d\x86\xf0\x96\xed\x62\xda\xcb\x6a\xe7\xae\xeb\xe3\xe7\xae\x3b\x83\x8c\x3b\x06\xd2\x3f\x82\xfc\x9d\x05\x37\x83\xec\xd7\x6f\xf1\x11\x71\x0a\xe1\xa1\x30\xaa\xd2\xcb\x59\x1b\xcf\xeb\xd9\xc3\x84\x3d\x5f\xb6\xf2\x32\xfc\x49\x78\xdc\x3a\x50\xad\x29\x73\xdd\x24\xf8\x68\x4e\x4e\xe1\x91\xa0\x0d\xd9\xf0\xc0\x10\xb2\x2d\x92\x29\xe9\xb8\x90\x64\xde\x63\xe0\x5e\x3a\x7d\xd3\xb2\x23\x17\x0e\xe9\x97\x39\x6b\x7c\x48\xe6\x93\x43\x99\x72\xd6\x38\x65\x70\xd6\x30\xa7\x25\xfe\xab\x44\xb9\xf5\x9e\xdd\xf1\x3d\x35\xcd\x70\x25\xd5\x5d\xc3\x26\xcd\xe5\x3d\xc9\x6d\x7c\xa1\x9d\x96\x43\xec\x67\x00\x00\x00\xff\xff\xd3\xb5\xa5\xa9\x3e\x07\x00\x00")

func internalTemplatedataFxTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataFxTmpl,
		"internal/templatedata/fx.tmpl",
	)
}

func internalTemplatedataFxTmpl() (*asset, error) {
	bytes, err := internalTemplatedataFxTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/fx.tmpl", size: 1854, mode: os.FileMode(420), modTime: time.Unix(1538677131, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataParametersTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x6a\xc3\x30\x0c\x86\xef\x79\x8a\x9f\x32\x46\x5b\x98\x7b\x1f\xf4\xb8\xed\xb4\x51\xb6\xbd\x80\x49\xd4\x34\x74\xb5\x3d\xdb\x69\x09\x42\xef\x3e\xea\x38\x59\xba\x96\xdd\x24\xa1\xff\x43\x9f\x98\x2b\xda\x36\x86\x30\x73\xda\xeb\x03\x45\xf2\x61\x86\x07\x91\x82\xf9\xd4\xc4\x1d\xd4\x73\xf3\x45\xa9\xbd\xab\xad\xdb\xd7\x78\x5c\x43\x6d\x74\xb9\xd7\x35\xa9\x17\x9b\x2b\x91\xa2\x60\xf6\xda\xd4\x04\xf5\x41\xfe\xd8\x94\x14\x12\xa6\x00\x98\x57\x4b\x6c\x06\x3c\x4a\x6b\x42\xf4\x6d\x19\xad\x0f\x58\xae\x86\x9d\x1c\x7e\xa5\xb8\xb3\x55\x9f\x05\x80\x6d\x6b\x4a\x30\xab\x37\x3a\xbd\xd3\x77\x4b\x21\x8a\xcc\x17\x00\xc3\x53\x6c\xbd\xc1\x3d\x73\x6d\x3f\x3b\x47\x50\x79\x01\xfd\xa5\x22\x2c\xc0\x0d\x4a\x70\xd6\x04\x4a\x98\xdb\x94\x7e\x61\x8a\x91\x74\x22\x99\xea\xc2\xe9\xe9\xe0\x62\x87\xf1\x71\x88\x9d\xa3\xa9\x52\xff\xc0\x6c\x94\x7d\x8e\xda\x63\x9e\xaa\x51\x79\x74\x3d\xcf\x54\x82\x8e\xaa\xc0\xfa\x7f\xc3\xab\xe0\x60\xf7\x37\x78\x25\x95\x8f\x18\xa4\xce\xdd\xe2\x42\x73\x22\x3c\x96\xbf\xb3\x9f\x00\x00\x00\xff\xff\x63\xc0\xf3\x97\x3a\x02\x00\x00")

func internalTemplatedataParametersTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataParametersTmpl,
		"internal/templatedata/parameters.tmpl",
	)
}

func internalTemplatedataParametersTmpl() (*asset, error) {
	bytes, err := internalTemplatedataParametersTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/parameters.tmpl", size: 570, mode: os.FileMode(420), modTime: time.Unix(1538674376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataServerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\x4f\x6f\xd4\x3e\x10\xbd\xe7\x53\x8c\x56\xfd\xfd\xd8\xad\x8a\xf7\x1e\x89\x03\xac\x54\xe0\x40\xb5\xb4\x70\x40\x88\x83\x49\x26\x59\xab\x59\x3b\xb5\x9d\x96\xca\xf2\x77\x47\x8e\x9d\xd8\x49\x5a\x75\x25\x10\x7b\x59\xff\x79\xf3\xfc\x66\xe6\x4d\x8c\x29\xb1\x62\x1c\x61\xa5\x50\xde\xa3\x5c\xc1\x6b\x6b\x33\x63\x1e\x98\x3e\x00\xb9\x64\x0d\xf6\xdb\xb3\x5a\xb4\xb7\x35\xe4\x6f\x80\xec\x69\x71\x4b\x6b\x24\xef\x45\x58\x59\x9b\x65\xc6\x48\xca\x6b\x04\x72\x83\xf2\x9e\x15\xa8\x7a\x1a\x00\x63\xce\xd4\x7d\xd1\xc7\x5d\xd1\x63\x0f\x75\x87\xdb\x73\x08\x40\x60\x5c\xa3\xac\x68\x81\x70\xbe\xf5\xd7\xdb\x2d\x18\xd3\x13\xa1\xb4\x16\x98\x02\x7d\xc0\xc0\x64\x2d\x28\x1f\xf8\x4a\x81\x57\x1c\x19\x48\x06\xa0\x1f\x5b\x9c\x86\x8f\xfc\x26\x03\x70\x8f\x07\xa5\x9f\x50\x1f\x44\x39\x08\xf5\x57\xac\x02\x21\x81\xec\x1a\x86\x5c\xdf\x68\x89\xf4\xc8\x78\x0d\x81\x2d\x1e\xc4\x18\x17\x15\x52\x5b\x8f\x47\x81\x8a\x0b\xbd\xe4\x4a\x43\xdd\xef\xdc\x98\x5a\x7c\x71\xaa\xc9\x35\xde\x75\xa8\x34\xf8\x62\x5b\x7b\x31\x21\x44\x5e\xce\x82\x5d\x9a\x3d\xef\x90\x6c\x0c\xd8\x24\x0a\x66\xe2\xad\x85\xf5\x32\x12\x50\x4a\x21\x5d\x18\x36\x0a\xad\xf5\x7b\xff\x6c\x52\x20\x77\xf9\x72\xf6\x85\xe0\x1a\x7f\x69\xb2\xf3\xff\x69\x1e\x27\xa5\xbb\x81\xf5\x04\xa7\x5a\xc1\x15\x46\x60\x10\x9b\x3d\x55\x9a\x74\x17\xdd\xb6\x97\xa2\xc0\xb2\x93\xe8\xb4\x29\x2d\xbb\x42\x33\xc1\x53\xcb\xbd\xeb\x58\x53\x1a\x43\x46\xa4\xb2\x36\x82\xbd\x07\xbf\xbd\xbd\xde\xef\xa0\x1d\x11\x50\x09\xf9\xa4\x39\x9d\x13\xab\x8e\x17\x4f\xb2\xae\x55\x6a\xd0\x0d\x7c\xff\xf1\x48\x65\x5b\x44\x4c\x70\xea\xc1\x8d\xcd\xff\x23\xf4\xe3\xb1\x6d\xac\x35\xde\xf4\x39\x28\x9f\xae\x44\xdd\x49\x0e\x3d\x43\x2b\x85\x16\x3f\xbb\x2a\x79\x6d\xe8\xca\x73\xf7\x7b\x2a\xe9\x51\x99\xb1\xf2\x61\x2a\x73\x30\xa6\x95\x8c\xeb\x0a\x56\xff\xdd\xad\x80\x5c\x7e\xbe\x4a\x1b\xf4\x95\x53\xf9\x98\x0f\xca\x47\xde\xfe\xf8\x79\xf2\x38\x7c\x9d\x03\x0e\x03\x48\xe6\xb6\x9e\xcc\x87\x47\x39\x8f\x2d\x44\x79\xe3\x5d\x4c\xe0\x1f\x28\x2f\x1b\x57\x9f\xa9\xb2\x2b\x7c\xe8\xc5\x85\xeb\xe9\x13\xf3\xf2\xa4\xc8\x65\x0e\xe9\x3b\x39\x1c\xc8\x38\x01\x17\x0b\xd0\x4e\x22\xd5\x98\x03\xc7\x87\xc1\x21\x23\x3a\xb8\x7f\xbd\x99\x87\xcd\x78\x26\xdb\x17\x3e\x09\xc9\xb5\x1f\xef\x65\x8b\xfc\xf9\x29\x3d\x52\x3d\xf2\x1f\x37\xc9\xcb\x3b\xa9\x4b\x13\xe8\x1f\xb4\xe9\x6f\xd4\x3b\x2c\x36\xfd\x37\x27\xe2\xe2\x32\x9e\xfd\x0e\x00\x00\xff\xff\xb3\xe9\xc6\x62\x71\x07\x00\x00")

func internalTemplatedataServerTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataServerTmpl,
		"internal/templatedata/server.tmpl",
	)
}

func internalTemplatedataServerTmpl() (*asset, error) {
	bytes, err := internalTemplatedataServerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/server.tmpl", size: 1905, mode: os.FileMode(420), modTime: time.Unix(1538677629, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataServer_implTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\x4d\x6f\xda\x30\x18\xbe\xf3\x2b\x9e\x55\xd5\x94\xa0\xd4\xbd\x57\xe2\x84\xba\x69\x87\x4e\x53\xd9\xbd\xf2\xcc\x4b\xb0\x9a\x38\xc1\x36\x14\x64\xf9\xbf\x4f\xb6\x43\x42\xe8\x90\x86\xd4\x8d\x0b\xf6\x9b\xf8\xf9\x4c\xe2\xdc\x92\x56\x52\x11\x6e\x0c\xe9\x1d\xe9\x17\x59\xb7\xd5\x0d\xee\xbc\x9f\x38\xf7\x26\xed\x1a\xec\x8b\xac\x28\x6e\x6f\xcb\xa6\x7d\x2d\xf1\x30\x03\xfb\xc1\xc5\x2b\x2f\x89\x7d\x6d\xba\x95\xf7\x93\x89\x73\x9a\xab\x92\xc0\x16\xa4\x77\x52\x90\x89\x30\x80\x73\xb7\x09\xfc\x5b\xdd\x56\xf1\xf8\xa2\xdf\x86\x73\xe1\x8e\xfb\x29\xd2\x10\x41\x00\xd5\xa4\x2c\xb7\xb2\x51\x98\xde\xa7\x5b\xec\xa1\xa5\x31\x92\xf7\x30\x56\x6f\x85\x85\x9b\x00\x40\xba\x02\xe7\x3a\xf8\xc8\xdd\xc1\x77\xc2\x9e\xc8\xae\x9b\xa5\x89\x57\xc2\x58\xae\xc0\xd5\x12\x6c\x5e\x49\x52\x76\x61\x35\xf1\x5a\xaa\xf2\x28\x70\x18\xdc\x75\x47\x80\xd5\x56\x09\x64\x6b\x4c\xcf\xb4\xe4\x81\xf8\x3b\xaf\xc9\xfb\xcc\x60\x7a\xe0\xba\x15\xad\x6e\x6c\xf3\x6b\xbb\x62\x09\x28\x81\xe6\x20\xad\x1b\xdd\x89\x0e\x3f\x4d\x76\xab\x15\xd6\x2c\xe1\xb1\x01\xe8\x73\x30\x73\x72\x36\x51\x39\x13\x47\x0f\x30\x3e\xef\x40\x8e\x86\xa8\x32\x04\xb9\x7a\xef\xe8\x1f\x1b\x30\x45\x18\x86\x6e\xaf\x70\x11\x37\xa7\x2e\x10\xa4\x07\x9c\x4f\x33\x28\x59\x9d\x30\xf4\x21\x91\xd6\xfd\xd0\x9f\x27\x68\xd8\x82\xd4\x32\xd3\x64\x2e\xe6\xf2\x9f\x8b\x15\xbb\x3e\x18\xc3\x9e\x49\x90\xdc\x51\x16\xf0\xe8\xed\x99\x36\x5b\x32\xd6\x7f\x84\xfb\x4d\x81\x97\xc0\x11\x08\x59\x36\x75\xae\x6c\x7e\x86\xf7\x85\x75\x24\x48\xaf\xee\x19\x97\xa6\x0d\x66\x17\xb9\xc6\x46\xe7\xdc\xd8\xc7\x60\x2f\xa8\x7f\xac\x5b\x7b\xe8\xf5\x17\x91\x36\xbf\x5c\xcb\x1f\x1e\x89\x28\xf9\xaf\x9f\x8b\x51\x8b\xd7\x35\x26\xec\x1e\xa2\x51\x96\xf6\x96\xcd\xd3\x7f\x81\x1a\xd1\x17\x7b\x22\x63\x78\x49\x39\xb2\xd1\xbe\x48\x45\xe6\xa3\x26\x8f\x11\xd7\x1f\x95\xaf\x92\x55\x71\x4d\xc8\xf5\x75\x09\x0b\xbb\x0f\xc5\x6c\xde\x65\xa8\x96\xfd\x67\xf9\xb8\x1e\x56\xc3\x72\x98\xfd\x0e\x00\x00\xff\xff\xa5\xe1\x05\x59\x23\x06\x00\x00")

func internalTemplatedataServer_implTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataServer_implTmpl,
		"internal/templatedata/server_impl.tmpl",
	)
}

func internalTemplatedataServer_implTmpl() (*asset, error) {
	bytes, err := internalTemplatedataServer_implTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/server_impl.tmpl", size: 1571, mode: os.FileMode(420), modTime: time.Unix(1538677912, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplatedataServer_streamTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x4d\x6f\xdb\x30\x0c\xbd\xe7\x57\xb0\x45\x0f\x71\xd0\x29\xf7\x01\x3d\x05\xdd\xb0\xc3\x3e\x90\xec\x3e\x78\x36\xed\x08\xb1\x25\x55\x92\xd3\x05\x82\xfe\xfb\x20\x5a\x56\x9c\xc6\xce\x7a\xd9\x29\x0c\x4d\xbe\xf7\xa8\x47\x3a\x57\x62\xc5\x05\xc2\xbd\x41\x7d\x44\xfd\xcb\x58\x8d\x79\x7b\x0f\x1f\xbc\x5f\x38\xf7\xca\xed\x1e\xd8\x27\xde\x20\xfd\x7d\xa8\xa5\x3a\xd4\xf0\xf1\x09\xd8\x8f\xbc\x38\xe4\x35\xb2\xcf\x32\x46\xde\x2f\x16\xce\xe9\x5c\xd4\x08\x6c\x87\xfa\xc8\x0b\x34\x04\x03\xe0\xdc\x43\x0f\x4f\xad\x3b\x0a\x43\x7d\xf8\xb2\x5e\x41\x9f\x80\x9e\x1a\xb8\xb0\xa8\xab\x3c\x74\xaf\xd6\x43\x55\xc4\xfd\x8a\x76\x2f\xcb\x01\x36\x7c\xe0\x15\x48\x0d\x6c\xd3\x70\x14\x76\x47\x08\x5c\xd4\x03\x49\x4a\xc4\xfa\xf5\x1a\x9c\x63\x7d\x76\x90\x01\xdc\x40\x1e\xc9\x43\x6b\xe2\x87\xce\x60\x09\x5c\x80\xdd\xe3\x79\x84\xd0\x30\x54\x30\x02\xb5\x27\x85\x53\xb0\x09\xc7\x51\x19\xc0\x46\x0a\x8b\x7f\xec\x32\x83\xa2\x8f\x58\xcc\x9c\x67\xb9\x1a\x64\x98\x14\x60\x8b\xc5\x71\xc9\x18\x3b\xe5\x5a\x15\x91\xec\xbb\xb2\x5c\x8a\x0c\x96\x2b\xe7\x6a\xf9\x33\x08\x61\x5b\x7c\xe9\xd0\x58\xe8\xcd\xf2\xfe\x11\x50\x6b\xa9\xb3\x48\x82\xa2\xbc\x7c\xbe\xb7\x4f\x35\xa2\xdc\xa1\x28\x2f\xa1\x8d\x92\xc2\xe0\x08\x7b\x46\x10\x51\x4e\x30\xfa\xab\xdc\x39\x9e\x59\x87\x56\x35\xd8\xa2\xb0\x79\x40\xfe\x0f\x3b\x31\x65\xdf\x97\x56\x35\xde\x07\x09\x5d\x61\x93\x7f\x51\xd1\x8a\x06\x56\x5a\x5a\xf9\xbb\xab\x2e\xfa\xe2\x8c\xf4\x53\x75\xa2\x80\xa5\x81\xd5\x24\x76\x36\xbf\x0d\x89\x4f\xa3\xed\xb4\x00\xc3\x7a\x62\x96\x3a\xc6\x34\x93\x6b\x13\x47\xfb\x97\x04\x5a\x29\xa9\xac\x99\xb3\xf1\x3d\x7b\x95\xe4\xb6\xa6\xa6\x5c\x38\xf2\xa4\x79\x8b\x05\xf2\x23\x2e\x9d\x63\xdf\xf0\x35\x62\x84\xe6\x40\xcb\x18\xcb\x62\x33\xaf\xa8\xf5\xee\x09\x04\x6f\x12\x64\x7a\x03\xc1\x1b\xc2\x8e\x79\x9f\x5e\xe8\xe5\x11\xe4\x21\x30\xb6\xa6\x66\xb7\xe4\x8e\x88\xee\xe4\x61\x86\xe1\xd2\xd9\x4d\x6e\xec\x73\x98\x31\xa8\x7f\x6e\x95\x3d\x8d\xf4\xb7\xa6\xce\xae\xd4\x10\x12\x89\x12\xbc\x99\x5c\xf8\xf9\xbb\x7b\xa7\x67\x74\x93\x1a\xa9\xe0\xc6\x5d\xde\x72\x95\x7c\x9b\xdd\xb2\x81\xe0\x8d\x45\x33\x83\x4c\x1e\xf2\x54\x78\xce\xfd\x0d\x00\x00\xff\xff\x05\xbb\x91\x3c\x74\x06\x00\x00")

func internalTemplatedataServer_streamTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplatedataServer_streamTmpl,
		"internal/templatedata/server_stream.tmpl",
	)
}

func internalTemplatedataServer_streamTmpl() (*asset, error) {
	bytes, err := internalTemplatedataServer_streamTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/templatedata/server_stream.tmpl", size: 1652, mode: os.FileMode(420), modTime: time.Unix(1538677950, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"internal/templatedata/base.tmpl": internalTemplatedataBaseTmpl,
	"internal/templatedata/bindata.go": internalTemplatedataBindataGo,
	"internal/templatedata/client.tmpl": internalTemplatedataClientTmpl,
	"internal/templatedata/client_impl.tmpl": internalTemplatedataClient_implTmpl,
	"internal/templatedata/client_stream.tmpl": internalTemplatedataClient_streamTmpl,
	"internal/templatedata/fx.tmpl": internalTemplatedataFxTmpl,
	"internal/templatedata/parameters.tmpl": internalTemplatedataParametersTmpl,
	"internal/templatedata/server.tmpl": internalTemplatedataServerTmpl,
	"internal/templatedata/server_impl.tmpl": internalTemplatedataServer_implTmpl,
	"internal/templatedata/server_stream.tmpl": internalTemplatedataServer_streamTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"internal": &bintree{nil, map[string]*bintree{
		"templatedata": &bintree{nil, map[string]*bintree{
			"base.tmpl": &bintree{internalTemplatedataBaseTmpl, map[string]*bintree{}},
			"bindata.go": &bintree{internalTemplatedataBindataGo, map[string]*bintree{}},
			"client.tmpl": &bintree{internalTemplatedataClientTmpl, map[string]*bintree{}},
			"client_impl.tmpl": &bintree{internalTemplatedataClient_implTmpl, map[string]*bintree{}},
			"client_stream.tmpl": &bintree{internalTemplatedataClient_streamTmpl, map[string]*bintree{}},
			"fx.tmpl": &bintree{internalTemplatedataFxTmpl, map[string]*bintree{}},
			"parameters.tmpl": &bintree{internalTemplatedataParametersTmpl, map[string]*bintree{}},
			"server.tmpl": &bintree{internalTemplatedataServerTmpl, map[string]*bintree{}},
			"server_impl.tmpl": &bintree{internalTemplatedataServer_implTmpl, map[string]*bintree{}},
			"server_stream.tmpl": &bintree{internalTemplatedataServer_streamTmpl, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

