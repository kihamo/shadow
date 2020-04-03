// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/manage.mo
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

var _templatesViewsManagerHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\x5f\x93\xdb\xb6\x11\x7f\xf7\xa7\xc0\x20\xd7\x8c\x14\x9b\xa2\xcf\x75\x33\x6d\x22\x9d\x27\x93\x76\xe2\x4e\xdd\x76\x26\x71\xdc\x07\xd7\xbd\x01\x89\xa5\x84\x2b\x05\x28\x00\x28\xe9\xa2\xd1\x77\xef\x00\xe0\x7f\x91\x12\x45\x9e\xee\xea\xdc\x83\x4d\x11\xc0\xee\xe2\xb7\xff\x80\xe5\xee\x76\x88\x42\xc4\x38\x20\x1c\x0a\xae\x81\x6b\x8c\xf6\xfb\x67\x53\xca\xd6\x28\x8c\x89\x52\x33\x2c\xc5\x06\xdf\x3c\x43\x08\xa1\xf2\xdb\xed\xed\x8a\x70\x88\xd3\x91\xc3\x51\xcd\x74\x0c\xa5\x51\x3b\x63\xf1\xea\x66\xb7\x43\xec\xfa\x8f\x1c\xe1\xef\x05\x8f\xd8\x3c\x91\x44\x33\xc1\x31\x9a\xa0\xfd\x7e\xea\x2f\x5e\xd5\x56\x94\x68\x86\x31\x10\x19\xb1\x2d\xbe\x99\xfa\x94\xad\x4b\x8c\x6b\x3f\x2b\x72\x64\xbb\xaa\xd1\x8d\x84\x5c\x22\x29\x62\x98\x61\xf3\x88\x11\x09\x8d\x1c\x33\xfc\x05\x46\x4b\xd0\x0b\x41\x67\x78\x25\x94\xc6\x88\xd1\x99\x81\x26\x62\x73\x85\x11\x25\x9a\x78\x5a\xcc\xe7\x66\xe1\x9a\xc4\x8c\x12\x2d\x64\x8d\x78\x2e\x84\xa3\xaf\x49\x50\x87\xaa\x32\x33\x89\x33\x69\x39\x59\x23\x4e\xd6\x9e\x26\x81\x42\x01\x91\xb7\xe6\x01\x17\x64\x62\xa6\xea\x1b\x29\xff\xed\x76\x48\x12\x3e\x07\x74\xc5\xc9\x12\x5e\xa0\xab\x35\x91\x0a\x7d\x33\x43\x93\x35\x91\x8c\x04\x31\x28\xa3\xdb\xb6\xe5\xd3\x98\xa5\xac\x56\x12\x14\x70\x9d\x6a\xc6\x68\x2c\x42\xf0\x8b\x23\x8b\x52\x30\x8c\x99\x64\x72\x1b\xec\xd6\x80\x77\x3b\x04\x9c\xa2\xfd\xbe\x5d\x44\xcb\x87\xa0\x85\x84\x68\x86\xbf\xd8\xed\x52\x9a\xfb\x3d\x46\x46\x44\xcf\x68\x4b\x8a\x58\xcd\x70\x65\x2c\x47\xa0\xa6\x01\xfb\xc6\x2e\x84\xed\x8a\x70\x0a\xd4\x2e\x6c\x11\x57\xcb\x04\x8c\x8c\xb1\x32\x54\x23\x12\x2b\xc8\x45\x3e\x02\x6b\x09\x5e\x6b\xb9\x8e\xee\x55\x46\xd9\xb3\x42\x64\xb2\x1e\xdf\xb9\x4f\xda\xd9\x4c\xfd\x98\x1d\xd5\xad\x13\xb4\xd9\x86\xfc\x24\xbe\x79\xd6\x3c\x54\xf2\x06\x4d\x02\xaf\xd9\x1f\x6a\xac\x06\x99\x51\x83\xe5\x97\x05\x30\x6f\x8e\x98\x94\xb3\x25\xc4\x78\xa1\x1a\xeb\x82\x65\x7b\x38\x61\x5f\xd5\x1d\xc7\xe0\x49\x50\x2b\xc1\x95\xb1\xd1\xd3\x6a\x9e\xda\x35\x15\x02\xc8\x91\x59\x88\x35\x34\xf9\x7a\x33\x95\x05\x10\xda\x75\xae\xec\x36\x31\x25\x8c\x36\x8c\xea\xc5\x0c\xbf\x7a\xf9\x3b\x5c\xc4\xd3\x7f\x90\x25\x60\x74\x65\xc3\xa8\x5e\xf4\x23\xf8\x87\x32\xc1\x0f\x24\x4e\x86\x52\xbc\xae\x50\xfc\x33\x44\x24\x89\xf5\x50\x9a\x95\x6d\xff\x8b\xe8\x70\x01\x52\xf5\x24\x5a\x16\x4e\x85\x92\xad\x5c\xcc\x3b\x8f\xd6\xd4\xef\xaa\x41\x43\xf3\x0c\xbb\x08\x04\xbd\xef\x36\xb7\x70\x5a\xf6\x02\x5d\x39\x7f\x32\x3e\xeb\x9c\xf7\x44\x5c\x2a\x11\x31\xf3\xff\x4a\xcd\x4a\x09\xab\x98\x84\x90\x11\x9b\x7c\x48\x9d\x7f\xf2\x37\xb8\x47\x78\x82\x11\xbe\xc5\x5d\x09\x9f\x69\xe1\x1d\xf1\x29\x89\xcd\xa2\x43\x39\xff\x42\x99\xf3\xde\x8e\x42\xe6\xfc\x63\x12\x40\x9e\x93\xd3\x8c\xe4\xd9\x97\x18\x45\x42\xba\x58\xe4\x70\x32\xd1\xc9\x66\xa4\xb9\x14\xc9\xca\x25\x1f\x9b\x23\xea\xc2\xfc\x60\xc6\x8b\xa4\x81\x4b\xc1\xcc\x50\x6b\xc2\xd8\x58\xa0\xe5\x7a\x36\x1a\x69\x86\x7b\xb8\x5d\xff\x1f\x6e\xb1\x3d\x19\x36\xee\xce\x3f\xc7\xa8\x9e\xdc\x02\x4b\x19\xcc\x9c\x4e\x1d\xf4\x1d\x33\x4f\xbb\x5c\x6f\x89\xfa\xc0\x60\x83\x30\xf0\x64\xd9\xd9\x77\x2b\x82\x29\x88\x21\xd4\x15\xd9\x52\x53\x41\x6e\xe8\x15\x46\x46\xeb\xce\x47\x9a\x55\x5e\xe4\xf3\xdc\x87\xce\xdf\x19\xaa\xc4\x3c\xf1\x02\x5d\x09\x1b\xba\x4d\xe4\x1a\x65\x8c\x7f\x00\x6d\x36\xfc\x4f\x37\x82\xdd\x0c\x85\xc7\x7d\xb6\x6e\xb7\x9f\xf2\x58\x9b\xcc\xe8\x3c\x81\x53\xd8\xe6\xbc\x5f\x9a\xbd\xe4\x67\x9b\x51\x6d\x70\x7c\x08\x88\x4d\xb1\xe6\xd4\xe3\xc0\x33\xc7\xd7\xec\xa9\x74\x9a\xce\x3c\xae\x46\xf0\x7a\xdc\xe4\x6c\x53\xdf\x0d\xf7\x86\xf4\x4c\xcf\xca\xb1\xf1\x9d\xe4\xbd\x8c\xd4\x06\xac\x26\x4b\xd5\xb0\xd5\x44\x02\xe9\x67\xad\xd9\xea\x46\x7b\x95\xa0\xd8\xaf\x46\x0b\xb7\x05\x13\x29\x36\x6a\x86\x5f\xf7\xb3\xe1\xa6\xd9\x99\x82\xa7\x7e\xc6\x64\x10\x3e\xe6\xc4\x5c\x67\xf1\xfe\x7e\x05\x08\x07\x42\xc4\xfd\x40\x62\x7c\x95\x68\xa4\xef\x57\x30\xc3\xe1\x02\xc2\xff\x06\x62\x9b\x1f\xd8\xef\x94\xa7\x36\x4c\x87\x8b\x5e\x90\x14\x57\xc7\xd1\x4a\x32\xae\x5b\xf0\x19\x23\x6c\xae\x66\xee\x46\x69\x24\x00\x9a\xdb\x21\xf2\x07\x01\x26\x24\x1a\xb5\x83\xc6\xb8\xc6\xe3\x13\x13\xbe\x7e\xdd\x2f\x5e\x54\x70\xe5\xc9\x32\x00\x89\x9b\xec\xb0\x0b\xb0\x45\xbc\x69\x35\xb0\x26\xf4\x57\x44\x6b\x90\x7c\x86\xff\xf3\xd1\x7b\xfe\xe9\xcd\xc7\x97\xde\x9f\x3e\x7d\x75\x85\x2f\x89\x68\x72\x12\xd2\xe4\x37\x83\xe9\x83\x41\xda\x0e\x56\x14\x0b\x62\xd0\xfa\xfc\xc1\x2a\x0c\x70\xf4\xef\x89\x7b\x18\xbf\xb9\x1c\x70\x34\xaf\x2b\x0e\x45\xce\xc4\xed\x27\xc6\x6d\x74\x00\xdc\x47\xe2\xfd\xfa\xe9\xf9\xf8\xf9\x50\x00\x0f\xf2\xed\x8a\x28\xb5\x11\x92\x0e\xc7\xad\xa0\xd4\x94\x7c\xb3\x51\x4f\x2d\xc4\xe6\x72\x48\x3e\x30\x3a\x9a\xcc\x55\x2f\x64\x8c\x50\xd4\x15\x3f\xde\xc3\x56\xdb\x8b\x79\xf3\x19\x35\x9d\xe5\x39\xb3\xbb\x98\xf1\x22\xb7\x95\x4b\xe1\xde\xf7\xec\x69\x40\x2f\xe3\xb4\xdf\xdb\x9b\x67\x19\x94\xd2\x05\xb4\x3c\xb3\xe9\xda\x99\x1f\x22\x06\x98\xc1\x67\x18\x3e\x7a\xee\xf6\xfc\x63\x7f\xed\xbb\x4b\x57\x46\x7d\x8a\x13\x6a\x45\xf8\x89\x03\xb6\x9d\xf2\x79\x97\x11\xec\x36\x0f\x14\x7a\xbb\x76\x85\x58\xeb\x0a\x47\x4c\x22\xad\xae\xf6\xbb\x4f\xb7\x54\x0a\x06\xe5\x83\xaf\xec\xdf\x63\xba\x1e\xca\x77\x42\x38\x3d\x52\x8f\x19\x99\xe1\x51\x73\x5d\x64\x5c\x0c\xd4\xa1\x1d\xf7\xae\x1c\xa0\xa7\x28\x57\x54\x11\xe9\x56\x91\x28\x8c\x68\x10\x4b\x54\xfa\x5e\xd6\xa9\x6c\x31\x74\x83\x3d\xab\x16\x0f\x43\x61\xa0\xc5\xa6\x24\x1e\x5a\x17\xfd\xf7\xd4\x37\x1d\x9c\x17\x82\x1f\xb5\x26\x9b\x7d\x26\xea\x97\x79\xfa\x20\x98\xfa\xfa\xe6\x05\xba\xda\x38\xe6\xe5\x63\x5f\x5f\x79\x72\x99\xb2\x33\x85\x2b\xdb\xdb\x7f\xb3\x73\x92\x2b\x45\xa5\x4c\x27\x3f\x89\x44\x86\x7d\x53\x24\x7a\x34\x6b\xe8\xc7\xe9\x31\x12\xf2\x4d\xeb\x17\x8f\x9f\x15\x99\xc3\xe0\x68\x56\x24\xac\x11\x6f\xf8\xdc\x96\x27\x2d\x2d\x13\x18\xb7\x67\xad\x3e\x86\x14\x48\xe4\x17\x5f\x40\xbf\x8b\x63\xb1\x01\xea\x8e\x9d\xea\x1b\xf7\x15\xf4\x7c\xa2\xc9\x99\x5f\x73\xd0\x53\xe4\xc6\x7a\x65\xf2\x20\x35\x8e\x5b\x46\xae\xfb\x1f\x02\xa6\x31\xbb\x69\xfc\x66\x70\xbc\xf5\xe3\xc4\x36\x86\x64\x9e\x56\x89\xd0\x97\x4b\x4a\xd4\xe2\xdb\xf3\x92\xf8\xb0\x7d\x3c\x66\xae\x42\x79\xd3\xcc\xe5\x79\x5d\x2e\x2f\x76\xef\x3b\x38\x4f\xf2\xa9\xdf\xb1\xf3\x60\xea\xdb\xd0\x74\xa2\x1f\xe7\xf8\x35\xf1\xc4\xf0\xa9\xc6\xa7\xc6\xc5\xe9\xeb\xe6\xae\xbc\x2c\x71\xf2\x5b\x25\x62\x46\x0f\xda\x09\x9b\x26\x9f\xfc\x0a\x5b\x69\x54\x14\xb1\xb7\xa4\xde\xd7\x28\x7d\x10\x51\xa4\x40\x7b\xbf\x3f\x72\x2d\x9b\x06\x89\xd6\x82\xa7\x45\x03\x09\x0a\x8a\xaa\x41\xa0\x39\x0a\x34\xf7\x18\x8f\x44\x5e\x13\x70\x53\x8a\xd8\xfd\xa3\x5b\xe2\xba\x28\x1d\xb1\xae\xdc\x54\x12\x2c\xd9\x21\x3b\x95\x84\x21\x28\x95\x73\x54\x64\x6d\x2e\xa1\x4c\x19\x8d\xd3\x19\xce\x9e\x4a\x42\xfc\x64\xa7\x9c\x96\xe1\xb8\xd6\xaa\xaf\x0c\xf2\x8d\xad\x9e\xe9\x63\xf6\x5f\x09\xff\xa5\xa0\x24\x46\x11\xa1\xe0\xaa\x23\xf6\xf7\xf7\x69\xac\xd2\x24\xb0\xa1\x6c\x86\xbd\xeb\xac\xb1\x90\x32\x12\x8b\x79\xda\x49\x68\x8f\x52\x31\xd0\xe0\xbe\xb2\xf2\x9d\xed\x84\x70\x82\xd8\x79\x0b\x46\x29\xf0\x99\xfb\x74\x75\xd8\x1f\x6b\x97\x7a\x29\xe5\xe6\xe6\x54\x37\xa5\xa5\x41\xf5\x60\xde\x02\x08\x6d\x6c\x3f\xab\x6a\xd3\xfd\xc0\x45\xd3\xac\x50\x59\xed\x80\x32\xb5\x64\x39\x41\xdc\xb4\x8d\x2f\x35\x5b\x82\xfa\xb6\x5d\x7d\xd3\xc5\xeb\xaa\x58\xae\xcf\xb7\x8e\xb3\x43\xab\xd6\xec\x2b\x97\xc8\x18\x11\x0a\x1b\x3b\x7f\x5f\xd7\x35\x7f\x68\x0c\x07\x98\x98\x28\x55\xe2\xf2\x9d\x04\x74\x2f\x12\xa4\x12\x09\x6f\x32\xc2\x1d\xc8\x44\x42\xe8\xf3\xa1\xcd\x1c\x25\x3b\x75\x37\x82\x5c\x82\xc0\x69\xe2\xb8\x77\x74\xf2\x4c\x6a\x4e\x4a\xf2\x14\xbf\x53\xce\xd8\xde\x44\x5d\xf5\xac\x22\x00\x3f\x2b\x35\x8b\x1b\x6b\xcc\x8b\x42\xbb\x1d\x52\x9a\x68\x16\xbe\x7d\xff\xf7\x77\x68\xe4\x9e\x7f\xfe\xf1\x1d\xc2\xbe\x39\x45\x04\x82\x48\xea\x13\xa5\x40\x2b\x7f\x0d\x9c\x0a\xa9\x7c\x23\xbc\x4d\x1e\x6a\xc2\x41\x7b\x81\xf2\x43\xe5\xde\xbe\x77\x6f\x03\x21\xb4\xd2\x92\xac\x26\x4b\xc6\x27\xa1\x89\x44\xb6\x65\x77\xfc\x80\x5c\x23\xb6\x05\xea\x1c\x2b\x93\xc0\xbe\x7a\x6b\x5f\x1d\x17\xa1\x19\x97\x3b\xf5\x80\xa8\xf8\x77\xca\xbf\xfb\x25\x01\x79\x3f\x29\x01\x63\x64\xb9\xbb\x04\x1a\x81\x32\x0c\x5b\x55\x70\x11\x9e\x25\x0d\xd4\x98\x97\x15\x71\x36\x7b\x17\x62\x32\xde\x77\xca\x5f\x12\x4e\xe6\x20\x2d\x15\x77\x91\x2a\xab\xf0\x7f\x01\x00\x00\xff\xff\x50\x8d\xf5\x9c\x06\x31\x00\x00"

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

