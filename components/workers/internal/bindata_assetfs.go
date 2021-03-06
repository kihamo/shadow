// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/manager.mo
// locales/ru/LC_MESSAGES/workers.mo
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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _templatesViewsManagerHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\x20\x58\x6c\x68\x1f\x14\x25\x6d\x07\x0c\xad\xed\x21\xe8\x36\xb4\x40\x56\x0c\x69\x86\x62\x4f\x01\x25\x9e\x2c\x26\x14\xa9\x91\x27\xd9\x86\x91\xef\x3e\xe8\x9f\x2d\xc9\x94\x1b\x6b\x4d\x91\x74\x79\x8a\x4c\xf2\xfe\xfd\xee\x78\xf7\xb3\xa3\xf5\x9a\x70\x88\x84\x02\x42\x43\xad\x10\x14\x52\x72\x7b\x7b\x34\xe1\x22\x27\xa1\x64\xd6\x4e\x69\xca\xe6\xe0\xa1\x40\x09\x74\x76\x44\x08\x21\xed\xcd\x72\xfd\x4a\x42\x84\xf5\x66\x79\x20\x7e\x35\x5b\xaf\x89\x38\xfd\x59\x11\xfa\x59\x9b\x1b\x30\x96\x92\x63\x72\x7b\x3b\xf1\xe3\x57\xb5\x12\x9f\x8b\x7c\x76\x34\xa0\xd0\x88\x79\xdc\xd1\xd8\x3a\x11\x69\x93\x78\x73\xa3\xb3\x94\xa4\x99\x94\x5e\xff\x6c\xff\xbc\x50\x69\x86\x95\x40\xef\x54\x79\x52\xb2\x00\xe4\xee\x7a\xb9\x57\x8a\x12\x5c\xa5\x30\xa5\x61\x0c\xe1\x4d\xa0\x97\xb4\xd1\x7b\x6d\x3d\xbb\x10\x18\xc6\x94\x08\x3e\xa5\x2c\x43\x6d\x20\x32\x60\x63\x4a\xfc\x19\xd9\x00\x70\xd6\xde\x28\x40\xd8\xf5\xc1\x77\x38\x51\x03\xe4\xf8\x58\x3f\x36\x08\xb6\x63\x0d\x25\x30\x13\x89\x25\x9d\xb9\x76\x8d\x5e\x10\xd4\xe9\x15\x0a\x09\xd6\x91\x4c\xa6\x44\xc2\x10\x38\x89\xa4\x48\x3f\xa8\xbf\x49\xa8\xa5\x27\xe7\xde\xeb\xf2\x21\xe1\xf5\x83\x4d\xea\x87\xa5\xf5\x4e\x5f\x0e\x64\xa9\xb0\xe1\x59\x64\x68\xf7\xa5\x26\xd4\x8a\xce\x26\x62\x93\x5a\x46\x22\xe6\x71\x60\x51\xf1\xf7\xa7\x32\x0c\x31\xeb\x21\xd1\x57\x12\xea\xac\xa8\xda\x22\x07\x52\x58\x04\x05\xc6\x7a\xd5\xe2\xec\xc4\x25\xdb\xae\xce\xf3\x46\xa2\x48\x0d\xad\xfc\xed\xd4\xa9\x1b\xfb\x07\x8a\x5c\x66\xc1\xd8\x71\xd0\x2d\xaa\x6b\x7a\x67\xe0\x5a\xd7\xfa\xd1\xc3\x86\xcc\xde\x8c\x84\xad\x14\xbd\x33\x68\x97\xc5\xe9\xc3\x20\xeb\xdc\xe3\x4e\x8d\x3b\x2e\xf0\xf2\x2a\x65\x0a\xe4\x00\x42\xcb\xab\x76\x17\xdf\x3a\xf8\xd2\x79\x03\x2b\x24\x88\xfb\xa6\x94\x7e\xbf\xec\x29\xca\x64\xa3\x45\xb1\x9c\x28\x96\x07\xcc\x54\xdd\x99\x94\x6e\x5d\xa1\xd6\xb2\xe8\xa0\xae\x2e\x2c\x66\x13\xb6\xc5\x57\x4a\x96\x5a\xf0\xa4\x50\x37\xbb\xe9\x0a\x63\xc8\x8d\x56\x5e\xd1\xcf\xab\x6c\xb1\xd9\xc4\x97\xc2\xa9\xb6\x11\xe5\x46\xa7\x5c\x2f\x94\xc3\x78\x79\x92\x91\xd8\x40\x34\xa5\xcf\x68\x5f\xc2\x43\x3d\x9f\x4b\xa0\x84\x33\x64\xf5\x87\x96\x3e\x62\x74\xf1\x39\xc8\x10\xb5\xa2\x84\x19\xc1\x3c\x58\xa6\x4c\x71\xe0\x85\xc7\xd2\xc2\x6e\x08\x0b\x03\x2a\x8c\xb7\xee\xbb\x7d\xda\x02\xba\xf1\x25\x01\x95\x35\x26\xcb\x67\xb7\x68\x83\xe9\xe0\x66\x27\xe6\x6b\x96\x33\x1b\x1a\x91\xe2\x9b\x5c\x0b\xfe\xfc\xe4\x45\x2f\xd8\x44\x73\x26\x9b\x35\x66\xe6\x80\x53\xfa\xac\xbd\x58\x3e\x57\x24\x61\x4a\x37\x25\xf3\x4e\xab\x48\x98\x84\x18\x48\x74\x0e\x84\x49\x49\x64\xb7\x8a\x3a\xe2\x81\xe6\xab\x96\xf4\x67\x21\x25\x09\xc0\x29\x4d\x60\x19\x42\x8a\x05\x77\x61\x99\xc4\xcd\xc6\x31\x39\x33\x40\x56\x3a\x23\x36\x33\xf0\x8b\xc3\x48\xc8\xa4\x0c\x58\x78\xd3\xba\x48\x17\xa5\x81\xe7\x2f\xde\xee\x01\x73\x83\x59\xbf\x75\x18\x66\x63\x8f\x49\xec\x5f\x97\x8b\x8d\xd7\x03\x43\xbf\xa3\x75\xa8\x04\xaa\xcd\xa1\x44\x4e\xfc\xcc\xc1\x5c\x76\xcf\xef\x9e\xdb\xcb\x19\xb6\x62\xdd\x8f\x9d\x56\xd2\xb0\xc5\x61\xbd\xc8\x02\x09\x9e\x01\x9b\x6a\x65\x45\xde\xef\x3b\xe5\xf1\xf2\x4c\x47\x80\x54\x62\xb1\xce\xc1\xd4\xcf\x16\x8d\x48\x81\x13\x8e\x2d\x6d\x44\xe9\x85\x61\x29\x25\x16\x57\x45\xd1\x2d\x04\xc7\xf8\xcd\xe9\xc9\xc9\x0f\x43\x37\x1c\x63\x60\x7c\x68\xcf\xec\xc1\x1f\xe3\x6d\x07\xff\xf0\x6b\xd3\xfe\x30\xbe\xa3\xc8\x47\x96\xc0\xc1\x42\xbf\xe5\xa0\xd0\x1e\x2c\xf6\xbb\x30\x30\x4a\xca\x22\x89\x84\x01\x7e\xb0\xec\x39\x1b\x2d\x7a\x16\xa2\xd0\xea\x8b\xee\x4e\xfc\xa1\xec\x14\x32\xce\x9c\x4e\xfc\xb2\x72\xfe\x0b\xa1\x6e\x71\xa2\x7b\x1c\xb4\x25\x61\xeb\xb7\x8e\xfe\x37\xa6\xa7\x39\xfb\x1d\xcf\xd9\x36\xf7\xb6\xb1\x5e\x8c\x98\x41\xb0\x82\x7e\x09\x7d\x8a\xf5\xa2\x9c\x98\xd8\x30\xce\xfb\x99\x40\x5f\x37\xfe\x58\x70\xd7\x94\xb8\x43\xfc\x9e\x95\xcc\xc6\x7d\x14\xde\x0b\x0e\xdf\x10\x85\x4d\x11\x89\x5c\x70\x30\x74\xe0\xb2\xb4\x24\x1e\x1e\x3f\x5b\xb4\x7b\xcf\x00\x71\xaa\xcf\x3c\x5a\xda\x44\x1e\x64\xba\x9a\x4b\xc0\x38\x1f\x01\x68\x2a\xb3\x9d\x39\x72\xc6\x79\x9d\xcf\x27\x0a\xfa\xff\xa6\xa0\xef\x0c\x30\x1c\xc1\xd0\x3e\x21\xc3\xec\x70\x3e\x79\xae\xc3\x9b\x11\xd6\x2e\x99\xbd\xf9\xfe\x48\x64\x35\x7b\xee\x8f\x42\xd6\xfa\xbb\x57\xff\x72\x3b\xf0\x9e\x08\xe4\xf7\x4d\x20\xef\x89\x08\xb4\x28\xd3\x00\x0d\x28\x4f\x3c\x5a\x12\xf0\x34\xb8\x1e\xc3\xe0\x1a\xf5\xdb\xc9\x9f\x46\x68\x23\x70\x75\xb0\xe0\x05\xa4\xc0\x46\xfc\xea\x52\xc9\x11\xa1\x10\x4c\xce\xe4\xe1\x83\x4f\x24\xa0\x33\xfc\x96\x63\xdd\x3c\x7c\x3a\x70\x86\x08\x49\x3a\x22\x1d\x67\x52\xea\x05\xb1\x45\x90\x23\x7f\x0a\xb3\x23\x01\x2a\x7f\x0c\x1b\x2b\xfc\xa0\x99\x4c\xdd\x66\xca\x19\x40\x22\xc6\xa1\xfb\xcd\xbd\xf8\xd2\x52\x74\x1c\xa1\x38\x2c\xa7\xd4\x3b\x6d\x46\x1f\x17\x4c\xea\x79\x3d\x6d\x63\xc1\x39\xa8\x29\x45\x93\xb9\xde\x25\xa8\x06\x4c\x2d\xe1\xee\x9d\xf5\x0c\xfa\x62\xff\xac\xce\x15\x68\x14\x5f\xec\x76\x01\xa9\x48\x40\xfd\x9f\xfd\x86\x11\x6c\x7a\xba\xb6\x0d\x89\xe0\xc2\x26\x62\xa3\xd0\x19\xc6\x8f\x28\x12\xb0\x6f\x27\x7e\xa5\xc6\x61\x2c\x7e\xdd\x75\xab\x66\x76\x8e\xef\x69\x5b\xca\xf6\x7a\x6f\xb2\xdc\xe1\x06\x9a\xaf\x5c\xc1\x46\xda\x24\x03\x15\xe3\x7c\xad\x62\x2f\x55\x61\x01\x48\x12\x69\xd3\xc9\x7d\xfd\x0f\xd0\x2d\x2d\x54\x68\xb4\xf4\xca\xc3\xad\x40\x3f\x66\x49\x00\x86\x28\x58\x34\xf1\xbe\x69\x02\xde\xf3\x16\x46\x69\xb7\xfd\x26\x86\x2a\xd5\x50\x92\x08\x35\xa5\xa7\xb4\x13\x40\x6d\x7a\xa7\x3a\x1b\x0f\x73\x26\x33\x28\xa4\x86\xae\xd0\x0e\xce\xd5\xf2\x2e\x86\x77\x4a\x49\xa4\x35\x1e\x5e\x81\x01\x2a\x12\xa0\xf2\xea\xff\x6e\xb9\x6b\xb1\x35\x06\xaa\x82\xad\x80\x1c\xac\xc2\x8e\x41\x9b\x05\x89\xc0\x1d\x83\x36\x0b\x43\xb0\xf6\x4b\x06\xcf\x38\xdf\x67\xee\xae\x9d\x65\xbd\x26\xa0\x78\xc1\xef\x8e\x5a\xaf\x21\x15\x97\x96\x36\xa4\x6f\xbd\x2e\xda\x29\x8a\xf0\xfd\xe5\x1f\xe7\xe4\x79\xf5\xfc\xd7\xc5\x39\xa1\x3e\x67\x36\x0e\x34\x33\xdc\x67\xd6\x02\x5a\x3f\x07\xc5\xb5\xb1\x7e\xe1\x7b\xd9\xef\xec\xb1\x02\xf4\x02\xeb\x87\xb6\x5a\xbd\xac\x56\x03\xad\xd1\xa2\x61\xe9\x71\x22\xd4\x71\x58\xc4\x5b\xb2\xff\x17\x5f\xd1\xea\x96\x6c\x35\x0e\x6c\x57\xf6\x3b\xe0\x46\xe5\xda\x7e\x45\x4c\xfc\x6b\xeb\x5f\xff\x93\x81\x59\x1d\xb7\x60\x29\x7c\xb9\xbe\x0f\x2c\x02\x5b\x18\x1c\x4c\xc0\xbd\xd8\xdc\xa2\xdd\xb3\xdd\x4a\xc3\x37\x30\x5e\xc7\x3e\x98\xfb\x83\xcc\xd7\xed\xac\x31\x7e\x6d\xfd\x84\x29\x36\x07\x53\xaa\x29\xa6\x51\xb7\x80\xfe\x0d\x00\x00\xff\xff\x15\x5a\xb4\x9f\xdc\x27\x00\x00"

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

	info := bindataFileInfo{name: "templates/views/manager.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsManagerMinJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x6d\x6f\xdb\x38\x12\xfe\x2b\x2c\xdb\x8b\x44\x58\x96\x93\xbd\xfb\x70\x90\x25\x05\x45\x8b\x05\x0a\x14\x77\x8b\x26\xc0\x7d\x48\x82\x80\x16\xc7\x36\x37\xb4\x68\x90\x54\xdc\xc0\xd0\x7f\x3f\x90\x94\x25\xf9\xad\x4d\x9a\x97\x05\xf6\x8b\x40\x91\xc3\xe1\x43\x72\xe6\xe1\x0c\xf9\x21\x64\xb2\xa8\x16\x50\x1a\x12\x2b\xa0\xec\x21\x9c\x56\x65\x61\xb8\x2c\x43\xb2\xfe\x10\xe2\xf7\x2b\xa9\xee\x40\xe9\xa1\x9e\xcb\x15\x26\x71\x21\x78\x71\x77\x44\x06\xc5\x86\xea\x3b\x27\x99\xcc\xa9\x0e\x79\x3c\xa5\x43\x78\x00\xd2\xf6\x23\x35\x89\xfa\x4a\xe7\x9c\xc1\x2f\x2a\x1d\x6a\x41\xf5\xfc\xb8\x6a\xca\x18\x9a\x54\xc6\xc8\xf2\xca\x3c\x2c\x21\xd3\xd5\x64\xc1\xcd\xcd\xc1\xd1\x62\xfa\x27\xfd\x1e\xae\xad\x5c\x82\xff\xf8\xef\xc5\x25\x8e\x2a\x25\x12\x3c\x6a\x94\x8d\xce\xa9\x93\xcd\x7a\xca\x71\xc4\xa8\xa1\xc9\xba\x90\x55\x69\x92\x9d\x91\x87\xae\x16\x93\xf8\x9e\x8a\x90\xd4\x91\xae\x8a\x02\xb4\x4e\xaa\x25\xa3\x06\x6a\x52\x93\xf1\x3d\x55\xc8\xd0\x89\x80\xaf\x5c\x1b\x28\x41\xe9\xcc\x2a\x11\x9b\x3f\xdf\x88\x49\x2c\xcb\x10\x33\x45\x57\x31\x33\x38\x6a\x51\x43\xa4\xc1\x18\x5e\xce\x34\x59\x6f\x4a\xf1\x9f\x5a\x96\x27\x27\x5b\x6a\x5a\x24\x06\xbe\x9b\x70\x4b\x32\x56\x50\x48\xc5\xf4\xa5\x34\x54\x90\x9a\xc4\x9f\xa9\xa1\x97\x76\xd4\x70\xad\x0d\x35\x70\x41\xef\x21\x79\x77\x1a\xb9\x9f\xcf\x95\xa2\x76\xe4\xe4\x34\x12\xb4\x9c\x55\x74\x06\xc9\xda\xaf\x12\xa3\x7a\x3e\x91\x54\xb1\x91\x5d\x12\x87\x5b\x8f\xf8\xd9\xbf\x4b\x37\xcc\xb9\x90\x05\x15\x90\xe1\xc1\x8a\x97\x4c\xae\x62\x3d\xa7\x4c\xae\xbe\xba\xda\x3a\xb2\x4b\xbf\x51\xb4\xbb\xdc\x76\x60\x7d\x02\xa5\xe1\xe6\x21\x6b\xa7\xe4\x57\xfe\x42\x15\x09\xb6\x05\x5c\x47\x85\x14\xd5\xa2\xd4\xc9\xd5\xda\xed\x09\xe6\x0c\xd7\x51\x53\x2e\xe9\x02\xba\x3f\xb8\x87\xd2\x68\x1c\x29\x28\x19\xa8\xa4\x5d\x4e\xdb\x1a\xd9\xfd\x8f\x94\x5c\x91\xb5\xdd\x9c\x42\x96\x06\x4a\x93\x61\x3c\x9e\x4a\x15\xda\x2a\xd7\xfb\x0b\x43\xbc\x44\xb6\x03\x51\x72\x15\x0b\x59\xdc\x01\x3b\x6f\xa4\x07\x59\x90\xea\x25\x2d\x51\x21\xa8\xd6\x19\x16\x74\x02\x02\xb9\xef\x90\x97\x53\x89\xf3\x60\x60\xbb\x5e\x35\xaa\x6e\x06\x38\x1d\xd9\x0e\x39\xc2\x49\x4f\x07\x45\x73\x05\xd3\x0c\xbf\xc7\xc8\x70\x63\x57\xef\x1b\x2c\xe4\x3d\x2f\x67\x68\xb3\x0e\xf8\xe8\x18\x0e\xdc\xd0\xc8\xd9\xcc\x76\x5c\x48\x46\xc5\xa6\x8e\xaa\x19\x98\x0c\xbf\xef\x57\xba\xf2\xb0\x19\xe6\x93\x2c\xa7\x5c\x2d\x90\xb2\xc3\x41\x3b\x18\x7a\x1f\x0c\xec\x6c\x39\x1b\x60\x34\x95\xcd\x52\x20\xbc\x33\x99\x60\x4b\x65\x41\x85\x98\xd0\xe2\x2e\xc3\xed\xde\xb9\x59\x40\x78\x1d\x74\xea\x82\x08\x05\x78\xd0\xa8\x18\xe0\x80\x8c\xaf\x71\xbe\xab\x18\xa3\xef\xe9\x88\xe6\x08\x8f\x15\x98\x4a\x95\x9b\xdd\xa9\xdb\x9d\x9d\x72\x05\x1a\xf7\x7f\xb5\xb9\xb5\x95\xec\x96\x9a\x83\x1b\x0e\x64\xdd\x28\xb3\x3f\xe7\xf6\x73\x29\x2f\x8c\xe2\xe5\xcc\x37\x27\x18\x77\xfa\x05\x7d\x29\x7d\x52\x31\x50\xd6\x49\x92\x77\x67\x9e\x44\xca\x4a\x88\x9f\x99\x64\xa3\xba\x67\x72\x18\x27\x41\xca\xf8\xfd\xc6\x0e\x26\xa6\x1c\xce\x94\xac\x96\xa8\x2d\x0d\xbf\x6b\x9c\xa7\x9e\x06\x91\xa3\x41\xec\x7f\x70\xaf\x93\x13\x67\xb4\x9c\x81\x72\x45\x5e\xd8\xf6\x57\xb5\xa1\xc0\xd9\x10\x15\x02\x35\x0e\xf9\x64\xb3\x09\xae\x03\x32\xc6\x79\xca\x37\x13\x99\x52\x34\xa5\x43\xa3\xa8\x9e\x0f\xa9\x30\xdb\x7e\xd3\x81\xd0\xbb\x23\xe7\xe9\x88\xe7\xe9\xc8\x2f\x4b\x9e\x8e\x18\xbf\xcf\x83\xba\xbe\x89\xdc\x36\x25\x57\x57\x67\x11\xa6\xba\xc0\x37\xd1\xd5\x6f\x4d\xe9\xa6\x26\x91\x63\xb9\xff\x79\xbe\xca\xfa\xc7\xd5\xb3\x68\x7b\x73\x84\xfc\x6d\x48\xbb\x69\x7c\x1a\x65\x17\x0a\xa8\x01\xf6\x08\x27\xdb\xf6\xaf\xce\x59\x2d\x86\xaa\xc7\x06\xde\x63\xf6\x15\x4e\x05\x9d\xb5\x0a\xed\xcf\x39\xfe\xea\x45\x13\xfc\xbb\x02\xe8\xfc\xdf\xb9\x28\x83\x29\xad\x84\xf9\xe4\xb9\xc7\xba\xf3\x93\xbc\x79\xfb\x5c\x79\x94\xef\x06\xe3\x6e\xb6\xd4\x45\x42\x27\x27\x61\xef\x9c\x78\x8c\x6b\x37\xc1\x87\x2b\x17\x5c\x15\x02\x50\x1b\x52\xb5\x3e\xad\xef\x32\x7c\x1d\xa0\x01\xe2\x68\x80\xae\x83\x7d\xcf\x82\x07\x68\x7d\xea\x62\x2e\x57\x4e\xc7\x75\xa0\x11\x03\x43\xb9\xd8\x75\xa4\x80\x44\x4f\x84\xf9\x1a\x0c\x74\xc7\x85\x40\xde\x06\x2d\xfb\xb8\x45\xb4\xe4\x71\x84\x6f\x1a\x6b\xed\xb1\x4d\xdb\xe3\x89\x74\xe3\x35\xfd\x9c\x5c\x7e\x6b\xc9\xe5\xb4\x23\x97\xf1\x1e\x9f\x20\x33\x91\xec\xa1\x61\x15\x17\xbf\xe2\xa8\x59\xc7\x2e\x3c\xee\xd3\x0c\x59\x43\xbc\x54\x8e\xe2\x3e\x7b\xa3\x0d\x7d\xc8\x39\xc9\x3e\x84\x66\xce\x35\x89\xa7\xbc\x64\x21\xe6\x98\xd8\xf3\x25\xeb\xf3\x59\xac\xe4\x2a\xdc\x88\x15\x42\x6a\xd0\x26\xc4\x46\x61\x42\xc6\x7c\x1a\x4e\xe2\x39\xd5\x9f\xec\x32\x84\xb8\x31\x0c\xe2\x4d\xdb\x99\x91\x25\x68\xbb\x6c\x21\x71\xd0\xc6\x93\xd8\x9f\x03\x3b\x3d\x62\xca\xd8\x56\x95\x0f\xe8\x3d\x9c\xb8\x98\x73\xc1\xc2\x20\xf5\xb3\x5f\x71\x66\xe6\x19\x3e\x3b\x3d\xfd\x07\xce\x53\xa3\xf2\xd4\xb0\x3c\xad\x44\x1b\x01\x71\x6d\xbc\xd3\xe0\x3c\x15\x7c\xbf\x7a\xc8\x0d\x2c\x70\xbe\x15\x9a\x2d\x2b\x21\x86\x8a\xcf\xe6\x06\x59\x9a\x1d\x2e\x2a\x03\x0c\xe9\x05\x15\x02\xe7\x29\x2c\xf2\x60\x60\xf1\xbb\xdd\x4f\x47\xb0\xc8\x9b\x40\x2d\xd5\x46\xc9\x72\x96\x7f\xf9\x9c\x8e\x9a\x62\x3a\x51\x68\x94\xa7\x23\xc1\x5f\x7e\x78\x1b\xbf\x1e\x06\xf0\x1f\xba\x80\x37\x81\xb0\x54\x5c\x2a\x6e\x1e\x0e\xc3\xf8\xa3\x69\x7d\x13\x28\x0a\x96\x40\x8d\x3e\x8c\xe4\x9b\x6f\x7c\x6d\x20\xac\x39\x54\x2f\xe5\x37\xa0\xcc\x1a\x68\x73\x08\xf5\x10\xde\xf2\xd2\x80\xba\xa7\x82\xfc\x08\xea\x97\x46\xe8\xaf\x45\x6c\xf8\x02\x64\x65\x8e\x20\xbd\xf4\xad\xaf\x0e\xb1\x7f\x9a\x3b\x58\x4d\x14\x70\x4b\x8f\x21\xfb\xe4\x05\x5e\x1b\x99\x47\xa3\x0d\x55\x1e\xcd\xf9\x3e\xd4\xae\xd1\xc6\xf8\x47\xe0\x5e\x78\xa1\x37\x71\x13\x1f\xff\x1c\xc5\x61\xaa\x57\x77\x12\x07\x83\x1a\x03\x8b\xe5\x31\x77\xfd\xd8\xb4\xbe\xbd\x69\x51\x21\xe4\xea\xd6\xed\xda\x71\xf3\xfa\x68\x85\x90\x13\x7a\x1b\x13\xf3\x29\xeb\x0f\x0d\x6d\x57\xe4\x07\xe6\xf6\xbb\x15\x45\xfa\x6d\x8c\xce\xa3\x73\x29\xf2\x0f\xf1\xef\x48\x78\xf8\xf8\x10\xfc\xaf\xf4\x27\xe8\x47\x95\xc8\xd3\x91\x8d\x03\x46\x36\x20\x18\xb9\x48\x21\xc7\x24\xb6\xa1\x50\x48\x6a\x10\x1a\xd0\xc1\xd8\x63\x13\x68\xec\x45\x20\xfd\xd8\x23\x9e\x73\x06\xe1\xd6\x45\xdd\x25\xd5\x77\x3e\xd7\xb3\x73\x79\x66\xa6\xe7\x54\xfc\x7d\xf2\x3c\x37\x9d\xe7\x5c\xcc\x6d\x62\x8c\xae\xa6\x39\xea\x77\x2b\xda\x93\x75\x3f\x9b\x2b\x75\x97\x1c\x1e\x3b\xf1\x4a\xdd\x4b\x13\x9b\x93\xef\x05\x34\x75\x87\xd5\xf3\xb2\x56\xf5\x68\x25\x3f\xbb\xaf\x7a\xf9\x14\x18\x6f\xe8\xbc\x53\xba\x4d\xa4\xcf\x98\xfa\x2e\xb1\xbd\xd4\x85\xdd\x0b\x69\x7c\x52\x92\xef\x35\xbe\xda\xd5\x9c\x4f\xde\x3d\xaf\xbd\xe4\x35\x9d\x55\xfb\x98\x24\xd9\xb9\xfa\x0b\xa4\xc8\x56\xcf\x53\x12\xe4\x7f\xf6\x6e\xdf\xfc\x3b\x4a\xd6\x7b\xc3\xd9\x7e\x4a\x71\x0f\x3a\xb1\x02\x21\x29\x0b\xb7\x2f\xeb\x0e\x35\x39\x66\xdf\x6e\xa8\x23\x5a\x19\xa9\x60\xaa\x40\xcf\x33\xbb\xdb\x2e\x2b\xef\x55\x1e\x7a\x49\xb2\x29\x73\x5c\xcc\xc1\xdd\xcf\xda\x4e\x59\x96\xf5\xba\x9c\x9c\x84\x1e\x79\x48\xb6\xd4\x6f\xb8\x17\xda\x8c\xa0\x91\x8b\xce\xe0\x5f\x84\x38\x63\x7b\xb7\xab\xa9\xe9\x54\x08\xa0\xaa\xed\xd6\x13\x21\x7b\x33\x20\x35\x89\x9a\x5e\x3b\x37\xab\xdd\x4a\x6e\x1a\xbe\x30\xb2\x2e\x64\xa9\xa5\x80\x58\xc8\x59\x48\xd5\xcc\xbd\x12\x6a\x12\x3d\xfe\xb9\xac\x7b\x85\x6a\xcc\xb5\x79\x33\xe3\x2c\xe9\xc6\x89\xfc\x45\x6c\xf2\x51\x29\xfa\x10\xd3\xe5\x52\x3c\x84\xce\xbb\xba\x21\x63\x2d\x78\x01\xe1\x59\xef\x25\xed\x91\x5b\x5f\xd7\xa4\xde\xcc\x79\xeb\x76\xa7\x9b\x31\x67\xbf\xf2\x02\xb8\x37\x21\xce\x8e\x83\x7b\xa2\xe9\xf5\x30\xf7\x9c\xed\x57\x11\xfb\x48\xe3\x95\xf1\xd6\x64\xfc\xff\x00\x00\x00\xff\xff\xc2\xd0\x9e\x56\x4a\x1e\x00\x00"

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

	info := bindataFileInfo{name: "assets/js/manager.min.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x3b\x6f\x13\x4d\x14\x86\xdf\xf8\xf3\x27\x24\x0b\x9a\x14\x54\x14\x87\x82\x08\x8a\x0d\xde\x84\x02\xad\xb3\x09\x82\x24\x12\x22\x16\x10\x19\xfa\xc1\x1e\xdb\xab\xd8\x33\xd6\xcc\x2c\x17\x29\x45\x2e\x45\xa0\x81\x06\x1a\x0a\xe8\xa8\x93\x10\x4b\x16\x81\xcd\x5f\x38\x23\x3a\x0a\x44\xc9\x7f\xa0\x41\xeb\xd8\x89\x90\xe0\x34\x7b\xde\xcb\xb3\x67\xbe\x4f\x16\xdf\x00\xc0\xff\x00\x2e\x00\x58\x04\x70\x06\x40\x07\xc7\xf3\x7c\xe4\xbf\x00\x30\x0b\xe0\x03\x80\x73\x00\xbe\x01\x98\x04\xf0\x13\xc0\xc2\x04\xf0\x0b\xc0\x12\x80\x4a\x01\xb8\x0f\x60\xb3\x00\x9c\x05\xf0\xb5\x00\x9c\x07\xf0\xa3\x00\x4c\x00\x28\x00\xf8\x6f\x74\x2f\x9f\xe2\xe8\x8b\xba\x56\xcd\xa4\x55\x5c\x94\x4d\x91\x76\x1c\x3d\xd1\x66\x4d\x1a\x4b\x75\x9d\x2a\x77\x12\xa6\x46\xb8\x44\x2b\x6a\x6a\x43\x2e\xa9\xaf\x49\x43\x89\xa2\x46\x62\x7b\xc2\xd5\xdb\xd2\x90\x6e\x8e\xc9\x31\x73\xd7\xb5\x4f\x55\xe0\xc4\xa3\xe2\xb8\xb0\x2a\x7b\xda\xb8\xa0\x6a\x5b\x49\x23\xb8\x99\xb6\x6c\x50\xd3\x11\x35\xe4\xe3\x1b\x6b\x49\x5b\x74\xf5\xb4\x49\x4b\x2b\xc2\xba\xa0\x66\x84\xb2\x1d\xe1\xb4\x89\xe8\xce\x30\xa2\x6a\x6a\x44\x57\x37\x34\xcd\xfd\xd1\x9f\x2f\xad\x08\xd5\x4a\x45\x4b\x06\x35\x29\xba\x11\x9d\xe8\x88\x56\x53\x6b\x13\xa1\x4a\xd5\xdb\xd5\xa5\xe0\xa1\x34\x36\xd1\x2a\xa2\x70\xba\x5c\xba\xa5\x95\x93\xca\x05\xb5\x67\x3d\x19\x91\x93\x4f\xdd\xd5\x5e\x47\x24\xaa\x42\xf5\xb6\x30\x56\xba\xf8\x41\x6d\x39\xb8\x7e\xda\xcb\xdf\xd3\x94\x26\x58\x52\x75\xdd\x48\x54\x2b\xa2\xd2\xbd\x4e\x6a\x44\x27\x58\xd6\xa6\x6b\x23\x52\xbd\xa1\xb4\xf1\x6c\x85\x8e\xd7\x58\x5d\x0a\xcb\x71\x1c\xd2\xd4\x14\xe5\x6b\xf9\x62\x1c\x86\xb4\x40\x65\x8a\x86\x7a\x3e\x9e\x19\x47\x73\xf1\xb5\x7c\xbd\x3c\xac\xcd\x85\x65\x5a\x5f\x3f\x46\xe6\xe3\x99\xf2\x15\x5a\xa0\x90\x22\x9a\xa9\x80\xdf\x72\xc6\x87\x3c\xf0\x3b\xdc\xf7\x9b\x7e\x8b\xf7\x39\x23\xce\x78\xcf\x6f\xf0\x2e\xef\x71\xe6\xb7\xfc\x0e\x0f\xf8\x13\x67\xbc\x4f\x7c\xc4\x59\xe0\xb7\xf9\x73\x0e\xf9\x1d\xde\xe5\x2f\x3c\xf0\x2f\xc1\xef\xb9\xef\x37\x78\xc0\x19\x1f\x10\x1f\xf0\xa1\x7f\x45\x7e\x6b\x88\xe5\xfe\x2e\xe5\xec\x01\x0f\xfc\x26\x1f\x71\x7f\xf8\xcb\xdc\xef\xff\xf3\x12\xf8\xb5\xdf\xf0\xdb\xfc\x91\x33\xee\x83\xdf\xfd\xa5\x35\xc0\xef\x00\x00\x00\xff\xff\x64\xe8\x78\x3c\xf7\x02\x00\x00"

func localesRuLc_messagesConfigMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesConfigMo,
		"locales/ru/LC_MESSAGES/config.mo",
	)
}

