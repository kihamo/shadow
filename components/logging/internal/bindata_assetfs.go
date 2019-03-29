// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/manager.mo
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

var _templatesViewsManagerHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4d\x6f\xe3\x36\x10\xbd\xe7\x57\x0c\xb8\x6e\x2d\x03\x89\x94\xe4\x54\x64\xad\x9c\x7a\xd8\x43\x5a\x14\xc5\xb6\xd7\x05\x2d\x8e\x6d\x1a\x34\xa9\x25\x69\x25\x86\xa0\xff\x5e\x90\x94\x2c\x59\xb5\x65\x79\xb1\x89\x2e\xe2\xd7\xcc\x7b\x9c\x79\x1c\xb2\x2c\x81\xe1\x92\x4b\x04\x92\x29\x69\x51\x5a\x02\x55\x75\x33\x67\xbc\x80\x4c\x50\x63\x52\xa2\xd5\x2b\x79\xbe\x01\x00\xe8\x8e\xbe\x7d\xcb\xa9\x44\x51\xcf\xfc\x7f\xd6\x72\x2b\xb0\x33\xeb\x57\xac\x1f\x9f\xcb\x12\xf8\xc3\x6f\x12\xc8\x8b\x5a\xad\x50\x1b\x02\x31\x54\xd5\x3c\x59\x3f\xf6\xd6\x76\xbc\x65\x02\xa9\x5e\xf2\x37\xf2\x3c\x4f\x18\x2f\x3a\x90\xbd\xee\x11\x83\x66\x3f\x3d\xbf\x4b\xa5\xb7\xa0\x95\xc0\x94\xb8\x26\x01\x9a\x59\xae\x64\x4a\x3e\x11\xd8\xa2\x5d\x2b\x96\x92\x5c\x19\x4b\x80\xb3\x94\x88\x9a\xe5\xb1\x93\x3e\x98\xa5\x0b\x81\x77\x1a\x4d\xae\xa4\xe1\x45\x7f\xdf\x07\x13\xbf\xee\xc8\x08\x82\xe9\x5a\x15\xa8\xeb\xb6\xb1\x9a\xe7\xc8\x80\xd9\x8e\x47\x90\xea\x55\xd3\x9c\x80\xb1\x7b\x47\xfd\x95\x33\xbb\x7e\x7a\xb8\xbf\xff\xe5\x0c\x56\xc0\x5b\x23\x65\x43\xf3\xfa\xfc\x64\xed\xa0\x4d\xd8\x5f\x54\x7b\x79\x84\x7c\xd9\xf5\x15\xa6\x7f\xd2\x2d\xfe\x90\xe1\x0b\x16\x28\x7e\xc8\xf2\x77\xcc\x51\xb2\x91\x84\xe7\xc9\x50\x24\x9c\xed\x85\x38\x2e\x14\xdb\x9f\x9f\x2f\x4b\x98\x70\xc9\xf0\x0d\x9e\x52\xb8\x77\xe7\x6b\x60\xa5\xa6\x72\x85\x30\xe1\xec\x16\x26\x41\x7d\xce\x2a\xae\x85\x78\xc1\xb8\x86\x49\x81\x32\xd6\x74\x1e\x86\x6c\x46\x48\x80\xf9\xa0\x2e\x01\xbf\x37\x84\xe2\x20\x05\x20\xae\x56\xfc\x2a\x17\x26\xff\x5c\x96\x80\xc2\x20\x54\x95\x63\x71\xbc\xcc\x8f\xa1\x64\x21\x0d\x03\x71\xec\xe0\x35\x2e\x9c\x72\xae\xb3\x2b\xc4\x37\x83\xf6\x62\xa4\x0f\x76\x06\x05\x66\xb6\x39\x94\xae\x24\xdc\xb9\xca\xa1\x95\x80\x30\xf5\x48\x40\xd2\x2d\xa6\xc4\xc7\xd7\xed\xa2\x5b\x1a\xee\xda\xa8\x57\xd5\xc0\x59\x6c\xbe\x36\xc3\xe2\xd6\xb3\x75\x4c\x27\xb1\x70\x42\x1f\xcc\xee\x11\x69\x95\xbb\x9a\x05\x05\x15\xbb\x9a\x98\x28\x44\xfc\xaf\xeb\x56\x15\x19\xe5\x23\x24\x95\x4a\x06\x91\xcf\x6c\x1d\xb7\xfb\x99\xef\x47\x4d\x02\xfc\x09\x9c\x75\xfc\xcf\xc6\x92\x6c\xbe\x10\x45\x64\x29\x69\x5a\xe3\x08\x76\x88\x1e\xd8\xa5\xc3\x6a\xee\x99\x05\xcd\x1d\x6a\x82\xdf\x83\x17\xd4\xc4\xdf\x76\x4b\xbe\x22\xe0\xd3\xc8\xe5\x8a\x78\x95\x85\xb8\x8e\xca\x62\x70\x7e\x59\x5f\x49\xd8\xf5\x65\x9f\x73\x93\x69\x9e\x5b\xb0\xfb\x1c\x53\x42\xf3\x5c\xf0\x8c\x3a\x3a\xc9\x86\x16\x34\x4c\x8e\x10\x98\xfb\x26\x11\x53\xd9\x6e\x8b\xd2\xce\x62\x8d\x94\xed\xa3\xe5\x4e\xfa\x6b\x0e\xa2\x19\x94\xa3\xc3\x3f\x89\xa6\x9f\x4e\xc8\x7c\x3a\x8b\x95\x8c\xa6\xf5\xf1\x78\x0a\xff\xe9\x2d\xb4\x20\x78\x0d\xca\x31\x92\x99\xce\x62\xb3\x5b\x6c\xb9\x8d\x66\x9f\x47\xfb\xa8\x46\xae\x1d\xb3\x6e\x9e\x84\x60\x5f\x28\x38\xd7\x96\xb2\xc3\x8d\x74\xb1\x9e\x0d\x5f\x47\x97\xb5\x37\x4f\x06\x2e\xa4\x79\xe2\x9f\x19\x27\x1e\x34\xc7\x8f\xa9\x30\xe4\xea\xe1\xc9\xf7\x56\xdd\xac\x7f\x2d\xa7\x9b\xce\x8b\xd2\xdd\x9a\xa4\x21\x5a\x96\x60\x2c\xb5\x3c\xfb\xf2\xf5\x8f\x17\x88\x42\xfb\x9f\xbf\x5f\x80\x24\x8c\x9a\xf5\x42\x51\xcd\x12\x6a\x0c\x5a\x93\x14\x28\x99\xd2\x26\x61\xd4\x52\xcf\xd6\xc4\x12\xed\xdd\xc2\x24\x99\x09\xa3\x5f\xc3\xe8\x42\x29\x6b\xac\xa6\x79\xbc\xe5\x32\xce\x8c\x21\xb0\xa4\xc2\xb4\x45\xea\x27\xa0\x2e\xf9\x1b\x32\xb7\x15\xd4\x0d\x03\x3f\xf4\xc5\x0f\x7d\x08\x85\xf6\x1d\xd8\x30\x68\x47\x86\x09\x9c\x4e\xcc\xc6\xfc\xc4\xb4\x24\x1b\x93\x6c\xbe\xef\x50\xef\xe3\x4e\x66\x1c\x97\xcd\x7b\xc4\x62\x61\x1c\xe0\x59\x0d\xbc\x0b\x66\x47\x02\x3d\xf0\xae\x12\xde\x0d\xbe\x4d\x76\x0f\xbd\xa3\x82\x0f\x00\xaf\x43\x7f\x56\x7a\x57\xc1\xd7\xd7\x6e\x03\xbe\x31\xc9\x96\x4a\xea\xea\xa4\x73\x63\xf5\xae\xa7\xdf\xff\x02\x00\x00\xff\xff\x6c\x6b\x67\xa7\xa9\x0e\x00\x00"

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