var _assetsJsManagerMinJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\xcd\x6e\x23\x37\x0c\x7e\x15\x9b\x1b\x64\x25\x78\x56\xde\x1c\x0a\x14\xb6\xe5\x1c\xdc\x3d\x14\x58\x14\x05\xb2\xed\x25\x08\x0a\xce\x88\xf6\xa8\xd1\x48\x53\x49\x63\x6f\x10\xf8\xdd\x0b\x8e\x27\x5e\x27\xde\x04\x6d\x2f\xbd\x18\x34\x7f\x3e\x8a\xfc\x48\xce\x85\x30\xa1\xea\x1a\xf2\x59\xaa\x48\x68\x1e\xc4\xba\xf3\x55\xb6\xc1\x0b\xf9\x78\x21\xe0\x5d\x15\xfc\xda\x6e\xd2\xc8\xfa\xb6\xcb\xb7\xd6\xdc\x15\xa3\xa3\x2e\x91\xa3\x2a\x9f\x28\x32\x7d\xcd\x18\x09\xd9\x0f\xa4\xaa\x6a\xf4\x1b\x3a\x45\xdc\x62\x1c\x91\xbe\x10\xb9\xb6\x49\x16\x31\xec\x34\xa9\x16\x23\xf9\x9c\x7e\xf3\xd9\x3a\x01\xb9\x0c\xe6\x01\x0a\xc8\x11\x64\x51\x75\x91\x6d\x9a\xd4\x16\x9d\x90\x85\xa1\xb5\xe6\x50\x65\x68\x8d\x9d\xcb\xbf\xa3\xeb\x68\x0e\x55\x4d\xd5\x7d\x19\xbe\x82\x66\xb8\x18\x5a\x01\xf9\xa1\x25\x90\xd7\xe2\x09\x01\x60\x7c\xb4\xf5\xee\x64\xe0\x1c\x6f\x75\xb0\xc8\x19\xdc\x7c\xfa\xfc\x69\xf5\x05\xf4\xc1\x9c\x71\xf3\x0b\x36\x74\x79\x29\x38\x82\xd4\xda\x7a\x23\x20\xb4\x5c\x15\x48\xb5\xb6\x2e\x53\x3c\xad\x33\x52\xee\xa2\x1f\x0d\x85\x0e\x79\x87\x24\x37\x7d\xd7\x38\xff\x5e\x1e\xea\xfa\x56\xa8\x36\xb4\xbe\x8e\x61\xa7\x22\x35\x61\x4b\x2b\x87\x29\x09\xa8\x31\x7d\xa0\x18\x43\x04\x39\x63\x23\x1a\x73\x6e\x29\x4e\xd9\x2a\xbb\x9c\x83\xbf\xe5\x2e\xe8\xd4\x95\x8d\xcd\xcc\xc7\xf0\x0c\x9b\xb0\x74\x64\xa0\xf8\xa8\xf5\x69\x50\x8e\xea\x04\x4f\x39\xf2\x9b\x5c\xcb\xfd\xeb\xc8\x91\x12\xf5\xc0\x95\xb3\xd5\xfd\x6b\x93\xf3\x02\xf6\xb5\xca\xfe\xc3\xfb\xc7\x1f\xdf\x7a\xdd\xb7\xb8\xb3\xe7\xf1\x18\x36\x7d\xed\x4d\x30\xe8\x56\x7d\x34\xcf\x9b\x06\x98\x1f\xb9\xfb\x7e\x09\xa3\x6c\x86\x65\xf0\xd8\xd0\x5d\x9f\x6b\xac\x6b\x6b\x0c\xf9\xd3\xe5\x78\x19\x73\xbe\x2c\x2f\x1c\x9e\x96\x07\xa4\x22\xac\xea\x57\xd7\xe6\xad\x79\xdf\x9e\x0f\xf9\x6c\x7b\x5c\x9f\x6a\xa2\x61\xe1\xec\x72\x91\x72\x0c\x7e\xb3\x84\xc9\x93\x37\x97\x02\x72\x02\x8b\xe9\x60\x9a\x8d\x60\xb2\xe5\xff\xce\x2e\x61\x2f\x8b\x66\x98\x79\xd5\x37\xec\x43\xbf\xa4\x52\xd5\xb9\x71\x02\x56\xfd\x9a\xa7\xd1\x61\x1f\xd2\x6c\x51\xc6\xd1\x74\xb9\xe8\xdc\x12\x26\x15\x63\xb0\xc4\x10\x7d\xac\x90\xc5\xf8\x6a\xa0\xed\xa4\xfb\xff\x86\x3a\x83\x19\xf5\xe3\xfe\x9f\x33\x65\xcd\xff\xc1\x13\x3f\xf3\xf6\x79\x87\xef\xf4\x1b\xe4\x9d\x53\x37\x10\xc7\xdd\x52\x6d\x48\x59\xc0\x3b\x28\x18\xb6\x38\x49\xba\xb3\xde\x84\x9d\x72\xa1\x42\xd6\xa8\x48\x2e\xa0\xe9\xa3\x8e\x8d\x7e\x7e\xc5\xfb\x1e\xb7\x98\xd2\x2e\x44\x73\xa7\x9e\xa4\x0f\xa9\x0e\x3b\x90\x2a\x78\x01\x2c\xaa\x32\x1d\x6d\xa3\xda\x1a\x3a\x55\x40\xf1\x9d\xba\xe1\x1d\x4c\x9e\xdf\x3c\x6b\x78\xac\xfe\xd8\xf2\xa1\x06\x39\xb7\x6b\x41\x4f\x87\xa5\x8f\xe2\xdb\xce\x0d\x15\x72\x3e\x08\xa4\xb8\x40\x01\x43\x88\x2c\x9e\x2b\x8a\x2c\xf7\x7b\x39\x7f\xc6\x39\x9f\x03\x90\xea\x27\xcc\xf8\x85\x65\xf1\xe8\xd0\x6f\x3a\xdc\xd0\xec\xb1\x8b\x6e\x06\x53\x83\xa9\x2e\x03\x46\x33\x65\xac\x3e\x20\x4d\xed\xd5\x8f\x5e\xfd\x99\x82\xbf\xe6\xde\x39\xd2\x30\x19\x7a\x99\x6a\x34\x61\xf7\xb9\xd7\xee\x8b\xf2\x57\xdc\x58\x8f\x99\x66\xe3\xab\xa2\xfc\xd9\xaf\x03\x0b\x88\x37\x21\x66\xeb\x37\xb3\xdb\xbb\xc2\x44\xdc\xad\xd0\xb9\x12\xab\xfb\xd9\x8b\xc6\x60\x6b\x0f\x5f\x11\x6c\xad\xe8\x3f\x7a\x49\x63\x6b\x15\x0b\xe2\xb1\xe5\x57\xc2\x70\xff\x61\x2f\x95\x0f\x86\x92\x90\x85\xc3\x94\xb5\xef\x9c\x9b\xb3\x73\x15\x5c\xd7\x78\xf1\xb1\x38\x0f\xe8\xbb\xf3\x72\x16\x63\xd8\x15\xf6\x90\x9f\x27\x4f\x5f\xb0\x66\xf0\x85\x4d\x0c\x5d\x0b\x72\xce\x29\xc6\x5a\xfb\xfe\xd3\xc6\xbf\x03\x35\x97\x97\xbd\x7b\x92\x8a\xfe\x12\x56\xaa\x92\xd6\x21\x92\x78\xbf\xc8\x71\x54\xf1\xe1\xd6\x03\xc4\x72\x91\xcd\xa8\x0a\x2e\xb5\xe8\x35\xfc\x00\xcb\xf7\x13\x86\xe1\x9d\xcf\x66\xb9\x98\xe6\xc8\x8b\x7f\xa8\x04\x1b\xda\x33\x75\x7b\x39\xff\x3b\x00\x00\xff\xff\x3e\x40\x31\xc0\x75\x08\x00\x00"

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

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x4d\x6b\x13\x4f\x1c\xc7\xbf\xcd\x3f\x7f\x84\x80\x88\x7a\x51\x10\xfc\x79\xb0\xd4\xc3\xd4\xdd\xe8\x41\x26\xd9\x44\xfa\x04\x62\x03\x52\xa2\xf7\x49\x32\xdd\x2e\xee\xce\x2c\x33\xb3\x45\xa1\x07\x51\x11\x3d\x79\x12\xc1\x93\x57\x6f\x15\x9f\x82\xda\xbe\x05\x67\xdf\x80\xaf\x45\xb2\x5b\x6d\x3c\xf8\x3b\x0c\xdf\x87\xcf\xcc\xfc\x7e\x9e\x69\xbe\x02\x80\xff\x01\x5c\x00\xb0\x06\xe0\x04\x80\x14\xf5\x3c\x07\x70\x12\xc0\x0b\x00\xa7\x01\xbc\x06\x70\x16\xc0\xbb\x23\x7f\x00\xa0\xbf\x00\xfc\x00\x70\x0e\xc0\xa9\x06\x70\x1e\xc0\x52\x03\xb8\x08\x60\xa5\x51\xe7\xa3\x06\xb0\x00\xe0\xbf\xa3\xbf\x9a\x00\x1a\x98\x9b\x55\xad\xb6\x93\xb8\x30\xc2\x25\x5a\x61\x5c\xb9\xe6\x9a\x1c\x15\x31\x65\x7a\x22\x8f\x93\x5d\x99\xea\x7c\x3e\x63\x4e\x8c\x9a\xb5\xc4\x96\xcc\xb5\x71\x6c\x60\xe3\x64\xc2\x56\x8a\xd8\xb2\xa1\xe6\x34\x91\xbb\x37\xef\x27\x3b\x22\xd3\xcb\xa6\x68\x6d\x0a\xeb\xd8\xd0\x08\x65\x53\xe1\xb4\xe1\x74\xbb\xaa\x68\x50\x18\x91\xe9\x89\xa6\xee\x5f\x7c\xaf\xb5\x29\x54\x5c\x88\x58\xb2\xa1\x14\x19\xa7\x3f\x9e\xd3\x56\x61\x6d\x22\x54\x6b\x70\x6b\xb0\xce\xee\x49\x63\x13\xad\x38\x85\xcb\x41\x6b\x55\x2b\x27\x95\x63\xc3\x87\xb9\xe4\xe4\xe4\x03\x77\x35\x4f\x45\xa2\x3a\x34\xde\x11\xc6\x4a\x17\xdd\x1d\x6e\xb0\x1b\xc7\xdc\x6c\x9f\x6d\x69\xd8\xba\x1a\xeb\x49\xa2\x62\x4e\xad\x3b\x69\x61\x44\xca\x36\xb4\xc9\x2c\x27\x95\x57\xd6\x46\xd7\x3a\x54\xcb\x48\x5d\x0e\x83\x28\x0a\x69\x71\x91\x66\x32\xb8\x14\x85\x21\xf5\x29\x20\x5e\xf9\x5e\xd4\xfe\x5d\x75\xa3\xeb\x33\xb9\x54\x61\xdd\x30\xa0\xbd\xbd\xfa\x4a\x2f\x6a\x07\x57\xa8\x4f\x21\x71\x6a\x77\xe0\xdf\xf8\x43\x7f\x50\x3e\xf5\x53\xff\xa1\x7c\x52\x3e\xf2\xfb\xe5\x33\x3f\x2d\x5f\xc2\xbf\xf5\x9f\xfc\x67\x3f\xf5\xdf\xc9\x1f\x96\x8f\xfd\x37\xbf\xef\x3f\xfa\xaf\x7e\x3a\xdf\xcc\x78\xff\xa5\x3a\xdf\x57\x54\xd5\xff\xeb\xc9\x5f\x01\x00\x00\xff\xff\xf4\x20\xc4\xc0\x7b\x02\x00\x00"

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

