// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// DO NOT EDIT!

package internal

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _templatesViewsManagerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\xdd\x93\xdb\xb6\x11\x7f\xf7\x5f\x81\x41\x2e\x99\xbb\xd8\x14\x6d\xd7\xcd\x4c\x13\x49\x9e\x8c\x93\xa9\x3b\x75\xdb\x99\xe4\xe2\x3e\xb8\xee\x0d\x48\x2c\x25\x5c\x21\x40\x06\x40\xe9\xae\x9a\xfb\xdf\x3b\xf8\xa0\x44\x52\xa4\x44\x52\x77\xe7\x3a\xf7\x70\x22\xf1\xb1\xbb\xf8\xed\x07\x80\xe5\x6e\x36\x88\x42\xc6\x04\x20\x9c\x4a\x61\x40\x18\x8c\xee\xee\x9e\x8c\x29\x5b\xa1\x94\x13\xad\x27\x58\xc9\x35\x9e\x3e\x41\x08\xa1\x72\xeb\xcd\xd5\x92\x08\xe0\xa1\x67\xbf\xd7\x30\xc3\xa1\xd4\xeb\x46\xcc\x5f\x4e\xdf\x48\x91\xb1\x59\xae\x88\x61\x52\x8c\xe3\xf9\xcb\xda\x90\x12\x91\x94\x03\x51\x19\xbb\xc1\xd3\x71\x4c\xd9\xaa\xc4\xa9\xf6\x5a\x61\x5c\x2c\xa3\x4a\x77\xb3\x41\x2c\x43\x23\x50\x4a\x2a\xbb\xc0\x36\x9e\x84\x83\x32\xc8\xfd\x8f\x28\x11\x33\x50\xc5\x0b\xd3\x0b\xa6\x35\x49\xf6\x56\xe5\x48\x24\xb9\x31\x52\x20\x73\xbb\x84\x09\xf6\x2f\x78\xb7\x0e\xa9\x01\x23\x4a\x0c\x29\xc8\x04\x46\x18\x11\xc5\x48\x34\x67\x94\x82\x98\x60\xa3\x72\xc0\xd3\x6f\x0c\x5b\x80\xfe\x61\x1c\x7b\x32\xfb\xcc\x36\x9b\xb0\x90\xd1\xcf\x8d\xcb\xa9\xa2\x13\x66\x80\xa0\x76\x60\x75\x64\x26\xd5\x02\x29\xc9\x61\x82\xed\x23\x46\x24\xb5\x6a\x99\xe0\xaf\x30\x5a\x80\x99\x4b\x3a\xc1\x4b\xa9\x0d\x46\x8c\x4e\xac\x81\x64\x6c\xa6\xc3\x4a\x8c\x9c\xcd\xec\xc4\x15\xe1\x8c\x12\x23\x55\x13\x2c\x16\x59\x4f\xdf\x90\xa4\x6e\x30\x95\x91\x39\x2f\xe0\x12\x64\x85\x04\x59\x45\x86\x24\x1a\x25\x44\x5d\xd9\x07\xbc\x23\xc3\x99\xae\x6b\xb7\xb6\x56\x65\x15\x87\xce\x04\x59\xc0\x33\x74\xb6\x22\x4a\xa3\xef\x27\x68\xb4\xb2\x60\x27\x1c\x74\x1d\xb1\x8a\x20\x9c\x05\x56\x4b\x05\x1a\x84\x71\x86\x8a\x83\x01\xc1\x27\x4f\x16\x05\x30\xac\xb3\x6c\x4d\x27\x35\x6c\x05\x78\x8b\x75\xbb\x88\x8e\x0f\x41\x73\x05\xd9\x04\x7f\xb5\xd9\x04\x9a\x77\x77\xc1\x1e\xac\x09\x2b\xc9\xf5\x04\x57\xfa\xb6\x08\xd4\x34\xe0\x5a\xdc\x44\xb8\x59\x12\x41\x81\xba\x89\x2d\xe2\x5a\x23\xb3\x32\x72\x6d\xa9\x66\x84\x6b\xd8\x8a\x7c\x00\xd6\x12\xbc\x85\x44\x87\xd7\x17\x93\x76\x62\xe3\x98\xb3\x83\x1a\x0c\xd6\xda\x3c\x37\xe7\xd3\x27\xcd\x5d\x25\x47\x36\x24\x89\x9a\x43\x41\x8d\xd5\x49\xc6\xd2\x60\xdf\x65\x01\x6c\xcb\x01\xc3\xf1\x16\x83\x98\xd8\x29\xc0\x39\x5a\x59\xeb\x47\xac\xa8\xba\x62\x0e\x91\x02\xbd\x94\x42\x5b\x4b\x3c\xae\xcc\xb1\x9b\x53\x21\x80\x3c\x99\xb9\x5c\x41\x93\x47\x37\x53\x99\x03\xa1\x5d\xc7\xaa\x6e\x03\x03\x61\xb4\x66\xd4\xcc\x27\xf8\xe5\xf3\xaf\xf1\xf4\xef\x64\x01\xe3\xd8\xcc\x87\x51\xf8\xe3\xd7\x78\xfa\x9e\xf0\x7c\x38\x89\x17\x96\xc4\x4f\x90\x91\x9c\x9b\xe1\x44\xec\x4a\xfe\x49\x4c\x3a\x07\xa5\x7b\x53\x99\xfe\x04\x3a\x55\x6c\xe9\x77\xcf\xae\x93\xc7\x71\x57\xd8\x2d\xcd\x1e\xca\x4c\x24\xbd\xed\x36\x76\xe7\x69\xec\x19\x3a\xf3\x4e\x60\x1d\xcd\x7b\xdc\x91\x60\x52\x22\x62\xc7\xff\x85\xda\x99\x0a\x96\x9c\xa4\x50\x10\x1b\xbd\x0f\x1e\x3b\xfa\x2b\xdc\x22\x3c\xc2\x08\x5f\xe1\xae\x84\x7b\x9a\x65\x47\x7c\x4a\x62\xb3\x6c\x5f\xce\x9f\x29\xf3\x2e\xd7\x51\xc8\x2d\x7f\x4e\x12\xd8\x6e\x97\x61\xb3\x88\x5c\x23\x46\x99\x54\x3e\x80\x78\x9c\x6c\x04\xb1\x6f\x4d\x18\xdd\xdd\x8d\x63\x37\xab\xf7\x6a\xc2\xe6\x71\x6f\x52\x3f\x84\x88\xed\x3b\x48\xa3\x74\x71\x1f\xa5\x7e\x76\x0b\x28\x85\x7d\x7b\x70\x8b\x66\x4a\xe6\xcb\x8e\xe1\xba\x5d\xae\xb7\x44\xbf\x67\xb0\x46\x18\x44\xbe\xe8\xec\x3b\x15\xc1\x34\x70\x48\x4d\x45\xb6\xa0\x6a\xe4\xbb\x5e\x62\x64\xb7\x36\x6f\xa3\xcd\x2a\xdf\x6d\x82\x3b\x1b\xee\x2d\x09\xaa\xc4\x1c\xf9\x0c\x9d\x49\x17\x34\x6d\xe4\x38\x2f\x18\xff\x19\x8c\x5d\xf0\x3f\x7c\x0f\xf6\x23\x34\xbe\x18\xb2\x74\xb7\xfc\xc0\x63\x65\xb7\x19\x7f\x08\x13\x14\x6e\xb6\xbc\x9f\xdb\xb5\x6c\x0f\x04\xe7\xb5\xce\x8b\x7d\x40\xdc\x7e\x65\x8f\x0a\x1e\x3c\x7b\xb2\x2b\x9e\x4a\x07\xcd\x3d\x3e\x2f\x9c\xe3\xf8\x97\xc1\xd0\xf5\xf4\xa0\x2d\x06\xb1\x97\x70\x90\x31\xba\xc0\x12\x8e\x4b\x75\x2c\x2e\x6f\x97\x80\x70\x22\x25\x1f\x66\x9a\x4c\x2c\x73\x13\x6e\x68\xe9\x1c\xd2\xff\x24\xf2\x66\x7b\x5a\xbb\xd6\x91\x5e\x33\x93\xce\x07\x99\xe7\xee\x76\x70\xbe\x54\x4c\x98\x16\x45\x5e\x20\x7f\xc5\x73\x97\x06\x2b\x01\xd0\x2d\xce\x28\x3e\x09\x30\xa9\xd0\x79\x3b\x68\x4c\x18\x7c\x71\x64\xc0\x77\xaf\x86\xd9\x7d\x05\x57\x91\x2f\x12\x50\xb8\xc9\xff\xbb\x00\xbb\xf3\x9b\x56\x4f\x68\x42\x7f\x49\x8c\x01\x25\x26\xf8\xdf\x1f\xa2\xa7\x1f\x5f\x7f\x78\x1e\xfd\xe9\xe3\xb7\x67\xf8\x21\x11\xcd\x8f\x42\x9a\xff\x6e\x30\xbd\x37\x48\xdb\xc1\xca\xb8\x24\x16\xad\x2f\x1f\xac\x9d\x01\x9e\xff\x6b\xe4\x1f\x2e\x5e\x3f\x1c\x70\x34\x64\xd2\x4e\x47\xce\xc0\x8d\xf9\xcc\xb8\x9d\xef\x01\xf7\x81\x44\xff\xfd\xf8\xf4\xe2\xe9\xa9\x00\xee\x9d\x70\x96\x44\xeb\xb5\x54\xf4\x74\xdc\x76\x94\x9a\x0e\x3d\x45\x6f\xa4\xe7\x72\xfd\x70\x48\xde\x33\x3a\x86\xcc\xf4\x20\x64\xac\x50\xd4\x5f\x90\x2f\xe1\xc6\xb8\x0b\x5e\xf3\x59\x2b\x8c\x8a\xbc\xd9\x3d\x98\xf1\x22\xbf\x94\x87\xc2\x7d\xe8\xd9\xca\x82\x5e\xc6\xe9\xee\xce\x67\x87\x4b\xa0\x78\x4e\xd5\x41\xbb\x33\xdf\x09\x0a\xff\x02\x03\xc5\xc0\xd5\xf6\x3f\xc0\x36\xe4\xce\xbb\x30\x1a\x72\x1d\xd6\x4b\x22\x1a\xef\xbd\x05\x0e\xe3\xd8\x0d\xf9\xb2\x2f\xbe\x6e\x99\x7b\x0a\xbd\x72\x66\x10\xd2\xd8\x07\x4c\x22\xe4\xda\x86\xdd\x00\x5b\xee\xb6\x27\x45\xfe\x6f\xdd\xdf\x63\xba\x5e\x98\x7e\x00\x9a\x47\x73\x8d\x7e\xe6\xf8\xa8\x19\x95\x22\x9f\x3a\xcc\x0b\x87\x20\x18\x12\x0b\xeb\x67\xe8\x6c\xed\x99\x97\x37\xbb\xa1\xf2\x6c\x65\x2a\xe2\xab\x4f\x9a\xb9\xff\xc5\xee\xe0\xb3\x65\x81\xe9\xe8\x57\x99\xab\x74\x68\xb8\x40\x8f\x66\x0d\xc3\x38\x3d\x46\x70\x6a\x8c\xc1\xbf\x69\x32\xeb\xed\xae\xde\x24\x89\xa0\xe8\x5c\x34\xa4\xa5\xb7\xc9\x3e\xa3\x72\xb8\xd8\x25\xa1\xaa\x59\xb7\x61\xf7\xc5\x44\xa1\x78\xfa\x23\xe7\x72\x0d\xd4\xef\xb1\xfa\xfb\xfe\x54\xf2\x9e\x69\x56\xf4\x39\x72\x6c\x07\xd2\x67\xb5\x96\x17\xc3\xf3\x78\x9c\xed\x27\xd5\x9e\xfb\x6c\xf4\x81\x0f\x97\x47\xc4\x3e\x65\x0b\x68\x95\x08\x7d\xb3\xa0\x44\xcf\x7f\x40\x2d\x49\xc0\x53\xe4\x1d\x96\x00\x3c\x25\x75\xd8\xd7\x02\xff\xdf\xb6\xb2\xee\x1f\xda\xfa\x49\x3e\x8e\x3b\x7e\x6a\x1b\xc7\x2e\xc6\x1c\xf9\x6a\x7c\xf8\x94\x7b\xa4\xfb\xd8\xe7\xf9\xc6\xc9\xa1\xb9\xb9\x42\xa4\xd8\xeb\xc4\x95\x96\x9c\xd1\xbd\x7a\x9f\xa6\xc1\x47\x3f\x7b\x54\x2a\x89\x24\x8f\x16\x34\xfa\x0e\x85\x07\x99\x65\x1a\x4c\xf4\x87\x03\xa7\xca\x6a\x45\x8f\x02\x0d\xbb\x4b\x4f\x62\x04\x4a\x8c\x88\x98\xc8\xe4\xf6\x4a\xe3\x87\x4c\x7f\xb1\x3f\xed\xb5\x3b\xcd\xe4\x75\x9e\x2c\xd8\x3e\x7d\x9d\xa7\x29\x68\xbd\x65\xa1\xc9\xca\x1e\x9a\x99\xab\x42\xa2\x13\x5c\x3c\xe1\xe9\xaf\x64\x05\x87\x99\x1e\xd6\x4b\xb5\xc9\x62\xdb\x58\x6d\x15\x1e\x8b\x9f\x12\xc2\x0b\x49\x09\x47\x19\xa1\xe0\xaf\x6f\xee\xfd\x4d\x28\x75\x30\x24\x71\xa1\x69\x82\xa3\x17\x45\x19\x0b\x65\x84\xcb\x59\xa8\x5b\x71\xe7\x1b\x0e\x34\xb9\xad\xcc\x7c\xe7\x3e\x0e\x7a\x41\x1a\x0a\xa5\x9e\xd4\xd5\xec\xa6\x46\x81\x72\x73\x7d\x98\x1f\xd2\x5c\x18\xb2\x3f\x6e\x0e\x84\x36\x96\x41\xf4\xaf\xf7\x72\x04\x87\xd5\x7b\x8d\xe7\xaf\xaa\x62\xf9\xda\xba\x3a\xce\x1e\x2d\x5f\x60\xa7\x16\xc8\x1a\x0b\x4a\x6b\xd5\x76\xaf\xea\xaa\xde\xd7\xfe\x1e\x08\x36\xf0\xe0\xe9\x8f\x0a\xd0\xad\xcc\x91\xce\x15\xbc\xee\x34\x2f\x93\xd2\xf4\x07\xaf\xb0\xfd\xe2\xb0\xdb\x08\xe3\xf4\x8d\x85\xf8\x00\x62\x5d\xbc\xcb\x97\xf7\xb5\x30\x68\x77\xa8\xf6\x5a\xc4\xaa\x77\x94\x6a\xee\x4a\x45\x96\xd6\xa2\xb6\x37\xcf\xcd\x06\x69\x43\x0c\x4b\xdf\x5e\xfe\xed\x1d\x3a\xf7\xcf\xbf\xfd\xf2\x0e\xe1\xd8\xee\xe9\x89\x24\x8a\xc6\x44\x6b\x30\x3a\x5e\x81\xa0\x52\xe9\xd8\x4a\xeb\x42\xbc\x1e\x09\x30\x51\xa2\xe3\x54\xfb\xd6\x4b\xdf\x9a\x48\x69\xb4\x51\x64\x39\x5a\x30\x31\x4a\x6d\xf8\x70\x45\x5e\x17\xf7\xc8\x35\x63\x37\x40\xbd\x73\x14\x12\xb8\xa6\xb7\xae\xe9\xb0\x08\xcd\xb8\x5c\xeb\x7b\x44\x25\xbe\xd6\xf1\xf5\xa7\x1c\xd4\xed\xa8\x04\x8c\x95\xe5\xfa\x21\xd0\x48\xb4\x65\xd8\xaa\x82\x07\xe1\x59\xd2\x40\x8d\x79\x59\x11\xbd\xd9\xfb\x78\x51\xf0\xbe\xd6\xf1\x82\x08\x32\x03\xe5\xa8\xf8\x7b\x4b\x59\x85\xff\x0b\x00\x00\xff\xff\x82\xe6\x45\x3d\x3e\x2c\x00\x00")

func templatesViewsManagerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsManagerHtml,
		"templates/views/manager.html",
	)
}

func templatesViewsManagerHtml() (*asset, error) {
	bytes, err := templatesViewsManagerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/manager.html", size: 11326, mode: os.FileMode(420), modTime: time.Unix(1507071611, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsManagerMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x51\x6f\xe3\x36\x0c\xfe\x2b\x0e\x5b\xf4\x24\xc4\x55\xd3\x01\x7b\x49\xa2\xf4\x21\xbb\x87\x01\xc3\x36\xe0\x6e\x7b\x09\x82\x41\xb6\xe8\x44\xa8\x2c\x79\x92\x9c\x5c\x11\xf8\xbf\x0f\x52\xdc\x9c\x93\x5c\x8b\x6d\xb8\x37\x5a\xd4\x47\xf1\x23\x3f\xd2\xb7\x44\xda\xb2\xad\xd1\x04\xca\x1c\x0a\xf9\x42\xaa\xd6\x94\x41\x59\x43\xe8\xe1\x96\xc0\x4d\x69\x4d\xa5\x36\x3e\x53\xa6\x69\xc3\x4a\xc9\x75\x9e\x9d\xce\x3c\x6a\x2c\x03\x50\x56\x6e\x85\xd9\xe0\x10\xba\x13\x2e\x43\x7e\x4b\xc2\x56\x79\x9a\x3b\xbb\xe7\xc8\x1a\xe1\xd0\x04\xff\x87\x09\x4a\x13\x08\x85\x95\x2f\x90\x43\x70\x40\xf3\xb2\x75\xd1\xc7\x91\xed\x84\x26\x34\x97\x58\xf1\x08\x65\x12\x2b\xd1\xea\xf0\xa7\xd0\x2d\xce\xa0\xdc\x62\xf9\x5c\xd8\x2f\xc0\x63\x38\x67\x1b\x02\xe1\xa5\x41\xa0\x4f\xe4\x35\x02\xc0\xe8\xe4\x4b\xd7\x51\xc2\x75\xbc\xe5\xd1\x43\xa7\xf0\xe9\xe3\x2f\x1f\x97\x9f\x81\x1f\xdd\x41\x6c\x7e\x15\x35\xde\xdd\x91\x88\x40\x56\x29\x23\x09\xd8\x26\xb2\x02\xca\x2a\xa5\x03\xba\x21\x4f\x87\xa1\x75\x26\xeb\x89\xf6\xef\xf6\x8f\x7c\x4a\xe5\x89\xef\x77\xf4\xc8\xeb\x2b\x51\x2e\xb1\x7a\x72\x76\xcf\x1c\xd6\x76\x87\x4b\x2d\xbc\x27\xb0\x15\xfe\x1e\x9d\xb3\x0e\xe8\x34\x3a\x85\x94\xd7\x9e\x7c\xd8\x96\xa2\x0d\xc1\x9a\x55\xac\x02\xf7\x6d\x51\xab\xb0\x86\x53\x1a\xca\x8b\x42\xa3\x84\x7c\xc2\xf9\x10\x14\x1c\x1b\xc4\x63\x1a\xcd\x26\x6c\x69\xf7\x76\x64\x87\x1e\x53\xe0\x52\xab\xf2\xf9\x2d\x89\x5c\x84\x7d\x8b\xd9\xff\xc8\x7f\x34\x79\x2f\xbb\xaf\xb8\xab\xf4\xa2\x0c\xeb\xc4\xbd\xb6\x52\xe8\x65\x42\x47\xbd\x71\x80\xd9\xa9\x77\xdf\xa6\x90\x05\xd9\xab\xde\x88\x1a\xd7\xe9\xad\x11\xdf\x2a\x29\xd1\x0c\xa7\xe0\x12\x73\x9a\x0a\x14\xe5\xf6\xcd\x99\x78\x4f\xcc\xbb\x6b\x05\x4f\x77\xa7\xd9\x28\xc7\x1c\xe6\x5a\x2d\xe6\x3e\x38\x6b\x36\x0b\x18\xbf\xde\x8e\x79\x02\x1d\xc3\xfc\xa1\x77\x4d\x33\x18\xef\xe2\xb7\x56\x0b\xe8\x68\x5e\xf7\x82\x66\xa9\x1a\xf7\x69\x02\x29\xdb\x86\x5a\x13\x58\xa6\x19\xf6\xd9\x51\xec\x7e\x3a\x2f\x5c\xf6\xb0\x98\xb7\x7a\x01\xe3\x32\xc6\x88\x56\x0c\x91\xb0\x84\xe6\xa3\xc7\xbe\x27\x83\xd2\xfe\x97\xbe\x48\x11\x04\x3f\x74\xff\xbe\x0d\x4a\x7e\xf7\x26\xc4\x1c\x56\xe7\xe5\x5b\xf3\x77\x3a\x73\xdd\x97\xbe\x2b\xb1\x14\xac\xb1\x3e\x10\xb8\x81\x3c\x86\xcd\x07\x8f\xee\x95\x91\x76\xcf\xb4\x2d\x45\x3c\x61\x0e\xb5\x15\x32\xa1\x4e\x55\x3c\x5f\xb4\xa9\x80\x8d\xf0\x7e\x6f\x9d\x5c\xb3\x57\xeb\xde\x6f\xed\x1e\x28\xb3\x86\x40\x34\x59\xe1\x4f\xbe\x6c\xab\x24\x0e\x0f\x20\xff\x06\x6f\xb8\x81\xf1\xf9\xb6\x52\x32\x6a\xe6\xaf\x5d\x5c\xb1\x40\x67\xaa\x22\xf8\xba\x12\x12\x2a\x6e\xe5\x80\x5f\x02\xa1\xb3\xde\x40\x16\x09\x12\xe8\x21\x34\x3f\x3f\xc8\x03\xed\x3a\x3a\x4b\xd8\x38\xc0\xe7\xbb\x27\x9e\x00\x65\x3f\x89\x20\x3e\x47\x9b\x1c\x8a\xdf\xc5\x46\x19\x11\x70\x3a\x7a\xcc\x8b\x9f\x4d\x65\xa3\x21\x9d\xd8\x2f\x85\xd6\x85\x28\x9f\xa7\x17\x44\x44\xa3\x8e\xfb\x5a\x34\x8a\xa4\xdf\x8b\xe7\xa2\x51\x2c\x1a\xe4\xd0\x88\x0d\x4e\xa1\xdf\xb4\xd0\x51\x66\xac\x44\x4f\x68\xae\x85\x0f\xdc\xb4\x5a\xcf\xe2\xe5\xd2\xea\xb6\x36\x64\x92\x5f\x03\x12\x9b\x4b\xed\x6c\x9c\x6d\x9b\x5c\xd1\x83\xaa\x48\x8c\x34\xe2\x3c\x1d\x1d\x53\x6a\x84\x0b\x9e\xdf\x1e\x6f\xd1\xbe\x60\xcc\x37\x5a\x05\x02\x0c\x68\x1e\xd5\xc5\xd3\xad\xbe\xba\x8b\x1f\x9e\xd2\xe7\xea\x71\x3d\x3d\x1a\x93\xf5\xac\x0f\x6c\xd2\x4f\xe8\x96\x44\x42\x94\xe1\xdf\x44\x51\x56\x60\x65\x1d\x92\x0f\xf3\xe0\xb2\x32\x6e\x55\x0e\xe9\x31\x58\xcc\x83\xcc\x4a\xab\x7d\x23\x0c\x87\x1f\x61\xf1\x61\x1c\xf1\x71\x66\x83\x5c\xcc\x1f\x82\x8b\x83\x7b\x24\x2f\x6a\xec\xba\xd4\x9e\xb3\xa6\xf4\x9b\x20\xca\x2a\xcd\x6a\xfa\x2f\xb3\x63\xf8\x4b\x15\xf5\x75\xfa\xcd\x49\x74\x3c\xb5\x93\xd9\x68\x13\x1a\x09\x4c\x38\xe7\xc3\x1b\xab\xc9\xfa\xee\x0e\x84\x2f\xe1\xd2\xf1\xb8\x7e\x1a\xa2\x57\x93\x1c\x24\xfa\x12\xd6\x94\xc5\xde\x13\x3a\xbd\x74\x8b\xa1\xb7\xa3\x1d\x9d\xfd\x13\x00\x00\xff\xff\xcd\x92\xb1\x96\xbf\x08\x00\x00")

func assetsJsManagerMinJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsManagerMinJs,
		"assets/js/manager.min.js",
	)
}

func assetsJsManagerMinJs() (*asset, error) {
	bytes, err := assetsJsManagerMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/manager.min.js", size: 2239, mode: os.FileMode(420), modTime: time.Unix(1507072603, 0)}
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
	"templates/views/manager.html": templatesViewsManagerHtml,
	"assets/js/manager.min.js": assetsJsManagerMinJs,
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
	"assets": &bintree{nil, map[string]*bintree{
		"js": &bintree{nil, map[string]*bintree{
			"manager.min.js": &bintree{assetsJsManagerMinJs, map[string]*bintree{}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"manager.html": &bintree{templatesViewsManagerHtml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: k}
	}
	panic("unreachable")
}