var _assetsJsManagerMinJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x41\x8f\xdb\x2e\x10\xc5\xbf\x4a\x96\x7f\x94\x80\x82\x88\x73\xf8\x4b\x15\x09\xd9\x43\xf7\x18\xa9\x52\xd5\x5b\xe4\xc3\x04\x26\x36\x5d\x16\x5c\xc0\x6b\x55\x91\xbf\x7b\x85\xbd\x6a\xd3\x6d\x7b\x1b\x31\xbf\x07\xf3\x78\xb3\xa4\x26\xe8\xfe\x05\x7d\x66\x22\x22\x98\xef\xf4\xda\x7b\x9d\x6d\xf0\x94\xdd\x5e\x21\x2e\x32\x5c\x1c\x9e\x42\xd3\x60\x4c\x6a\x49\xc9\x7f\x6e\xae\xe7\x06\x61\xe2\x09\x32\x7c\x29\x35\xbd\x75\xd0\xe0\x09\x7d\x93\x5b\xf9\x7f\xc5\x1d\xf8\xa6\x87\x06\xe5\xad\x8f\x4e\x92\xad\x81\xd4\x5e\x02\x44\xb3\x35\x90\x61\x92\xa7\xad\xdd\x7d\xf0\xe2\x6b\x0a\xfe\xd1\x05\x0d\x0e\x15\xd9\x0c\xd6\x9b\x30\x88\xd4\x82\x09\xc3\x69\x3a\x1d\x79\x88\x06\xa3\x3c\x9f\x2b\x4e\x20\x69\x52\xd7\x5c\x07\xd7\xbf\xf8\x27\xbc\x26\x79\xbe\xbd\xda\x64\x2f\x0e\xe5\xc3\x8e\x67\x88\x0d\xe6\x24\xab\xb1\xe6\x26\xc2\xf0\x11\x9c\xbb\x80\x7e\x96\x3f\x8d\x25\xcc\xd9\xfa\x26\xcd\x06\xa1\xb3\x2a\xb7\x36\x09\xe8\x2c\x65\x3c\x86\x21\x29\xe8\xac\x28\xc5\x6c\x49\x12\xdd\xc7\x88\x3e\x93\x91\x09\x1f\x0c\x26\xca\xb8\x83\x94\x95\xef\x9d\xdb\x17\x78\x1e\x86\x56\xfc\x4f\x41\x31\x4b\x99\x40\xd0\xed\xaf\xbf\x6d\x62\xe8\x3b\x6e\xe7\x09\x96\xf0\x39\x0c\x6a\x49\xcb\x8b\x4c\xe0\x37\x6a\xd9\xbe\x5c\xff\xa0\xd4\xc4\xad\x56\x74\x42\xc4\x05\xaf\x21\x22\x5d\x1f\x72\x5c\x68\x07\x29\x29\x32\x01\xe4\x78\xd8\x66\x73\x3c\x64\xb3\xd0\xc1\xa5\x0e\xbc\x22\xeb\xcd\xac\xd1\xad\x75\x26\xa2\xa7\x4c\xb8\x29\x9a\xcd\x9a\x1c\xd7\x9b\x49\xb7\x21\xb3\x6e\x9b\xe3\x91\xbc\x59\x9a\x1a\x6c\x64\xe3\xc8\xf6\xf7\xd9\x8b\xe0\x29\xd1\xce\xea\x67\xc2\x49\x8e\x62\x7e\x98\xbf\xdb\x96\x37\xdf\x9f\x4a\x5a\xea\x77\x79\x39\xa2\xec\x5c\xd5\xfb\x4a\x29\x75\x0f\x9e\xab\x7a\xb5\x9a\x62\x7d\xdf\xd8\xd5\x8f\x7f\xb9\xa4\x2c\x81\xc1\xb2\x05\x4c\x94\x84\x29\x93\xff\xa0\xe0\x1e\x1a\xd9\xc8\xf6\x3f\x02\x00\x00\xff\xff\x39\xed\x9a\x13\xf0\x02\x00\x00"

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

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x4d\x8c\x53\x55\x18\x3d\xb4\xb5\xad\x55\x54\x40\x41\xfc\xcb\x65\x01\x42\xf0\x61\xdb\x91\x48\xca\x14\x8c\xc3\x8f\x08\x03\x13\xa6\xea\xc6\x84\x5c\xda\x37\x9d\x17\x5e\xef\x6d\xde\x7b\x45\x8d\x2c\x86\xc1\x18\x11\x12\x34\x62\xa2\x46\x40\x94\xb5\x29\xe3\x14\x2b\x43\x3b\x1b\xd7\xe6\xbb\xc6\x85\x2b\x13\x13\x13\x37\xec\x5d\x98\x18\x73\xdf\x7d\x7d\xd3\x19\x5d\x68\x37\xfd\x7e\xce\x3d\xdf\xb9\xe7\x7b\xf7\x97\xb5\xa9\x8f\x01\xe0\x49\x00\x4f\x00\x48\xad\x02\xb6\x01\xb8\xbb\x0a\xe1\xef\x83\x04\xb0\x1a\xc0\x87\x09\xe0\x01\x00\xd7\x12\xc0\xfd\x00\xbe\x4e\x00\x6b\x01\x2c\x24\x80\x47\x01\xfc\x90\x00\xd6\x03\xf8\x3d\xc2\xfd\x15\xe1\x56\x27\xcd\xff\x63\x49\xe0\x25\x00\x4f\x27\x81\xfb\x00\x34\x93\x06\x77\x26\x09\x6c\x00\x70\x21\x09\x3c\x0c\xe0\xab\xa8\xdf\x4d\x02\xeb\x00\x7c\x1f\xe5\x3f\x27\xcd\xbc\xdf\x22\xbe\x3f\x92\xc0\x1a\xad\x37\x65\x78\xd6\xa7\x4c\x7d\x73\x0a\x78\x1c\x40\x31\x65\x74\x1f\x4d\x99\x73\xaf\xa7\x0c\x8f\x8c\xf0\x6f\x47\xf5\xf7\x53\xc0\xf4\x2a\xe0\x4a\x94\xa7\xd3\xa6\xbf\x31\xfa\xdf\x9a\x36\xba\x9e\x4f\x03\x9b\x00\x1c\x49\x03\x3b\x00\x88\xa8\x7e\x25\x6d\xe6\xce\xa5\xcd\xf9\xc5\x34\x70\x18\xc0\x8f\x51\xfd\xa9\x0c\xf0\x10\x80\xed\x19\xa0\x04\x60\x34\x63\x7c\x7e\x23\x03\xa4\x00\x5c\xce\x18\xff\x3e\xcd\x00\x39\xcd\x93\x31\x3e\xf6\x32\xe6\xfc\x4f\x51\xfd\xd7\x8c\xd1\x73\x37\xea\xff\x99\x01\xf2\x00\xd6\x66\x0d\x6e\x2c\x6b\xee\x3d\x91\x05\xd2\x00\xdc\xac\x99\xe3\x67\x81\x8d\x7a\x8f\x59\x40\xaf\x54\xc7\x49\xb3\xda\x50\xff\xba\x28\xd6\x67\xb4\x9e\x35\x51\xae\xcf\xea\xfb\x64\x01\xdc\xab\xf7\x07\xe0\x41\x18\x0d\xf7\xc0\x78\xab\xe7\x6a\x4f\x73\xd1\xf7\xb3\x21\x3a\x9b\xc1\xf2\x5f\x62\x28\xd6\x5e\x68\x9d\x8f\x0c\xd5\xf4\xfd\xd7\x0f\x92\xaa\x14\x53\x4e\x3d\x35\xc6\x5d\xd7\xf6\xe2\x4c\x0a\x5f\xba\xf6\x20\xdd\x67\x9f\x6c\xd5\x97\x92\xd3\xb6\x2b\x9b\x0d\x5b\x04\xff\x52\x62\x4d\x2e\x9c\x6a\xdc\x68\x79\x3c\x70\xa4\x60\x53\xd2\x6b\xf0\x18\xbf\x5f\x54\x65\x6d\x69\xda\x7e\xcf\x93\x71\x72\x80\x07\xdc\x8d\x13\xc7\x76\x6b\x3e\x73\x06\x04\x6c\x4a\x17\x4e\x08\xde\xb0\xcb\x61\x58\x38\x71\x9a\xbb\x2d\xfb\x99\x30\x29\x0e\x35\x8a\xa6\x11\x13\xb5\xdc\x98\xf4\xa0\x2d\x6c\x6f\x69\xc6\x41\x57\x9e\xe4\x2e\x73\x65\x9d\xb9\xfa\x1e\x83\xfa\x21\x61\x66\x3a\x52\x2c\x81\x5f\x9e\x3c\x76\x74\x10\x8f\x3b\xae\xeb\xf8\x76\x55\x8a\x9a\x1f\xd7\x64\x2d\x9e\x79\x94\x0b\xb9\xa2\x3d\x31\x6c\xce\x84\x27\x6b\xad\xaa\xe6\x1f\x54\x26\x97\xa3\x27\xa7\xa5\x17\x7b\x36\x19\xf0\xea\xa9\xc0\xe3\x55\xfb\x9f\x52\x27\x03\xcf\x11\xf1\x82\x2a\x4e\xc3\x5e\x61\x78\xe5\xad\x66\x2c\xeb\x35\xee\x89\x21\x34\xaf\xd5\x18\x37\xbe\xe2\xb8\xdd\x94\x5e\x60\x8d\xfb\x75\xa7\x66\xbd\xd8\xaa\xfb\x56\x45\x96\x58\xcd\x3e\xfd\xc2\x29\x67\x9a\x37\xe4\x0e\xaf\x95\x9b\x38\x56\xb1\xc6\x3c\x3b\xf4\xc5\xda\xc7\x03\xbb\xc4\x8a\xf9\xc2\x2e\x2b\x3f\x62\x15\x47\x58\x71\xa4\xb4\x73\xe7\xf6\xfc\x48\x3e\x9f\x3b\xc2\xfd\xc0\xaa\x78\x5c\xf8\x2e\x0f\xa4\x57\x62\x87\x43\x0e\x36\xde\xf2\x78\x43\xd6\x24\x1b\x5d\x46\xbc\x27\x77\x84\x8b\x7a\x8b\xd7\x6d\xab\x62\xf3\x46\x89\xc5\x79\x89\x1d\x6f\xf9\xbe\xc3\x45\x6e\xfc\xd0\xf8\x7e\xeb\x55\xdb\xf3\x1d\x29\x4a\xac\xb0\x23\x9f\x1b\x93\x22\xb0\x45\x60\xe9\x0b\x96\x58\x60\xbf\x19\x3c\xdb\x74\xb9\x23\x76\xb3\xea\x34\xf7\x7c\x3b\x28\xbf\x52\x39\x60\xed\x5a\xc2\x69\x3d\x53\xb6\x67\x85\x9f\xa0\x23\xea\x25\x96\x9b\x70\x5b\x1e\x77\xad\x03\xd2\x6b\xf8\x25\x26\x9a\x61\xea\x97\x47\x76\x33\x13\x96\xc5\xe6\x42\xbe\x5c\x2e\xb0\x2d\x5b\x98\x0e\xf3\x9b\xca\x85\x02\xdb\xcb\xf2\xac\x14\xe6\x7b\xca\xc5\x41\x6b\xb4\xfc\x9c\x0e\xb7\x86\xb0\xd1\x42\x9e\x9d\x39\x63\x8e\xec\x29\x17\xf3\xdb\xd8\x5e\x56\x60\x25\x56\xdc\x0d\xfa\x84\x7a\xd4\x55\xef\x52\x97\xda\x6a\x96\xfa\x6a\x06\xf4\x19\xf5\xa9\xa7\xce\x52\x9f\x16\xd4\x45\xd0\x55\x35\x4b\x0b\xd4\xa6\x79\xba\x4d\x6d\xd0\x17\xd4\xa6\x6f\xd5\x0c\xb5\xe9\x26\xf5\xd5\xac\x29\x5e\xa3\xb6\xe6\xd1\x09\x0b\x7b\xcb\x11\x5d\xd0\x0d\xcd\x4d\x77\xf4\x14\x46\x5d\xea\xa9\x59\xea\xa8\x19\x9a\xa3\xb6\x26\x67\x34\xa7\x66\xa8\x43\x77\xa8\xa3\x79\x8c\x86\x79\xea\xaa\x19\xea\xd3\x9c\x3a\xaf\xa9\xb5\x92\xf7\xa8\x4b\x37\xcd\xc8\x1b\xa1\xe0\xb6\xd6\x48\x3d\x75\x81\xbe\xd3\x2a\xb4\xe4\x4b\x8c\xe6\x98\x7a\x67\x69\x1e\x75\xfe\xff\x6b\x35\x5c\x03\xe2\xab\xea\x2c\xf5\xb4\x92\xb0\xd0\x01\x7d\x44\x0b\xd4\xa7\x9b\xc3\xe3\x99\x3a\x67\xd4\xea\x2b\xa8\x8b\x2c\x44\x7c\x33\xb8\x82\xf1\x47\x5d\x0a\x0d\x1f\x12\xa7\x8d\xef\x53\x2f\x1a\x14\x3e\x6a\xfa\x9c\xba\xb4\x40\x0b\xd4\x55\x67\xa9\x43\xb7\xd5\x39\xea\xd1\xbc\xba\xa0\xad\xef\xd0\x2d\xea\xd2\x1d\xd0\x95\x90\xb0\xbf\x12\x31\xb4\x07\x0c\xbd\x6a\xba\xbe\x1c\x76\x9d\xfa\x74\x5b\x6f\x48\x9d\x0f\x0d\x37\xd3\x87\x1e\xf6\x7f\xbc\xcb\x75\x35\x1b\x96\xa2\x85\x0c\x6d\x78\xc5\x3a\xbf\xa4\x2e\x2d\x82\xae\x85\xd5\x79\x75\x8e\x16\xc3\xe8\x16\xcd\x1b\x84\xf6\xf4\x72\xe4\xe8\x1c\x75\xd5\xac\x1e\xba\x18\xae\xa0\x83\xbf\x03\x00\x00\xff\xff\x28\x1f\x51\x90\xaa\x08\x00\x00"

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