var _localesRuLc_messagesManageMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x4f\x4b\x6f\xf3\x44\x14\x3d\x6e\x3e\x5e\xf9\x80\x0f\x2a\x24\x24\x84\xd0\x74\x41\x05\x0b\x97\xb8\xb0\x40\x6e\xdd\x50\xfa\x10\x88\x06\xa1\x12\xca\x96\x51\x32\x4d\x2c\x1c\x3b\x9a\xb1\x03\x95\x2a\x91\x34\x85\x0a\x82\x28\xe2\xa1\xae\x00\xf5\x1f\x84\xd2\x50\xab\x6d\xd2\x15\xfb\x3b\x42\x62\x85\xf8\x21\xac\x90\x1f\x7d\x30\x8b\xb9\xf7\x9c\x7b\xee\x99\x33\xff\x4c\xdf\xfb\x11\x00\x1e\x07\xf0\x3c\x80\x5d\x00\x4f\x03\xf8\x1b\xd9\x61\x06\xf0\x00\xc0\x8c\x01\x3c\x01\xc0\x32\x80\x87\x00\x94\x0d\xe0\x39\x00\x6b\x06\xf0\x08\x80\x8f\x0c\xe0\x3e\x80\xa6\x01\xdc\x03\x10\xe6\xba\x9d\x1c\x7f\x96\xe3\x7d\x03\x78\x14\xc0\x57\x46\xe6\xf1\xbd\x01\xcc\x02\xf8\x73\x0a\x78\x06\xc0\xbf\x53\xc0\x93\x00\x1e\x14\x80\xb7\x00\xbc\x50\x00\x9e\x05\xd0\x28\x00\x4f\x25\x3e\x79\xfd\x36\xaf\xc7\x05\x60\x1a\xc0\x49\x8e\xff\xc8\xeb\x5f\x05\xc0\x40\x96\xed\xfa\x24\xef\x17\xf2\xfe\x7e\x5e\x1f\xcb\xff\x9e\x64\x7a\x18\x59\xd6\x62\x3e\x9b\x4a\xae\x65\xcf\x0b\x3e\x11\x75\xd6\xe1\x5e\x24\x94\x8d\x65\x29\xd8\x4e\x10\x31\x15\x49\x51\xc6\x8a\x17\x28\x81\x95\xc0\xdf\x76\x65\x8b\x29\xde\x11\xac\x96\x80\x46\x24\x79\xe8\x06\x3e\x56\xc5\x36\x8f\xbc\x10\xab\x42\xd5\xa4\xdb\x4e\xb9\x77\x79\x4b\x60\x53\x28\x11\xe2\x7d\xde\x11\xd8\x4a\xac\xf1\x21\x0f\x6b\x4d\x21\x15\x36\x45\x3b\x90\xa1\x59\x51\x0d\xb7\x6e\xbe\x19\x35\x94\x59\x0d\x6c\x56\x17\x9d\x37\x3e\x76\x9b\xbc\x15\xcc\xc9\xa8\xb8\xc1\x55\x68\x56\x25\xf7\x95\xc7\xc3\x40\xda\xec\x9d\x74\xc4\x2a\x91\xe4\xad\xa0\x1e\xb0\xc5\xff\xe9\x97\x8a\x1b\xdc\x6f\x44\xbc\x21\xcc\xaa\xe0\x2d\x9b\xdd\x60\x9b\x6d\x46\x4a\xb9\xdc\x2f\x56\xde\xae\xac\x99\x5b\x42\x2a\x37\xf0\x6d\x66\xcd\x95\x8a\x2b\x81\x1f\x0a\x3f\x34\xab\x3b\x6d\x61\xb3\x50\x7c\x1a\xbe\xd2\xf6\xb8\xeb\x2f\xb0\x5a\x93\x4b\x25\x42\xe7\x83\xea\xba\xf9\xfa\xad\x2e\xc9\xb3\x2d\xa4\xb9\xe6\xd7\x82\xba\xeb\x37\x6c\x56\x7c\xcf\x8b\x24\xf7\xcc\xf5\x40\xb6\x94\xcd\xfc\x76\x0a\x95\xf3\xea\x02\xcb\x5a\xc7\x7f\xd1\x2a\x39\x8e\xc5\x66\x67\x59\xd2\x96\x66\x1c\xcb\x62\x65\x56\x62\x76\x8a\x97\x9c\xf9\xeb\xd1\xa2\xf3\x5a\xd2\xbe\x94\xca\x16\xad\x12\xdb\xdd\xcd\x56\x96\x9c\xf9\xd2\xcb\xac\xcc\x2c\x66\xb3\xf9\x05\xd0\x0f\x34\xa1\x2b\xdd\xd7\x3d\xbd\x47\x31\x5d\xea\x01\x8d\x18\x9d\xd1\x98\x86\xfa\x80\x46\x34\xa6\x58\x1f\xda\xa0\xef\xf4\x80\xe9\x3e\x9d\xd0\x48\x77\x13\x5a\x0f\xca\xa0\x23\x1a\xd2\xb9\xee\xea\x81\xde\xd3\x5f\x83\x7e\xa1\x09\x9d\xea\xbd\x5c\xf4\x3b\x9d\x66\xfb\x34\x62\xba\x47\x13\xfd\xb9\xee\xd2\x90\xc6\x37\x24\x9d\xd3\x84\xc6\x7a\x9f\x62\xfa\x4d\xf7\x93\xa1\xfe\x82\x62\x8a\x53\x23\x53\xf7\xe9\x92\x26\x74\xa1\x0f\xd2\xa5\x58\x7f\x03\xfa\x99\xae\x28\xd6\xbd\x8c\xa0\x11\xe8\x27\x1a\xd2\x19\x9d\xdc\x12\xc7\xf4\xab\xee\xd2\x44\xf7\x28\xce\x32\x1d\xdf\x79\x38\xa7\x8e\xee\x7e\x2e\x5b\xba\xa0\x11\x9d\xea\x43\xfd\x65\x4a\xfc\x17\x00\x00\xff\xff\xa4\x0a\xee\x46\xe7\x03\x00\x00"

func localesRuLc_messagesManageMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesManageMo,
		"locales/ru/LC_MESSAGES/manage.mo",
	)
}

func localesRuLc_messagesManageMo() (*asset, error) {
	bytes, err := localesRuLc_messagesManageMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/manage.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"locales/ru/LC_MESSAGES/manage.mo": localesRuLc_messagesManageMo,
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
				"manage.mo": &bintree{localesRuLc_messagesManageMo, map[string]*bintree{}},
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
