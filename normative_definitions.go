// Code generated by go-bindata.
// sources:
// NormativeTypes/artifact_types
// NormativeTypes/capability_types
// NormativeTypes/data_types
// NormativeTypes/group_types
// NormativeTypes/interface_types
// NormativeTypes/node_types
// NormativeTypes/policy_types
// NormativeTypes/relationship_types
// DO NOT EDIT!

package toscalib

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

var _artifact_types = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x93\x41\x6b\xe3\x30\x10\x85\xef\xfe\x15\xef\xb8\x7b\x58\x6f\xf6\xea\x5b\x76\x97\x42\x0e\xa1\xa5\x49\x73\x29\x41\xa8\xf6\x38\x1a\x90\x25\x21\x29\x21\xfe\xf7\xc5\x72\x9a\x90\xa2\x12\x07\x7a\x33\xcc\x7b\xdf\xbc\x37\xc8\xd1\x86\x5a\x8a\x86\x5a\x36\x1c\xd9\x9a\x20\x0e\xe4\x03\x5b\x53\x61\x1c\x05\xee\x9c\x26\xd1\xcb\x4e\x8b\x3f\x62\x26\x66\x45\x21\x7d\xe4\x56\xd6\x51\xc4\xde\x51\xa8\x0a\x00\xa3\xb8\xfc\x98\x84\xf2\xd9\xda\x38\x4e\x80\x86\x42\xed\xd9\xc5\x44\x5d\x2b\xc2\xfa\x71\xf5\x6f\x8e\xf9\x49\x8c\x75\xef\x08\x52\x6b\xd8\xa8\xc8\xe7\xa6\x01\x0d\x79\x3e\x10\x5a\x6f\xbb\x22\xbb\xf0\x81\x35\x5d\x16\x0e\xe2\x46\x0c\xea\x2a\x1b\x2d\xcf\xf8\x4f\x4e\xdb\xbe\x23\x13\xef\x20\xe5\x3a\xa6\x06\x6f\x32\x10\x86\x13\xa1\xb5\x1e\xcd\x99\x8d\x33\xe0\x56\x8a\x72\xd1\xc9\xdd\xc4\x56\x17\xd7\x34\x6a\xb9\x59\xde\x0b\x1e\x8d\xb9\xc2\x1b\xf6\x71\x2f\x35\x96\xb2\x56\x6c\x08\x3f\x36\xcb\x9f\x18\xd5\xd9\x30\x8b\xe1\x4d\x0d\x48\x99\xfc\xdf\x7e\x6c\xbe\xe2\xdf\x3a\xf8\x75\x9a\xf2\xaf\x0c\x6a\x5a\xa4\x6b\x63\x2e\xdc\x2a\x7d\x9e\x03\xa4\x70\x51\x11\x5e\x0c\x1f\x31\x2c\x42\x50\xa4\xf5\xc9\xda\x71\x47\xe9\xa7\xaa\x20\x9d\xd3\x5c\x27\xee\xef\xe3\xaf\xa0\x4e\x8a\x96\x35\x09\x3a\xc6\x0a\xaf\x08\x0a\xdb\x49\x85\x9e\xfa\xa8\xa6\x5e\xf9\x76\xa5\xf9\xe7\x2e\x6c\x22\x79\xe7\x29\x52\x83\x71\x15\xb4\x34\xbb\xfd\xe5\xad\x7c\xd9\xcb\x25\x79\xa6\x9b\xeb\xb1\x2d\xde\x03\x00\x00\xff\xff\x67\x2b\xbe\x8b\x9c\x04\x00\x00")

func artifact_typesBytes() ([]byte, error) {
	return bindataRead(
		_artifact_types,
		"artifact_types",
	)
}