var _localesRuLc_messagesManagerMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbf\x6b\x14\x5f\x14\xc5\x4f\xe6\x9b\x6f\xc4\x55\x44\x52\x58\x59\x5c\x0b\x83\x22\x2f\xce\xec\x1a\x08\x2f\x99\x44\xcc\x0f\x10\x77\x63\x08\xa3\xfd\x23\xfb\x9c\x0c\xee\xbe\xb7\xbc\x99\x0d\x0a\x11\xa2\x60\x95\x80\x8d\xa2\x90\x46\xf0\x1f\x08\xd1\x85\x11\x93\x6d\xac\xec\xee\x94\x36\xe9\xfc\x3f\x64\x77\x27\x8a\xb7\x39\xe7\xf0\x3e\xef\xdc\x7b\x3a\x39\xfe\x0e\x00\xce\x01\xb8\x0a\x20\x02\x70\x01\xc0\x3e\x46\x93\x03\x38\x0f\xe0\x2b\x80\xff\x01\x7c\x2f\xd9\x1f\xa5\x16\x00\xc6\x01\xfc\x04\x30\x01\xe0\x14\xc0\xd6\x18\xf0\x0b\xc0\x24\x80\x35\x0f\xb8\x04\x40\x79\xc0\x45\x00\x1d\x0f\xb8\x02\xe0\x85\x37\xe2\xdf\x7a\xc0\x65\x00\x07\x1e\x30\x56\xee\xf4\xca\xee\xf1\x32\x4f\x94\xbb\x07\xf3\x5f\xa9\x58\xd6\x1d\x6d\x9a\xda\x64\xa8\xeb\x6d\xdd\x42\xdd\xc6\xb1\x76\xe9\x50\x13\x13\x63\x4d\xb5\x35\xd6\x95\x1b\x10\x1b\xba\x63\x5d\x26\x1a\x69\x9c\x34\xc5\xbd\x6e\x9c\x8a\xc8\x4a\x6a\xea\xed\xbb\x4f\x93\x2d\xd5\xb6\xd3\xae\x5b\x59\x7f\x18\x89\x25\xa7\x55\x96\x58\x23\x96\x55\xa6\x25\x55\xfd\x60\x56\xf8\x35\x51\xad\x51\xb5\x26\x67\x66\x6e\xf9\x35\xdf\xaf\xd4\x55\x9a\x89\xc8\x29\x93\xb6\x54\x66\x9d\xa4\x07\xc3\x0e\x6a\x74\x9d\x6a\xdb\xa6\xa5\xf9\x7f\x8a\x17\x2a\x75\x65\xe2\xae\x8a\xb5\x88\xb4\x6a\x4b\xfa\x93\x25\x6d\x74\xd3\x34\x51\xa6\xd2\xb8\xdf\x58\x11\x8f\xb5\x4b\x13\x6b\x24\x05\xd3\x7e\x65\xc9\x9a\x4c\x9b\x4c\x44\xcf\x3b\x5a\x52\xa6\x9f\x65\xb7\x3b\x2d\x95\x98\x39\xda\xdc\x52\x2e\xd5\x59\xf8\x28\x5a\x15\xb3\x7f\xb9\xc1\x3d\x4f\xb4\x13\x2b\x66\xd3\x36\x13\x13\x4b\xaa\xac\xb7\xba\x4e\xb5\xc4\xaa\x75\xed\x54\x92\xe9\x0c\x63\x1a\xd6\xe6\x68\x64\x43\x73\x3d\xf0\xc3\x30\xa0\xa9\x29\x1a\x58\xff\x5a\x18\x04\xb4\x48\x3e\xc9\x61\x5e\x08\xab\x67\x4f\xf3\xe1\x9d\x81\xbd\x31\xc4\xe6\x03\x9f\x76\x76\x46\x5f\x16\xc2\xaa\x7f\x93\x16\x29\x20\x49\xd5\x39\xf0\x7b\x3e\xe4\x23\xce\x8b\x97\x9c\xf3\x71\xb1\x57\xbc\x06\x7f\x2a\x76\xb9\xcf\x47\xdc\xe3\x93\x62\x1f\x7c\xc0\x7d\xfe\xcc\xbd\x62\xb7\xd8\x3b\x0b\x79\x49\x1c\xf2\x09\xe7\xdc\x03\x7f\xe0\xe3\xe2\x0d\xf8\x23\xf7\xf9\x0b\xe7\xc5\x2b\xee\xf1\xb7\x62\x1f\xbf\x03\x00\x00\xff\xff\x88\xcf\x80\x3d\xac\x02\x00\x00"

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