func localesRuLc_messagesConfigMo() (*asset, error) {
	bytes, err := localesRuLc_messagesConfigMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesManagerMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x6c\x5c\x57\x15\xfe\x52\xe6\x4d\xc7\x63\x07\x4a\x43\x4b\x29\x7f\xb7\xa5\x49\x13\xda\x71\x3d\x86\x45\x35\xc9\xc4\x18\x93\x10\x43\x1c\x85\xd8\x21\x1b\x24\x78\x9d\xb9\x76\x9e\xfc\xe6\xbd\xd1\x7b\x6f\xe2\x44\x64\xe1\x19\x0b\x12\x34\xd0\x90\x96\x2a\x05\xa9\xa9\xd2\x4a\x08\x55\x2c\x26\x96\x27\x19\x27\xf6\x58\xe2\x47\x82\xdd\xb9\x48\x6c\x59\xb1\x04\xb1\x60\x05\x0b\x84\xce\xbd\x77\xec\x71\x62\x23\x9a\x91\xfc\xee\x3d\xf7\xfc\x7d\xe7\x3b\xe7\xf8\x2f\x4f\xa6\xde\x02\x80\x51\x00\x9f\x01\xf0\xeb\x3d\xc0\x04\x80\x23\x1f\x81\xfe\x9d\x48\x01\x8f\x03\x98\x4c\x01\xfc\xf4\xad\x14\x90\x05\x70\x36\x05\x0c\x02\xf8\x9e\x3d\xfd\x14\x90\x01\x70\xd1\xca\xf5\x14\xe0\x00\xf8\x49\xca\xc4\x7d\x33\x05\x3c\x03\xa0\x95\x02\x9e\x05\xf0\x47\x1b\xf7\xaf\x29\x20\x0d\xe0\x6f\xd6\xfe\x9f\xd6\xff\x5f\x29\x60\x2f\x00\xc7\x01\x3e\x0a\x60\x9f\x03\x3c\x06\xe0\x79\xc7\xe4\xdf\xef\x00\x43\x00\x46\x1c\x60\x00\x40\xd1\x31\x71\xbe\xee\x00\x29\x00\xa7\x1c\x60\x1f\x80\x69\xc7\xe0\x9a\xb3\x7e\x55\x07\xf8\x18\x80\xef\x3b\x26\x7f\xd3\xc6\x7f\xd3\xca\xef\xd9\x38\x1f\xd8\x38\xb7\x1d\x83\xeb\x8e\xd5\xaf\x39\x86\x9f\xdf\xd9\x38\x99\xb4\xc1\xfb\x4c\xda\xe0\x7d\xc1\xca\xa3\x69\x83\xeb\x68\xda\xe0\x3c\x91\x36\xf6\x67\xad\x2c\xd3\x26\x6f\x64\xed\xea\xd6\xbf\x69\xdf\x7f\x9e\x06\x9e\x00\xf0\x2b\x1b\x6f\x35\x0d\x8c\xed\x01\x7e\x6f\xdf\x9d\x8c\x39\x9f\xce\x00\xfb\x01\xbc\x98\x01\x0e\x01\x98\xca\x00\x9f\x07\xb0\x90\x31\x71\x6e\x64\x0c\xff\xbf\xb4\xf2\x6f\x33\xc0\x24\x80\x3f\x65\x80\x31\xee\xf9\x00\x30\x05\xe0\x3b\x03\x46\xff\x81\x3d\x57\x07\x0c\xce\x3f\x0c\x00\x9f\x04\xf0\xe7\x01\xe0\x53\x00\xfe\x3e\x00\x08\x00\x8f\x67\x4d\x3f\x0e\x66\x81\xcf\x02\x78\x29\x6b\xf2\x4e\x66\x81\x27\x01\x78\x59\xe0\x13\x00\x2e\x67\x0d\x9f\xd7\xb3\xc0\xab\x00\xde\xb6\x7a\x95\x05\x9e\xe2\xfe\x5b\xbf\x7f\x67\x4d\xde\x67\x07\x81\x17\x00\x1c\x1a\x34\xf5\x9d\x1a\x34\x38\xbe\x6b\x4f\xdf\x9e\x97\x06\x8d\xfd\x0f\x07\x81\x37\x00\xbc\x35\x68\xf2\x1d\x1c\x32\x7d\x2e\x0e\x99\x79\x3b\x31\x04\x3c\x0d\xc0\x1d\x32\xf6\x17\xad\x7c\x65\x08\xf8\x02\x80\x9b\xf6\xfd\x37\x43\xa6\x0e\x65\xe5\x7f\x58\xbb\xff\xd8\xf3\xa9\xbd\x46\xff\xf2\x5e\xa3\xff\xc6\x5e\x60\x8f\x59\x11\xbc\x02\xd3\x03\xe6\xe7\x20\xcc\xbc\xf5\xff\x3e\x07\x33\x43\x8c\xfb\x31\xfb\xc6\x9c\x3c\x0f\x83\x31\x0f\x53\xf3\x70\x9f\x0f\xf7\x72\xa4\x4f\xd6\x18\x60\xb8\x7f\xc2\xbe\xed\xb3\x3b\xcb\x75\x1c\x00\xf0\x1c\x80\x4f\x5b\x1d\xef\x33\xcf\xd5\x17\x61\xf6\x8d\xb9\x61\xfc\x8c\xfd\x65\x6b\xc3\x73\x65\xd7\x5c\xcf\x1e\xf7\x99\x67\xfe\xe3\x00\x72\xf6\xfd\x45\x7b\x32\xa7\x2f\xc1\xec\x40\xef\xc7\x7d\xe4\x99\xe6\xfe\x61\xbc\x94\x78\x61\x10\x63\xbc\x5c\xe6\x3f\xb1\x10\x46\xf3\x32\xea\xbb\xc6\x18\xf7\xfd\x70\x41\xc4\x89\x1b\x25\x18\x4f\x12\x59\xa9\x26\x31\xc6\x6b\x49\x18\xc9\xd9\x48\xc6\xe7\x31\xe1\x87\xb1\xc4\x44\x18\xcc\x7a\x51\x45\x44\xb2\x12\x5e\x90\xc2\xf5\x7d\xe1\x7b\x71\x22\x03\x8e\xb1\x83\x32\x71\xe3\xf9\x1d\x15\xbd\xbc\x13\x91\x74\x13\x59\xc6\xb1\x0b\x32\x48\x62\x1c\xf7\x22\xa9\xbf\x71\x22\x66\xbd\x48\x96\xed\x5d\x03\x93\x65\x9c\xf0\xca\xfd\x81\x27\xbf\x86\x93\xee\xa6\xa9\xbe\xf6\x2c\x4f\x6e\xc2\x3a\x19\x96\xe6\x65\x19\xa7\xdc\x8a\xc4\xa9\x5a\xe5\x35\x19\x89\x40\x2e\xf4\x10\x14\x70\x3a\xf2\xc2\xc8\x4b\x2e\xe1\xcc\x26\x3c\x9c\x91\x55\xe9\x26\xc2\x0b\x12\x19\x5d\x70\x7b\x72\x8c\xe9\xf3\xe1\x42\x5f\xfe\x69\x9b\x6c\x3a\x71\x93\x5a\x8c\x19\x37\x9e\xd7\x9f\x18\x33\x5e\x45\x86\xb5\x04\xe7\x3c\xdf\x17\xaf\xc9\x1d\x09\x13\xf2\x62\x49\x56\x13\x51\x96\xb3\x6e\xcd\x4f\x36\x15\xc3\x62\x3c\x92\xe2\x52\x58\x13\x71\x2d\x92\x63\x88\x13\x37\x89\x53\x5b\x05\x19\xd9\xe4\x31\xf7\x73\x96\x4d\x46\x95\x9a\x70\x83\x92\xf4\xcd\xfd\xb8\xeb\xd9\xdb\xe9\x28\x2c\xc9\xd8\x9a\x98\x72\xce\xb9\x5e\x62\xe4\xe9\x5a\x69\x4b\x79\x36\x28\xcb\x59\x2f\x90\x65\x23\x6a\x2b\x43\x56\x2f\xb4\x95\x7a\x21\xad\xb8\xe5\x67\x1f\xb4\xe7\x19\x59\x0d\xa3\x24\x37\x15\xcf\x79\xe5\xdc\x57\x6b\x73\x71\x6e\x26\x2c\x88\xb2\xbc\xf0\x95\x79\xef\xbc\x5b\x09\x87\xa3\x5a\x96\x1b\x97\x9b\x89\xdc\x20\xf6\xdd\x24\x8c\x0a\xe2\x9b\x5a\x25\xa6\x6a\x91\x5b\x09\xcb\xa1\x38\xb2\xcd\xfe\x68\xf6\xa4\x1b\xcc\xd5\xdc\x39\x99\x9b\x91\x6e\xa5\x20\x36\xe5\x82\x38\x53\x8b\x63\xcf\x0d\xb2\x53\x93\x53\xc7\x72\xdf\x96\x51\xec\x85\x41\x41\xe4\x87\x47\xb2\x13\x61\x90\xc8\x20\xc9\xcd\x5c\xaa\xca\x82\x48\xe4\xc5\xe4\x95\xaa\xef\x7a\xc1\x61\x51\x3a\xef\x46\xb1\x4c\x8a\x67\x67\x8e\xe7\x5e\xdd\xb2\x63\x3c\xb3\x32\xca\x1d\x0b\x4a\x61\xd9\x0b\xe6\x0a\x22\x7b\xda\xaf\x45\xae\x9f\x3b\x1e\x46\x95\xb8\x20\x82\xaa\x16\xe3\xe2\x97\x0e\x0b\x73\x2d\x06\xfb\xf3\x23\xc5\x62\x5e\x1c\x38\x20\xf8\x3a\xf2\x5c\x31\x9f\x17\x63\x62\x44\x14\xb4\x7c\xb4\x38\xda\x53\x1d\x29\x7e\x99\xaf\x07\xb5\xd9\x91\xfc\x88\xb8\x7c\xd9\xb8\x1c\x2d\x8e\x8e\x1c\x12\x63\x22\x2f\x0a\x62\xf4\x30\xe8\x67\xd4\xa6\x55\x55\x57\x0d\x5a\xa6\x8e\xba\xc6\x0f\x5d\xba\x4d\x2d\x2d\x36\xd4\x8f\x1f\x7a\x10\x2c\xaa\x45\x6a\xd1\x6d\xea\xaa\x86\xba\x42\x1d\xba\xf7\xff\x59\x51\x97\x96\x41\xd7\xe9\x3e\x75\xe8\x0e\xb5\x68\x55\x5d\xa5\x0e\xad\x0a\xba\x4b\x2d\xda\x50\x4b\xaa\xce\x91\xde\xa5\x2e\x6d\xa8\xa6\x6a\xd0\x3d\xea\x80\x7e\x4a\xcb\xaa\x41\x5d\x1d\x7e\x9d\x23\xd0\x7d\x6a\xd3\x3a\x75\xa8\x0d\xba\x41\x2d\xba\xa7\x16\xd9\x9a\xa1\xb2\xeb\x8a\x2e\xa5\xad\x16\xe9\x0e\xad\xf4\x2c\x85\x5a\xa2\x15\x6a\xf5\x5c\xd5\x35\x41\xcb\xaa\x4e\x6d\xf5\x03\xa1\xea\x74\x5f\x2d\xa9\xab\xd4\x52\x0d\x6a\x6b\x8b\xd5\x47\x8b\xa4\xcb\x58\xa1\x96\xba\xf2\x88\xfe\xbb\x52\x76\x8b\xba\x74\x57\xbb\xad\x53\xd7\x88\xb7\x35\x43\xba\x63\x6f\xa8\x26\xdd\x65\x43\xd5\xe4\xc4\x9c\x70\x59\x35\x99\xd7\xe5\x9e\xe6\x81\xf7\x6d\x7c\xdf\xda\x62\xd0\x62\xe9\xab\x84\x3a\xfc\x5f\x90\xcb\x61\x9e\xa8\x4d\x2b\xba\x8c\x07\x63\x3f\xa4\x7d\x20\xc3\x76\x86\x3b\x66\x08\xba\xdc\x60\xb5\xa8\x7b\x7a\x8f\x5a\xa0\xb7\x69\x8d\xeb\xf9\x05\x75\xd9\x48\x5d\xa1\xb6\x1d\xcc\xae\x30\xad\x57\xcd\xff\x49\xd3\xbb\x6a\x91\x3a\xd4\xe5\x2f\x67\x52\x0d\xd0\xfb\x96\xec\x4e\x5f\x79\x9c\x68\x5d\x1b\x2c\xd2\x32\x6b\x05\x6d\xe8\xe8\x0d\xed\xdb\x32\xdd\xb3\xa2\xe1\x94\xa1\xb6\xb8\xa8\x5d\x59\xd2\xa3\xc8\x15\xff\x48\x37\xb6\x05\xba\xa5\x1a\xda\x7e\x49\xd5\x8d\xd6\x98\xb6\xfa\x85\x0e\xe8\x3d\x5e\x04\x5a\xa3\x96\x5a\x62\xc0\xd7\x79\x3e\xf8\xba\x6d\x50\x54\xb3\x97\xf4\xa1\x71\xed\x68\x1c\x82\x3a\xcc\x35\xdd\x57\xaf\x33\x6f\x66\xd6\x68\x6d\xc7\xe9\xd6\xe5\xe6\xd4\x12\xad\x31\xd1\x0c\x49\x4f\xe2\xeb\xc3\x82\x47\x89\xf3\x9a\xa9\xd5\x79\xc7\x1e\x6e\x9f\x5e\x90\x1b\x5b\xb3\x7e\x73\xb7\x7e\x6c\x6e\xa7\xa6\xb9\xad\x1a\xaa\xce\xfd\x7d\x87\xda\xa6\x36\xc3\x06\x8f\xef\x86\x6e\xf9\xba\xba\xb6\x65\x75\x93\xee\x50\x47\x33\xd0\x56\x8d\xed\x1d\xb2\x9b\x03\x7a\x5f\xd5\x69\x83\xda\xea\xaa\xd9\x8b\x77\xb8\x29\xfc\xcf\x63\x51\x8f\x62\xdb\x2e\xd9\x3a\x0f\xfd\x03\x01\x77\x05\xb7\x0b\x9a\x0f\x13\xfa\xbf\x01\x00\x00\xff\xff\xc4\x9e\x7a\xb9\x59\x0d\x00\x00"

func localesRuLc_messagesManagerMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesManagerMo,
		"locales/ru/LC_MESSAGES/manager.mo",
	)
}

