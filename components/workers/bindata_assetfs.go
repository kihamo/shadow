// Code generated by go-bindata.
// sources:
// templates/views/index.html
// public/js/dashboard.js
// DO NOT EDIT!

package workers

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

var _templatesViewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\xdd\x8e\xdb\xb6\x12\xbe\xdf\xa7\x20\x74\x80\x93\xe4\x42\xeb\xec\x39\x29\xd0\x26\xb6\x83\x6d\x02\xb4\x45\x82\xa0\x48\x8a\xe6\x9a\x96\x46\x16\x77\x29\x51\x20\x29\xff\x60\xb1\xef\xde\x21\x45\xc9\xb2\x4d\xfd\xb8\xeb\x26\x9b\x22\xbe\xd8\x95\xc9\xe1\x70\x38\x33\xdf\x70\x66\xac\xbb\x3b\x12\x43\xc2\x72\x20\x41\x24\x72\x0d\xb9\x0e\xc8\xfd\xfd\xc5\x34\x66\x2b\x12\x71\xaa\xd4\x2c\x90\x62\x1d\xcc\x2f\x08\x7e\xda\xa3\x91\xe0\x21\x5f\x86\x57\xff\x73\x73\x76\x3e\xbd\xaa\xa7\x0b\xba\x84\x30\x05\x1a\x83\x6c\x11\x98\xcf\x67\x21\x6f\x41\xaa\xbd\xb1\x36\xe3\x85\xce\xc3\xa5\x14\x65\x41\x8a\x92\xf3\x50\xb2\x65\xaa\x0f\x58\xd8\x25\x8b\x52\x6b\x91\x13\xbd\x2d\x00\x17\xd9\x2f\x41\x8b\x07\x31\x7c\x58\x9e\x08\xfb\xa0\xb2\x80\x28\xbd\xe5\x48\x9a\x51\xb9\x64\x79\xa8\x45\xf1\x92\xfc\x58\x6c\x02\xc2\xe2\x59\x40\x4b\x2d\x24\x24\x12\x54\x1a\x90\x98\x6a\x8a\xf3\xcb\x25\x6f\x71\xa6\x92\xd1\xb0\x40\x02\x05\x48\xaf\x65\x09\x38\x86\xab\x22\x91\x15\x1c\x34\x52\x8a\x24\xf1\xc8\x69\x3e\xd7\x3b\xee\xc7\xe7\x98\x54\x3b\xec\xaf\x9c\x4e\x50\x23\x2d\xc5\x4e\xd2\x2b\x67\x82\x6a\xc2\xfd\x1b\x69\xa7\x8d\x3a\xb0\xd3\xf1\xfc\x8b\x03\xd1\xdb\x24\x05\xcd\x81\x13\xfb\x17\x2d\x03\x90\xfb\xcc\x71\x48\x6f\x8d\xcf\xf2\x65\x87\x4a\xfc\x82\x0f\x51\x3a\x61\xff\xdf\x43\x6e\x97\xb0\x7a\x41\x42\x49\x42\xc3\x18\x68\x62\xfe\xff\xb0\x09\xe6\xd3\x09\xeb\xd9\x6b\x5f\xed\x23\x44\xf9\x89\x68\xd8\xe8\x4e\x2f\xed\x5a\x9d\x96\x4b\xa8\x3c\x8f\x33\x85\xa8\x43\x40\x84\x91\x28\x11\x7d\xf3\xe7\x03\x42\xd4\xac\xe6\xef\xeb\x95\x43\x52\x77\x4f\x77\x4c\x75\x0d\x53\x92\xa2\x1b\xcf\x82\xff\x34\x42\x8f\xb0\x6e\xe5\x0d\x89\x10\xfa\x28\x14\xec\x2d\x51\x48\xd8\xac\x31\xc8\xe7\x90\xa0\x3a\xfe\x64\xb0\x26\x6f\x41\x53\xc6\xf1\xa0\x86\xe8\x14\x1e\xce\x2e\x87\x0e\x41\x25\xfa\x5c\x18\x31\x19\x71\x68\x68\xd0\x31\x06\x37\x68\x3b\x00\x07\x2a\x13\x66\x5d\xea\x74\x05\xd3\x01\xbc\xef\x7f\xfd\xfb\x78\xdd\x02\xe7\x5e\x78\x3d\x6a\xc0\x96\x0a\x7d\xeb\xb1\x22\x76\x5d\x5d\x60\x27\xe3\xd5\x5d\x7c\x5f\x1e\xad\x4e\xe0\xef\x58\x7d\xe4\x58\x95\x10\x7f\x6b\x40\xd5\x54\xdd\x3e\x5a\xa0\x5a\xe1\xc2\x35\x65\xfa\x74\xac\xe2\x22\x62\xd7\x7f\x79\xb8\xda\x6d\xbf\x83\xf5\x9f\x06\x6b\x4f\x26\xdd\x06\xa5\x2a\xa3\x08\x93\xfe\x83\x64\xcd\x93\x68\x77\x61\xb3\x49\xd3\xbc\x21\x62\x44\xb5\x33\xaa\xd2\xc1\x22\x8e\x96\x5c\xdb\xe7\x8d\x22\xb1\x14\x45\x2c\xd6\xb9\x2b\x63\x0e\x6a\x9a\x7a\xd6\x55\x35\xb0\x41\xd9\x63\x53\xd6\x24\x94\x2b\xf0\xf8\xd2\x75\xa4\x99\xc8\xd5\xb1\xba\xdb\x7e\x11\x51\x09\xd6\xdc\xc7\xa6\xee\x28\x72\x4a\x5e\xaf\x6d\xe4\xcd\x20\x2f\x89\xe2\x0c\xc5\x31\x02\x7a\x2c\xcc\x3b\x82\x4c\x03\xa0\x1b\xba\xa2\x2a\x92\xac\xd0\x2f\x57\x82\xc5\x4f\x9f\x3f\x3b\x38\x7e\x26\x62\xca\xeb\x31\xac\x06\x41\x23\xea\xda\x83\xf6\x39\xd4\x4c\x1b\xea\x37\x22\x4f\x98\xcc\x88\x84\x4c\xac\x80\x50\xce\xc9\xce\x0d\xda\xf4\x0b\x11\x6f\x67\xc1\x67\x86\x04\x0b\xf0\x92\x13\xd8\x44\x50\x68\x52\x1b\xab\x9e\xb8\x24\xd7\x12\xc8\x56\xe0\xc9\x4b\x09\xaf\xf7\xb8\x46\xc8\x61\x41\xa3\xdb\x96\xf3\x7d\xb4\xac\x9f\x3e\x7b\xd5\x87\xfa\x06\x98\x4b\xbe\x2d\x52\x86\xc5\x3d\x69\x9e\x42\x2d\x29\x16\xba\x16\x99\xe4\x63\x23\x68\x07\xd0\xa8\x0f\x66\x87\x56\x98\x4e\x4a\xde\x07\x33\xfb\x78\x84\x16\xa3\xb2\x8e\xba\x54\xd3\x85\x09\x21\xa0\x0a\x74\x3c\xb6\x3a\xf4\xca\xa9\x9d\xdf\x23\x26\xd5\x92\x14\x4f\xe3\x8b\x87\x53\x6d\xb0\xe9\x1b\x97\x1d\x0e\xa5\xd3\xf9\x07\x9a\xc1\x74\x82\x0f\x9d\x14\x6f\x24\x50\x0d\xf1\x00\x91\xb9\x7e\x88\x0b\x25\xfd\xa4\xef\xa9\x1a\x49\x59\x31\x4d\x30\x7e\x0f\x6d\x6f\x79\x8e\x21\x74\x38\xf7\x13\xe1\xa8\x47\x53\x86\xb6\x43\xaf\xc6\xba\xde\x05\xc7\x13\x38\x68\x8c\xf7\x90\x40\xbd\xa6\x32\x37\x91\xb7\x9d\xa3\x9f\x10\xa6\x0f\xdb\x52\xdf\x83\xb4\x1b\xfd\x72\x41\xba\x5d\x5c\xa9\xb4\x3f\x75\xed\x0d\x6f\xb0\x85\x50\x14\xa6\x4f\x65\x23\xdc\x27\xe4\x65\x03\xb1\xcd\xac\x1e\x12\xe5\xce\x73\xb6\x14\x15\xf7\xa0\xb3\x45\x5c\x18\x0f\xb0\x87\xfb\x15\x99\x9d\xf5\x70\x8d\xb5\xd9\x8a\xd9\xee\xed\x39\xd5\x70\x96\x7b\x58\x81\xb6\x27\xae\x41\xee\xbf\x2e\xdd\xec\x47\x43\xfe\x90\xbb\xb2\x6e\x0b\xd7\xb7\xa5\xdb\xfd\xeb\xb8\xd1\x19\xf4\x77\x6b\x12\x14\x0c\x25\xed\x0c\x65\x94\x26\xdf\xe1\xc2\x73\x25\x1d\xef\xac\x10\x67\xd1\xe2\x57\xf6\xd7\x1a\xd4\x34\xf6\x15\xef\xe3\x34\x53\xf0\x52\x39\xc5\x5c\xc7\xb1\xb3\xc6\x37\x9e\x8d\x9d\x2b\x03\xfb\xed\xed\x19\xf2\xaf\x4f\x9a\xea\x72\x20\x9d\xfa\x03\xe3\xe7\xbf\x37\x3b\x8a\x69\xbe\x44\x70\xec\xfa\x22\xc3\xa9\x91\xb7\x07\xf2\x0d\xe6\xf0\x43\x1e\x34\x9c\xe3\x8f\x71\x9f\xdf\x25\x13\x92\xe9\xed\x80\x0b\x69\x0d\x59\xa1\xc7\xd4\x00\x20\xa5\x90\x0f\xf6\xfd\xb8\xe5\xb5\x3e\xc5\x3d\x3e\xaf\xb5\x51\x16\xab\x95\xf8\xa0\xdd\x6e\x02\xac\x71\x0f\x86\x59\xf0\x66\x16\x84\x57\x01\x91\xc2\x26\xca\x8c\x72\xb1\x74\x69\x32\xe6\x56\x31\xe4\xee\x07\xda\x63\x97\xad\xae\x36\xb7\xc2\xef\xb4\xee\xf6\x73\x3f\x86\x77\x77\x6e\x2b\x3a\xef\x4f\xdc\x96\xb4\xaf\x20\xa8\x12\xb8\xea\x26\x89\x99\xca\x58\xc3\xd0\x7b\x8c\xff\x6a\x96\x81\x7a\xe5\x4f\xd0\xed\x66\xe9\x8b\x7d\xb1\xec\x9d\x1f\xcc\x77\xb7\x09\xda\x3f\x7d\xd1\xdb\x1a\xf3\x9f\xef\x00\xdd\x0d\x61\x22\x64\x36\xdc\xa2\x34\x54\x55\xe5\xd4\x77\x37\x72\xba\xc0\x38\x85\xb4\x7b\xc6\x76\x3d\xdb\x5d\x53\x38\xd7\xd2\xbc\x73\x60\x88\x83\xf9\x87\x32\x5b\x80\x24\x39\xac\xeb\x03\xbe\xc4\x6b\xd0\xcc\xf5\x5d\xc2\x79\x51\x6a\x67\x91\xdc\x32\x08\x48\xc6\x50\xcd\x57\xc1\x9e\xc4\x6e\xaf\x23\xff\xab\xdb\xc8\xa7\x75\x24\x8f\x55\x35\x4a\xf3\x9d\x8d\xdd\x53\x4a\x4d\xbf\x8f\xcd\xdf\x18\xff\xeb\x71\xa7\xbd\x1d\x54\xb9\xc8\x98\x3e\xda\xa1\x69\x8c\x7a\x77\x40\xcf\x1b\xf9\x66\x43\x47\x4c\xb8\xbb\x23\x80\x49\xea\xfd\xfd\xc5\x45\xeb\x1d\x95\x1b\x55\xbd\x9e\x52\xe5\x69\x4e\x40\x5a\x14\x9c\x45\xd4\x84\xb9\xc9\x2e\x89\x73\x9a\x5b\x51\x59\xf7\xdb\xea\x4e\xac\x09\xf8\x64\x46\x9e\x20\xdf\x4b\xdf\xd4\xfd\xfd\x93\x57\x28\x46\xc5\x66\xde\x6c\xa6\x64\x34\x0b\x26\xce\x1d\x26\x37\x6a\x12\x63\x3e\xbb\x10\x54\xc6\x97\x37\xea\xf5\x6a\x66\xd8\x5d\xef\x44\xb9\xfc\x05\xf4\xcf\x25\xe3\xe6\x0c\xb6\xe0\x76\xec\x9a\x83\xfd\x15\x00\x00\xff\xff\x84\x00\x95\x50\x7b\x23\x00\x00")

func templatesViewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsIndexHtml,
		"templates/views/index.html",
	)
}

func templatesViewsIndexHtml() (*asset, error) {
	bytes, err := templatesViewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/index.html", size: 9083, mode: os.FileMode(420), modTime: time.Unix(1500372034, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicJsDashboardJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5a\x5f\x6f\xdb\x38\x12\x7f\xcf\xa7\xe0\x39\xc5\x49\xbe\xd6\x72\x7a\x77\xbd\x07\x3b\x4e\x51\xa4\x87\x43\x71\xc1\x6e\xd1\x06\xe8\x43\x5b\x04\xb4\x44\xdb\xdc\xc8\x92\x41\x51\x71\x8d\x36\xdf\x7d\x39\x94\x28\x91\x12\x25\x33\x5e\x67\x1b\xac\x80\x22\x32\x35\x33\xfc\x0d\x39\xf3\x1b\xfe\xe9\x22\x4f\x42\x4e\xd3\x04\x2d\x09\xff\x94\xb2\x5b\xc2\x3e\x72\xcc\xf3\xec\x17\xbc\x26\x7e\x26\x5f\x87\xe8\xfb\x09\x12\x4f\xb6\xa5\x3c\x5c\x35\x1a\xe1\x09\x71\x46\xd0\xd9\xa4\xfa\x0d\x0f\x23\x3c\x67\x09\xf2\xb6\x98\x72\x6f\x6a\x7c\x9a\x33\x82\x6f\xa7\xa6\xf6\x4b\xbb\xf6\x86\xa5\x21\xc9\x32\x07\x03\xff\xb4\x1b\x98\xe7\xd9\xae\x5f\x3b\x22\x0b\x9c\xc7\xdc\xae\x9e\x27\xb7\x49\xba\x4d\xba\x2d\xdc\x9f\xdc\x9f\x9c\xe8\x63\x78\x8d\xb3\xdb\xbf\xd4\x08\x66\x79\xe8\x68\xe0\x5f\x76\x03\x0b\x4c\x63\x07\xed\x7f\x77\x6b\xa3\xf9\x0e\x71\xba\x26\x69\xee\x32\x10\xaf\xec\x86\x6e\x69\xec\x02\xe3\x3f\xdd\x93\x20\x7e\x6c\x08\xde\x03\xe1\x98\xe1\x94\x6f\x22\xcc\x89\xaf\x02\xe5\x59\x80\x7f\xc3\xdf\xfc\x3a\x6a\xf8\x6e\x43\x26\xc8\xfb\xdf\x7f\xaf\xbd\x17\x55\x63\xce\x62\xd1\x36\xde\xca\x5c\xce\xc6\xa0\x32\x7e\x8d\xa5\xc1\x19\x04\x5e\xa6\xc9\x96\x73\x3b\x41\xaa\x4b\x9f\x0d\xbf\x1b\xd8\xee\x30\x43\xfc\x0a\xcd\xd0\x33\xdf\x3b\x8d\x69\xc6\x49\x22\xac\x22\x3e\x4f\xa3\x9d\x37\x0c\xc8\x7a\xc3\x77\xfe\x70\xda\xd6\xf9\x54\xea\x94\x38\x1c\x34\xae\x4b\x0d\x2e\x32\xc8\x22\xdf\x52\xa8\xd1\xcc\x10\x0b\xaa\x5f\x37\x61\x9a\x27\x1c\xfd\xf8\x81\xce\xda\x9d\x28\x34\xa0\x51\xbe\xf7\xca\x4b\x2c\x37\x72\xf2\x67\xf0\xdd\x10\xa0\x0b\xe4\xbf\x61\x0c\xef\x02\x9a\xc9\xbf\xbe\x06\x63\xa8\xa7\xb7\x7a\x16\x29\x43\x3e\xd8\xa5\x88\x26\x3a\x68\x9b\x70\xd3\x4f\xd3\xcd\xcf\xf4\x6b\x03\x8e\xa1\x73\x29\xa4\xbd\x73\xce\x2e\x3c\xab\xcc\x73\xf8\x18\x5d\x78\xe2\x45\x59\x0c\x12\xc1\x57\xf2\xc3\x18\xbe\xec\x53\x83\xc8\xbc\x4e\x3f\x72\x46\x93\xa5\x5f\xd9\x08\x45\x28\x73\x12\xdd\x60\x3e\x74\x37\x55\x6b\xc3\x4c\xdc\xc0\x98\xdf\x94\x91\xe9\x6e\xa4\xc6\x10\xe3\xcc\x34\x22\xd0\xa0\xd7\x1d\x80\x6d\xc2\x43\x24\x12\xc8\xd3\x1c\x40\x87\x78\x00\xc4\x45\xa2\x3f\xe6\x40\x61\xc3\x0d\x7f\x25\xdb\x82\x6f\xed\x79\x6a\x6d\xb5\x36\x42\x9c\x9b\x51\xf2\xb7\x99\xe2\xb9\xab\xb2\x1d\x8a\x5d\x57\x10\xc3\x23\x02\xf2\xf9\xac\x70\xf6\x3c\xa2\x77\x28\x14\xc0\xb3\xd9\x60\xce\x93\xd1\x92\xa5\xf9\x06\x55\x6f\xa3\x6c\x3d\xe8\xc0\xac\x8f\xda\x3c\xe7\x5c\x50\x24\x50\xa0\x30\x23\x7f\x0c\x34\xab\xd2\x5e\x84\x93\xa5\xc8\x1b\x78\xa5\xa1\x10\x56\x4e\x8c\x18\x59\xa7\x77\x64\x00\x63\x8a\x47\x3c\x5d\x2e\x63\x61\x64\x9d\x46\x38\x56\x6d\x98\x89\x5a\x3e\x1b\x9c\xea\x8d\xf2\x7d\xc4\x29\x07\xe9\xcb\x34\x59\x50\xb6\x46\x85\xa9\x3a\x49\xad\x29\x65\x18\x08\x71\x1c\xcf\x71\x78\x3b\x1b\x54\xc9\xfc\x41\x1a\xf1\xbf\x78\x56\xed\x2f\xde\x70\x3a\xb8\x38\xa7\xca\xbd\x65\xbc\xdb\xac\xa4\x43\xd5\xdb\x88\x33\x9c\xad\x06\xa8\x04\x57\xd8\x13\x3a\x63\x2a\xfe\x15\xa3\xe3\x30\xa6\x63\x31\x33\x17\x3d\x31\x73\x8f\x48\x2c\x4a\xa4\xdb\x2c\x4b\x33\xf6\x28\xbb\x77\x8f\x3d\x65\x71\x0c\x54\xf6\x80\x98\xe5\x57\x01\xde\x6c\x48\x12\xf9\xf1\xe5\xb0\xad\xd7\x46\xd0\x51\x4d\x82\x98\x24\x4b\xbe\x3a\x31\x95\xf7\x96\x81\xb2\xb6\x38\x15\x01\x25\xdb\x53\x02\x0a\x11\xbd\x6a\x49\xfa\xef\x12\x07\x3e\x10\xc2\x85\x68\x00\xbf\xba\x65\xb7\xa2\x52\xf4\x14\x0a\x78\x34\x92\x2a\x4d\xd2\x7d\xbc\xd6\x50\x33\xb8\xab\xb4\x51\x96\x8a\x7d\x3c\xd5\xb0\x64\xdb\xa5\x94\x06\xd5\xaa\xba\xb2\xd7\x51\x1f\x61\xbe\x60\x4c\xfa\xe8\x6a\xab\x05\xb2\x0b\xd3\xa8\x72\x05\xef\x21\x65\x61\x4c\xe4\x24\x8c\xb2\x55\xba\xad\x28\x25\x13\x29\x0f\x3e\x50\x49\x09\x7b\x12\x9a\xec\xc8\x28\x15\x01\xdc\x48\xe2\xde\xb4\xda\x97\x9e\x5b\xd7\xf4\xb4\x36\x6f\x8f\x45\xe1\xae\xf4\x4d\x93\x45\x5a\x93\x77\x31\xcb\x82\xba\x33\xc2\x8f\xc4\xdc\xc2\x92\x4a\xad\xd3\x56\x74\x77\x70\x76\x99\x7f\x1f\x40\xb9\x24\x6c\x43\xcd\x85\xac\x19\x59\x88\xce\x75\xba\x06\x9f\x5c\xd9\xfa\xd0\xea\x57\x0e\x20\xec\x83\x8e\x32\x7e\x60\xe8\xd0\xe1\xfb\xbf\xd0\x3d\x74\xf4\xcc\x52\x07\x96\x1e\x34\x74\xfb\x8a\x5c\x25\x57\x54\x1c\x7b\x75\xf9\xa4\xaa\xcb\xf6\xb2\xb9\x35\x51\x8f\x0b\xcd\xd4\x76\x7a\x6b\x33\xf0\x33\xca\xf8\x0e\xfc\x8d\x68\xb6\x89\xf1\x6e\x92\xa4\x89\x58\xc0\xd0\x68\x36\x90\x44\xa3\xf1\xca\xfe\x3a\xcf\x23\x14\xa6\x71\xb6\xc1\xc9\x6c\xf0\xca\x45\x21\x8f\xd5\x84\x40\x61\x2c\x92\xdc\x45\x2f\xa6\x6d\xbd\x11\xe5\x44\xd0\xc3\x39\x74\xaf\xbe\x6e\xf2\x38\x1e\x31\xba\x5c\x71\xc4\xc9\x37\x3e\x5a\xe7\xa2\x32\xa0\x6c\x8d\xe5\xdc\x92\xb5\xa4\x7e\xf0\xb3\xaa\x3c\xa2\xed\x7c\x0c\x26\x84\x21\xce\xd2\x64\x79\xf1\xee\xad\x68\x28\x5e\xcf\xe7\x0c\x8d\xc5\xf7\x98\xfe\x14\x8c\xf5\x46\xaa\x8d\x12\x0a\xd6\x13\xc0\xd9\x3e\xa8\x92\xc8\x8d\x12\xda\x06\x5f\x88\x3f\x01\xf8\x12\xec\x86\xd1\x94\x51\xbe\xeb\x42\xfb\xbe\xfc\xfe\x54\xf0\x62\xce\xe1\x30\x23\xeb\xc2\xfb\xa6\xfc\xfe\x04\xf0\x16\xd1\x20\xf7\x96\x84\x31\xb1\x5e\x7d\x8d\x9a\x2d\x62\x8f\xf9\xf7\x75\x24\x98\x78\xea\x75\xc6\xcb\x95\x10\x47\x52\xfc\x09\xf8\x94\x90\x2d\x7a\x0b\xc7\x69\xd2\x13\xb5\xfa\x0c\x78\x7a\x95\x8a\xf2\x44\xca\xc5\x69\xa7\x2f\x97\x85\xc2\x61\x8e\x8c\xf3\xd8\x45\xaa\xb7\x26\xd5\x52\x7d\x6b\xf5\xa1\xe3\x86\xab\xdd\x62\x39\x1a\x3b\x70\xf3\x53\x1f\x9a\x39\xed\x7f\x34\xf1\x9e\x2d\x50\xb9\xa7\xd1\xc5\x3b\xf7\x40\xfc\x5a\x95\x55\xf7\x9d\x8d\x51\x5d\x5c\x77\x23\x0d\xba\x7f\xc0\x26\x66\x2f\xfb\x3e\x08\x82\x49\x85\x0f\x52\x35\x59\xc9\x55\xf5\x20\x82\x38\x68\xb7\x68\x64\xab\xbb\x9d\x9f\xb1\x43\x91\x0b\xb1\xe3\x1f\x2d\xc9\xc8\x3f\x6d\xc4\x68\xc7\xf2\x1a\x24\x8c\xc3\x24\x4d\xe5\xd1\x8f\x91\x1e\xbc\xba\x76\x3a\x98\x31\x8e\xe0\xf5\xf4\x77\x60\x27\xe3\xd2\x62\x24\x0f\x67\x3d\x41\xf9\xa2\x46\xf8\xf5\xd1\xfb\xb4\xa5\x52\xd2\x9f\xa9\xa0\x8e\x69\xda\xe2\x12\x52\x31\x79\x4a\x58\x03\x2d\xbc\x45\xf2\x45\xb6\x35\x7d\xae\xf5\x41\xc6\xec\x51\x23\xc6\x69\xdb\x2f\xc5\xd7\x41\x75\xd0\x20\xf4\xc2\x98\x86\xb7\x7e\x75\x81\xe4\xdb\x18\x15\xd8\x94\xc8\xbb\x16\xbe\xa2\x4d\x7f\x94\xc0\x5c\x08\x90\x60\x41\x81\x43\xa9\x6d\xa2\x0a\x02\x07\x21\x08\x44\xdf\x03\x14\x36\xb9\x56\x03\x54\x8c\x79\xb0\xc2\xd9\x25\x84\xa1\xef\xb5\x8f\x3c\x3c\x6b\xd9\x80\x67\x1e\x14\x39\xd1\xa3\x6a\xaf\x08\xf3\x00\x47\x91\x55\x2b\x8c\xd3\x8c\x74\xa9\xa9\xd9\x29\xb6\x57\xc3\x00\x46\xd9\xf7\x16\x82\xe4\xac\xb1\xdb\x77\xfe\xd2\x0b\xbd\x17\x44\x37\xf6\x3e\x8f\x9b\xd0\x57\x34\x22\x3d\xd0\xcd\x2c\xd2\x24\x8a\x2f\xd0\xa2\xdf\x4c\x36\x4f\xae\x69\xb4\xef\x92\xf2\xfd\xaf\x1f\x5d\x6f\x29\xeb\x8c\x2d\x86\x4c\x53\x83\x58\x9b\x34\x46\x98\x46\x13\xf1\xaf\x06\x6c\xb9\xde\x2c\x2e\x52\xad\x9e\x68\x8c\x79\x54\x27\xb4\x62\xf0\xa8\xf8\xf5\x03\x95\xa3\x3a\xa0\x48\x50\xde\x9b\xff\x09\x1e\x14\x27\x6a\x8f\xe2\x82\x3c\xe9\x7b\x0c\x1f\x9e\x55\x54\x5b\x5f\xd2\x6b\xf5\xc3\x81\x94\xed\x4c\x3e\x11\xf4\xe8\xd3\xa0\x9d\xed\xc3\xca\x5c\x99\xa2\xf7\xaa\x34\xe8\xfd\x42\xae\x1f\xaf\x5f\x49\x4e\x6e\x1d\x0b\xa2\x42\xc5\x32\xe1\xb3\x5c\x37\x65\xf9\x7c\x4d\xf9\xd7\x7d\x58\x9a\x33\xdd\x3d\xdb\xae\x33\x2e\x80\x34\xb4\x6c\x73\x0e\x8f\xac\xb8\x93\xa6\x17\x55\x21\xbe\xc3\xb1\x6f\xee\xa9\xee\x4d\xbb\xb6\xe8\x50\xe3\x63\x8c\x93\xfa\xcf\x1c\x45\x33\x94\x4f\x9c\xf3\xb4\x3c\x0a\x86\xdb\x1a\x51\x6b\xd3\x6d\x20\x22\xf5\x5d\xc2\x09\x83\x8e\x0b\x95\x17\xe8\xe5\xd9\xd9\x19\xfa\x87\xf8\x53\x2a\x03\x58\x4d\xb9\x35\xbc\xc6\xe8\x42\xbd\x35\x7a\x9a\xcd\x50\x22\x36\xcc\xcd\x1a\x6b\xe2\x53\xcf\xc1\x18\xa5\xef\xb6\x82\x58\xda\x08\x63\x82\x59\x65\x45\xeb\xa6\x17\x00\x00\xb7\xd6\xa6\xe1\xf4\xf7\x00\x00\x00\xff\xff\x51\x39\x0d\x81\xc9\x26\x00\x00")

func publicJsDashboardJsBytes() ([]byte, error) {
	return bindataRead(
		_publicJsDashboardJs,
		"public/js/dashboard.js",
	)
}

func publicJsDashboardJs() (*asset, error) {
	bytes, err := publicJsDashboardJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/js/dashboard.js", size: 9929, mode: os.FileMode(420), modTime: time.Unix(1500405762, 0)}
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
	"templates/views/index.html": templatesViewsIndexHtml,
	"public/js/dashboard.js": publicJsDashboardJs,
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
	"public": &bintree{nil, map[string]*bintree{
		"js": &bintree{nil, map[string]*bintree{
			"dashboard.js": &bintree{publicJsDashboardJs, map[string]*bintree{}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{templatesViewsIndexHtml, map[string]*bintree{}},
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