func artifact_types() (*asset, error) {
	bytes, err := artifact_typesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifact_types", size: 1180, mode: os.FileMode(511), modTime: time.Unix(1476627076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _capability_types = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x57\xc1\x6e\xdb\x38\x10\xbd\xe7\x2b\x06\xd8\x6b\x6b\x24\x57\x1f\x16\x48\x9c\x62\x37\x40\x37\x0d\xea\xec\x5e\x16\x0b\x61\x4c\x8d\xac\x41\x28\x92\x1d\x0e\xd3\xba\x5f\xbf\x90\x4c\xd9\x8a\xed\xd4\x8a\x13\x20\x37\x81\x9c\x79\x1a\xbe\x79\x6f\x44\xa9\x8f\x06\x8b\x92\x2a\x76\xac\xec\x5d\x2c\x1e\x49\x22\x7b\x37\x85\xf5\x56\xe4\x26\x58\x2a\x56\xd8\xd8\xe2\xa2\x38\x2f\xce\xcf\xce\x0c\x06\x5c\xb0\x65\x5d\x15\xba\x0a\x14\xa7\x67\xb0\x0e\x9e\x6c\x76\x98\xe2\xe4\x52\x15\x4d\xdd\x90\xd3\x36\x00\xa0\x24\xe1\x47\x2a\x8b\x4a\x7c\x33\x3d\x94\xf0\xd5\x7b\x3d\x3b\x8c\x35\xf3\x4d\x48\x4a\x2f\x01\x6a\x23\x83\xf8\x40\xd2\xae\xad\x33\x01\x1c\x36\xd4\x3f\x03\xb4\xe5\x4f\x21\xaa\xb0\x5b\x6e\x16\x85\xbe\x25\x16\x2a\xa7\x50\xa1\x8d\xd4\x27\xa6\xa6\x30\x21\xc5\xdd\x64\x76\x4a\x4b\x92\x23\xd9\x00\xc6\xbb\xa8\x82\xec\x74\x00\x01\xf0\x11\x96\x42\xa8\x24\x85\x97\x82\xbe\x25\xb4\x53\xb8\xc8\xfb\x26\xa4\xa2\x6a\xf1\xc8\x99\xd5\x5e\xd1\x06\x2d\xca\xc7\xe4\x58\x27\x9b\xa0\xb7\xac\xe2\x7c\x72\x01\x7f\xfc\xf9\x33\x47\x95\x1c\x1f\x8a\xc8\x3f\xf7\xc9\x1b\xd4\xd1\xee\xbf\x69\x09\xf0\xd7\x55\x0e\x69\xa8\x79\x9f\xd7\x1f\xd6\xa3\x59\xeb\x71\x32\xf3\x4e\x91\x1d\xc9\x48\x65\x66\x1d\x3f\x83\xfa\xc9\x95\xc1\xf3\xcb\xfc\x72\x58\xe6\x41\xbc\x7a\xe3\xed\x68\xa9\xab\xa4\x2d\x47\x25\x55\x98\xac\x4e\x41\x4d\xe8\x01\xbd\xe8\x2e\xd8\x9d\x17\xbd\xa6\xea\x08\xe5\x91\x4c\x92\xbd\xbe\x2d\xbc\xb7\x84\xee\x68\xbb\x36\xa5\x0c\x97\x93\xd8\x22\xa0\xd6\x27\x19\xb9\x3d\x49\x71\xfa\x18\x20\xfd\xee\xe5\xe1\x74\x80\xc1\x91\xee\xbe\xde\xfc\x73\x79\xff\x29\x6f\x74\x03\x18\xd5\xcb\x6b\x61\xa3\x4f\x62\x46\x08\xfe\x11\x2d\x97\xc5\x23\xda\x44\x71\x0a\xff\xe6\xbc\x0f\xa0\x28\x4b\xd2\x0f\x10\x88\x04\xfe\x1b\xb0\xb6\x37\xfb\x1a\x0c\xa7\xfb\xad\x61\x57\x58\x72\x4b\xad\xb7\x13\x0f\x80\x9c\xca\xaa\x88\xa6\xa6\x06\x87\xf1\x5b\xc5\xcd\x03\x99\x6e\x03\x55\x85\x17\x49\xb7\xaa\xe7\x50\x60\x59\x0a\xc5\xbd\x4a\x33\x87\xbf\xb6\xdd\xe4\xb2\x6c\xd8\x8d\x34\x5f\x9f\xd4\x45\xff\x06\xb3\x1a\xdd\x92\xa0\x5f\xcd\xaa\x07\x76\x25\x9b\xb6\xa9\xa0\xbe\xf3\x18\xb4\x80\xc0\x1a\xfb\x7e\x81\xaf\x06\x8c\xed\x3b\x79\x9c\x7d\xb6\x8e\x1d\xfa\xf8\x59\xee\xf3\x80\xeb\x82\x8f\x70\x72\x8d\x8a\x0b\x8c\x63\x3f\xbd\x1b\x5a\x8e\xc0\xde\xa5\x85\x65\x73\x0a\xd7\xfb\x14\x6d\xd8\xd7\x9a\x36\xb4\x0e\x8d\xda\x92\x9f\xe2\x7a\xbf\x62\x89\x0a\xa1\x7b\x7d\x1f\x04\x95\x4f\xae\x7c\xb9\xc3\xb7\x4e\xfe\xfb\xea\xf3\xcd\x6c\x34\xef\x4f\xc2\x2b\xeb\x51\xd9\x2d\xa7\x03\xd8\x68\x84\x83\x76\xb7\xb0\xdf\x07\x00\x59\x4d\x14\x41\x6b\xd4\xee\x38\xf9\x20\x59\xf6\x10\x6b\x9f\x6c\x09\x0b\x02\xb4\xd6\xb7\xb1\xe5\x5a\x71\x08\xc1\x7b\xdb\x89\x2d\xbf\x0f\x6e\xee\x32\x0e\x0a\x01\xc6\xe8\x0d\x77\xf1\xdf\x59\xeb\x0e\x3b\x33\x31\x19\xab\xbd\xa7\xc6\x8f\x8a\x9a\xe2\x14\xe8\x47\x20\xe1\xf6\x1a\x88\xb6\xbf\x4c\xb8\xb8\x43\xef\x93\x23\xdf\xd7\x04\xbe\x7b\x46\x0b\x7d\xff\x84\x96\x1c\x95\x64\x5d\xde\xf5\xed\xfc\xb4\x29\x79\xb0\xaa\xc3\x5a\xbd\x5d\x1f\xff\x5d\xee\x9c\x87\x2b\xea\x1b\x72\xc5\xae\xc4\x85\x1d\xeb\xc9\x5b\x5f\x1e\x43\xfc\xcc\xee\xe1\x6d\x10\xdb\xad\xd7\xdf\xf7\xbf\x04\x92\x4e\xa5\xf3\x55\x54\x6a\x5e\xdf\x03\x14\x53\xb3\x92\xd1\x03\xc3\x74\x94\x74\xba\xd0\x53\x12\x4b\x8e\xeb\xef\x54\xab\xed\x53\x00\xfa\x1f\xb2\x9d\xdc\xbc\xfc\x42\x09\xb5\x14\xf5\x6c\xee\x58\xee\xfe\xcb\x7c\x76\x09\xe2\xbd\xc2\x6c\xf3\x87\x07\xf7\xab\xd0\x0d\x13\xf0\x5a\x93\xe4\xa0\xf6\x83\xb0\x1b\x14\x73\x7f\xba\x79\xf3\xcc\xdb\xe7\xed\x6d\x7d\xbc\xcc\x7e\xd1\xcf\xf6\xf6\xc0\x2e\x2a\x3a\x43\x47\xff\xc9\x36\x13\xaa\xbf\x67\x34\xf8\xe3\x15\xd9\x79\xe1\x38\xc2\x33\x24\xa8\x17\x5c\xbe\xcf\xbf\xec\xff\x01\x00\x00\xff\xff\xf7\x7a\xaf\xf5\xef\x0f\x00\x00")

