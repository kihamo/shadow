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
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

var _assetsJsManagerMinJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x5d\x6f\xdb\x3a\xd2\xfe\x2b\x3c\x6c\xdf\x48\x84\x65\xb9\x39\xef\x5e\x2c\x64\x49\x41\xd1\xe2\x00\x05\x8a\xdd\x83\x26\xc0\x5e\x24\x41\x40\x8b\x63\x9b\x27\xb4\x68\x90\x54\x5c\xc3\xd0\x7f\x5f\x90\x94\x25\xf9\xab\x4d\x9a\x8f\xc5\xb9\x11\x28\x72\xf8\xf0\x21\x39\x33\x9c\x21\xdf\x87\x4c\x16\xd5\x02\x4a\x43\x62\x05\x94\xad\xc3\x69\x55\x16\x86\xcb\x32\x24\x9b\xf7\x21\x7e\xb7\x92\xea\x1e\x94\x1e\xea\xb9\x5c\x61\x12\x17\x82\x17\xf7\x27\x64\x50\x6c\xa8\xbe\x77\x92\xc9\x9c\xea\x90\xc7\x53\x3a\x84\x35\x90\xb6\x1f\xa9\x49\xd4\x07\x9d\x73\x06\xbf\x08\x3a\xd4\x82\xea\xf9\x69\x68\xca\x18\x9a\x54\xc6\xc8\xf2\xda\xac\x97\x90\xe9\x6a\xb2\xe0\xe6\xf6\xe8\x68\x31\xfd\x8b\x7e\x0f\x37\x56\x2e\xc1\x7f\xfe\xfb\xf2\x0a\x47\x95\x12\x09\x1e\x35\x60\xa3\x0b\xea\x64\xb3\x1e\x38\x8e\x18\x35\x34\xd9\x14\xb2\x2a\x4d\xb2\x37\xf2\xd0\xd5\x62\x12\x3f\x50\x11\x92\x3a\xd2\x55\x51\x80\xd6\x49\xb5\x64\xd4\x40\x4d\x6a\x32\x7e\xa0\x0a\x19\x3a\x11\xf0\x95\x6b\x03\x25\x28\x9d\x59\x10\xb1\xfd\xf3\x8d\x98\xc4\xb2\x0c\x31\x53\x74\x15\x33\x83\xa3\x96\x35\x44\x1a\x8c\xe1\xe5\x4c\x93\xcd\xb6\x14\xff\xa5\x65\x79\x76\xb6\x03\xd3\x32\x31\xf0\xdd\x84\x3b\x92\xb1\x82\x42\x2a\xa6\xaf\xa4\xa1\x82\xd4\x24\xfe\x4c\x0d\xbd\xb2\xa3\x86\x1b\x41\xcb\x59\x45\x67\x90\x6c\xfc\x42\x30\xaa\xe7\x13\x49\x15\x1b\xd9\x59\x3b\x6a\x7a\xc4\xcf\xff\x59\x3a\xa4\x0b\x21\x0b\x2a\x20\xc3\x83\x15\x2f\x99\x5c\xc5\x7a\x4e\x99\x5c\x7d\x75\xb5\x75\x64\x57\x77\x0b\xb4\xbf\xa2\xda\x50\xa3\xcf\xa0\x34\xdc\xac\xb3\x96\xb5\x5f\xdc\x4b\x55\x24\xd8\x16\x70\x1d\x15\x52\x54\x8b\x52\x27\xd7\x1b\xb7\xec\x98\x33\x5c\x47\x4d\xb9\xa4\x0b\xe8\xfe\xe0\x01\x4a\xa3\x71\xa4\xa0\x64\xa0\x92\x76\xc5\x6c\x6b\x64\xb7\x38\x52\x72\x45\x36\x76\xfd\x0b\x59\x1a\x28\x4d\x86\xf1\x78\x2a\x55\x68\xab\x5c\xef\x2f\x0c\xf1\x12\xd9\x0e\x44\xc9\x55\x2c\x64\x71\x0f\xec\xa2\x91\x1e\x64\x41\xaa\x97\xb4\x44\x85\xa0\x5a\x67\x58\xd0\x09\x08\xe4\xbe\x43\x5e\x4e\x25\xce\x83\x81\xed\x7a\xdd\x40\xdd\x0e\x70\x3a\xb2\x1d\x72\x84\x93\x1e\x06\x45\x73\x05\xd3\x0c\xbf\xc3\xc8\x70\x63\x57\xef\x1b\x2c\xe4\x03\x2f\x67\x68\xbb\x0e\xf8\xe4\x18\x8e\xdc\xd0\xc8\xd9\xcc\x76\x5c\x48\x46\xc5\xb6\x8e\xaa\x19\x98\x0c\xbf\xeb\x57\xba\xf2\xb0\x19\xe6\x93\x2c\xa7\x5c\x2d\x90\xb2\xc3\x41\x3b\x18\x7a\x17\x0c\xec\x6c\x39\x1b\x60\x34\x95\xcd\x52\x20\xbc\x37\x99\x60\x07\xb2\xa0\x42\x4c\x68\x71\x9f\xe1\x76\xef\xdc\x2c\x20\xbc\x09\x3a\xb8\x20\x42\x01\x1e\x34\x10\x03\x1c\x90\xf1\x0d\xce\xf7\x81\x31\xfa\x9e\x8e\x68\x8e\xf0\x58\x81\xa9\x54\xb9\xdd\x9d\xba\xdd\xd9\x29\x57\xa0\x71\xff\x57\x9b\x3b\x5b\xc9\xee\xa8\x39\xba\xe1\x40\x36\x0d\x98\xfd\xb9\xb0\x9f\x2b\x79\x69\x14\x2f\x67\xbe\x39\xc1\xb8\xc3\x17\xf4\xa5\xf0\xa4\x62\xa0\xac\x91\x24\xbf\x9d\x7b\x3f\x51\x56\x42\xfc\x4c\x25\x1b\xe8\x9e\xca\x61\x9c\x04\x29\xe3\x0f\x5b\x3d\x98\x98\x72\x38\x53\xb2\x5a\xa2\xb6\x34\xfc\xae\x71\x9e\x7a\x4f\x87\x9c\xa7\xc3\xfe\x07\xf7\x3a\x39\x71\x46\xcb\x19\x28\x57\xe4\x85\x6d\x7f\x55\x1d\x0a\x9c\x0e\x51\x21\x50\x63\x90\x4f\x56\x9b\xe0\x26\x20\x63\x9c\xa7\x7c\x3b\x91\x29\x45\x53\x3a\x34\x8a\xea\xf9\x90\x0a\xb3\x6b\x37\x1d\x09\xbd\x3f\x72\x9e\x8e\x78\x9e\x8e\xfc\xb2\xe4\xe9\x88\xf1\x87\x3c\xa8\xeb\xdb\xc8\x6d\x53\x72\x7d\x7d\x1e\x61\xaa\x0b\x7c\x1b\x5d\xff\xde\x94\x6e\x6b\x12\x39\x2f\xf7\x1f\xef\xaf\xb2\xfe\x89\xf4\x2c\xcf\xbc\x3d\x25\xfe\x4e\x7e\xb9\x69\x7c\x9a\x57\x2e\x14\x50\x03\xec\x11\x76\xb4\x6b\x42\x9d\x3d\x5a\x0e\x55\xcf\xe0\xbd\x51\x1c\x02\x4e\x05\x9d\xb5\x80\xf6\xe7\x02\x7f\xf5\xa2\x09\xfe\x43\x01\x74\x26\xee\xac\x90\xc1\x94\x56\xc2\x7c\xf2\xee\xc5\x5a\xec\x93\x0c\x76\xf7\xe8\x78\x94\x79\x06\xe3\x6e\xb6\xd4\xc5\x33\x67\x67\x61\xef\x28\x78\x8c\xf5\x36\x21\x84\x2b\x17\x5c\x15\x02\x50\x1b\x18\xb5\x66\xab\xef\x33\x7c\x13\xa0\x01\xe2\x68\x80\x6e\x82\x43\xe3\x81\x35\xb4\x66\x73\x39\x97\x2b\x87\x71\x13\x68\xc4\xc0\x50\x2e\xf6\x6d\x25\x20\xd1\x13\x69\xbe\x86\x93\xb9\xe7\x42\x20\xaf\x83\xd6\xc1\xb8\x45\xb4\xfe\xe1\x84\x4b\x69\xb4\xb5\xe7\x50\xda\x1e\x4f\xf4\x28\x1e\xe9\xe7\xfe\xe3\xf7\xd6\x7f\x7c\xe8\xfc\xc7\xf8\xc0\x65\x20\x33\x91\x6c\xdd\x38\x0e\x17\x85\xe2\xa8\x59\xc7\x2e\xc8\xed\x7b\x12\xb2\x81\x78\xa9\x9c\x17\xfb\xec\x95\x36\xf4\x81\xe3\x24\x7b\x1f\x9a\x39\xd7\x24\x9e\xf2\x92\x85\x98\x63\x62\x8f\x90\xac\xef\xb2\x62\x25\x57\xe1\x56\xac\x10\x52\x83\x36\x21\x36\x0a\x13\x32\xe6\xd3\x70\x12\xcf\xa9\xfe\x64\x97\x21\xc4\x8d\x62\x10\xaf\xda\x4e\x8d\xac\x0f\xb6\xcb\x16\x12\x47\x6d\x3c\x89\xbd\xab\xdf\xeb\x11\x53\xc6\x76\xaa\x7c\x58\xee\xe9\xc4\xc5\x9c\x0b\x16\x06\xa9\x9f\xfd\x8a\x33\x33\xcf\xf0\xf9\x87\x0f\xff\x87\xf3\xd4\xa8\x3c\x35\x2c\x4f\x2b\xd1\x06\x39\x5c\x1b\x6f\x34\x38\x4f\x05\x3f\xac\x1e\x72\x03\x0b\x9c\xef\x44\x5f\xcb\x4a\x88\xa1\xe2\xb3\xb9\x41\xd6\x93\x0e\x17\x95\x01\x86\xf4\x82\x0a\x81\xf3\x14\x16\x79\x30\xb0\xfc\xdd\xee\xa7\x23\x58\xe4\x4d\x2c\x96\x6a\xa3\x64\x39\xcb\xbf\x7c\x4e\x47\x4d\x31\x9d\x28\x34\xca\xd3\x91\xe0\x2f\x3f\xbc\x0d\x51\x8f\x13\xf8\x17\x5d\xc0\x9b\x50\x58\x2a\x2e\x15\x37\xeb\xe3\x34\xfe\x6c\x5a\xdf\x84\x8a\x82\x25\x50\xa3\x8f\x33\xf9\xe6\x1b\x5f\x9b\x08\xab\x14\xb5\x56\x76\x25\xbf\x01\x65\x56\x41\x9b\x43\xa8\xc7\xf0\x8e\x97\x06\xd4\x03\x15\xe4\x47\x54\xbf\x34\x42\xff\x5b\xc6\x86\x2f\x40\x56\xe6\x04\xd3\x2b\xdf\xfa\xea\x14\xfb\xa7\xb9\xa3\xd5\x44\x01\x77\xf4\x14\xb3\x4f\x5e\xe0\xb5\x99\x79\x36\xda\x50\xe5\xd9\x5c\x1c\x52\xed\x1a\x6d\x18\x7f\x82\xee\xa5\x17\x7a\x13\x33\xf1\xf1\xcf\x49\x1e\xa6\x7a\x75\x23\x71\x34\xa8\x31\xb0\x58\x9e\x32\xd7\x8f\x4d\xeb\xdb\xab\x16\x15\x42\xae\xee\xdc\xae\x9d\x56\xaf\x8f\x56\x08\x39\xa1\xb7\x51\x31\x9f\x95\xfe\x50\xd1\xf6\x45\x7e\xa0\x6e\x7f\x58\x51\xa4\xdf\x46\xe9\x3c\x3b\x97\x05\xff\x90\xff\x9e\x84\xa7\x8f\x8f\xd1\xff\x4a\x7f\xc2\x7e\x54\x89\x3c\x1d\xd9\x38\x60\x64\x03\x82\x91\x8b\x14\x72\x4c\x62\x1b\x0a\x85\xa4\x06\xa1\x01\x1d\x8d\x3d\xb6\x81\xc6\x41\x04\xd2\x8f\x3d\xe2\x39\x67\x10\xee\x5c\xb7\x5d\x51\x7d\xef\xd3\x39\x3b\x97\x67\x26\x73\x0e\xe2\x6f\x95\xca\x39\xc6\xcf\xb9\x5e\xdb\x86\x11\x5d\x4d\x73\x9a\xef\x57\xb4\x87\xe7\x61\xc2\x56\xea\x2e\xff\x3b\x75\xa8\x95\xba\x97\x09\x36\x87\xdb\x0b\x20\x75\xe7\xd1\xf3\x12\x53\xf5\x68\x90\x9f\xdd\x3a\xbd\x7c\x96\x8b\xb7\x1e\xbb\x03\xdd\xf5\x95\xcf\x98\xfa\xbe\xef\x7a\xa9\x6b\xb7\x17\x42\x7c\x52\x1e\xef\x11\x5f\xed\x82\xcd\xe7\xe7\xde\x75\xbd\xe4\x65\x9b\x85\x7d\x4c\x1e\xec\x4c\xfd\x05\xb2\x60\x8b\xf3\x94\x1c\xf8\xff\x7b\x77\x68\xfe\xc1\x23\xeb\x3d\xb6\xec\xbe\x79\xb8\x97\x97\x58\x81\x90\x94\x85\xbb\x57\x6e\xc7\x9a\x9c\xf3\xde\x6d\xa8\x23\x5a\x19\xa9\x60\xaa\x40\xcf\x33\xbb\xdb\x2e\xf1\xee\x55\x1e\x7b\xf2\xb1\x59\x71\x5c\xcc\xc1\xdd\xb2\xda\x4e\x59\x96\xf5\xba\x9c\x9d\x85\x9e\x79\x48\x76\xe0\xb7\xbe\x17\xda\xa0\xbf\x91\x8b\xce\xe1\x1f\x84\x38\x65\xfb\x6d\x1f\xa9\xe9\x54\x08\xa0\xaa\xed\xd6\x13\x21\x07\x33\x20\x35\x89\x9a\x5e\x7b\xf7\xa3\xdd\x4a\x6e\x1b\xbe\x30\xb2\x29\x64\xa9\xa5\x80\x58\xc8\x59\x48\xd5\xcc\x3d\xe7\x69\x12\x3d\xfe\x5d\xab\x7b\x2e\x6a\xd4\xb5\x79\xdc\xe2\x2c\xe9\xc6\x89\xfc\x75\x6a\xf2\x51\x29\xba\x8e\xe9\x72\x29\xd6\xa1\xb3\xae\x6e\xc8\x58\x0b\x5e\x40\x78\xde\x7b\xf2\x7a\xe4\xd6\xd7\x35\xa9\xb7\x73\xde\xb9\xc0\xe9\x66\xcc\xd9\xaf\x3c\xd5\x1d\x4c\x88\xb3\xd3\xe4\x9e\xa8\x7a\x3d\xce\x3d\x63\xfb\x55\xc6\x3e\x98\x78\x65\xbe\x35\x19\xff\x37\x00\x00\xff\xff\xf1\x3b\xc0\x7c\xf3\x1d\x00\x00"

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

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xbd\x6f\x13\x59\x14\xc5\x4f\x1c\xaf\x56\x6b\xed\x36\x29\xb6\xda\xe2\x6e\xb1\xd1\xae\x56\x2f\xcc\xd8\x44\x8a\x26\x99\x04\x91\x0f\x09\x11\x2b\x21\x32\xf4\x0f\xfb\xd9\x1e\xc5\x9e\x67\xbd\x79\xc3\x87\x94\x22\x1f\x45\xa0\x81\x06\x1a\x0a\xe8\x28\xa8\x92\x10\x4b\x16\x81\xc9\xbf\x70\x5f\x4b\x41\xcd\xff\x00\x05\x1a\x27\x4e\x84\x04\xb7\x79\xf7\xdc\x7b\x7e\xf7\xbc\x4f\x13\xc5\xe7\x00\xf0\x0b\x80\xbf\x00\x2c\x01\xf8\x15\x40\x07\x67\xf5\xe8\x7c\xfe\x18\x40\x05\xc0\x6b\x00\x7f\x00\xf8\x08\x60\x02\xc0\x67\x00\xed\x31\xe0\x0b\x80\x65\x00\x8d\x02\x70\x0b\xc0\x9b\x02\xf0\x3b\x80\xaf\x05\xe0\x4f\x00\xbf\x8d\x03\x63\x00\x0a\x00\xc6\xcf\xf3\xf2\x2a\x9e\xbf\xa8\xeb\xb8\x19\xb5\x8a\x4b\xaa\x29\xd3\x8e\xa5\xfb\xda\x6c\x2a\x93\x50\x5d\xa7\xb1\xbd\x58\xa6\x46\xda\x48\xc7\xd4\xd4\x86\x6c\x54\xdf\x54\x86\xa2\x98\x1a\x51\xd2\x93\xb6\xde\x56\x86\x74\x73\x44\x8e\x98\x35\xdb\xbe\x54\xc2\xca\xbb\xc5\x91\x61\x43\xf5\xb4\xb1\xa2\x9a\xb4\xa2\x86\xb8\x9e\xb6\x12\x51\xd3\x01\x35\xd4\xbd\x6b\x9b\x51\x5b\x76\xf5\x94\x49\x4b\xeb\x6b\x35\xb1\x68\xd4\x30\x55\x2c\x49\xab\x02\x2a\x7b\xfe\x8c\xf0\x2a\xa2\x5c\xa1\x72\x25\x98\x9e\xfe\xdf\xab\x78\x5e\x69\x55\x26\x56\xd4\x8c\x8c\x93\x8e\xb4\xda\x04\x74\x73\x78\x83\xaa\xa9\x91\x5d\xdd\xd0\x34\xf7\xdd\xe1\xf9\xd2\xaa\x8c\x5b\xa9\x6c\x29\x51\x53\xb2\x1b\xd0\x85\x0e\x68\x23\x4d\x92\x48\xc6\xa5\xea\x8d\xea\xb2\xb8\xa3\x4c\x12\xe9\x38\x20\x7f\xca\x2b\x2d\xea\xd8\xaa\xd8\x8a\xda\xc3\x9e\x0a\xc8\xaa\x07\xf6\x4a\xaf\x23\xa3\x78\x96\xea\x6d\x69\x12\x65\xc3\xdb\xb5\x15\x31\x73\xe9\xcb\xff\xd3\x54\x46\x2c\xc7\x75\xdd\x88\xe2\x56\x40\xa5\xf5\x4e\x6a\x64\x47\xac\x68\xd3\x4d\x02\x8a\x7b\x43\x99\x84\x95\x59\x3a\x6b\xc3\xf8\x1f\xdf\x0b\x43\x9f\x26\x27\x29\x6f\xbd\xbf\x43\xdf\xa7\x05\xf2\x28\x18\xea\xf9\xb0\x3c\x5a\xcd\x85\x57\xf3\xf6\xdf\xa1\x6d\xce\xf7\x68\x6b\xeb\x0c\x99\x0f\xcb\xde\x7f\xb4\x40\x3e\x05\x54\x9e\x05\xbf\xe0\x8c\x4f\x78\xe0\xf6\xb9\xef\x76\xdc\x2e\x1f\x71\x46\x9c\xf1\xa1\xdb\xe6\x03\x3e\xe4\xcc\xed\xba\x7d\x1e\xf0\x3b\xce\xf8\x88\xf8\x94\x33\xe1\xf6\xf8\x7d\x0e\xb9\x7d\x3e\xe0\x0f\x3c\x70\x4f\xc0\xaf\xb8\xef\xb6\x79\xc0\x19\x1f\x13\x1f\xf3\x89\x7b\x4a\x6e\x77\x88\xe5\xf3\x03\xca\xd9\x63\x1e\xb8\x1d\x3e\xe5\xfe\xf0\x64\x3e\xef\xff\x34\x09\xfc\xcc\x6d\xbb\x3d\x7e\xcb\x19\xf7\xc1\x2f\x7f\xe0\x1a\xe0\x5b\x00\x00\x00\xff\xff\xc1\x0e\x6f\x8e\x20\x03\x00\x00"

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