func localesRuLc_messagesManagerMo() (*asset, error) {
	bytes, err := localesRuLc_messagesManagerMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/manager.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesWorkersMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbf\xaa\x13\x41\x18\xc5\x4f\x2e\x8a\xb0\xa5\x85\x95\xc5\x67\xe1\x45\xc1\xb9\xee\x46\x0b\x99\xec\x24\xa2\x24\x20\x66\x21\x86\x55\xb1\x1c\xb2\xe3\x66\xc9\xee\xcc\x32\x33\x2b\x0a\x29\xc4\xc6\x27\xb0\xb0\xf1\x19\x4c\x17\x2c\x7c\x86\xc9\x0b\xf8\x2c\x92\x8d\x7f\x4f\x75\x7e\x9c\xdf\x07\xdf\x8f\xab\x97\x3e\x01\xc0\x19\x80\xeb\x00\xee\x00\xb8\x0c\x20\xc5\x29\x0b\x00\x57\x00\x3c\x03\x30\x19\x00\xaf\x00\x5c\x03\xf0\x79\x00\x0c\x7e\x39\x67\xf8\x27\x2f\x8d\xdd\x28\xeb\xb0\x54\xad\xb1\x9e\x65\xae\xac\x0a\xf6\xa8\x2b\x1d\xcb\x0d\xa7\x42\xbd\x79\xb8\xa9\xd6\xb2\x31\x17\xb6\x8b\xe6\xd2\x79\x96\x5b\xa9\x5d\x2d\xbd\xb1\x9c\x9e\xf6\x13\x65\x9d\x95\x8d\x29\x0c\xa5\xff\xf9\xe3\x68\x2e\x75\xd9\xc9\x52\xb1\x5c\xc9\x86\xd3\x1f\xe6\xb4\xec\x9c\xab\xa4\x8e\xb2\x27\xd9\x94\xbd\x50\xd6\x55\x46\x73\x4a\x2e\xe2\xe8\xb1\xd1\x5e\x69\xcf\xf2\x77\xad\xe2\xe4\xd5\x5b\x7f\xb7\xad\x65\xa5\x47\xb4\x5a\x4b\xeb\x94\x17\xcf\xf3\x19\x7b\xf0\xd7\x3b\xfe\xf3\x5a\x59\x36\xd5\x2b\x53\x54\xba\xe4\x14\x2d\xea\xce\xca\x9a\xcd\x8c\x6d\x1c\x27\xdd\xf6\xe8\xc4\xbd\x11\x9d\xaa\xd0\x37\x93\x58\x88\x84\xce\xcf\xe9\x58\xe3\x1b\x22\x49\x68\x42\x31\xf1\x9e\xc7\x62\xf8\x7b\x4a\xc5\xfd\x63\xbd\xd5\x6b\x69\x12\xd3\x76\x7b\x3a\x19\x8b\x61\x7c\x9b\x26\x94\x10\xa7\xe1\x08\xe1\x4b\xd8\x1d\xde\x87\xaf\x61\x17\xbe\x1f\x3e\x1c\x3e\x86\x7d\xf8\x16\xf6\xf8\x19\x00\x00\xff\xff\x21\xfd\x73\x78\xb0\x01\x00\x00"

func localesRuLc_messagesWorkersMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesWorkersMo,
		"locales/ru/LC_MESSAGES/workers.mo",
	)
}

func localesRuLc_messagesWorkersMo() (*asset, error) {
	bytes, err := localesRuLc_messagesWorkersMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/workers.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
	"locales/ru/LC_MESSAGES/manager.mo": localesRuLc_messagesManagerMo,
	"locales/ru/LC_MESSAGES/workers.mo": localesRuLc_messagesWorkersMo,
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
	"locales": &bintree{nil, map[string]*bintree{
		"ru": &bintree{nil, map[string]*bintree{
			"LC_MESSAGES": &bintree{nil, map[string]*bintree{
				"config.mo": &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
				"manager.mo": &bintree{localesRuLc_messagesManagerMo, map[string]*bintree{}},
				"workers.mo": &bintree{localesRuLc_messagesWorkersMo, map[string]*bintree{}},
			}},
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
