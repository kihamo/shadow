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

var _templatesViewsManagerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\x20\x54\x60\x68\x1f\x14\xf5\xdf\x5e\x5a\x5b\x43\xd0\x0d\x68\x81\xac\x18\xda\x0e\xc5\x9e\x02\x4a\x3c\x59\x74\x28\x52\x23\x4f\xb2\x8d\x20\xdf\x7d\x20\x25\xd9\x92\x4d\xc9\x49\x9b\x76\x69\x91\xa7\xd0\xe2\xfd\xe3\x8f\x77\xbf\x3b\x45\x57\x57\x84\x41\xc6\x25\x90\x20\x55\x12\x41\x62\x40\xae\xaf\x4f\x66\x8c\xd7\x24\x15\xd4\x98\x79\x50\xd2\x05\x84\xc8\x51\x40\x10\x9f\x10\x42\x48\x7f\xd3\x3d\xbf\x10\x90\x61\xbb\xe9\x04\xf2\x17\xf1\x67\xa5\x2f\x41\x9b\x59\x94\xbf\x68\xb5\x22\xc6\xeb\xf8\x64\xc4\x82\xe6\x8b\x7c\x60\xa2\x27\x91\x29\x5d\x84\x0b\xad\xaa\x92\x94\x95\x10\xe1\xbe\xec\xbe\x3c\x97\x65\x85\x8d\xc2\x9e\x94\x93\x14\x34\x01\x71\xf8\xdc\xed\x39\x55\x82\x9b\x12\xe6\x41\x9a\x43\x7a\x99\xa8\x75\xd0\xd9\x5d\x9a\xd0\xac\x38\xa6\x79\x40\x38\x9b\x07\xb4\x42\xa5\x21\xd3\x60\xf2\x80\x44\x31\x39\xdb\xfd\x3e\x74\x1a\x79\xbc\xb6\x88\x78\x7e\xb6\xcb\x0e\xb2\xfe\xe1\x52\x01\x54\x67\x7c\x1d\xc4\xbe\x5d\xad\x56\x04\x55\x79\x81\x5c\x80\xf1\x5c\x17\x95\xbc\xa0\x08\x8c\x64\x82\x97\xef\xe4\x3f\x24\x55\x22\x14\x8b\xf0\xa5\x5b\x14\xac\x5d\x98\xa2\x5d\xac\x4d\xf8\xec\xf9\xc8\xb5\x58\x1f\xa1\x41\x8a\x66\xea\x2e\x52\x25\x83\x78\xc6\xb7\x77\x49\x49\x46\x43\x06\x34\xb3\x7f\x7f\x75\xc7\xe0\xf1\x1e\x12\xfb\x46\x52\x55\xd9\xbc\xb4\xa0\x0b\x6e\x10\x24\x68\x13\x36\x0f\xe3\xa7\x3e\xdd\xfc\x45\x7c\xde\x09\xee\x32\xd0\x0f\xf2\x3d\x85\xa8\x32\xa0\xcd\x97\x61\xb4\x6a\x4a\xef\x18\x42\x07\x15\xfa\x43\xe1\x83\xd4\x5c\x7e\x21\x3e\x4e\xf5\x18\x3a\x9f\xac\xd0\x31\x6c\x06\x25\x38\x48\x4f\x4f\xed\xad\x2f\x4a\x2a\x41\x8c\x40\xb1\xbe\xe8\x53\xec\x2e\x92\xe7\xde\xe2\x69\x8e\x4c\xfa\x49\xfe\x7c\x4f\xb3\x12\x9d\x9a\xa4\x35\x91\xb4\x4e\xa8\x6e\xa8\x93\xb8\x38\x2e\x50\x29\x61\xe9\xcd\x47\x91\x3c\x9e\xd1\x1d\x72\x42\xd0\xd2\x40\x28\xb8\xbc\x3c\xbc\x88\x34\x87\x5a\x2b\x19\x5a\xb2\x6d\xee\x81\xc6\xb3\x48\x70\xaf\xd9\x4e\x95\x69\x55\x32\xb5\x92\x1e\xe7\x4e\x92\x92\x5c\x43\x36\x0f\x1e\x05\xfb\x1a\x21\xaa\xc5\x42\x40\x40\x18\x45\xda\xfe\xe8\xd9\x23\x5a\xd9\xdf\x49\x85\xa8\x64\x40\xa8\xe6\x34\x84\x75\x49\x25\x03\x66\x23\x16\x06\x0e\x8f\xb0\xd2\x20\xd3\x7c\x17\xbe\x3f\xa6\x1d\xa0\xdb\x58\x0a\x90\x55\xe7\xd2\xad\xfd\xaa\x1d\xa6\xa3\x9b\x83\x33\x2f\x69\x4d\x4d\xaa\x79\x89\xaf\x6a\xc5\xd9\xe3\xa7\x4f\xf6\x0e\x5b\x28\x46\x45\xf7\x8c\xea\x05\xe0\x3c\x78\xd4\x7f\xe8\xd6\x4d\xcb\x9e\x07\x6f\x94\xcc\xb8\x2e\x88\x86\x42\xd5\x40\xa8\x10\x64\x97\xa5\x7d\xf9\x44\xb1\xcd\x3c\xf8\xcc\x85\x20\x09\x78\xc5\x09\xac\x53\x28\xd1\x4e\x0c\xb4\x12\xb8\xdd\x38\x25\x67\x1a\xc8\x46\x55\xc4\x54\x1a\x7e\x1b\x58\x4d\xa9\x10\x09\x4d\x2f\x7b\xb5\xf1\xc1\x99\x7e\xfc\xe4\xf5\x04\x5c\x5b\x54\xb6\x57\xb5\x10\x9b\x32\xb7\x5c\x40\xb6\xab\x10\x35\x35\xed\xbd\x91\x0f\xdb\x80\xa7\x61\x1e\xbb\xdf\x66\x73\xec\x96\x66\x51\xe5\x99\x19\x0e\xe5\x0f\xe5\x26\x9b\xf7\x4e\x6d\xf8\x73\x40\x0c\xdd\x60\x36\x6e\x17\x69\x22\x20\xd4\x60\x4a\x25\x0d\xaf\xf7\x59\xc4\x89\x3b\x99\x81\x02\x69\xd4\x72\x55\x83\x6e\xd7\x06\x35\x2f\x81\x8d\x95\x25\xe6\x40\xd9\xd8\x9e\x9e\xc0\x15\xf3\xf8\x8f\x1a\x24\xce\x22\xcc\xa7\xc5\xde\xd3\x02\xc6\xa5\x66\xd1\x98\x1b\xab\xe3\x0d\x6e\x16\xb9\xa3\x7d\xcd\xe8\xd5\x6b\xaa\xdf\x90\xd7\x5d\xc7\x6f\x93\x79\xd7\x9b\x1f\x68\xfd\x27\xa6\xf5\xfe\xb4\x66\x72\xb5\xfa\x6a\x42\x84\x0d\x84\xaa\x04\xd9\xa6\xd1\xc7\x5c\xad\x1c\x85\xbb\x91\xe7\x5b\xf0\xe2\xdd\x62\x90\x73\xe6\xe3\xae\xdb\x63\x90\x0a\xe5\x32\xc2\x82\xf0\x96\x33\xf8\x2e\x20\x6c\xf3\x88\xd7\x9c\x81\x0e\x46\xea\xa5\xa7\x71\x6f\x26\x82\x8e\xdd\xfc\x9d\xbb\xdd\xfd\x11\xfb\x36\xb9\x97\x57\xd3\xe5\x3b\x65\x63\xad\x76\xe0\x6c\x12\xcb\x52\x54\x5d\xd7\x38\x63\xac\xbd\xc9\x87\x11\xe8\x3e\x8e\x40\xef\x7e\x3f\x3e\xff\xbc\xd1\x60\xdf\xae\x8f\x0b\x7e\x44\x8a\x95\x39\x2e\x67\x5f\x64\x8f\x4b\x9d\xa5\xc8\x95\x9c\x30\xf7\xff\x4d\x5e\x8e\xb6\xbf\xe1\xdc\xd5\xda\x77\x15\xd4\xbd\xf5\x3f\x4c\x5d\x3f\xf1\xd4\x75\xb7\xad\xb3\x49\x1f\x7f\xe3\x74\x7b\x3f\x62\xdb\x7c\xe0\xfa\xef\xc1\xf5\xd3\xef\xba\xe4\x96\x44\xff\x97\xe6\x4a\x73\xdc\xdc\x80\xec\x11\xa1\x28\xf1\x06\x36\xcf\xa9\x41\x02\x5a\x2b\x7d\x87\x9d\xeb\x9e\xf6\x9a\x36\x57\x5c\x0d\x93\x8c\x32\x18\xbe\x9b\xd8\x59\xcd\xa6\x0d\x97\x0c\xd6\xf3\x20\x7c\xd6\x71\x15\xe3\x54\xa8\x45\x4b\x8f\x39\x67\x0c\xe4\x3c\x40\x5d\xf9\x3e\x97\x35\x04\xd1\x6a\xf8\x0b\xa0\xe5\x90\xa3\x45\xd0\xc8\x59\x34\xec\x3c\x7b\x08\x48\xc3\xda\xed\xb7\xac\x8e\xc2\xb7\x85\x69\x5f\x8f\x1a\xce\x62\xdc\x14\x7c\x6b\xd0\x7b\x8c\x5f\x90\x17\x60\x5e\xcf\xa2\xc6\x8c\xc7\x59\xfe\x72\x18\x56\xdb\x7b\x77\x03\xa9\x6d\xaa\x2f\x27\x6f\xc7\x7f\xbe\x44\xb1\x8d\xef\x74\x99\xd2\xc5\x48\x8a\x78\xbf\x1c\x4e\x36\x13\x9a\x80\x20\x99\xd2\x83\xcb\x6e\x3f\x10\xec\x1a\xb7\x44\xad\x44\xe8\x84\x83\xf8\x7d\x55\x24\xa0\x89\x84\x55\x77\xc0\x57\xde\x6f\x7c\x03\x47\xfd\xaf\x8b\xd2\x19\x08\x48\xc1\xe5\x3c\x78\x16\x0c\x22\x6e\x7d\x1d\xe4\x5f\x17\x52\x4d\x45\x05\x56\x6b\xac\x48\x0e\x80\x6d\x1e\x1f\x82\x76\xa3\x3b\xc8\x94\xc2\xdb\xe7\x58\x82\x92\x24\x28\xc3\xf6\x9f\xc5\xfe\x6c\x8b\xdf\xd8\x4c\x9c\x48\xac\x81\x07\x53\x25\x05\xc7\x03\x0f\xa6\x4a\x53\x30\x66\xc4\xc3\x19\x63\x7e\xfb\x37\x65\x87\xab\x2b\x02\x92\x91\xeb\xeb\x93\x93\xde\xd7\x72\x5b\x78\xee\x53\xb9\x55\xb8\xba\x22\x06\x29\xf2\xf4\xed\xa7\x3f\xcf\xc9\xe3\x66\xfd\xf7\x87\x73\x12\x44\x8c\x9a\x3c\x51\x54\xb3\x88\x1a\x03\x68\xa2\x1a\x24\x53\xda\x44\x36\x58\xc7\x59\xe6\x54\x02\x86\x89\x89\x52\xd3\x3c\xfd\xd4\x3c\x4d\x94\x42\x83\x9a\x96\xa7\x05\x97\xa7\xa9\x3d\xa0\x1b\xb9\x9e\x58\xaf\xfe\xa0\x96\xe6\x0e\x43\x8a\x96\x26\x5a\xfe\x5b\x81\xde\x9c\xf6\xa2\xb2\xb1\x2c\x07\xa1\xdc\x1d\x00\xcb\x89\xf3\xdf\xca\x67\x5b\x33\x9d\xc7\xa5\x89\x0a\x2a\xe9\x02\xb4\x33\x63\x49\x6d\x08\xe2\x7f\x01\x00\x00\xff\xff\xfd\x99\xf7\x92\x06\x21\x00\x00")

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

	info := bindataFileInfo{name: "templates/views/manager.html", size: 8454, mode: os.FileMode(420), modTime: time.Unix(1513024043, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsManagerMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x51\x6f\xdb\xba\x0e\xfe\x2b\x9e\xb6\x9b\x58\xa8\xe3\x6c\xf7\xde\x27\xc7\xf6\x70\xd1\xbd\x0c\x18\xee\x19\x4e\x0b\x9c\x87\xb6\x38\x50\x2c\x36\xd1\x2a\x4b\x81\x44\x37\x2b\x02\xff\xf7\x03\x49\x8e\xe3\xb4\xee\xb6\xae\xeb\xde\x24\x53\x22\x3f\x52\x24\xfd\xf1\x4d\xcc\x75\xd5\xd4\xa0\x90\xa6\x06\x18\xbf\x8b\xaf\x1b\x55\xa1\xd0\x2a\xa6\xbb\x37\x31\x79\xbd\xd5\xe6\x06\x8c\x9d\xd9\xb5\xde\x12\x9a\x56\x52\x54\x37\x8f\x9c\x89\x52\x64\xf6\xc6\x9f\xcc\xd6\xcc\xc6\x22\x5d\xc9\xbb\xcd\x5a\x54\x5a\xcd\xe0\x0e\x66\x7a\x03\x8a\xf6\x3a\x68\x4b\x93\xa1\x81\xb5\xe0\xf0\x4c\x03\x95\xd4\x16\x1e\xb7\xc0\x38\x8f\x96\x0d\xa2\x56\x17\x78\xb7\x81\xc2\x36\xcb\x5a\xe0\xd5\xa8\xd1\x94\x7d\x61\x5f\xe3\x9d\x3b\x97\x91\xcf\x7f\x9c\x9d\x93\xa4\x31\x32\x23\xf3\x4e\xd9\xfc\x3d\xf3\x67\x8b\x81\x72\x92\x70\x86\x2c\xdb\x55\xba\x51\x98\xdd\xb3\x3c\xf3\x5f\x09\x4d\x6f\x99\x8c\x69\x9b\xd8\xa6\xaa\xc0\xda\xac\xd9\x70\x86\xd0\xd2\x96\x2e\x6e\x99\x89\x90\x2d\x25\x7c\x12\x16\x41\x81\xb1\x85\x53\x22\xf7\xbb\x20\x24\x34\xd5\x2a\x26\xdc\xb0\x6d\xca\x91\x24\x3d\x6a\x48\x2c\x20\x0a\xb5\xb2\x74\xb7\x5f\xa5\x5f\xac\x56\x93\xc9\x91\x9a\x1e\x09\xc2\x57\x8c\x8f\x4e\xa6\x06\x2a\x6d\xb8\x3d\xd7\xc8\x24\x6d\x69\xfa\x81\x21\x3b\x77\x56\xe3\x9d\x0b\x48\xb6\x1b\x0f\x82\x45\x86\x76\x02\x0a\x05\xde\x15\xbd\xa1\x10\x8f\x33\x53\x65\xc4\x2d\x48\x9b\x54\x5a\x36\xb5\xb2\xd9\xc5\xce\x47\x8a\xc0\x2d\x28\x24\x6d\xd2\x6d\x15\xab\x81\xb4\x57\x2d\x4d\xbc\xa7\x7f\x05\x2b\xc5\x30\x03\x9e\x15\x82\xfd\x73\xbc\x70\x00\x3a\xe1\xf7\xdd\x17\xfc\xe0\x7b\x65\x80\x21\x70\x92\x18\x50\x1c\x4c\xd6\xfb\xe4\xd2\x83\xee\x0c\x60\x63\x54\xe4\x36\xe7\xfa\x0c\x8d\x50\xab\x20\x69\x7b\x0d\x0e\x43\x63\x7b\x8d\xaa\x91\x32\xe1\x70\xcd\x1a\x89\xa7\x5a\x21\x28\xcc\x88\x93\x6a\xc3\xc1\x38\x97\xb2\x57\xef\x92\xc3\xd1\x4a\x32\x6b\xff\xcf\x6a\xc8\x48\x70\x60\x16\x7c\xb3\xa3\x90\x18\xdd\xb9\x74\xad\x82\xe2\x62\x9a\x73\x71\x1b\x79\x15\x05\x59\xa2\x9a\xad\x8c\x6e\x36\x51\xbf\x9a\x7d\xb5\xa4\x9c\x2e\x0e\x5e\x30\x5f\xca\x93\x49\xdc\x69\x38\x29\xa6\x79\xa8\xcd\xc8\xd7\x26\x09\x1b\x32\xd0\xe9\xb5\x75\x65\xe3\xd7\x95\x30\x95\x84\xa8\xef\x09\xc4\x2b\x9e\xb9\x7d\x41\x2e\xa7\xd1\x49\x24\xa2\x93\xe8\x72\x4a\xca\x5c\xec\xf5\xf4\x0d\x23\x7a\xd8\x9b\x48\x84\x02\x25\x14\xe4\x6c\xad\xb7\x5e\xed\xe5\xd4\x46\x1c\x90\x09\x69\x49\x99\xcf\x45\x99\xcf\x03\xae\x72\x4a\x93\x27\x22\xe7\x4c\xad\xc0\xf8\xa5\xb3\xba\x07\xab\x57\x2b\x67\xb2\xd6\x9c\xc9\xde\x01\xb3\x02\x2c\xc8\xeb\xe1\x47\xbf\x9e\x75\x00\x4f\xb5\xba\x16\xa6\x8e\x6e\x84\x94\x51\x78\xad\xe8\xf5\xf4\xc4\xc7\x55\xf0\x93\xe9\xd1\x9d\x8a\x49\xb9\x64\xd5\x4d\xd1\xbd\xab\xfd\x13\x6a\x7d\x0b\xf1\xe5\x74\x70\xe3\x72\x4a\x17\xdf\x8b\x13\x1a\x66\xd7\x7d\x90\x82\x96\xce\xfa\xbd\xe8\xe4\x73\x2e\x6e\xcb\x69\xdb\x5e\x25\x3e\xdb\xb2\x8b\x8b\x7f\x27\x84\xd9\x8a\x5c\x25\x17\x6f\xbb\xd5\x55\x4b\x17\x0f\x4a\x3b\xc2\xa5\xe6\x77\x5d\x81\xfb\xb6\x4c\x92\x2e\xa2\x87\xe6\x3f\xac\x78\xba\x83\x74\x63\x7c\x23\xf9\x10\x92\x3d\x0e\x9d\x74\x59\xbc\x89\x71\x2d\x2c\x4d\xaf\x85\xe2\x31\x11\x84\x26\x46\x6f\x8b\x61\x6b\x49\x8d\xde\xc6\xfb\x63\xfe\xef\x61\x31\x26\x68\x08\xa5\x0b\x71\x1d\x2f\xd3\x35\xb3\xa7\x2e\x20\x31\x19\x49\x17\x1a\x6a\xc0\xe7\x9b\xd1\xdb\xd4\x05\x33\xa6\x1e\xe6\x62\x99\x1a\x1f\x9f\x6f\xdc\x4e\x19\xe7\xa3\x62\x0f\x24\xa0\x4d\xab\xb5\x90\x3c\x9e\xe6\x21\x38\x5b\xc1\x71\x5d\x90\x77\x6f\xdf\xfe\x8b\x94\x39\x9a\x32\x47\x5e\xe6\x8d\xdc\xbf\x9a\x6b\xbd\xa1\xe0\x48\x99\x4b\xf1\xf0\xf3\x4c\x20\xd4\xa4\xcc\xed\x86\xa9\xbd\x74\xd3\x48\x39\x33\x62\xb5\xc6\xc8\x35\xc4\x59\xdd\x20\xf0\xc8\xd6\x4c\x4a\x52\xe6\x50\x97\xd3\x13\xe7\x92\x4f\x93\x7c\x0e\x75\x99\xcf\xdd\xf5\x32\xb7\x68\xb4\x5a\x95\x1f\x3f\xe4\xf3\x6e\x99\x2f\x4d\x34\x2f\xf3\xb9\x14\xbf\xde\xbc\xfb\x39\x8c\x03\x70\x4d\xeb\xb7\x40\x08\x1d\x76\x1c\xc4\x99\x97\xfd\x16\x18\x1b\x23\xb4\x11\x78\x37\x0e\xe4\x73\x27\xfd\x2d\x50\x18\x22\xd4\x1b\x7c\x24\x26\xff\xeb\xa4\x2f\x0d\x25\xf6\x58\x24\xb3\xf8\x37\x18\xa3\xcd\xfb\x7b\xfb\x8c\x4c\x6a\xce\xec\x7a\x41\xe8\x38\xd0\x4f\xcc\x62\xe4\x8f\xbe\x34\xd4\xa3\x1f\xb8\x87\xd9\xfd\xf8\xe9\x09\x19\x43\x76\x1a\xa4\xa3\xb0\xe6\x8d\x2c\xf3\xb9\xeb\x00\x73\xd7\x0a\xe6\xbe\x47\x94\x84\xa6\xae\x47\xc6\xb4\x05\x69\x21\xfa\x66\x23\xea\x3a\xcd\xa3\x9d\x28\x34\xaa\x43\x23\x4a\x1d\x49\x8f\x8f\xb8\xea\x39\xb3\x37\x81\xa2\x39\x6f\x9e\x49\xd0\xbc\x8a\x97\xa6\x67\xde\xc8\xd3\xc8\x59\x20\xa6\x8f\x10\x2d\xb2\xaf\xc7\xc3\x97\x7d\x59\x7c\x9f\x8c\x3d\x9f\xfc\x3d\x42\xe7\xc6\x39\x5b\xd0\xf4\x43\x6c\xed\x67\x48\x4d\xe0\x62\x21\xe1\x7e\x09\xc1\x09\xaa\xbc\xda\x1f\x21\x38\xfe\x69\x7f\x29\xbd\x71\x1a\x9f\x42\x6e\xfe\x73\x20\x37\x49\x18\xed\x8a\xc1\x58\x79\x3c\xdd\xf9\x19\x33\x35\x20\x35\xe3\xf1\xf1\xcc\x33\x26\xf2\x95\x76\x2c\x68\x13\xd6\xa0\x36\x70\x6d\xc0\xae\x0b\xf7\xee\x9e\x51\x0d\x3e\x8e\x0d\xb7\x8e\xee\xa4\xd5\x1a\xaa\x1b\xe0\xef\xdd\xa5\xa2\x28\x06\x57\x26\x93\x38\x20\x8f\xe9\x91\xfa\xad\x50\x5c\x6f\x53\x0b\xf8\x51\x21\x18\x37\xc7\x86\x73\xc9\x3b\xf8\x2f\xa5\x3e\xed\x5e\xdd\xd7\xd4\x5d\xaa\x24\x30\xd3\x5f\x1b\x1c\xa1\x0f\x3c\x70\x53\x7b\x77\xab\x9f\x24\xc3\x63\x1c\x22\x29\xf8\x53\x46\xf4\xc3\xe4\xdb\x65\x66\x37\xa7\x0b\x9e\x09\x7e\x98\xc4\x1f\x7f\x9c\x96\xb6\x7b\x48\x47\x14\xfa\x67\x01\xed\xe7\xd0\x6f\xc2\xb9\x9f\x36\x4f\xcc\x8c\x01\xe6\x41\x55\xfc\x2c\xe2\xd0\x98\x5f\x18\x6f\x4b\x17\xff\x04\x00\x00\xff\xff\xea\x72\x9e\x86\x8f\x12\x00\x00")

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

	info := bindataFileInfo{name: "assets/js/manager.min.js", size: 4751, mode: os.FileMode(420), modTime: time.Unix(1514899969, 0)}
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
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