var _localesRuLc_messagesManagerMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x5f\x6c\x5b\x57\x19\xff\x75\xd8\xc6\x71\x12\x18\x2b\x8c\xf2\xff\xec\x4f\xbb\x8d\xe1\xcc\x4e\x98\x54\xb9\x75\x42\xc8\x5a\x5a\xd1\x94\x34\x71\xe8\x2b\x77\xf1\x49\x72\x15\xfb\xde\x70\xef\x75\xd2\x8a\x3d\xc4\x8e\xa6\x16\x0c\x2b\xeb\x98\x06\x48\xdb\x68\x27\x24\x84\x90\x70\xa3\xb8\x75\xd2\xc4\x01\x09\x84\x78\xe2\x3b\xe2\x9d\x07\xc4\x1b\x7b\x41\x3c\x31\x24\x84\xbe\x73\x8e\x13\xa7\x6d\x10\x9b\xa5\xdc\x73\xbe\xf3\xfd\xfb\x7d\xbf\xef\xfb\xf2\x97\x47\x62\xaf\x03\xc0\x20\x80\xcf\x02\xf8\xd5\x21\x60\x0c\xc0\xc9\x0f\x41\xff\xce\xc4\x80\x0f\x03\x38\x1b\x03\xf8\xe9\x42\x0c\x48\x01\x98\x8e\x01\xbd\x00\xbe\x65\xcf\x52\x0c\x48\x02\xb8\x64\xe5\x6a\x0c\x88\x03\xf8\x41\xcc\xc4\x7d\x2d\x06\x1c\x01\xd0\x88\x01\x9f\x06\xf0\x47\x1b\xf7\x6f\x31\x20\x01\xe0\x5d\x6b\xff\x4f\xeb\xff\xaf\x18\xd0\x0f\x20\x1e\x07\x3e\x02\xe0\x70\x1c\x78\x08\xc0\xe3\x71\x93\xff\x68\x1c\xe8\x03\x90\x89\x03\x3d\x00\xf2\x71\x13\xe7\x6b\x71\x20\x06\xe0\x7c\x1c\x38\x0c\x60\x2a\x6e\x70\xcd\x59\xbf\xc5\x38\xf0\x51\x00\xdf\x89\x9b\xfc\x75\x1b\xff\x35\x2b\xdf\xb4\x71\x7e\x69\xe3\xdc\x8a\x1b\x5c\xb7\xad\x7e\x2b\x6e\xf8\xf9\x9d\x8d\x93\x4c\x18\xbc\x47\x12\x06\xef\x93\x56\x1e\x4c\x18\x5c\xc3\x09\x83\xf3\x4c\xc2\xd8\x4f\x5b\x59\x26\x4c\xde\xc0\xda\x55\xad\x7f\xdd\xbe\xff\x24\x01\x3c\x0c\xe0\x17\x36\xde\x46\x02\x98\x3f\x04\xfc\xde\xbe\x0f\x24\xcd\x39\x92\x04\x8e\x02\x98\x48\x02\xcf\x00\x58\x4a\x02\x5f\x00\xf0\xb3\xa4\x89\xf3\x9b\xa4\xe1\xff\x4f\x56\x7e\x37\x09\x9c\x05\xf0\xef\x24\x30\x02\xe0\x54\x0f\x30\x0e\xe0\xe5\x1e\xa3\xff\xb3\x3d\xff\xda\x63\x70\xfe\xa3\x07\xf8\x24\x80\xff\xf4\x00\x9f\x02\xf0\x68\x0a\x10\xcc\x7b\xca\xf4\xe3\x42\x0a\xf8\x1c\x80\x42\xca\xe4\x0d\x52\xc0\x23\x00\xae\xa7\x80\x8f\x03\xf8\x79\xca\xf0\xb9\x99\x02\x8e\x03\xf8\xad\xd5\xbf\x97\x02\x3e\x01\xa0\xbf\xd7\xf8\x3d\xd1\x6b\xf2\x8e\xf5\x02\x4f\x02\x98\xec\x35\xf5\x2d\xf7\x1a\x1c\x57\xed\xf9\xba\x3d\x6f\x5a\xfb\x5f\xf7\x02\xd7\x01\xb4\x7b\x4d\xbe\x0b\x7d\xa6\xcf\xb3\x7d\x66\xde\xbe\xdd\x07\x3c\x0a\xe0\x7b\x7d\xc6\xfe\x86\x95\x1b\x7d\xc0\x13\x00\xfe\x60\xdf\xff\xde\x67\xea\x78\xcf\xca\x47\xfa\x8d\xdd\x53\xf6\x1c\xee\x37\xfa\xe9\x7e\xa3\x0f\xfb\x81\x43\x66\x45\xf0\x1c\x4c\x0f\x98\x9f\xa7\x61\xe6\xad\xfb\xf7\x79\x98\x19\x62\xdc\x0f\xd9\x37\xe6\xe4\x71\x18\x8c\x59\x98\x9a\x07\xba\x7c\xb8\x97\x99\x2e\x99\x31\x30\x5f\xcc\xfd\xc3\xf6\xed\xb0\xdd\x59\xae\xe3\x18\x80\xc7\x00\x7c\xc6\xea\x78\x9f\x79\xae\xbe\x08\xb3\x6f\xcc\x0d\xe3\x67\xec\x5f\xb2\x36\x3c\x57\x76\xcd\xf5\xec\x71\x9f\x79\xe6\x3f\x06\x20\x6d\xdf\x9f\xb2\x27\x73\xfa\x2c\xcc\x0e\x74\x7e\xdc\x47\x9e\x69\xee\x1f\x46\x67\x22\xd7\xf7\x42\x8c\x16\x8b\xfc\x27\x96\xfd\x60\x41\x06\x5d\xd7\x10\xa3\xa5\x92\xbf\x2c\xc2\xc8\x09\x22\x8c\x46\x91\x2c\x2f\x46\x21\x46\x2b\x91\x1f\xc8\xd9\x40\x86\xf3\x18\x2b\xf9\xa1\xc4\x98\xef\xcd\xba\x41\x59\x04\xb2\xec\x2f\x49\xe1\x94\x4a\xa2\xe4\x86\x91\xf4\x38\xc6\x03\x94\x91\x13\x2e\x3c\x50\xd1\xc9\x3b\x16\x48\x27\x92\x45\x9c\x5a\x92\x5e\x14\xe2\xb4\x1b\x48\xfd\x0d\x23\x31\xeb\x06\xb2\x68\xef\x1a\x98\x2c\xe2\x8c\x5b\xec\x0e\x7c\xf6\x05\x9c\x73\x76\x4d\xf5\xb5\x63\x79\x6e\x17\xd6\x39\x7f\x66\x41\x16\x71\xde\x29\x4b\x9c\xaf\x94\x5f\x94\x81\xf0\xe4\x72\x07\x41\x0e\x13\x81\xeb\x07\x6e\x74\x19\x93\xbb\xf0\x30\x29\x17\xa5\x13\x09\xd7\x8b\x64\xb0\xe4\x74\xe4\x10\x53\xf3\xfe\x72\x57\xfe\x29\x9b\x6c\x2a\x72\xa2\x4a\x88\x82\x13\x2e\xe8\x4f\x88\x82\x5b\x96\x7e\x25\xc2\x45\xb7\x54\x12\x2f\xca\x07\x12\x26\xe4\xa5\x19\xb9\x18\x89\xa2\x9c\x75\x2a\xa5\x68\x57\x31\x20\x46\x03\x29\x2e\xfb\x15\x11\x56\x02\x39\x82\x30\x72\xa2\x30\xb6\x57\x90\x91\x4d\x1e\x73\xbf\x68\xd9\x64\x54\xb1\x31\xc7\x9b\x91\x25\x73\x3f\xed\xb8\xf6\x36\x11\xf8\x33\x32\xb4\x26\xa6\x9c\x8b\x8e\x1b\x19\x79\xaa\x32\xb3\xa7\x9c\xf6\x8a\x72\xd6\xf5\x64\xd1\x88\xda\xca\x90\xd5\x09\x6d\xa5\x4e\x48\x2b\xee\xf9\xd9\x07\xed\x39\x29\x17\xfd\x20\x4a\x8f\x87\x73\x6e\x31\xfd\xd5\xca\x5c\x98\x2e\xf8\x39\x51\x94\x4b\x5f\x59\x70\xe7\x9d\xb2\x3f\x10\x54\x52\x13\xdf\x28\xa4\xf5\x20\xb8\xbe\x97\x7e\xc1\x89\x64\x4e\x0c\x66\xb2\xc7\xd3\x99\xa1\xf4\xe0\x90\x18\x1c\xca\x3d\xff\xfc\xb3\x99\xa1\x4c\x26\xc5\x1d\x4e\x17\x02\xc7\x0b\x4b\x4e\xe4\x07\x39\xf1\x75\x1d\x43\x8c\x57\x02\xa7\xec\x17\x7d\x71\x72\x5f\xe0\xe1\xd4\x39\xc7\x9b\xab\x38\x73\x32\x5d\x90\x4e\x39\x27\x76\xe5\x9c\x98\xac\x84\xa1\xeb\x78\xa9\xf1\xb3\xe3\xa7\xd2\xdf\x94\x41\xe8\xfa\x5e\x4e\x64\x07\x32\xa9\x31\xdf\x8b\xa4\x17\xa5\x0b\x97\x17\x65\x4e\x44\xf2\x52\xf4\xdc\x62\xc9\x71\xbd\x13\x62\x66\xde\x09\x42\x19\xe5\xa7\x0b\xa7\xd3\xc7\xf7\xec\x18\xcf\xac\x0c\xd2\xa7\xbc\x19\xbf\xe8\x7a\x73\x39\x91\x9a\x28\x55\x02\xa7\x94\x3e\xed\x07\xe5\x30\x27\xbc\x45\x2d\x86\xf9\xa1\x13\xc2\x5c\xf3\xde\xd1\x6c\x26\x9f\xcf\x8a\x63\xc7\x04\x5f\x33\x8f\xe5\xb3\x59\x31\x22\x32\x22\xa7\xe5\xe1\xfc\x60\x47\x75\x32\xff\x65\xbe\x3e\xad\xcd\x4e\x66\x33\xe2\xa5\x97\x8c\xcb\x70\x7e\x30\xf3\x8c\x18\x11\x59\x91\x13\x83\x27\x40\x3f\xa2\x26\x6d\xa8\xaa\xaa\xd1\x1a\xb5\xd4\x35\x7e\x68\xd3\x2d\x6a\x68\xb1\xa6\xbe\x7f\xdf\x83\x60\x51\xad\x50\x83\x6e\x51\x5b\xd5\xd4\x15\x6a\xd1\xe6\xff\x67\x45\x6d\x5a\x03\xbd\x4a\x77\xa9\x45\xb7\xa9\x41\x1b\xea\x2a\xb5\x68\x43\xd0\x1d\x6a\xd0\x8e\x5a\x55\x55\x8e\xf4\x36\xb5\x69\x47\xd5\x55\x8d\x36\xa9\x05\xfa\x21\xad\xa9\x1a\xb5\x75\xf8\x6d\x8e\x40\x77\xa9\x49\xdb\xd4\xa2\x26\xe8\x0d\x6a\xd0\xa6\x5a\x61\x6b\x86\xca\xae\xeb\xba\x94\xa6\x5a\xa1\xdb\xb4\xde\xb1\x14\x6a\x95\xd6\xa9\xd1\x71\x55\xd7\x04\xad\xa9\x2a\x35\xd5\xcb\x42\x55\xe9\xae\x5a\x55\x57\xa9\xa1\x6a\xd4\xd4\x16\x1b\x1f\x2c\x92\x2e\x63\x9d\x1a\xea\xca\x07\xf4\x3f\x90\xb2\x1b\xd4\xa6\x3b\xda\x6d\x9b\xda\x46\xbc\xa5\x19\xd2\x1d\xbb\xae\xea\x74\x87\x0d\x55\x9d\x13\x73\xc2\x35\x55\x67\x5e\xd7\x3a\x9a\x7b\xde\xf7\xf1\x7d\x63\x8f\x41\x8b\xa5\xab\x12\x6a\xf1\xbf\x4b\x2e\x87\x79\xa2\x26\xad\xeb\x32\xee\x8d\x7d\x9f\xf6\x9e\x0c\xfb\x19\x6e\x99\x21\x68\x73\x83\xd5\x8a\xee\xe9\x26\x35\x40\x3f\xa6\x2d\xae\xe7\xa7\xd4\x66\x23\x75\x85\x9a\x76\x30\xdb\xc2\xb4\x5e\xd5\xff\x27\x4d\x6f\xab\x15\x6a\x51\x9b\xbf\x9c\x49\xd5\x40\xef\x58\xb2\x5b\x5d\xe5\x71\xa2\x6d\x6d\xb0\x42\x6b\xac\x15\xb4\xa3\xa3\xd7\xb4\x6f\xc3\x74\xcf\x8a\x86\x53\x86\xda\xe0\xa2\x0e\x64\x49\x8f\x22\x57\xfc\x5d\xdd\xd8\x06\xe8\x86\xaa\x69\xfb\x55\x55\x35\x5a\x63\xda\xe8\x16\x5a\xa0\x9b\xbc\x08\xb4\x45\x0d\xb5\xca\x80\x5f\xe5\xf9\xe0\xeb\xbe\x41\x51\xf5\x4e\xd2\xfb\xc6\xb5\xa5\x71\x08\x6a\x31\xd7\x74\x57\xbd\xc2\xbc\x99\x59\xa3\xad\x07\x4e\xb7\x2e\x37\xad\x56\x69\x8b\x89\x66\x48\x7a\x12\x5f\x19\x10\x3c\x4a\x9c\xd7\x4c\xad\xce\x3b\x72\x7f\xfb\xf4\x82\xbc\xb1\x37\xeb\x6f\x1d\xd4\x8f\xdd\xed\xd4\x34\x37\x55\x4d\x55\xb9\xbf\x6f\x52\xd3\xd4\x66\xd8\xe0\xf1\xdd\xd1\x2d\xdf\x56\xd7\xf6\xac\xde\xa2\xdb\xd4\xd2\x0c\x34\x55\x6d\x7f\x87\xec\xe6\x80\xde\x51\x55\xda\xa1\xa6\xba\x6a\xf6\xe2\x4d\x6e\x0a\xff\xf3\x58\xd1\xa3\xd8\xb4\x4b\xb6\xcd\x43\x7f\x4f\xc0\x03\xc1\x1d\x80\xe6\xfd\x84\xfe\x6f\x00\x00\x00\xff\xff\x76\x75\x5a\xb4\x82\x0d\x00\x00"

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

