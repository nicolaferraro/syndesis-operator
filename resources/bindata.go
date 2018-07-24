// Code generated by go-bindata. DO NOT EDIT.
// sources:
// bindata.go
// template.yaml
package resources

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

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(436), modTime: time.Unix(1532094278, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xbd\x7b\x73\xdb\x38\xb2\x28\xfe\xbf\x3f\x05\x7e\x8a\xcf\x4f\xc9\x59\x53\xb2\x9d\xf1\x3c\x38\xeb\xbb\xa5\xc8\x8c\xa3\x19\x5b\xd2\x8a\x72\x66\x72\xf7\xce\x55\x41\x24\x24\x21\xa6\x08\x0e\x00\xca\xf6\x66\xf2\xdd\x6f\x01\xe0\x03\x24\x41\x49\x7e\xc6\x3e\xb5\xaa\xa9\x9a\x98\x00\x1a\x8d\xee\x46\xa3\xbb\xd1\x00\x5e\xed\xbc\x02\x27\x04\x84\x84\x03\xe4\x63\xbe\x07\xf8\x02\x33\x80\x19\x80\x60\x8e\x42\x44\x21\x47\x3e\x98\xe1\x00\xb5\x00\x18\x13\x40\x51\xfa\x75\x0f\x00\x1a\x87\x36\x68\xb5\x93\x2f\x84\xb6\x69\x1c\xb6\xd8\x62\xe7\xd5\x0e\x8c\xf0\x47\x44\x19\x26\xa1\x0d\x38\x5a\x46\x01\xe4\xa8\x45\x22\x14\xb2\x05\x9e\xf1\x16\x26\xed\xd5\xc1\xce\x25\x0e\x7d\x1b\x8c\x93\xe2\x9d\x25\xe2\xd0\x87\x1c\xda\x3b\x00\x84\x70\x89\x6c\xc0\x6e\x42\x1f\x31\xcc\x76\x04\x92\x69\x3d\x30\x0b\xe0\x8a\x50\x5b\x7e\xec\x04\x01\xb9\x02\x01\xf1\x60\xb0\x20\x8c\x03\x8a\x66\xcc\x06\x33\x18\x30\xb4\xf3\x0a\x5c\x30\x04\x7c\xe2\x5d\x22\x0a\xf0\x12\xce\x91\x56\xe4\x26\xa0\xc1\x18\xce\x6d\x20\xe0\x32\xbe\xf3\x0a\x0c\x29\x59\x22\xbe\x40\x71\x52\xb0\x3a\x6c\x1d\xb4\xf6\x45\x01\x61\x7c\x4e\x11\xfb\x33\x50\x05\x3f\xb5\x8e\x76\x5e\x81\x41\x27\xe6\x8b\x21\x25\xd7\x37\x49\xf5\x03\x55\x7d\x07\x80\x00\x4e\x51\xc0\xc4\x58\x00\x80\x51\xa4\x0d\x46\x7c\x49\xff\x10\x94\x58\x5f\xca\x6f\x22\x64\x03\x1c\xce\x28\x64\x9c\xc6\x1e\x8f\x29\xda\x89\x20\x85\x4b\xc4\x11\x65\xb6\xc0\x2d\xfb\x0b\x5c\x61\xbe\x00\x21\x01\x3e\x9a\xc1\x38\xe0\x0c\xbc\x5e\x40\x06\x38\x01\x53\x24\xfe\x8b\x28\x59\x61\x1f\xf9\xe0\x6a\x81\xc2\x8c\x31\x92\xdb\x51\x14\x60\xe4\xbf\xb1\x77\xac\x84\xf8\xa3\xc1\xc5\xd8\x99\x7c\x18\xb8\xe3\x7e\xe7\xdc\xd9\x01\xc0\x47\xcc\xa3\x38\xe2\x92\xa9\xe3\x05\x02\xe8\x9a\x23\x1a\xc2\x00\x08\xd2\x8b\x46\xa2\x23\xe8\x79\x88\xb1\x8c\xbe\x3b\x00\x50\xf4\x67\x8c\x29\xf2\x6d\xc0\x69\x8c\x0c\xf8\xa6\xc8\xe6\x7d\x0f\x86\x4e\xdf\xfd\xd0\x7b\x3f\x9e\x9c\x77\xdc\xb1\x33\x2a\xf7\x3e\x8c\xa7\x01\xf6\xc0\x20\x42\xa1\x2b\x04\x0a\x2c\x21\xe3\x88\x02\xe8\xfb\x14\x31\xd1\xe9\x0a\x06\x31\xb2\xc1\x82\xf3\x88\xd9\xed\x76\x26\x21\xf6\x8f\xdf\x7d\xf7\xb6\x8a\x54\xb5\xe3\xee\xa0\xef\x0e\xce\x9c\xc9\xc5\xe8\xcc\x34\xf6\x8b\xd1\x99\x18\x2d\x5f\x20\x0d\x0b\x8f\x84\x8c\x04\x48\xd4\xc7\x2c\x0a\xe0\x4d\x5f\x01\xcd\x2a\x74\x55\x05\xd1\xda\xd0\xe5\x70\x34\xf8\xc5\xe9\x8e\x4d\xdd\x49\xea\x92\x59\xa9\xbf\x88\x92\xcf\xc8\xe3\xb9\x30\x63\x06\xa6\x08\x87\x73\xe0\xa3\x28\x20\x37\xc8\x07\x38\xe4\xa4\x55\x8f\x50\x0a\x81\x93\xa4\x89\x6c\xb0\x0d\x7d\x06\x9d\x8b\xf1\x87\x49\xf7\xac\xe7\xf4\xc7\x13\xd7\xe9\x8e\x9c\x0a\xe2\x79\x37\x72\x9e\x00\x2f\xc0\x28\xe4\x80\x21\x8f\x22\xbe\x03\x32\xfd\x62\x03\x74\x1d\x09\xbe\x61\x12\xee\x00\x30\xa3\x64\x69\x83\xc6\xbf\xa0\xf5\xef\x8e\xf5\xbf\xf7\xad\x9f\xfe\xf8\xf2\xfd\x77\x5f\x1b\x26\xa4\x0a\xdd\x9d\xc3\x6b\xbc\x8c\x97\x00\x2e\x49\x1c\x72\x41\xad\x25\x5a\x12\x7a\x23\x89\x96\xcc\x5f\xf7\x9f\x67\x82\x4b\x1c\xe2\x10\x51\xe0\xc1\x10\xc4\x0c\x55\xe8\x73\xae\xda\x9d\xe1\x25\xe6\x99\x32\x1a\x0e\xdc\xf1\xe9\xc8\x71\xff\x79\x36\x39\x77\xce\x07\xa3\x4f\x93\xb3\xde\x79\x6f\x9c\xcb\xda\xe1\xd1\xd1\x39\x2e\x23\x35\x2e\x30\x4c\x80\x67\x11\xf4\x90\x98\x80\x14\x95\x31\xeb\x09\x25\xe5\x72\x8a\xe0\x12\x50\xc4\xb0\x8f\x58\x05\xb7\x0c\x84\x09\xb1\xde\x79\xe7\xd4\x99\xb8\xe3\x91\xd3\x39\x9f\x88\x69\xeb\x0e\x3b\x5d\x27\x47\x31\xd3\xc0\x65\x34\x2f\x98\x98\xc9\x4b\x04\x66\x84\xea\x18\xc5\x0c\x51\xc0\x17\x90\x83\x2b\x1c\x04\x42\x85\xc4\x4c\xac\x07\x84\x26\x13\x5d\x08\x9b\x18\x85\x50\xdc\x53\x68\xa0\xa5\x06\xac\x4b\xc2\x10\x79\xa2\xc3\xac\x3f\xd3\x20\x2e\x5c\x39\xe1\x13\x94\x33\xc5\x58\xc2\x78\x08\x19\xbb\x22\x54\xe1\x52\xe5\x70\xda\x91\x18\xc0\x96\x38\xa5\x10\x33\x11\x6c\xea\x22\x78\xf0\xfd\xd7\x66\xbd\xcc\x56\x06\x31\xec\xb8\xee\x6f\x83\xd1\xc9\x46\xa9\xed\x6b\x33\x5b\xc3\x2b\xa5\x67\x42\x66\xe4\xaf\x1b\xc4\x49\x5a\xb9\x5f\x43\xd2\x93\xce\xb8\xf3\xae\xe3\x3a\x55\x6c\x36\xd2\xf9\x23\x09\xe2\x25\x02\x4a\x6a\xe1\x0a\xe2\x00\x4e\x83\x8a\x9c\x08\x74\xf7\x00\x6a\xcd\x5b\xe0\xe8\xe0\xf0\x1c\xef\x81\xc3\x53\xbc\x0e\xe5\x04\x6c\x17\x46\xd0\xc3\xfc\xc6\x84\xf4\xc7\xc1\xd9\xc5\xb9\x33\xe9\x76\x86\x9d\x6e\x6f\xfc\x69\x0d\xee\x07\xa7\x95\x79\xb7\x4e\x3c\x18\x5c\x46\x01\xf2\xa7\x1b\x85\xc3\x95\x15\x4f\xde\x3d\x9a\x94\xb8\x9d\xf3\xe1\x99\x73\xf2\x6e\x7b\x71\x71\x42\x41\x7e\x06\x84\xbd\x62\xb1\x38\x8a\x08\xe5\x00\x85\x7e\x44\xb0\xd0\x79\x21\x98\x42\xef\x12\x85\x3e\xe8\x0c\x7b\xe5\x91\x8d\x11\xe3\xc0\x4d\xda\x28\x40\x7e\x86\xd5\xd8\x71\xc7\x13\xf7\x62\x38\x1c\x8c\xc6\x13\xa7\xdf\x79\x77\xe6\x18\xb0\xc9\x48\xde\x90\x86\x54\xa3\x0e\x3d\xc6\x21\xe5\x42\x37\xc4\x51\xba\xc6\x2f\x89\x94\x92\x32\x52\x27\xa2\x40\x08\x70\x05\xa3\x13\xe7\x7c\x20\x25\xf7\xee\xe8\x8c\xd0\x1c\x33\x4e\x6f\x24\xab\x52\x9d\x4b\xc0\x0c\x71\x6f\xa1\xad\x9a\xd2\x36\x2c\x23\x96\x15\x4b\xad\x9c\x81\xca\xd0\x73\x3f\xf5\x4f\x1c\xb7\xe7\x4e\x46\xce\x69\xcf\x1d\x8f\x3e\xe5\xd8\x34\x95\xd1\xd9\xc2\xa4\x59\xc6\xc8\x5d\x90\x38\xf0\x93\xc5\x76\x89\xd4\x3a\x85\x43\x8e\xe6\x14\x8a\x1a\x62\x01\x07\x48\x51\xe2\x1f\x65\x94\x14\x85\x40\x2f\xaf\x0e\x4e\x32\x40\x19\x62\xdd\x41\x7f\x3c\x1a\x9c\x9d\x39\x23\x77\xd2\xeb\x8f\x9d\xd3\x51\x67\xdc\x1b\xf4\x35\x2a\xa6\x68\x0a\x1a\x56\x30\xcc\xd7\xa8\x64\xa1\x14\x6c\x94\x24\x02\x4c\x2e\x4d\x15\x4a\x29\x02\x25\xeb\x56\x75\x7d\xda\xb4\x28\x35\xab\x44\x92\x16\x82\xe0\x54\xcc\x24\xc3\x50\xe8\xd1\x9b\x88\x03\x02\xa5\x1d\x41\xc8\x25\xae\x32\x4c\x19\x19\x5d\x59\x98\x80\xc8\x70\x48\xac\x95\xc1\xe0\xd7\x9e\x93\x5b\x2b\x6b\x0d\x90\xc2\xbc\x7e\x7b\xf8\xb5\x99\x19\x41\x19\xe7\x9d\x7e\x77\xf4\x69\x38\x9e\xfc\xea\x7c\x32\xda\xc7\x0a\x6b\xc1\xa6\x4b\x74\xa3\x56\xce\x7c\x30\x6d\x1f\xa9\x41\x31\x4e\x28\xf2\x13\xb3\x88\x3d\xbc\x5d\xb4\x4e\x83\xe7\xee\xce\x96\x1a\x3c\x6f\x50\xab\xc1\x47\x83\x73\x67\xfc\xc1\xb9\x70\x0d\x1a\x5c\x53\xd8\xf7\x33\xe7\x72\x3c\xee\x6a\xce\xe5\x68\x9a\xcd\x39\x49\x85\xfb\x10\xf7\x1c\x71\xb8\x25\x59\x65\xd5\x3a\x82\x9e\x3b\xe3\xce\x2d\x17\xc3\x7b\xd1\x36\xb5\x04\x2c\xe1\x90\x03\x86\xe8\x0a\x7b\x08\x2c\xf1\x7c\xc1\xb7\x27\xaf\x44\xfa\x3e\x84\xdd\x0e\x45\x81\x1d\xa2\x77\x45\xd2\x75\x46\x1f\x9d\x51\x0d\x9a\x3f\xee\xef\xdf\x02\xcd\x30\x5e\x4e\x11\xad\x28\x73\x61\x1f\x07\x48\x99\xd1\x42\x3c\x3d\x8a\x20\xaf\x38\x88\xeb\x81\xe4\x34\xed\xfc\xae\x2b\x74\x77\x32\x74\x46\x25\x6b\xb9\x71\xb0\x59\x1f\xfc\xaa\xe9\xa2\x08\xd1\x19\xa1\x4b\x20\xf4\x2a\x0a\x39\xf6\xd4\xaa\x42\x66\x99\xb3\x86\x7d\x94\x69\x28\x0e\x79\x95\xb2\xdd\x42\x45\xc8\x51\x19\xd8\x25\xba\x59\xa3\x54\x37\x18\x4b\xa9\x83\x39\xee\x8c\x9d\x89\x50\xe2\x4e\x7f\xdc\xeb\xaa\x05\x4d\xa9\xdd\xdb\x0f\x56\xd3\xcb\xf7\x1b\x68\x51\xc1\x3f\xd0\x20\x93\x45\xa5\x76\x80\x4b\xc4\x18\x9c\x23\x1b\xfc\x65\xed\x80\x82\xeb\x9f\x39\xfd\x9c\x80\xdd\x2f\xc5\xf8\xcd\xd7\xd6\x0e\x99\x0a\x57\x5f\x46\x59\xf4\xf0\x9c\x5c\xda\x2b\xb1\x39\x00\x54\x74\x4e\x73\x49\x77\x00\xd0\x43\x74\xe5\x20\x5d\x32\x17\x65\x89\x1e\xfc\x32\x85\xbf\x36\x05\xc0\xb6\x08\x81\x55\xab\x79\x64\x19\x91\x10\x85\xdc\x84\x13\x8b\x90\xa7\xf0\xe1\x70\x9e\x60\x66\x29\x7e\x25\xa0\xd2\x11\x9f\x48\xdb\x4d\x8e\x3b\x2b\x51\x23\xdd\xfd\x52\xb1\xf9\xbe\xb6\xd3\xae\xda\xa5\x3e\xed\x24\xa6\xa8\x00\xe0\xa5\x30\xba\x87\x24\xc0\xde\x4d\xde\x21\xf3\x16\xc8\x8f\x03\x4d\x87\xe7\x9d\x35\x54\xfb\xc6\x63\x71\x2b\xc6\xcf\x8b\x53\x12\x9f\x27\xe4\x52\x8c\x9f\x3b\x87\x44\xe1\xf3\xe2\x51\x82\xd1\x13\x72\x49\xf4\xf8\x4c\xf9\x24\x7d\x11\x2b\xa2\xe4\xfa\xe6\x79\x71\x49\x22\x96\xe2\xf5\x68\xbc\xca\x08\xd7\xd6\x08\x61\x27\x1b\x20\x77\x66\x95\x6a\xff\xd0\xac\x8a\x32\x6f\xe1\x79\x71\xaa\x80\xd7\xa3\x71\x4a\xf4\xd2\xce\xbb\xb2\x93\x3d\xad\xbb\xf3\xe8\x50\xf1\xe8\xd1\xec\x88\xc3\x67\xb2\x34\x1d\x62\xeb\x33\x5c\x3d\xb1\xc6\x63\x87\x0f\xb7\x30\x95\x38\xa4\xb1\x23\x0b\x8f\xac\xe5\x84\x36\xb3\x2d\x15\x74\xb1\xb2\xcd\x9b\x27\xe2\x0f\xe3\x14\x87\xf3\x93\x0c\x43\x89\x92\x8a\xf1\xa8\x31\x08\xc2\x1a\xa2\x3b\x5f\xef\x3b\x76\x65\xc5\x7d\xf3\xf1\x2a\xef\xc4\x15\xee\x46\xa7\xe0\x56\xfd\x8a\x6e\xc4\xd0\x37\xf8\x47\x5f\xcb\x40\x9c\xcc\x65\x31\x01\x28\xfa\x1e\xf7\xa6\xe1\x3c\x20\x53\x18\x58\x1e\x09\x67\x78\xfe\xed\x68\x98\x36\xd1\x66\x86\xf8\x1c\x65\x5b\xfa\x36\x68\xfc\xd4\x3a\x6a\xe4\x02\xa6\x16\xb3\x7c\x39\x92\xd5\x73\x0d\x9a\xeb\x40\x59\x00\x29\x5c\xb2\xc4\x1f\x13\xbf\xa2\xe7\x75\x5c\x71\xc5\x92\x6a\xe5\x4d\xee\xe3\xdd\x2f\xe5\x4f\xd5\xaa\xc9\x1e\x71\xa1\x6e\xf2\xad\x5a\xd9\xb0\x47\x5b\x68\x68\x28\x4f\x81\xd4\xec\x76\x1e\xef\x7e\xa9\x29\x31\x34\x34\x07\x7e\x8b\x20\xcc\x75\x0c\xc0\x2e\x5c\x49\xa1\xd2\x17\x43\xc5\x74\x2b\xa5\x58\x39\xfd\x6a\x68\x90\x6e\x8e\x15\x1b\xa4\x5f\x0d\x0d\x4a\xb1\xb8\x62\xbb\x52\xa1\xa1\x79\x65\xd3\xa7\x08\xa0\x52\x9c\x82\x30\xed\xd0\x1c\xef\x7e\x31\x7d\x4e\x9b\x54\xb6\x50\x8e\x77\xbf\x54\xbe\xa5\x95\x2b\xcb\xd3\xb1\x69\xc9\x4a\x2a\x6f\xd8\x64\x38\xde\xfd\xb2\xa1\x46\x0a\xa8\x56\x48\xd6\x4b\x86\x41\xe5\x1f\xd7\x2c\x04\xa5\xd1\x69\x41\x7b\x7d\x80\xda\xe7\x8c\x67\xe6\x10\xb1\x60\x97\xb9\xc4\xd0\xd0\x20\x2b\xb5\x85\x6b\xf5\xad\x0c\x70\xbe\x70\xa7\x5e\x18\x32\x99\xf5\x24\xfe\xb0\xc1\x8f\xa9\x21\x1a\x51\xc2\x89\x47\x02\x1b\x8c\xbb\xc3\xe4\x1b\x87\x74\x8e\xf8\x30\xa9\x98\x54\x65\x28\x40\x1e\x27\xf4\xa1\x86\x57\x83\x77\x91\x15\x30\x8a\x58\xad\x6d\x9b\xef\xc1\x75\xd3\x95\xae\xc8\xa5\xe7\xc5\x0f\xb3\xd4\xe4\x5c\xa2\x28\x0a\xb0\x07\x99\x0d\x0e\x9e\x92\xe0\xb2\x1a\xa7\x90\xa3\x79\x66\xe7\x52\x12\x04\x38\x9c\x0f\xd5\x0a\x9b\x19\xbf\x38\xe4\x88\xae\x60\xe0\x22\x8f\x84\x7e\x86\xa8\xf8\x2d\xe1\xb5\x1b\xd3\xb9\x4c\xc7\xf9\x2f\xfd\xeb\x45\x98\xed\xce\x14\xcb\x38\x5e\x22\x12\xf3\x0c\xd6\xf7\xfb\xfb\x59\x59\x1c\xf9\x90\xa3\x21\xa2\x98\xf8\x95\xce\x28\x62\x24\xa6\x1e\xd2\x10\x0b\xf0\x12\x73\xed\x6f\x90\x6c\x5a\xd8\xa0\x71\x78\xf4\xfd\x39\x6e\x64\x25\x14\xfd\x19\x23\x56\x57\x77\x3f\xaf\xaa\xb8\x3d\x52\x84\x50\x2e\x48\x92\xb6\x97\x36\x2d\xca\x5a\x55\xde\xea\x78\xb6\x0d\xdf\x6e\x21\x7b\xb7\x60\xb3\x2e\x6d\xe2\x97\x6d\xe1\x69\x38\x5b\x35\xba\x2d\x91\x80\xa5\x0c\x81\x37\x41\x73\xa7\xfc\x75\x18\x07\x41\xe2\x2d\x81\xde\xac\x4f\xf8\x90\x22\xa6\xf6\xc7\x73\x36\xad\x50\x88\x18\x1b\x52\x32\x45\x3a\x99\x80\x4c\x18\x3c\x45\xbc\xf8\x51\xd8\x78\x7c\x61\x83\x46\xbb\x51\xfe\x5e\xd4\x4d\x19\x22\x21\xe6\x18\x06\x27\x28\x80\x37\x99\xdc\xbc\xd5\xeb\x50\x04\x7d\xfc\xf4\x38\x1c\xec\x14\xdb\x15\x84\xc4\xca\x19\x31\x34\xc1\x5c\xc9\x0d\xca\x73\x12\x87\xe5\x76\x72\x8f\x6e\x28\xf1\x6b\xc7\x8c\xb6\xd9\x02\x52\xd4\x0e\xe7\x38\xbc\x6e\x2f\xf8\x32\x68\x6b\x2e\x40\xfa\x53\xec\x55\x05\x96\x02\xad\x95\xbf\x02\xae\xca\x00\x98\xc6\x94\x71\xb9\xa7\x2a\xf3\x48\x20\x08\xc8\x55\xba\x15\x38\x23\x84\x47\x14\x87\xb2\xa2\x4c\x39\x01\xaf\x8f\xf6\xc1\x39\x7e\x53\xa0\x74\x65\x9a\x02\xe3\x54\x05\xda\x14\x54\xb9\x7c\x7a\x89\x69\xc2\xea\x2d\x8e\xf6\xb5\x06\x6a\x38\x05\x69\x56\x03\x3d\x87\x91\x6d\xa0\x82\x26\xe4\x56\x85\x54\x75\x84\xe2\x14\xcf\xe7\xd9\x9c\xb1\x94\xf4\x77\x17\x30\x9c\xa3\xb2\xc2\x84\x31\x27\x4b\xc8\xb1\x57\x88\x16\x68\x13\x4f\x66\x6d\xe8\xe8\x9a\xa6\x5d\x31\xde\x61\x88\xec\x8c\x61\x15\xef\xda\x80\x7b\x5a\x45\x6e\x9d\x0b\x4f\x74\xbd\xc1\xa7\x34\x4f\x2f\x1f\xe3\x4e\x32\x6c\x55\xa0\x56\xde\xa4\xa4\xd6\x8c\xea\xa6\x4c\xd8\x6c\x48\x3d\xb9\xd7\x5a\xae\x56\xab\x3f\x73\xa4\x15\x8a\xad\xcf\x4c\x0c\xf3\xaf\x04\xc6\x97\x8c\xbe\x0d\x18\xe1\x77\x90\xa1\x86\x0d\x1a\x69\x22\x74\xc5\x19\x6d\xec\x15\xea\x3b\x49\x2a\x99\x68\xd3\x86\x11\x6e\xaf\x0e\xf4\x1a\x1c\xf3\x40\xc2\x4b\xf7\x1d\xf5\xc2\x24\xfd\xf9\x82\x06\xa2\x86\xee\x68\x6a\xc9\xd4\x85\x0e\x93\xf4\xe3\x72\xf5\xd4\xa1\xd5\xab\x8a\x51\x2f\x61\x14\x21\xda\xb0\xb5\x31\x02\xd0\x98\x42\x86\xce\x61\x14\xe1\x70\x9e\x98\xc9\x09\x06\xb5\x63\x4e\x06\xd6\x86\x3c\x80\xac\xad\xf5\x92\x40\xfb\x05\xae\x60\x2f\x14\x6b\x14\xc7\x24\xbc\x13\xd0\xcf\x70\x05\x0d\x90\x7f\x3f\x3f\xbb\x27\xe0\xeb\x65\x60\xc2\xd8\x1d\xf4\xef\x8b\x31\x23\xa1\xb6\xb6\x7c\xd5\x68\x3f\x43\x50\x08\x29\x6b\x80\x12\xe9\x03\x32\x9f\xe3\x70\xde\x48\xcf\x55\x98\x1a\x4f\x29\x0c\x7d\x55\xa9\xdc\x96\xfc\xb6\xc0\x1c\xbd\x13\x45\x0d\xc8\x18\xe2\xac\xad\x72\xf1\xf2\x70\xa8\xa8\x65\xb1\xd5\xdc\xba\x12\x55\x5b\x6c\x35\x2f\x0e\x5e\x94\x9f\x40\x7a\x79\x2f\x20\xd8\x23\x61\x2d\x2a\xf3\x40\xfc\xcd\x26\x02\xc8\x84\xfd\x19\x43\x8a\x5a\x51\x68\x80\x50\x87\xc6\x56\x00\x60\x14\x09\x1d\x5c\x33\xb5\x04\x0f\xe0\x4a\x74\xf2\xf6\x50\xce\xcc\xe4\x2f\xeb\xed\xe1\xf5\xdb\xc3\x2a\xb4\xa4\xf8\xe0\xfb\x42\xe5\x83\xef\xaf\x0f\xbe\xaf\x56\xe6\x24\xf6\x16\x3d\x8f\x84\xc9\x9c\x8f\x02\x64\xc9\x6f\x96\x68\x55\xad\x1f\x51\xe2\xc7\x1e\x7f\x17\xe3\xc0\x2f\x73\x3e\xd5\xd6\x5f\xd7\xc4\x9e\xeb\x95\xf0\xb3\xd1\xb3\xfe\xd4\xb4\x28\xa4\x79\xc3\xf9\xd2\x90\xa3\x0e\x7d\x3f\x29\xb7\xfc\x69\x8b\x2d\x72\x7d\xfc\xea\xff\x6b\x4f\x71\xd8\x9e\x42\xb6\x48\xbe\xc4\x21\xc7\x01\x10\x1f\x80\xe5\x81\x46\xc4\xfe\x0c\x80\xb5\x00\x07\x87\x3f\xb4\xf6\x5b\xfb\xad\x03\x60\x5d\x80\xdd\x52\xa0\x0b\x58\x7f\x02\xcb\x2f\x7c\x4e\x23\x54\x02\x48\xd3\x75\xce\x9c\xee\x18\x1c\x34\x1b\x3f\x03\x9f\x64\xec\x42\xde\x82\x80\xc6\x6f\x10\xcb\x64\x5c\x2d\x69\x9b\x01\x15\xe0\x6e\xb5\x5a\xf9\x9c\x67\x01\x42\x51\x66\x23\xfa\x24\x4c\xe9\xa5\xc0\xfc\xb7\xf8\xa9\x14\x26\x01\x2d\x25\x47\xda\x5e\x8e\xe3\xef\x7f\x77\x06\xef\x33\x80\xdd\x91\xd3\x19\x3b\x20\xc3\x34\x6d\xf2\x73\xb9\x86\x1c\x62\x96\x97\xfd\x5b\x6f\xfc\x01\xa4\xe1\x2f\xd0\xdc\x5d\x17\x1d\x6b\xe6\xb0\x4e\x47\x9d\xfe\x18\x74\xce\xce\xc0\x70\xd4\xfb\xd8\x3b\x73\x4e\x1d\x17\x0c\xfa\xd5\xee\xa5\xb9\x58\x42\x25\x47\x5b\xf1\xc3\xcf\x6b\x5b\x17\xf9\xbf\xff\xfe\xf7\xa6\x33\x78\xdf\x2c\xe3\x3f\xee\xbc\x3b\x73\x40\xef\x3d\xe8\x0f\xc6\xc0\xf9\xbd\xe7\x8e\x5d\x65\x5c\x79\x1c\xbc\x9e\x61\xca\xf8\x44\x1e\xaf\xf8\xd8\x19\x75\x3f\x74\x46\x7b\x20\x80\x95\x4f\x42\x10\x61\x78\xa3\xd5\x41\xd0\x9f\x28\xe3\x55\xab\x25\x33\xc8\x26\xc2\x2f\x15\x03\x73\xde\xe4\xe3\xef\xf5\x5d\x67\x34\x06\xbd\xfe\x78\x90\x75\xfe\xb1\x73\x76\xe1\xb8\xe0\x75\xf3\x17\x82\x9a\x7b\xcd\x5f\xa0\x77\xc9\x48\xd8\xdc\x6b\x8e\x90\x0f\x3e\x40\xde\xdc\x6b\xfa\xd3\xe6\x9e\x17\x53\x8a\x42\x3e\x11\xce\x30\xe3\x70\x19\xbd\xa9\xb0\xc8\x34\x44\x4e\x7c\x02\x5e\x63\x1f\xb8\xce\xa8\xd7\x91\x74\x3f\xef\x8c\x3e\x81\x5f\x9d\x4f\x7b\x80\x43\x76\x59\x1c\x5d\x80\x38\xf2\x81\x0c\x06\x3a\xa3\x6a\x0f\x83\x11\x18\x39\xc3\xb3\x4e\xd7\x01\xef\x2f\xfa\xdd\x71\x6f\xd0\x17\xf3\x6a\x22\xe8\xf0\x5a\xd3\x40\x8a\x9e\x30\xf4\x27\x39\x15\x57\x90\x7a\x0b\x48\x75\x45\x95\xd2\xd3\x50\x14\x2d\x48\x68\x6c\x83\x96\x10\x07\xa6\x02\x9d\x17\xb5\xc5\x1c\xf2\x98\x99\x8a\xa9\x9a\x31\x49\xc9\x1b\xdd\x67\x1d\x39\xe3\x8b\x51\xdf\x05\x2b\x82\x7d\xed\xf3\x59\xa7\x7f\x7a\xd1\x39\x75\x40\x33\x0a\xa2\x39\xfb\x33\xd0\x1c\xdd\x8e\x0b\x76\xdf\x0d\x4e\x3e\xed\x66\x5f\x4e\x9c\xee\x59\x67\xe4\x68\xed\x25\xf1\x93\xfe\x72\x42\xbf\x73\x4e\x7b\xfd\x72\x2d\xfb\x58\x48\x8b\x07\xf9\x6b\x7d\x14\x7f\xfd\x25\xbc\xeb\x3d\xd0\x3c\x43\xd0\xb7\xc1\x30\x40\x90\xa1\x4c\xac\x9a\x7b\x26\x2e\xec\x81\xa6\xca\xb4\x6f\x8a\xe6\x09\xfd\xc5\xc7\x15\x86\x8a\xe6\xb6\x2a\x92\xff\x4e\x0b\x24\xcd\x93\x02\xf9\xef\x3d\xd0\x6c\x65\x5d\x03\xcc\x34\x98\x1a\x1b\x64\xad\x91\x24\x6c\xd2\x58\x51\x59\x7c\x6f\x6a\xc2\x25\x3c\x61\x86\x28\x97\x27\xdb\xa4\xc4\xbe\x16\xc3\xde\xcb\x04\xf2\x4d\x36\x49\xe4\xf7\x7d\xad\xad\xd3\x3f\xc9\xff\x50\x34\xff\x79\x67\x1b\xb1\x4d\x66\x69\x59\x72\x07\x17\x63\xa0\x69\x03\x8e\xae\xf9\x5e\xa9\x38\x17\x69\x53\x69\x2a\xd3\xc6\x96\x9a\x88\x8a\xf2\x37\x06\x29\x73\x9d\xf1\xe0\x3d\xa0\xc8\x53\x27\x58\x72\x89\xd2\xfe\xd8\xdd\x2d\x78\xa8\xc9\xca\x92\xa3\xad\x29\xaf\xbd\x9c\xc9\x5a\xef\x85\xe6\xef\x47\x83\xf3\x54\x6c\x7e\xae\xed\x25\x17\x77\x21\xea\xe0\xe3\xe0\xac\x33\xee\x9d\x39\xba\x72\x36\x2c\x45\x99\x56\x56\xe4\xf6\xf3\x7d\x3d\x97\x43\xca\x37\x2c\xc3\xed\x15\xa4\xed\x00\x4f\xdb\x72\x7e\xb5\x53\x60\xed\xf2\x52\x0e\xfe\xff\xff\x05\x40\x3b\xa2\xc4\x6b\x1f\xb4\x67\x7e\xfb\x60\xed\xbe\xfa\x56\x71\xfa\xc4\xbe\x78\x5e\xf6\x4e\x4d\x9c\x3e\x49\xa1\xc9\x36\x4b\x13\x90\x21\xf1\x91\x8a\x12\x65\x11\x7c\xf9\xd7\xd1\x77\x6f\x0f\xd3\x0f\x1b\x42\xfa\x59\xd5\xa7\x88\x30\x27\x24\x67\x2a\xf5\xb6\x33\x9b\xe1\x10\xf3\x1b\x1b\xf4\x53\x2b\x27\x89\x23\x04\x31\xe3\x88\xf6\x86\x72\x1f\x59\x68\xc2\xc4\x34\x25\xd0\x7f\x07\x03\x18\x7a\x88\xda\xe0\xcb\x9a\xad\x9a\xa1\xf8\xc6\x38\x0a\xb9\x4a\xe8\xef\x06\x10\x6f\xcc\x79\x79\xe6\x02\xa1\x0e\x21\x9e\x13\x1f\x65\x62\x31\x42\xd0\xff\x8d\x62\x8e\x06\x61\x32\xe3\x2b\x81\xb6\x6a\xc8\x8c\x71\x42\x65\xe0\x76\xfd\x96\xe9\x43\xee\xbd\xbc\x2c\x42\x7f\x8b\xbd\x97\x74\x66\x94\xf6\x5e\x92\x8d\x07\x94\x1d\x59\x30\xf2\xf8\x11\xf6\x3c\xbe\xe5\xfe\xc6\x1a\xfa\x6c\xb7\x69\xe1\xc1\x08\x4e\x71\x80\x39\x46\x4c\xea\x89\xbc\x1f\x14\xae\x8a\x51\xf3\xba\x93\xcf\xf9\x2f\x39\xd5\x51\x9b\xf6\x50\x03\x47\x3b\x4d\xba\x1e\x56\x39\xc5\xa0\x06\x9e\x76\x7c\x78\x3d\xbc\x72\xd2\x44\x0d\x3c\xd3\xb1\xd7\xf5\x80\x6b\x93\x22\x40\x71\x33\xa8\xf4\x71\xab\xbd\xa0\x19\xf2\x6e\xbc\xa0\xb4\x07\x93\x19\x11\xe5\x78\x3f\xba\xd6\x85\x20\xfd\x79\x64\xb9\x84\xa1\x5f\x2d\xb0\x80\x34\x3b\x32\xa3\x43\x2f\xb1\x3c\x53\xf5\x1a\xab\x44\xb7\x6a\x0a\xe8\xd7\x6e\x65\x19\x37\x80\xbe\x2f\x6e\x12\x71\x2f\x72\x89\x77\x69\xd8\x6c\x2a\x2f\xe6\x20\xcf\xa6\x2d\x9b\x02\x60\xab\xcd\xa4\x12\x2c\xb3\x79\x00\x36\x6c\x8d\x99\xa8\x5f\x43\xfb\x3a\xca\x5b\xc0\xc2\xd5\x4f\x65\x56\x58\xe0\x61\xc3\x32\x9b\x39\x73\xa4\x55\x79\x05\x4e\xde\x81\x7f\x12\x17\x78\x01\x64\xf2\x6c\x4f\xe3\x34\x86\x14\x86\x1c\x21\xbf\x01\x5e\xa7\x9a\x14\x1c\x1f\x27\xfa\xf7\x4d\xa1\x75\x9f\x70\x64\x83\x41\x08\x06\xee\x00\x70\x79\xec\x19\x33\x10\x12\x90\x43\x51\xa0\xf7\x00\xe6\x0c\xc0\xe0\x0a\xde\xb0\x7c\xab\xad\xc0\x8c\x3b\xec\x9e\x6d\x4c\x12\x4b\x61\xaf\xdf\x52\xdb\x0a\xcc\x76\x9b\x92\xc5\x49\x95\x9c\x43\xcf\x7f\x15\x23\xc1\x2a\x55\x59\x07\x2d\x9d\xa2\xeb\x20\x56\x43\x87\x3a\xf6\xeb\xb6\xbd\xab\xb8\x44\x26\x03\xb3\x48\x43\x4f\x7c\xea\x1b\x2d\x1f\xb0\x6e\x37\x32\xb9\x89\x47\x18\x7a\x36\x38\x3a\x38\xb8\xcb\x90\xb6\xaa\xf8\xa8\x7b\x97\x06\x05\x75\xa7\xad\x4b\x2d\x41\xf4\xa7\xd6\x51\xa9\x38\xdb\xb6\xdc\x3a\x99\xb1\xba\x83\x09\x6e\xbd\x81\x59\xe7\x5f\x3e\x1b\x7b\x36\x3b\xb4\x64\x3e\xcc\xf4\x02\x73\xc1\x24\xe6\x0f\xeb\xed\xbd\x80\x93\x5d\x8f\xe0\xf1\x99\xce\xb1\x3f\xac\xaf\xf7\xdc\xc8\xb9\x79\x16\x7c\x0b\x7f\x2f\x43\xae\x92\x6d\xf7\x94\x09\x6d\xba\x5f\xf9\x82\x33\xda\x32\x62\x96\xdd\xc3\xe4\x4a\x82\x8e\xe7\x09\xc3\xa1\x5f\x7b\x54\x1a\x6c\x9b\xff\xa6\xf5\x04\xea\xbd\xc9\x5f\x3a\x1f\x3b\x93\xce\x70\x38\x39\xe9\x19\x5d\xc9\x76\x7e\xf3\x0c\x33\x34\x3f\x1b\x74\x4e\x9c\xd1\xe4\xc3\xe0\xdc\xe8\xec\xe9\xad\xdb\xe8\x9a\xd7\x21\x30\x90\x67\x48\x5c\x13\x88\x86\x75\xf2\x19\xae\x60\x2b\x44\xbc\x15\x51\x34\x43\xb4\x37\x5c\x7d\xe7\x72\xe8\x5d\x1e\x8b\xf5\x1d\x58\x27\xf2\x5a\xa6\x05\x59\xa2\xe3\x36\x5f\x46\x0d\xb3\xa7\x77\x07\x57\xef\xce\x39\x77\xed\x05\x82\x01\x2f\x3b\x12\xc9\xa2\x75\xf0\xe3\x41\xa9\x80\x79\x0b\x24\x48\xf1\x61\x3c\x1e\x6e\xb6\xfb\x0f\xf6\xb7\x73\xe7\xbe\x1d\x86\x6f\xf7\x8b\x4e\x63\x54\xcc\x87\x3d\x2c\x96\xce\x20\x0e\x62\x8a\xc6\x0b\x8a\xd8\x82\x04\x7e\xd1\xb3\xb9\x75\xd2\x61\xaa\x43\xc5\xf0\xb7\x72\x1f\x2b\x00\x7f\xfa\xe1\x87\x9f\x0c\x00\x4b\x87\x42\x6f\x0d\xf6\xc7\x1f\x7e\xf8\xd1\x00\xf6\x33\x09\xc8\x25\x86\x5b\x7a\xba\x77\x72\xae\x2a\x77\xab\xdc\xce\xad\x3a\xfc\x71\xbf\x90\xdb\xf8\x0a\xb0\x88\xe2\x70\x6e\x4d\x09\xe1\xb9\xb9\x0d\x83\xe0\x06\x44\xd8\xbb\x64\x20\x8e\xd4\xfd\xa0\xea\xe4\x59\xeb\x66\x19\xa8\x2d\xb9\x56\x35\xaf\xf3\x8a\xd0\x4b\x1c\xce\x4f\x30\xad\xd5\x35\xf5\x1e\xdb\xfa\x94\x50\x50\xf0\xc1\x74\x4d\x54\xc1\x22\x05\x85\xae\xf9\x6d\xe0\xe8\x1a\xad\xd6\x33\x33\xc2\xbc\xa3\x4f\x56\xd0\xe9\x9b\xc6\xbf\x5d\x06\xa9\x00\xf9\x64\x9e\x96\x79\x20\xf7\xcb\x13\xad\x1e\xf9\x07\x8f\x97\x29\xfa\xa2\xb3\x94\xd6\x59\x99\xc6\x1c\xa5\xe2\x1c\xd6\x4e\x2c\x26\x57\x96\x64\xf4\xae\xa4\x94\xbf\x02\xbf\x21\x40\xc2\xe0\x06\x5c\xc1\x90\xab\x3b\x98\xe4\xee\xd8\x9e\xbc\x39\x5a\xfc\x3d\x8b\x83\x40\x76\xd6\x02\x1f\x50\xe8\x21\xc0\x90\x17\x53\xcc\x6f\x00\x09\xf7\x00\x43\x21\xc3\x1c\xaf\x10\x20\xb3\x59\x2b\x83\xea\x22\x94\xdd\xd1\xeb\x13\x8f\xb5\x94\x2a\x12\x23\xd6\x94\x92\x2c\x6a\x27\x39\x2a\x6d\x69\x34\x88\x1e\x54\x6a\x79\x92\x8e\x86\x49\x68\x89\x05\xfe\xc6\x5a\x92\x10\x73\x22\xc1\x88\x0a\x59\x5f\xef\x09\x05\x3e\xe2\x10\x07\x29\x1f\x96\x30\x84\x73\x24\xe6\x7e\x65\xe8\xfa\x9a\x99\x0e\xc4\x2e\x98\x5f\xf2\x8a\xbe\xc2\x4c\x49\xaf\x61\xd4\xa6\x88\x5a\x96\xf5\x86\x19\x21\xd2\xcc\xb9\x97\xeb\xeb\x17\x2e\x9a\x00\x00\x86\x21\xe1\xea\xca\x2a\x5b\x93\x2b\xec\xa1\x16\x0c\xa2\x05\x2c\xfa\x75\xb2\x28\x9c\x5b\x1e\xa2\x3c\x39\x65\x6d\x99\x8e\x9f\xab\xd3\xe7\x3c\xe1\x59\x6d\x8d\x8d\xb1\x05\x75\xed\x33\xd8\x26\xba\x90\x56\x7d\x0a\x67\x4c\x1b\x41\x51\x10\x28\x89\x0d\x57\xa6\xa7\xb2\x31\x12\xa5\xdf\x4a\x32\x0c\x8c\x4e\x52\xbf\x4d\x8c\x26\x2b\xc1\x69\x74\x65\xc1\x28\xb2\xe4\xa8\x6c\xd0\x10\xb3\xa6\x61\xe0\x68\x81\x8d\xf2\xbe\xee\xea\x4d\x56\x19\x8b\xed\x75\xcc\xe3\x39\x21\x70\x28\x67\x30\x72\xfc\x39\x1a\x23\xba\xc4\xa1\x44\x3e\xf5\x16\x46\xc8\xc7\x14\x79\xe9\x6a\xc3\xf3\x1a\x36\xa0\x28\xb9\xe3\x4b\xc1\x24\x29\xc8\xf2\x0c\x5d\x2f\x9b\xff\x23\x63\x1c\xa5\xd9\xbf\xdd\xcc\xfc\x16\xf1\x8e\x12\xa2\xff\x89\x7a\xdc\x3f\xea\x51\x22\xe9\x9d\xce\xf3\x55\x60\x80\xba\xad\x5c\x48\xe7\x25\x67\xc6\x02\x96\x95\xbc\x65\x40\x8f\xf3\x4b\xcd\xcb\x55\xd4\xbd\x17\x16\xf6\x8f\xd9\x0d\xe3\x68\x69\x27\xeb\x11\x54\x41\x19\xdb\x74\xe6\xc4\x2e\x22\x98\xc0\xa8\x03\xad\x16\xae\x6d\x2f\x55\xc8\x9b\xc7\x91\xba\x4b\xf7\x58\x18\x3f\x76\xbb\x7c\x79\x9c\x3c\x99\x71\xeb\x46\xea\x78\xcc\x2d\xda\xc5\x86\x4e\x78\xc0\xe4\x9a\x7c\xdc\x46\xdc\x6b\xf3\x80\xb5\x23\x8a\x57\x90\x23\xf1\xef\x96\x47\xab\xa4\x10\x2d\x2e\xd1\x8d\xb9\x81\xba\x17\xb1\x44\x3b\xfd\x2a\x99\xe3\xdd\xd7\x86\x53\xfb\x6f\xaa\xec\x86\x8c\x59\x2a\x2a\x6d\x71\x72\x89\xc2\x4a\x0d\x76\x89\xa3\x4c\x2a\xac\x69\xcc\x39\xa9\xa9\x24\x19\x4b\xd1\x1c\x5d\x1f\xa7\x67\x60\xd8\x15\x14\x6e\x52\xeb\xbf\xb7\x6e\x81\x43\x1f\x5d\xeb\x16\xe6\xe6\x26\x1e\x45\x3e\x0a\x39\x86\x01\x6b\x0b\xff\x7a\x0a\xbd\xcb\xad\x1b\xaf\xd4\xfa\xb1\xa6\x7e\x44\xd1\x2c\xc0\xf3\x45\x95\x45\xd9\x14\xb1\x3c\xa8\xf8\x14\x5d\x62\xc9\x2b\xc1\x6a\x81\x8c\x35\x8d\x43\x3f\x40\x46\x06\x17\x5b\xaf\xa0\x7c\xe7\xa5\x9d\xdc\x57\xdc\xbe\x8c\xa7\x88\x86\x88\x23\x96\x99\x75\xd9\x0c\x6b\x7b\x70\x03\x44\x06\xe9\xf1\x97\x46\xe6\xdd\x35\xec\x9a\x83\x60\x8d\x54\x4d\x37\xec\x46\x44\x7c\xd6\xd8\x6b\xac\x10\x9d\x36\xec\xc6\x1c\xf1\x46\x71\x82\xbd\x02\x27\x98\x49\xf3\x5c\x9e\x78\x08\xc9\x95\x9d\x0a\x50\xcc\x84\x68\x20\x48\x11\x55\x52\x04\x20\x53\x2f\xde\x24\xb7\x6c\x32\xe1\xdf\x53\x28\x1c\xfa\x25\x96\xa9\x82\xe0\x6a\x81\xbd\x85\x74\x7e\x4a\xbd\x78\x30\x04\x53\x04\xe6\x78\x85\x42\x30\xbd\x01\x10\x78\x2a\x7b\xd0\x82\xfe\x12\x87\x05\x6f\xc1\x18\xac\x35\xdf\x4a\x9d\xff\x64\xc8\xf4\x7d\xc9\x9b\x06\xca\x23\xa1\x88\xff\x8a\x6e\x46\x68\x56\xcd\x39\xb9\xd5\x2d\x4e\xfa\xef\x12\xdd\xd8\xd5\x1b\x96\xf4\x30\xc7\xe6\xe8\x5d\x6e\x67\xeb\xc8\x44\xf2\xc1\x96\x42\xc1\x5d\xd2\x3f\xea\x62\xa0\xe6\xae\xb3\xe0\xa8\x1c\x52\x12\x22\xfd\x77\x99\x94\x5a\x28\xb4\x18\xb0\x36\x47\x6b\x8f\x0a\x75\xca\x97\x07\xdc\x37\x98\xfb\x3c\x07\xb2\x5d\x8e\x45\x49\xfd\x1b\xa4\x60\x9d\x67\xa7\x7e\x77\x3b\xac\xbd\xbf\x7f\xcb\xc3\xda\x87\x7a\x83\xb5\x3b\x35\x46\x23\x60\x73\xd6\x46\xed\x00\xd5\xcc\x2b\xe2\xa5\xbe\x99\x3a\x2e\x42\x78\x9a\x50\x9e\xc1\x2e\xbb\x53\x40\xaf\xf6\x5a\x48\xf0\x0d\x0e\x7e\xbb\x05\x1e\xbf\xe0\x9b\x8c\x1f\x68\x84\xda\xfd\xe2\x2f\x72\x98\xcf\xd7\x2d\xd6\x64\xa7\x4e\xaa\x5e\x60\x12\x8c\x91\x2d\xff\x43\x02\x19\xb7\xe2\xd8\xb7\x08\x60\x68\x08\xfe\x27\x78\x71\xff\xe0\x45\x21\xf9\xe2\x91\x93\x36\x4a\x6d\x1e\x2b\x6d\xe3\x51\x93\x2e\xd2\x4e\xf4\xa7\x7b\x4a\x3d\x98\x7c\x94\x19\x46\x81\x6f\xf4\x4f\x64\x89\x32\x1a\x53\xb1\x69\x85\xda\xb3\x41\xe5\x8e\x9d\xfe\xc9\x70\xd0\xeb\x8f\xdd\x49\xcd\x23\x55\x95\xf1\xae\xbf\x15\x51\x87\xbd\xf9\xb1\x24\x03\xf0\x2d\xaf\x36\x04\x4f\x72\xfa\xc1\xda\xe2\xa5\x8c\xce\xd9\xe9\x60\xd4\x1b\x7f\x38\x37\x0a\xc7\x87\x25\xf4\xdc\x0f\x9d\x03\x13\xd3\x37\xbf\xbf\xb1\x85\x28\xdc\xca\x5d\xad\x5e\xbc\xab\xff\xa4\x8b\xba\xee\x52\xdc\x4d\x63\xd0\xae\xb8\x5d\x4f\x95\x8e\xe3\xb6\xbb\xef\xba\xed\xe1\xaf\x5d\xf7\x68\x08\x7d\x79\x87\xc9\x2d\xa0\x3f\x07\xea\x14\x6e\xfb\xdd\x84\xfb\xb8\x67\x94\xc4\x86\x51\x2e\xd6\xbd\x45\x53\x01\xb1\xfb\xa5\xb6\x7a\xcd\x39\x9e\xf2\xd7\xc7\xb9\xd4\xcd\x9c\x6f\x94\x79\xdb\x6b\x02\x6f\x0a\xe6\x07\x04\xfd\x82\xfe\x2f\xd2\xa8\xe3\x79\x28\xaa\x32\x29\x7b\x7d\x0d\x5d\xf3\x76\x14\x40\x1c\x6e\x71\x3a\xe4\x21\x73\xb0\xee\x71\xf7\x9c\x0a\x3d\xd4\xdc\x40\x57\xce\x33\xdb\x9c\xee\xf6\x50\xd9\x60\xf7\xca\xf9\xba\x7b\x66\xd7\xf3\x48\x76\x52\x57\xe5\xe5\xa7\x85\x38\x29\x9d\x16\x52\x16\x21\x38\x3e\xce\x2c\xbe\xe2\x69\xa1\xf1\x02\x33\xe0\x13\xc4\xc2\x26\x97\x63\x02\x44\x9d\x1c\x92\xb1\xd1\xe4\xf0\x10\x04\x33\x7c\x8d\x7c\x79\xc9\x02\x29\x34\x97\xa9\x60\xb2\x0f\xd1\x75\x6a\x91\x82\xd7\x3f\xee\xff\x17\x48\x72\x36\x82\x9b\x37\x2d\xd0\x4c\x7b\x6f\x0a\x78\x78\x1e\xca\xb7\x96\x64\x07\x05\xe1\xbc\x53\x4e\x9c\xe1\x29\xaf\x5b\x66\xc5\x09\xdb\x78\xa7\xc8\x36\x83\x65\x77\xbf\xe4\xac\x44\x7b\x3f\x79\x7a\x56\xc9\x0c\xbd\x5f\x82\x96\xe9\x7d\x23\xf0\xbc\x52\xb4\x0c\x19\x0a\x98\xb4\xd2\x11\xb4\xe3\x68\x4e\xa1\x8f\xac\xa5\x3c\xe7\x74\x89\x50\xf4\xb2\x3c\xd2\xdb\x25\x77\xe5\xca\x43\x73\xce\x08\xf4\x2d\x1f\x2d\x89\x3c\x5c\x26\x18\x56\x7b\x2d\xb8\x47\xf4\x15\x0e\x06\x01\xb9\x42\xfe\x80\xe2\x39\x0e\x59\xfe\xc0\x7a\x35\x4b\x23\x6d\x0e\xbd\x85\xb6\xc6\x24\x3b\x24\x76\x45\xbc\xd2\x5b\xcf\x92\x0b\x18\xb2\xf2\x25\xbc\x2e\xe6\x5d\x71\x2a\x8f\x56\x1f\x64\x0b\x62\x92\xa1\x91\xd7\xba\x54\x2f\x1b\xac\xbb\xda\x5b\x25\x97\xe5\x4d\xfe\x8d\xa3\x4b\x1c\x1a\x13\xbc\x8a\x37\xdb\x09\x72\x29\x05\xa5\x57\x8e\x69\x60\x83\xcf\xfe\xd4\xb3\xb5\x33\x63\xda\x46\xab\x3f\xb5\x8f\xbe\x7b\x7b\x98\x7d\xf8\x07\x63\x81\x10\xbe\x63\x5f\x6d\x54\xe9\xa0\x92\x97\xb1\x37\x9c\xfa\x8e\x92\x67\x88\xb7\x38\xd0\xed\x53\x2c\x65\x46\x2c\x0e\x49\x8a\x15\xa1\xf3\x56\x8e\x69\xeb\x44\xd6\x48\x49\x53\xc9\x76\x9b\x42\x86\xbd\x8d\xb4\x79\x16\xc9\x74\x69\xc1\x67\x46\x42\x7f\xba\xa9\x8f\x2c\x76\xa5\xc9\xb7\xba\x32\xf3\x42\x30\xd4\xf0\x1e\x43\x9b\x28\xa3\x30\xab\x5f\x50\x7a\x15\x3d\x97\x58\xaf\x4a\xa9\xf6\xb7\xd7\x8f\x00\x4c\x63\x1c\xf8\xc9\x43\x32\x99\x4e\xd6\x15\x41\xf9\x99\x18\x7d\x9e\xab\xc7\x2e\x47\x6a\xc5\x3b\xc7\x72\x87\xa4\xb6\x9a\x7c\x12\x53\x54\x3a\x3a\x38\xd4\x66\xdd\x0a\x85\x83\x48\x69\x50\xd0\xb0\x7e\xff\xdd\xfe\xdb\x05\x43\xa7\x07\xa7\x5d\x90\xfe\xe1\xaa\x67\x36\x90\x1f\x67\x9a\x07\x58\xbf\x2f\xaf\xdf\x1e\xec\x2f\x53\x0b\xd1\x87\x44\x9b\x97\x52\x6b\x2b\xd6\x64\xca\x25\xe4\x94\x04\x41\xc1\x8a\x5e\xc2\x6b\xed\x89\x63\x36\x44\xf4\x82\x21\xba\xa5\x1b\xb1\x84\xd7\x79\xe8\x71\xcb\xb6\x4f\x11\x6d\xe6\x24\xbb\x96\xaf\x46\xfd\x67\x35\xd4\x1c\x4d\xfe\xdc\x10\x2c\x7e\xa2\x70\xf1\x96\xd8\x7f\x9b\x1c\xc6\xfb\x93\x56\x65\x1a\x8a\xcf\x56\x4d\xba\xa1\xb2\xc3\x33\x3e\x18\x53\x0f\x53\x0a\x6f\x9f\x7a\xd8\x11\xab\x69\x0a\x4a\xcf\x3b\x44\x7e\x72\x3c\x78\x9b\x94\x43\x6d\xfc\x57\x08\xcf\x17\x5c\xad\x8e\x0f\xf6\x3e\xd7\x7d\x58\x52\x43\xf0\x80\x90\xcb\x38\x2a\x3e\x6a\x15\x10\x0f\x06\xb9\x22\xd7\x6f\x5f\xca\x5f\xdd\x92\x36\x32\x47\xba\x5d\xcc\x85\x62\x4c\x74\x61\x69\xcc\x52\x8b\xd6\x8e\x59\xde\x60\xfa\x70\xbb\x13\x9b\xe7\x2f\xe3\x5d\xb2\x5c\x62\x9e\x5d\x0c\x93\x3b\x3b\xd9\x97\x38\x13\x0e\x17\x51\x0c\x55\x72\x51\xd1\xe2\x98\x63\x6d\xc1\x8a\x29\xb6\x41\x33\x35\xc3\xe6\x98\x2f\xe2\x69\xcb\x23\xcb\xcc\xd6\x10\xa6\xa8\x10\x6c\x74\x2d\xcf\xe1\xb7\xe6\x98\x37\x0b\x66\xf8\x29\x56\x6b\x08\x89\x79\x14\xe7\x92\x5d\x51\xdc\x35\xfe\x81\x1a\x6d\x53\xf4\x91\xac\x48\x0a\x7c\x79\xab\x42\x8d\xc1\x2d\x7d\xbd\xa3\x33\xd2\x8c\x16\x91\xfd\x43\x6b\xbf\x69\xf6\x3f\xca\x29\x81\x6a\xa0\x6e\x7e\x3b\x5d\xd1\xed\xaa\xf1\x3c\x32\x79\xcb\xfd\x14\x4d\xea\xaa\x2e\xcc\x37\xde\x1d\xdb\x24\x7f\x0f\xb2\x75\x75\x0b\x5d\xfb\x82\xf6\xaa\x0c\xfe\x22\xd0\xed\x43\x35\xc4\xc4\xcd\x6e\x89\x11\xb6\x94\x3a\x15\xee\xc8\x8d\xb4\x53\x6d\xd0\xfc\x57\x23\xab\x63\xe5\x05\x8d\xbd\x46\x72\x11\x4d\x63\xaf\x61\x79\x8d\xbd\xc6\xae\x3b\xee\x4d\xdc\xee\xa8\x37\x1c\xbb\x93\x61\x67\xfc\xa1\x1d\x33\x38\x47\x8d\x3f\x72\x69\x56\x97\xec\x92\x70\x9c\xde\xc7\x6a\x83\x30\x0e\xf2\x4c\xc3\x7b\x6c\xae\x69\x1c\xdc\x86\x8b\xdb\xe4\xf6\x8a\x49\x52\xda\xd8\x02\x5a\xcc\x64\x3c\x38\x19\x4c\x4e\xde\x4d\x54\x8c\xa6\x54\x29\x8b\x87\x6a\xce\xd2\x06\x38\x62\x59\xae\x85\x62\xba\x8c\xa5\x0a\xa2\x12\xa9\xbe\x3d\x08\xe1\x6a\xd5\x81\xd8\x7e\xf7\xc6\x10\xdf\xee\xc8\xdb\x78\x0a\x75\x2a\xeb\xbc\xd6\xb4\x9c\x2d\x5d\x1b\x42\xab\x0b\xa2\xd5\xc5\xc0\x52\x60\xe6\xb8\xd9\xba\x56\x95\xb0\x2e\xd8\x2a\xb4\x0b\xea\x0f\xfb\x56\x17\xca\xf4\xa7\x99\x4d\xe7\xea\x65\xf4\x2c\x66\xba\x6a\x6b\x85\x56\x40\xe6\x9b\x1a\x26\x0c\x78\x8f\x35\x8f\xdc\x0f\x59\xfa\x3d\x09\x4d\xbc\xc7\x54\xf3\xba\x28\x92\x0f\x8f\xd4\x30\x2f\x7d\xa6\x93\xaa\xbd\xe4\xe4\x7e\x1d\x2b\xfb\x9c\x57\x4c\xdc\xe2\x2e\x09\x39\xba\xe6\x85\x61\x6a\x98\x9e\x52\xe8\x95\xdf\x23\x4a\xde\x95\x11\x2b\xaf\xee\x06\xdf\x75\x89\x2b\x87\x1d\xeb\x03\x8f\xf5\xa1\xc7\xa4\xbf\xa2\xbc\x96\x17\xfa\x8d\x4b\xbd\x2e\xf7\x65\x67\xb7\xba\xfc\x96\x56\xdf\x98\x2f\x08\xc5\xff\x56\x01\xb1\x7a\x17\x24\x40\xef\xb0\x7c\x9e\x60\x43\xc6\x94\x8d\x7c\xcc\x09\x7d\xc2\x07\x84\x29\x09\x50\xb6\x23\x98\x1c\x3d\xf6\xa5\xa9\xc6\xe2\xf4\x25\x7f\x41\xe5\x9a\xf4\xaf\xba\x64\x96\x47\x26\xd3\x0a\xa3\x2b\xf4\x8d\xc9\x24\x70\xb8\x33\x99\xb4\x2d\xa2\x07\xca\xb5\x7b\x01\xaf\x4f\xbf\xe8\x63\xd0\x25\xfa\xd6\x52\xdf\x14\x35\xcf\x4b\x4b\x41\x73\xf5\x74\xad\x96\x19\xe4\x51\x18\xa1\x49\xfa\xf2\x9c\x2d\x3f\x1e\xe5\x8a\x1e\x89\x95\x5f\xce\x22\xad\xce\x11\x4b\xa3\xe0\x49\x73\x85\x41\x61\x8f\xe6\x33\x99\x4e\xd6\x6c\x4e\x0a\x2f\x18\x7b\xd5\x86\x89\x86\x95\xe1\x07\xc3\x52\x2b\x7d\x69\x19\xd5\xf8\x69\xff\xa7\xfd\x9d\x82\xd1\xcc\x29\xf6\x26\x14\x49\xfe\x99\x00\x5b\x89\xab\x36\x49\x38\x0c\xfe\x35\x91\x18\x4e\x26\x7f\x94\xd6\xe4\x39\xba\xb6\xc1\x9c\x4c\x5e\xb7\xfe\x56\x3c\xb0\x03\x3d\x15\xc6\xf0\x29\x89\xee\x0e\x59\x58\x01\x8f\x05\x3b\x44\xfc\xb1\x40\x47\x94\x78\x88\xb1\x47\x04\x9f\x88\xc9\x63\xf5\xc0\x99\x3f\xdd\x00\xdb\x28\xc0\x5a\x52\xb1\x15\x11\x9f\xe9\x62\x97\x9f\x12\x9a\x30\xdf\x2c\x75\x42\x93\xdb\x20\x22\x7e\xc5\x04\x90\x5e\x7d\x45\xca\xc3\xaa\xcd\xa1\x00\x19\xdf\x46\xde\xd1\x87\x79\x2b\xd9\x17\x6a\x6f\xa2\xe1\x1f\x11\x7f\x92\x3b\x8c\x13\x8d\x1f\x98\x4c\xd4\x44\xff\xc3\x48\xb9\x6c\xa3\xb1\x4c\xef\xa2\x51\x75\x6f\x1c\x22\xc8\x17\x66\x0c\x28\x8a\x02\x58\xba\xa0\x5f\xe9\x10\xd5\x99\x0d\x64\x57\x14\x7b\x4c\x42\x99\x4c\x4c\xf8\x96\x44\xc3\x84\x2f\xf4\x7d\x2a\xe6\xc0\x64\x0f\xdc\x16\x79\x42\xf9\xf6\xc8\xa7\x18\xfd\xeb\xff\xda\x7f\xfc\xed\xcd\xeb\x7f\xd8\xf6\xff\xf1\xff\xf6\xe6\x1f\x3f\xbf\x16\xff\x2b\xd5\x94\xad\xe5\xfe\x14\xd8\x3d\xb0\x77\x0f\xd7\x52\x21\x1b\x40\x61\xa0\x29\x2a\xb2\xda\x12\x1a\xb9\x69\x1e\xaf\x92\xb8\x0a\xe5\xee\x0e\x50\x23\xe0\xeb\x6c\x85\xdc\xc8\x97\x32\xa4\x6c\x76\xdd\x55\x5e\x4c\xb0\xee\x20\xcb\xa2\xed\x03\xa0\x90\x82\x7a\xd4\x05\xef\xf3\xf2\xfa\xb1\x34\xef\xe7\xd5\xf2\xf9\xae\x4a\xc2\x44\x13\xd3\xd9\xd4\x09\x43\x11\xa4\x90\x13\x6a\x83\xa6\x5d\x8e\x46\xc8\xfe\xbd\xc4\xb1\x7d\x4d\xe8\x7c\x02\x23\xe8\x2d\xd0\xc4\x83\x4b\x14\x4c\x9c\x6b\x4f\x7a\x70\x6c\x4c\x38\x0c\xfe\xaa\x2f\x7f\x0f\x71\x80\xfc\xbf\x84\x96\x4d\x04\x3e\xa9\x21\xef\xc5\xce\x42\x65\x86\x0a\x67\x90\xf1\x14\x4c\x37\x7d\xf7\x65\xdb\x06\xef\x93\x64\xba\xec\x69\xa4\x7a\xe5\x7e\xef\x87\x41\x5e\x80\xaf\xb0\xe9\xa1\x90\xb2\x15\x7b\xcb\x63\x2d\xd2\x5e\x05\x4f\x74\xb2\xa2\x38\xac\xa7\x7b\xee\xe3\x85\xb1\xf9\x31\x9e\xff\xd8\xfa\x15\xfc\x07\x7d\xfe\xe3\x85\x11\xfe\x5b\x9c\x36\x2a\x91\xe8\x3f\xcf\x82\xdc\xff\xc4\x51\x45\x27\xde\xe2\xd4\x91\x21\x2a\xb0\xee\xe4\x91\x31\x88\x70\xbf\xc4\x7a\xd3\xf5\x2b\x4d\x2b\x89\xa6\xb4\x66\x38\x40\xc9\x95\x12\x59\xd7\xed\x62\x58\xa5\x59\x6d\x9c\x68\x82\x96\x70\x35\x5b\x14\x09\xf5\x89\x49\x78\xfc\x76\xdf\xd7\x2b\xdf\x35\xa5\x3f\x5b\x43\xf2\x8f\xeb\xb3\xcc\x4b\x0d\xee\x7c\x05\x41\x09\xce\xcb\x7f\xc6\x21\xd7\xd2\xf7\x7a\xc6\x61\x33\x98\x4d\x79\xf2\xa6\x78\x5e\xe5\x19\x07\x3d\x65\xbe\x26\xd5\x7f\x8b\xe8\xa0\x19\x5e\x51\xc2\x77\x8a\x78\xaf\x3b\x02\x58\x8f\xf0\x1d\x2f\x0d\x35\x60\x71\xab\x71\x6d\x97\xaa\x6e\x86\xf0\xb8\xef\x36\x54\x07\x76\xb7\x77\x1b\x32\x38\xf6\xea\xf0\x5b\x5c\x3c\x90\x2e\x55\xb5\x66\xcb\x38\xa9\xf0\xad\x82\xea\x25\x5e\x27\xf9\xf0\xca\xd6\xab\xa6\x42\x14\x46\xe1\x63\x16\x05\xf0\x26\x49\x24\xce\x1e\xab\x06\x17\x0a\x46\x96\xf8\x89\x98\x47\x71\xa4\x7c\xa3\x2c\xaa\x0e\xd4\x59\x13\x2c\xcf\x95\x80\xa4\x5b\x3d\x75\x5d\x5e\xd8\x93\xc2\x6c\x99\xfa\x0f\x48\x38\xb7\xea\xa0\xff\xb6\x40\xa1\xba\xce\x27\x65\x81\x3c\xba\x22\xe0\xab\xa3\x26\x21\x80\x61\xd6\x2f\x27\x00\x82\x10\x5d\x65\x1d\x82\xe4\xa8\x97\x68\x14\x21\x3a\x23\x74\x89\xfc\x56\x06\x7d\x10\x06\x37\x12\xf1\x22\x45\x41\xb6\xe4\x33\x00\xa9\x18\x96\x0f\x39\xf2\xf7\x84\xa2\xa6\x71\x18\xe2\x70\xae\x87\x46\xe5\xe8\x49\xec\x2d\x74\xd0\xef\xd0\x8c\x50\x04\x7c\x22\x2a\x17\x68\x03\xa6\xd0\xbb\x8c\xa3\x02\x4a\x00\x86\xbe\xdc\x78\x96\x27\x68\x70\x08\x3c\xc8\x10\x20\x33\x00\x01\x25\xea\x5a\xa9\x56\xe6\x5a\xcd\x99\x0d\x1a\x29\xab\xf7\xe2\x22\x9b\xb0\x47\xc2\xae\x58\x41\x6c\xf9\x4f\x8b\x2e\x2a\x57\x43\x94\xe8\x9f\x5d\xb0\x55\x12\x84\xe4\x3d\xe1\x3d\xd0\x0b\xbd\xec\x49\xe9\xa2\xe4\x10\x2f\x16\x96\xba\x0a\x10\xcb\x94\xf9\xec\x6d\x7a\x4d\x6a\x8d\x6d\x59\x1c\x89\x15\xb6\xd4\x4a\xf9\x25\x2d\x8a\xfc\x05\xe4\x2d\x8f\xc8\xb4\xe3\x48\xe8\x21\xc4\x13\x25\x95\x6a\xc6\xec\x30\xc0\x47\x67\xe4\xf6\x06\xea\x11\xdc\x82\x18\x35\xc7\x0b\x94\x09\x00\x27\xba\x94\x10\x0a\x1a\x6a\xff\xb9\x21\x05\x54\x30\x28\x44\x57\x88\xf1\xb4\x81\xb2\x5a\x92\xec\x0f\x6d\xab\x3a\xb9\x55\x2a\x4f\x3b\x5f\xaa\x54\x83\x4c\x6a\x33\xd1\x4b\xbb\x4b\xe2\x34\x40\x66\x15\x20\x7f\x07\x00\xa2\x6f\x63\x56\xbc\xd0\xcc\x0f\x4d\xe2\xe6\x65\xb3\xd8\x3c\xd5\x2d\xed\x78\x44\x42\x11\xa5\xf6\x74\xcb\x74\xfd\xc5\x38\x11\x92\x01\x9f\x9d\x92\x56\xd7\xf2\x0d\x54\xd7\xba\x72\x01\xb9\x31\x9a\x1d\xc6\x29\xa3\x66\xd7\xa1\x06\xea\x8f\xcf\x67\x0d\x2e\x86\xa7\xa3\xce\x89\x93\x86\xfd\x0b\x0b\x5c\xed\x01\xdc\x75\x67\xd5\xb7\x3d\xad\xbe\x9d\x31\x5d\x36\xa5\x2d\xd0\xb0\x2c\x35\xc1\x8b\x87\x6c\x1b\x6d\x12\xf1\x76\x5a\xf2\xaa\xd4\x82\x78\x56\x40\xe6\x38\x2c\x97\xec\x7e\x19\x74\x27\x67\x83\xd3\x5e\xff\x6b\xa3\x64\xa1\x94\x2d\xab\xe2\xd5\x4d\x79\x67\xba\x57\x13\x4f\x55\x71\xa5\x44\x11\x5d\x7d\xb6\x7c\x9c\xca\x40\xc9\x16\xb2\xea\xea\x6d\x61\xf9\xac\x7d\xc0\xaa\x94\x72\xd3\x47\x2b\x44\xff\x5f\x00\x00\x00\xff\xff\xd0\xc9\xbf\x36\xc7\xb4\x00\x00")

func templateYamlBytes() ([]byte, error) {
	return bindataRead(
		_templateYaml,
		"template.yaml",
	)
}

func templateYaml() (*asset, error) {
	bytes, err := templateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.yaml", size: 46279, mode: os.FileMode(436), modTime: time.Unix(1532094241, 0)}
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
	"bindata.go": bindataGo,
	"template.yaml": templateYaml,
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
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"template.yaml": &bintree{templateYaml, map[string]*bintree{}},
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