func capability_typesBytes() ([]byte, error) {
	return bindataRead(
		_capability_types,
		"capability_types",
	)
}

func capability_types() (*asset, error) {
	bytes, err := capability_typesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "capability_types", size: 4079, mode: os.FileMode(511), modTime: time.Unix(1480694932, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _data_types = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x95\xcf\x6e\xdb\x30\x0c\xc6\xef\x79\x0a\x3e\x40\x16\xb4\x18\xb2\x83\x6f\x43\x7b\xe9\x65\x1b\xd6\xdc\x86\x42\xe0\x2c\x3a\x11\x22\x4b\x1a\x49\xa7\xc8\xdb\x0f\xfe\x93\xc5\x5d\xd4\xc6\x2d\x02\xb4\xa7\x00\x24\xf5\x91\xdf\x4f\x16\xa3\x51\x4a\x34\x96\x2a\x17\x9c\xba\x18\xc4\xec\x88\xc5\xc5\x50\x40\x9f\x12\x57\x27\x4f\x66\x8f\xb5\x37\xd7\xe6\xca\x5c\xcd\x66\x16\x15\x8d\xee\x13\x49\x31\x03\x80\xbe\x70\xd1\x46\xbb\xe0\xe2\x67\x8c\xda\x67\x00\x2c\x49\xc9\x2e\x69\xa7\xb8\xda\x10\xac\xbe\xdf\xdf\x7c\x05\x8e\x51\xe1\x16\x15\x61\xb5\x4f\x04\xe8\x3d\x44\xdd\x10\x0f\xe9\xdf\x28\x74\x4c\x0b\x58\x62\xb7\x23\xa8\x38\xd6\xb3\x6c\xcb\x1b\x26\x4b\x41\x1d\xfa\x63\xe3\xf6\x88\x35\xed\x99\x22\x3b\xe2\x50\x98\x38\x26\x62\x75\x07\x37\x43\x4c\x63\x19\xfd\x31\x02\xd0\x1e\x2c\x40\x94\x5d\x58\x8f\xc2\x4c\x7f\x1a\xc7\x64\x0b\xa8\xd0\x0b\xfd\xcb\x68\xdc\x52\xe8\x28\x4d\x10\xb1\x54\x61\xe3\xb5\x80\x84\x22\x8f\x91\xed\x53\x99\xb3\x0a\x5b\xda\xcb\x69\x51\x8d\x69\xc2\xa0\x00\x14\x94\xf7\x46\xca\x0d\xd5\x38\x56\x79\xa6\x59\x23\xc4\x6f\x00\x93\xbd\xb7\x95\xab\xe9\x2e\x28\xf1\xee\x42\x37\x27\x8a\xac\x46\x5d\x9d\xc1\xde\x46\x45\xb1\xce\x53\x51\x6e\x8e\x50\x28\xd8\xb7\x8b\x64\x9d\x06\xd2\xc7\xc8\xdb\xc5\xb7\xfe\xf7\x2e\x54\xf1\x22\x86\x07\x5d\x13\x30\x37\xed\x7f\xb7\x72\x28\x76\xf6\x6c\x29\x5a\xcb\x24\x42\x99\xef\xca\x3b\xd1\xd7\x7f\x3e\x2f\x52\xf9\x11\x59\x2f\x86\x24\x45\xd6\x69\x3c\xba\xca\x09\x30\x5e\xc1\xad\xc6\xd2\x0c\xec\x3e\x20\xe3\x5b\xaa\xf2\x88\x5d\x50\x5a\x13\x0f\xb9\x32\x06\x51\x46\x17\x74\x34\xda\x27\x70\xc1\x30\x86\x35\x15\xf0\x0b\xae\xe7\xf0\x65\xb9\xfc\xbc\x84\x87\xf3\x5d\xef\x13\x95\xef\xb7\x97\x9f\x3c\xec\xd1\xae\xd5\x72\xfc\x86\xb3\x96\x7b\xdb\x3b\xf4\xce\x9a\x1d\xfa\x86\xa4\xb5\xde\xd8\x34\x6f\x4f\xcf\xc1\xad\xeb\x04\x0f\xc7\x65\x8d\xbc\x26\x3d\x1d\x6e\x20\x3f\xe5\x5f\xa3\x53\x18\x28\x9f\xe8\x74\xe1\x49\x2b\xfd\x05\x37\xcf\x5c\xe2\xa1\x40\x62\xc3\x65\xa6\xf7\x74\x0f\xbd\xc2\x7b\x78\xf8\x1b\x00\x00\xff\xff\x2f\x6c\xda\xd7\xcd\x08\x00\x00")

func data_typesBytes() ([]byte, error) {
	return bindataRead(
		_data_types,
		"data_types",
	)
}

func data_types() (*asset, error) {
	bytes, err := data_typesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data_types", size: 2253, mode: os.FileMode(511), modTime: time.Unix(1476627095, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _group_types = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8f\xbb\x6a\xc5\x30\x0c\x86\x77\x3f\x85\x9e\xc0\x9c\xae\x67\x2b\x1d\x3a\x16\x7a\xb2\x1b\x63\xcb\x8d\xc0\xb1\x8c\xa4\x06\xfc\xf6\x25\x37\x32\x54\xe3\xf7\x5f\x24\x19\x6b\x8a\x21\x63\xa1\x46\x46\xdc\x34\xac\x28\x4a\xdc\x9e\x70\x48\x4a\x4b\xaf\x18\x46\x5c\x6a\x78\x0b\x8f\xf0\x70\xee\x47\xf8\xb7\x07\x1b\x1d\xf5\xe9\x00\xe0\x70\xfa\x1d\xab\xff\x66\xb6\x03\x03\x64\xd4\x24\xd4\x6d\xef\x9b\x66\x84\xe9\xeb\xf5\xf1\x0e\x9f\x9b\x13\xa6\xd1\x11\x62\xad\xc0\x36\xa3\xfc\x93\x14\x32\x0a\xad\x08\x45\x78\x39\xfb\xa8\x19\x4a\x89\xe9\x5a\xbc\xcd\xcb\x62\xcb\x51\xf2\x4d\x00\xb6\xdb\xce\x07\xfc\x9d\xf1\x8d\x33\xfa\x4a\x05\xd3\x48\x15\xfd\x95\x74\x7f\x01\x00\x00\xff\xff\xed\x2a\x99\xb7\x05\x01\x00\x00")

func group_typesBytes() ([]byte, error) {
	return bindataRead(
		_group_types,
		"group_types",
	)
}

func group_types() (*asset, error) {
	bytes, err := group_typesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "group_types", size: 261, mode: os.FileMode(511), modTime: time.Unix(1476627106, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _interface_types = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x93\xcd\x6e\xe3\x3a\x0c\x85\xf7\x7e\x0a\xbe\xc0\x35\x7a\xb7\xdd\x0d\xba\x9a\x55\x81\x69\xf6\x02\x63\xd1\x31\x01\x59\x14\x28\xc6\x85\xdf\x7e\xe0\xbf\x8c\xe3\xba\x18\x61\xb6\xd1\xe1\x77\x3e\x53\x91\x49\x6e\xd0\x79\x6a\x39\xb2\xb1\xc4\xec\x06\xd2\xcc\x12\x5f\x61\x39\xca\xdc\xa7\x40\x6e\xc4\x3e\xb8\xff\xdd\x8b\x7b\xa9\x2a\x8e\x46\xda\x62\x43\xce\xc6\x44\xf9\xb5\x82\x25\x5b\x3f\x0e\x72\xfd\x4b\xc4\xa6\x03\x00\x4f\xca\x03\x79\xd7\xaa\xf4\x2b\xb4\xa6\x68\x6c\xe3\x1c\x5a\x33\xb9\x51\x4e\x36\xf7\x5e\x3a\x82\xcb\xfb\xc7\xdb\x0f\x50\x11\x83\x9f\x1b\x14\x2e\x63\x22\xc0\x10\x40\xac\x23\x5d\x33\x57\xcc\x74\xc8\xe4\xb5\x13\xa6\xca\xea\xcc\x2e\x8a\xa7\x3a\x70\x4b\xcd\xd8\x04\xaa\x3f\x0c\xa3\x47\xf5\xdf\x0b\x1f\xbe\x6c\xce\x35\x4a\x68\xb4\xcc\x1c\x3e\x61\x03\xc2\xa3\x63\x4d\x83\x24\x52\x9c\x42\xf5\xc2\x90\xd8\xf2\xed\xae\xe5\x98\x6d\xe0\x48\xca\x86\x6a\xa5\x94\x39\xfc\x95\x20\xa9\x1c\x20\xe9\x38\xef\x29\x50\xf9\x3e\x96\xf4\x9e\x71\x76\x51\x4a\x61\x3e\xcd\x1d\xa7\xfa\xed\x79\x59\xa5\xf7\x94\x94\xdc\x63\x6d\x2e\xcb\x5d\x9b\x73\xcb\xf7\xcd\x05\x4c\xa6\xa9\xff\xfe\x2c\xdb\x3a\x82\x65\x12\x28\xfa\x24\x1c\xad\x3e\x81\x1b\xea\x8d\xce\x6f\xe1\x2f\xf0\x65\xf2\x08\x97\x6c\xff\xa2\x2e\xd9\x8a\xdc\x9f\xf1\xc5\xf2\x5f\xf1\xa7\xf6\xe8\x7d\x31\x33\x8a\x71\x3b\xee\x55\xa7\x27\x0a\xd2\x02\x6e\xf0\xf9\x87\x2b\x71\xbc\x4d\x64\xf2\x30\x30\x02\xc2\xd3\x3f\xe4\xd1\x5b\xba\xaa\x5d\xef\xbe\x66\xee\xdd\x8b\x7c\x76\xdc\x74\xc0\x19\xa2\x7c\x02\x0e\xc8\x01\xaf\x81\xbe\x55\x58\x58\xae\xe9\x30\xde\xc8\x17\x6b\xac\x8d\x59\x7a\x82\xa4\xd3\xd3\xb0\x11\x44\x01\xcd\x94\xaf\x77\x9b\xc5\x76\xae\x2b\x7f\xc6\x2b\xf5\x32\x94\x5f\xe2\x12\x7f\xde\x6e\x5d\xfd\x0e\x00\x00\xff\xff\xca\x8d\x54\x44\x0e\x06\x00\x00")

func interface_typesBytes() ([]byte, error) {
	return bindataRead(
		_interface_types,
		"interface_types",
	)
}

func interface_types() (*asset, error) {
	bytes, err := interface_typesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "interface_types", size: 1550, mode: os.FileMode(511), modTime: time.Unix(1476627117, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _node_types = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x59\xcd\x8e\xdb\x38\x0c\xbe\xcf\x53\x08\xe8\x65\x17\xe8\x04\x53\xa0\xd8\x43\x6e\x93\x49\x77\xb7\x40\xe7\x07\xcd\x14\x3d\x14\x03\x83\x91\x98\x44\x5b\x59\x74\x25\x3a\xd3\xec\x69\x5f\x63\x5f\x6f\x9f\x64\x21\xff\xc5\x89\x9d\xd8\x71\x3a\xed\xa5\x19\x9b\xa4\x28\xf2\xfb\x28\x8a\x66\xf2\x12\x22\x85\x0b\x6d\x35\x6b\xb2\x3e\x5a\xa3\xf3\x9a\xec\x58\xe4\xaf\xbc\x8e\x13\x83\xd1\x06\x62\x13\xbd\x89\xae\xa2\xab\x8b\x0b\x4b\x0a\x23\xde\x24\xe8\xc7\x17\x22\x17\x1b\x85\x67\x7e\x34\x63\x72\xb0\xc4\xd1\xc4\x90\xfc\x5a\xfc\x11\x64\x84\x50\xe8\xf4\x1a\x55\xb4\x70\x14\x8f\x77\x74\x3e\x12\x71\x26\x92\x38\x4a\xd0\xb1\xce\xcd\x86\x7f\x5e\xff\x8d\xe5\x6f\x21\xc2\x8a\x63\xe1\x25\x18\x70\x97\xa9\xd5\x3c\x0a\xef\xab\xd7\x92\xac\x67\x07\xda\xb2\xdf\xea\x08\x71\x29\x96\x0e\x81\xd1\x45\xe4\x22\xfc\x96\x82\x19\x8b\x37\xe2\x76\x52\x88\xac\xc9\xa4\x31\x46\x5a\x35\xd6\x61\xa7\xed\xb2\x7a\xe8\xf0\x5b\xaa\x1d\xaa\xb1\x58\x80\xf1\xe5\xaa\xde\x42\xe2\x57\xc4\xc3\xf4\x25\x24\x30\xd7\x46\xd7\xb7\x0c\xcc\x20\x57\x31\x5a\xde\x37\x98\xc7\xac\xae\x33\xba\xae\x84\x2f\xf6\x12\x71\x43\x71\x92\x72\xff\xd8\x03\xb3\xd3\xf3\x94\xb7\x8e\x24\x4e\xaf\x81\x31\x02\xa5\x1c\x7a\x7f\x74\x7b\x49\x3a\x37\x5a\xf6\x12\xb5\xc8\xcf\xe4\xbe\x36\x84\x62\x48\xaa\x27\x68\xd9\x6d\x22\x2f\x57\x18\x43\x3d\x95\xf5\x38\x28\x60\xc8\x20\x38\x2a\x2c\x8e\xee\xf2\xff\xdf\xdb\x05\x95\x6e\x91\xe3\x17\x58\xe8\x81\x1c\x57\xab\x14\x79\x0d\x39\xa8\x96\xba\x14\x86\x24\x98\xc8\xd7\x09\x50\x80\xb4\x4c\xdf\xa6\x2b\xa1\x5b\x9d\x90\xaa\xdd\xb4\xd5\xe9\x55\x13\x74\x68\x20\xe3\xf0\x4a\x27\xa5\x42\xfd\x59\x69\x1f\xfd\x23\xd5\xd4\x48\xca\xd4\x39\xb4\x12\xfd\x58\x7c\xb9\x7a\x2d\x3e\xdd\x4d\xee\x3f\xdd\x4d\xdf\x4d\x9f\x0e\xc2\x74\x45\xbe\x0f\x40\x6f\xc8\x32\x68\x8b\xae\x12\x5d\x83\xd1\x2a\xf2\x94\x3a\x59\xd6\x10\xf1\x65\xa7\x84\xd0\x82\x9f\xc1\x61\x40\x30\x59\xb4\xfc\x74\x51\x26\x4b\x25\xa4\x7b\xf1\xe2\x5d\x21\x3a\xba\x56\xb1\xb6\x85\x3c\x35\xa0\xd0\xa2\x79\x9f\xa0\x03\xd6\x76\x39\xdb\x78\xc6\xb8\xa4\x79\xa8\x39\x73\xd3\xa8\x45\x2d\x06\x66\x85\x68\x21\x39\xd7\x56\x69\xbb\xec\xa1\x58\xa2\x6b\xa2\xad\xca\x0c\x34\x18\x5d\x84\x72\x74\x9d\x24\x46\xcb\x2c\xab\xbd\xf9\xdd\x0e\xd3\xdd\x2c\x76\xa2\xb3\x99\xcd\x36\x70\x6e\x1d\xfd\x98\x5a\xd6\xf1\x69\x08\xfd\x93\x3c\xa3\xba\xb7\x95\x8f\xa7\x93\x68\x97\x18\x97\x65\xc9\x39\xc1\xc2\x3b\xab\x1e\x02\x7e\x0e\xe7\xa0\xd8\x5a\x57\xfc\x1b\x60\xfe\xd1\x84\x1a\x82\xcc\xbd\x4d\x4d\x81\x61\x0e\xfe\x9c\x73\xda\x42\xdc\x3c\xa7\x77\xcf\x3f\x85\x5e\x3a\x9d\x70\xde\x59\xac\x50\x18\x5a\x6a\x09\x26\xd3\x15\xb4\xc8\x9e\xa9\xc2\x95\x5a\xfd\xde\x37\xab\x2d\xe3\xb2\x86\xc0\x86\xdd\xa0\x94\xfd\x48\xad\x42\x67\x36\xda\x2e\x2b\xbb\xc2\xa3\x5b\x6b\x89\xe2\x59\x1b\x23\x8c\xf6\x8c\x56\x30\x89\x05\xb9\x4c\xa6\xb0\x9a\x7a\x74\xa7\x6e\x87\xb2\xdf\x60\x32\x65\x01\x52\x52\x6a\x39\xdf\x5c\xb0\x3e\x9d\x08\x08\xb5\x48\x87\x16\x25\x48\x76\x34\x16\x09\x78\xff\x4c\xae\xab\xab\x38\xec\x46\x69\x20\x5b\x3d\xbc\x99\x4e\x76\x5c\x3b\xba\xfe\x4f\xac\x16\xd3\xc9\xed\x6c\x78\x81\x68\x63\x52\x99\xec\x68\xc8\x81\x51\x92\xa1\x41\x92\xc9\xed\x6c\x18\xd9\x9b\x6c\x71\x44\x1c\xf5\x4c\x70\x3b\x3a\x8e\x25\x3e\x58\x6f\xcb\xfe\xed\xac\x44\xff\x09\xec\xea\xbd\x7c\xc5\xba\xfa\x42\x87\x68\x26\xc0\xaa\xcc\x34\x7a\xf6\x3f\xa9\xc5\x68\x2d\x79\xe2\x69\x3f\xcf\x1f\x08\xd4\x04\x0c\x58\x59\x96\x80\x41\x05\x11\xcc\x92\x9c\xe6\x55\x3c\x30\xbd\x9e\x81\x53\x3f\x16\xf8\x3d\x41\xa7\x03\x0f\xc1\x1c\x0c\x94\x34\xba\xdf\x75\xa1\x42\xf9\x43\xd6\xaf\x57\x0a\x5d\x0d\x60\x6b\xc6\x7f\x37\x94\x35\x4a\xe2\x97\xf7\x0f\xbf\x16\x3e\xfc\xf7\xcf\xbf\x5e\x90\xcd\x11\x91\xad\x51\x1e\xbe\x42\x82\x0d\x37\x34\x8b\x92\x05\x1f\x6b\x9d\x61\xbf\xbb\xe9\x55\x6c\xca\xad\x9d\x54\x4a\x3e\x52\xb8\xf0\x9c\xd4\x0b\xb7\x04\xe3\x26\xdf\x95\xa6\x0c\xe3\x64\x51\x90\x13\x31\xb9\x70\xc4\x81\x12\xf3\x1c\x4d\xaa\xbe\x31\xbf\x0f\xbb\xbd\x3b\xcc\x19\xc8\xd3\x49\x75\x81\x1f\xcc\xed\x05\xa4\x86\xc7\xe2\x6d\xf7\xd5\x3a\xa7\xda\x1a\x4c\x9a\x93\xec\xed\x6b\xf1\x9b\x28\x23\x25\xb5\xea\x3a\x46\x0f\xdc\xaa\x19\x1c\x47\x3a\x19\xa4\x8c\x56\x0d\x55\x5d\x02\xe3\x33\x6c\x86\xaa\x17\x49\x8c\x7a\x74\x43\xc7\x0d\x0c\x1d\x46\xe0\x32\xab\x14\x01\x61\x43\x6d\x94\x2e\x64\x2a\x43\x0c\x24\xab\x8d\x0f\x7d\x5d\xd4\xe8\xba\xcf\x1a\x8b\x18\x6d\x1b\xa6\x8e\x5c\xa2\x3e\x68\xfb\xb5\xad\xd7\xad\x5f\xe1\xcf\x23\x59\x9f\x59\xc7\xa1\x18\x91\x53\xcd\x06\xf3\x30\x3d\xd9\xa5\x2d\xec\xbc\xea\x66\x67\x73\xf0\x55\x2a\x69\x1f\x95\x86\xf6\xdc\x98\x13\x19\x84\xae\xf6\xb4\xe6\x47\xfd\xb1\x4e\x22\x07\x76\x89\x51\xc6\xe0\x41\xb1\xa9\x4c\xa0\x1d\x82\xe0\x96\x33\xe5\x72\x0f\x3c\x1d\x27\x49\x03\x42\xdb\xe5\xba\xce\x93\xba\x6a\x75\xac\x5c\x36\x67\x00\x3d\x3d\xa8\x26\x01\xa7\x7b\x10\x54\x83\x07\x07\xc6\xb3\xf7\xf3\xbf\x50\xf2\xf9\xf3\xd9\xce\x4a\xf7\x52\x03\xdc\x2b\xf1\x47\x39\xc0\x8d\xe1\xfb\x8b\xaf\xd2\x56\x91\x8a\xb9\xc4\x90\x5b\xc6\x7e\x5a\x42\x84\x0f\x67\x01\x2d\x6b\xde\x6c\xd3\xb0\xd3\x7c\x3c\xae\x50\x3c\xde\xcf\x6e\xae\xc5\x1d\x29\x14\x8f\x9b\x04\x05\x18\x23\x88\x57\xe8\x8a\x37\x59\xa7\x5b\xbd\xf6\xc5\x22\x22\xac\x91\x19\x6c\xce\x7e\xf3\xc9\x7f\xc7\x11\x92\x0b\x75\x43\x80\x81\x8f\x49\xb4\xc5\x76\x81\xc0\xa9\xeb\x33\xd4\x08\xdb\x3a\xd2\x4c\x2a\x4c\xd0\x2a\xb4\x72\x73\x42\x2f\x59\xd9\x2c\x40\xde\xb8\xb3\x56\xa9\xe8\x4b\xcb\x69\xe6\x86\xbf\xb7\x07\x3b\x4d\x51\x6f\x35\x8b\x0e\x2a\x9c\x08\x6e\x01\x72\x1b\x97\x19\x83\x55\xd0\xbc\x34\xe6\x6b\x6e\xe5\x33\x3f\x47\x46\x2f\x50\x6e\xa4\xc1\x51\xa9\xd7\xa8\x07\xfb\x37\xd6\x33\x6a\xc1\x2b\xa1\x28\x06\x6d\x2f\x7d\x82\x52\x2f\xb4\x14\xbe\xb0\x2e\x64\x69\x5e\x14\xad\x69\xd9\x20\x96\xcf\x0f\xb5\xac\xbb\xe2\x87\x8e\x8c\x6c\xb0\x12\x49\x87\x2a\x50\x05\x4c\x7b\x74\xb6\x73\xfc\x9b\x4a\xf2\xd4\xa3\xe4\xe5\x46\xa6\xd9\xd7\x9a\x93\x30\x55\xcd\x41\xf6\x72\xfa\x19\xe7\x43\x06\xc4\xcd\x84\x4a\xb2\x8c\xdf\x39\x72\x55\x75\xea\xcd\x5f\x48\x92\x41\x75\xf1\x27\x87\xfd\x33\xce\x67\xe8\xd6\x3b\x52\x67\x05\x3e\xb7\xf6\xe3\x86\xc2\xaf\xc4\x43\xfe\x15\xee\xb5\x30\xb0\x41\x27\xde\x56\x5f\x41\x7c\x21\x12\x60\xbd\x8d\x75\x8f\xab\x71\xce\x96\x5e\x1a\x3b\x9f\x4f\x5e\x70\x1c\xb3\x0b\x59\xf1\x74\xf1\x7f\x00\x00\x00\xff\xff\x68\xd7\xc5\x36\x95\x1e\x00\x00")

func node_typesBytes() ([]byte, error) {
	return bindataRead(
		_node_types,
		"node_types",
	)
}

func node_types() (*asset, error) {
	bytes, err := node_typesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node_types", size: 7829, mode: os.FileMode(511), modTime: time.Unix(1480697558, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _policy_types = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x92\x41\x6a\xf3\x30\x10\x85\xf7\x3e\xc5\x3b\x41\xc8\xbf\xcd\xee\xa7\x07\x68\x68\xd2\xb5\x10\xd2\x38\x19\x90\x35\xea\x8c\x6c\xf0\xed\x8b\xe5\x96\x14\x6a\x68\xc9\xa2\xd9\xce\xd3\xbc\xef\x93\x50\x15\x0b\xde\x45\xea\x39\x73\x65\xc9\xe6\x26\x52\x63\xc9\x07\xac\x91\xf1\x50\x12\xb9\xd9\x0f\xc9\xfd\x73\x7b\xb7\xef\xba\x22\x89\xc3\xec\xea\x5c\xc8\x0e\x1d\x80\xf5\xe8\xae\xcd\x99\x6c\xf7\x22\x52\xd7\x00\x88\x64\x41\xb9\xd4\x56\x79\xbe\x12\xce\xcf\xa7\xa7\xff\x38\xb6\x0e\x9c\xe7\x42\xf0\x29\x41\xea\x95\xf4\x7b\x66\x88\xa4\x3c\x11\x7a\x95\xa1\xdb\x62\x1d\x93\x0f\x34\x50\xfe\x02\x5c\x16\xa2\x5b\x36\x0e\x5b\x66\xbf\x17\xbb\x3d\x0b\xea\xd5\x57\xb0\x61\x34\x8a\xa8\x82\x8b\x4c\xa4\x19\xe5\x93\x0e\xe9\x3f\x0a\xb2\x44\x32\x88\xe2\xa2\x32\x16\x5b\x82\x36\xda\x6d\xea\x9f\x82\x4f\x9c\x2f\x0f\x91\xb7\x95\x7d\xaf\xfa\x6b\x89\xbe\xd2\x43\xcc\xc7\x86\xbe\x57\xfc\x48\xda\x8b\x0e\x3e\x87\x3f\xb6\x8f\x14\x92\x57\x42\xb9\x09\x40\xe9\x6d\x64\x6d\x7f\xc8\xd0\x8b\xfe\x78\xa3\xf7\x00\x00\x00\xff\xff\x8b\xf6\x26\x41\xb2\x03\x00\x00")

func policy_typesBytes() ([]byte, error) {
	return bindataRead(
		_policy_types,
		"policy_types",
	)
}

func policy_types() (*asset, error) {
	bytes, err := policy_typesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy_types", size: 946, mode: os.FileMode(511), modTime: time.Unix(1476627148, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _relationship_types = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x54\x4d\x6f\x13\x41\x0c\xbd\xe7\x57\xf8\x0f\x10\xb5\xd7\xbd\x95\x80\xc4\x01\x51\xa9\xe4\x86\xaa\x91\x33\xe3\x24\x56\x67\xed\xc1\xe3\x04\xe5\xdf\xa3\xdd\xcd\x17\x94\x2c\x04\x29\x3d\xee\xfa\x3d\x3f\xbf\x37\xb2\x5d\x6b\xc4\x90\x68\xc9\xc2\xce\x2a\x35\x6c\xc9\x2a\xab\x34\x30\x94\x2a\xb7\x25\x53\xd8\x61\x9b\xc3\x7d\xb8\x0b\x77\x93\x89\x51\xc6\x1e\xbb\xe6\x12\x7c\x57\xa8\x36\x13\x18\xe0\xd3\xf3\x5a\x9d\x3e\xb8\x63\x5c\x53\x9d\x6b\x87\x00\x48\x64\xbc\xa5\x14\x96\xa6\x6d\xf3\x47\xc6\x93\xaa\xf7\xd0\x2d\x66\x4e\xc1\xd1\x56\xe4\x7b\x11\xf8\xb6\xa7\x44\x2c\xb8\xe0\xcc\xce\x74\xd0\x68\x49\x1c\x9e\x7b\x66\x31\x2d\x64\x5d\x6d\x10\x05\xc8\x1a\x7b\x89\xc3\x37\x40\xd7\xb0\x81\xea\xc6\xb2\x3a\xfe\x8c\x2a\xd5\x0d\x59\xbc\x9e\x90\x00\xef\xa0\x65\x09\x99\x64\xe5\xeb\x06\xee\xf7\x95\x44\x5b\x8e\xf4\x97\x8e\x46\xdf\x37\x6c\x94\x1a\x58\x62\xae\x34\xb9\x10\xd3\x4c\x45\x28\xfa\x4d\x63\xfa\x28\xa9\x28\x8f\x84\x14\x8d\x12\x89\x33\xe6\xdf\x4d\x0d\xdd\x12\x3a\xf6\x02\xd3\xd9\x11\x79\xad\xd1\x0f\x54\x48\x52\x7d\x94\xdb\xf9\xfc\xa2\x89\xe0\xf9\xd2\x00\x9f\xb4\x3a\xa5\x5b\xea\xcf\x54\x1c\x59\xc8\x2e\x0f\x21\xe4\x3f\xd4\x5e\xa6\xef\x59\xd2\x15\x6f\x7e\x0c\xef\xba\x81\xce\xd5\x70\x91\x47\xc2\x39\x20\x3f\xb3\xbc\xbc\xdd\x5c\x9d\xda\xf8\x5c\xdd\x23\x1c\x86\xa9\xd1\xb8\xf4\xbb\x0c\xf3\x35\xc1\xfc\xf1\xeb\xec\x01\x4c\xd5\xe1\xe9\x8c\x03\xf3\x5d\x21\xc0\x9c\x41\x7d\x4d\xb6\x87\x2d\xb0\xd2\x6b\x58\xdd\x5b\x84\xce\x61\xaf\x82\xee\xc6\x8b\x8d\x9f\x76\x63\x38\x85\x9c\x46\xd7\x7d\x00\x09\xb6\x63\x57\x81\xc5\xc9\x96\x18\x4f\xbd\x67\x2a\x4b\x5e\x6d\xec\x15\x6b\x88\xe2\x44\xf8\x25\x95\xe9\x91\x76\x39\xb5\xce\xc1\xbf\x3f\xe3\xe9\x08\xfd\xf7\x61\xf9\x19\x00\x00\xff\xff\xf9\x61\xcf\xc6\x4d\x06\x00\x00")

func relationship_typesBytes() ([]byte, error) {
	return bindataRead(
		_relationship_types,
		"relationship_types",
	)
}

func relationship_types() (*asset, error) {
	bytes, err := relationship_typesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "relationship_types", size: 1613, mode: os.FileMode(511), modTime: time.Unix(1480697573, 0)}
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
	"artifact_types":     artifact_types,
	"capability_types":   capability_types,
	"data_types":         data_types,
	"group_types":        group_types,
	"interface_types":    interface_types,
	"node_types":         node_types,
	"policy_types":       policy_types,
	"relationship_types": relationship_types,
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
	"artifact_types":     &bintree{artifact_types, map[string]*bintree{}},
	"capability_types":   &bintree{capability_types, map[string]*bintree{}},
	"data_types":         &bintree{data_types, map[string]*bintree{}},
	"group_types":        &bintree{group_types, map[string]*bintree{}},
	"interface_types":    &bintree{interface_types, map[string]*bintree{}},
	"node_types":         &bintree{node_types, map[string]*bintree{}},
	"policy_types":       &bintree{policy_types, map[string]*bintree{}},
	"relationship_types": &bintree{relationship_types, map[string]*bintree{}},
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