var _localesRuLc_messagesWorkersMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4f\x6b\x14\x31\x18\xc6\x9f\x2d\x8a\x30\x47\x0f\x9e\x3c\xbc\x1e\x2c\x8a\xa6\x26\x33\x16\x4a\x76\xb3\x15\x6b\x0b\x62\x17\xd7\x32\x2a\x1e\x43\x37\xce\x0e\xdd\x49\x86\x24\x23\x0a\x3d\x88\x17\x3f\x81\x57\x3f\x83\xbd\x2d\x1e\x3c\xf8\x09\xb2\x5f\xc0\xcf\x22\x3b\xeb\xdf\xe7\xf4\xfc\xc8\x2f\x0f\xef\x8f\xab\x97\x3e\x01\xc0\x16\x80\xeb\x00\xee\x02\xb8\x0c\x60\x84\x4d\xa6\x00\xae\x00\x78\x06\x60\x3e\x00\x5e\x01\xb8\x06\xe0\xfb\x00\x18\xfc\x72\xb6\xf0\x4f\x5e\x3a\x7f\x66\x7c\xc0\x89\x69\x9d\x8f\x6c\x12\xaa\x7a\xc6\x1e\x76\x55\x60\xa5\x93\x34\x33\x6f\x1e\x9c\xd5\x73\xdd\xb8\x1d\xdf\x65\xd3\xa7\x25\x3b\xf0\x46\xc7\xda\x59\xf6\x48\x47\x23\x29\xe7\x62\x8f\xf1\x82\xe5\x05\xe5\x85\xdc\xdd\xbd\xc3\x0b\xce\xb3\x63\x1d\x22\x2b\xbd\xb6\x61\xa1\xa3\xf3\x92\x9e\xf4\x1b\x34\xe9\xbc\x6e\xdc\xcc\xd1\xe8\xbf\xe1\x71\x76\xac\x6d\xd5\xe9\xca\xb0\xd2\xe8\x46\xd2\x1f\x96\x74\xd2\x85\x50\x6b\x9b\x4d\x1e\x4f\x0e\xd9\x0b\xe3\x43\xed\xac\x24\xb1\xc3\xb3\x03\x67\xa3\xb1\x91\x95\xef\x5a\x23\x29\x9a\xb7\xf1\x5e\xbb\xd0\xb5\x1d\xd2\xe9\x5c\xfb\x60\xa2\x7a\x5e\x1e\xb1\xbd\xbf\xde\xfa\x9e\xd7\xc6\xb3\x43\x7b\xea\x66\xb5\xad\x24\x65\xd3\x45\xe7\xf5\x82\x1d\x39\xdf\x04\x49\xb6\xed\x31\xa8\x62\x48\x9b\xaa\xec\x4d\xc1\x95\x12\xb4\xbd\x4d\xeb\xca\x6f\x28\x21\x68\x9f\x38\xc9\x9e\xc7\x2a\xff\xfd\x34\x52\xf7\xd7\xf5\x56\xaf\x8d\x04\xa7\xf3\xf3\xcd\x97\xb1\xca\xf9\x6d\xda\x27\x41\x92\xf2\x21\xd2\xe7\x74\xb1\x7a\x9f\xbe\xa4\x8b\xf4\x6d\xf5\x61\xf5\x31\x2d\xd3\xd7\xb4\xc4\xcf\x00\x00\x00\xff\xff\x4e\x6a\x2e\xd8\xd9\x01\x00\x00"

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
	"templates/views/manager.html":      templatesViewsManagerHtml,
	"assets/js/manager.min.js":          assetsJsManagerMinJs,
	"locales/ru/LC_MESSAGES/config.mo":  localesRuLc_messagesConfigMo,
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
				"config.mo":  &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
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
