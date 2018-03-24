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

var _templatesViewsManagerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\xeb\x73\x1b\xb7\x11\xff\xee\xbf\x02\x83\xa8\x19\x32\xf6\xf1\x2c\xd7\xcd\xb4\x09\x29\x4f\x26\xcd\xc4\x9d\xba\xed\x4c\xe2\xb8\x1f\x5c\x57\x83\x3b\xec\x91\x50\x8f\x00\x03\xe0\x48\xa9\x1c\xfd\xef\x1d\x00\xf7\x3e\x1c\x1f\x47\x51\xaa\xe3\x0f\xf2\x11\x8f\x7d\xfc\xb0\x8b\xc5\x02\xbb\xdd\x22\x0a\x09\xe3\x80\x70\x2c\xb8\x06\xae\x31\xba\xbf\x7f\x36\xa5\x6c\x8d\xe2\x94\x28\x35\xc3\x52\x6c\xf0\xd5\x33\x84\x10\xaa\xb7\xde\x5e\xaf\x08\x87\x34\xef\xe9\xf6\x6a\xa6\x53\xa8\xf5\xda\x11\x8b\x57\x57\xdb\x2d\x62\x97\x7f\xe4\x08\x7f\x2f\x78\xc2\xe6\x99\x24\x9a\x09\x8e\xd1\x04\xdd\xdf\x4f\xc3\xc5\xab\xd6\x8c\x1a\xcd\x38\x05\x22\x13\x76\x8b\xaf\xa6\x21\x65\xeb\x1a\xe3\xd6\xcf\x86\x1c\x85\x56\x4d\xba\x46\x8a\x04\x4d\x40\x4a\x21\x8d\xbe\x7d\x3c\x49\x0a\x52\x23\xfb\x37\xa0\x84\xcf\x41\x16\x3f\x98\x5a\x32\xa5\x48\xd4\x51\xd2\x92\x88\x32\xad\x05\x47\xfa\x6e\x05\x33\xec\x7e\xe0\x4a\x0f\xa1\x00\x23\x4a\x34\x29\xc8\xe4\x8c\x30\x22\x92\x91\x60\xc1\x28\x05\x3e\xc3\x5a\x66\x80\xaf\xbe\xd4\x6c\x09\xea\xdb\x69\xe8\xc8\x74\x99\x6d\xb7\xb9\x22\x93\x1f\xbc\xea\x34\xd1\xc9\x67\x00\xa7\x66\x60\x73\x64\x22\xe4\x12\x49\x91\xc2\x0c\x9b\x4f\x8c\x48\x6c\x16\x67\x86\xbf\xc0\x68\x09\x7a\x21\xe8\x0c\xaf\x84\xd2\x18\x31\x3a\x33\xf6\x92\xb0\xb9\xca\x35\xd1\x62\x3e\x37\x13\xd7\x24\x65\x94\x68\x21\x7d\xb0\x18\x64\x1d\x7d\x4d\xa2\xb6\xfd\x34\x46\x66\x69\x01\x17\x27\x6b\xc4\xc9\x3a\xd0\x24\x52\x28\x22\xf2\xda\x7c\xe0\x8a\x4c\xca\x54\x7b\x75\x5b\xba\x4a\xb3\x70\xe8\x82\x93\x25\xbc\x40\x17\x6b\x22\x15\xfa\x66\x86\x26\x6b\x03\x76\x94\x82\x6a\x23\xd6\x10\x24\x65\x39\xab\x95\x04\x05\x5c\xe7\xe6\xea\x0c\x08\x7e\x75\x64\x51\x0e\x86\xf1\x9d\xd2\x74\x62\xcd\xd6\x80\x4b\xac\xfb\x45\xb4\x7c\x08\x5a\x48\x48\x66\xf8\x8b\xed\x36\xa7\x79\x7f\x9f\xdb\x83\x31\x61\x29\x52\x35\xc3\x8d\xbe\x12\x81\xd6\x0a\xd8\x16\x3b\x11\x6e\x57\x84\x53\xa0\x76\x62\x8f\xb8\xc6\xc8\x8c\x8c\xa9\x32\x54\x13\x92\x2a\x28\x45\xde\x01\x6b\x0d\xde\x42\xa2\xdd\xfa\x85\xa4\x9f\xd8\x34\x4c\xd9\xce\x15\xcc\xad\xd5\x3f\x37\x4b\xaf\x9e\xf9\xbb\x6a\x8e\xac\x49\x14\xf8\xb7\x82\x16\xab\x93\x8c\xc5\x63\xdf\x75\x01\x4c\xcb\x0e\xc3\x71\x16\x83\x18\xaf\x16\xc0\x3a\x5a\x7d\xd5\xf7\x58\x51\x53\xe3\x14\x02\x09\x6a\x25\xb8\x32\x96\xb8\x7f\x31\xa7\x76\x4e\x83\x00\x72\x64\x16\x62\x0d\x3e\x8f\xf6\x53\x59\x00\xa1\x87\x8e\x95\x87\x0d\xcc\x09\xa3\x0d\xa3\x7a\x31\xc3\xaf\x5e\xfe\x0e\x57\xa1\xe4\xef\x64\x09\x18\x5d\xd8\x08\xa2\x17\xc3\x08\xfe\xa1\x4e\xf0\x03\x49\xb3\x53\x29\x5e\x36\x28\xfe\x19\x12\x92\xa5\xfa\x54\x9a\x0d\xb5\xff\x49\x74\xbc\x00\xa9\x06\x12\xad\x0b\xa7\x62\xc9\x56\x6e\x67\x3b\x8e\xd6\x34\x3c\x74\x05\x0d\xcd\x23\xec\x22\x12\xf4\xee\xb0\xb1\x95\xd3\xb2\x17\xe8\xc2\xf9\x93\xf1\x59\xe7\xbc\x7b\xf6\xa5\x1a\x11\x33\xfe\x2f\xd4\xcc\x94\xb0\x4a\x49\x0c\x05\xb1\xc9\x87\xdc\xf9\x27\x7f\x85\x3b\x84\x27\x18\xe1\x6b\x7c\x28\xe1\x23\x2d\xfc\x40\x7c\x6a\x62\xb3\xa4\x2b\xe7\x0f\x94\x39\xef\x3d\x50\xc8\x92\x7f\x4a\x22\x28\x23\x6f\x1e\x77\x02\xdb\x88\x51\x22\xa4\xdb\x8b\x1c\x4e\x66\x77\xb2\x71\x67\x2e\x45\xb6\x72\x21\xc6\x58\x53\x47\x98\x1f\x4d\x3f\xba\xa8\xb6\xba\x6a\x33\x33\xd4\x7c\x18\x1b\x0b\xb4\x5c\x8f\x46\x23\x8f\x63\x0f\xa7\xf5\xff\xa1\x8a\xfd\xc1\xd0\xab\x5d\x78\x8c\x51\x3d\xb9\x05\xd6\x22\x98\x39\x83\x3a\xe8\x0f\x8c\x3c\xfd\x72\xbd\x25\xea\x03\x83\x0d\xc2\xc0\xb3\xe5\xc1\xbe\xdb\x10\x4c\x41\x0a\xb1\x6e\xc8\x96\x9b\x0a\x72\x5d\xaf\x30\x32\xab\xee\x7c\xc4\xbf\xe4\x55\x3c\x2f\x7d\xe8\x78\xcd\x50\x63\xcf\x13\x2f\xd0\x85\xb0\x5b\xb7\xd9\xb9\x46\x05\xe3\x1f\x41\x1b\x85\xff\xe1\x7a\xb0\x1b\xa1\xf0\x78\x88\xea\x56\xfd\x9c\xc7\xda\x44\x46\xe7\x09\x9c\xc2\x6d\xc9\xfb\xa5\xd1\xa5\x3c\xdb\x8c\x5a\x9d\xe3\x2e\x20\x36\xc4\x9a\x53\x8f\x03\xcf\x1c\x52\x8b\xaf\xda\x99\xb9\xc3\xe7\xd2\x3a\x8e\xfb\x31\x18\xba\x23\x3d\xa8\xc4\x20\x74\x12\x0e\x32\x46\xbb\x31\xe5\x27\xbf\x36\x16\xef\xef\x56\x80\x70\x24\x44\x3a\xcc\x34\x19\x5f\x65\x3a\x4f\x36\xe3\x05\xc4\xff\x89\xc4\x6d\x79\xf0\xbc\x51\x81\xda\x30\x1d\x2f\x06\x99\x67\x95\xe8\x8c\x56\x92\x71\xdd\xb3\x90\x63\xe4\xb2\x55\x9b\xff\x18\x09\x80\x96\x38\xa3\xf0\x24\xc0\x84\x44\xa3\x7e\xd0\x18\xd7\x78\xbc\x67\xc0\xd7\xaf\x87\xd9\x7d\x03\x57\x9e\x2d\x23\x90\xd8\xe7\xff\x87\x00\x5b\xf9\x4d\xaf\x27\xf8\xd0\x5f\x11\xad\x41\xf2\x19\xfe\xf7\xc7\xe0\xf9\xa7\x37\x1f\x5f\x06\x7f\xfa\xf4\xd5\x05\x3e\x27\xa2\xd9\x5e\x48\xb3\xdf\x0c\xa6\x0f\x06\x69\x3f\x58\x49\x2a\x88\x41\xeb\xf3\x07\xab\x32\xc0\xd1\xbf\x26\xee\x63\xfc\xe6\x7c\xc0\xd1\xf2\x6a\xf0\x54\xe4\x34\xdc\xea\x27\xc6\x6d\xd4\x01\xee\x23\x09\xfe\xfb\xe9\xf9\xf8\xf9\xa9\x00\x76\x4e\x38\x2b\xa2\xd4\x46\x48\x7a\x3a\x6e\x15\x25\xdf\xa1\xa7\xe8\x0d\xd4\x42\x6c\xce\x87\xe4\x03\xa3\xa3\xc9\x5c\x0d\x42\xc6\x08\x45\x5d\x12\xff\x1e\x6e\xb5\x4d\x30\xfd\x67\xad\x7c\x54\xe0\xcc\xee\x6c\xc6\x8b\x9c\x2a\xe7\xc2\x7d\xe8\xd9\xca\x80\x5e\xc7\xe9\xfe\xde\x5d\x74\xd7\x40\xa9\x25\x52\xf5\x91\xbe\xf4\xa9\x3c\x44\x9c\x60\x06\x9f\xe1\xf6\x31\x50\xdb\xe3\x8f\xb5\x9e\xc7\x81\x43\x18\x0d\x49\xb2\xd5\x8a\x70\x6f\x36\x5c\xe0\x30\x0d\xed\x90\xcf\x3b\x1d\xb6\x6a\x76\x16\xf4\x7a\xed\x2e\x14\xad\x2b\xec\x30\x89\xfc\x96\x70\x58\x5e\xd8\x93\xf1\x9e\x14\x0f\xbe\xb2\xff\x1e\xd3\xf5\x50\xa9\x09\xe1\x74\xc7\xbd\xc2\xc8\x74\x8f\xfc\xf9\xfd\xb8\xea\x68\x43\x3b\x1e\x9c\x01\xa3\xa7\x48\xbb\x9b\x88\x1c\x96\x59\x57\x46\x74\x12\xcb\x82\x6d\x27\xfd\x3e\x55\x91\x81\xd9\xf7\xc3\x50\x38\xd1\x32\x73\x12\x0f\x8d\xf9\x70\x9d\x86\x6e\xfb\xc7\x6d\xb5\x8f\x7a\x87\x58\x3c\x6b\x0c\x8b\x30\x43\x10\xcc\x7d\x7a\xf3\x02\x5d\x6c\x1c\xf3\xfa\xf1\x6e\xa8\x3c\xa5\x4c\xc5\xd9\xc1\x5d\x33\xdb\xbf\xc5\x79\xc8\xdd\x0f\xe7\x4c\x27\x3f\x8b\x4c\xc6\x43\x43\x21\x7a\x34\x6b\x18\xc6\xe9\x31\x02\xef\x55\xef\x0d\xfd\x2f\x8a\x98\x25\xee\x1e\x31\x07\x18\xaa\x8d\x3c\xdc\xf3\x3c\x54\x06\x27\x2d\x33\x18\xf7\x47\xa7\x21\x86\x14\x49\x14\x56\x2f\x76\xdf\xa5\xa9\xd8\x00\x75\xc7\x4b\xf5\x8d\x7b\xb5\x3b\x9e\x68\x76\xe4\xeb\x03\x7a\x8a\x18\xb8\x23\xf6\xb5\x5a\x2e\x87\x5f\x6f\xa7\xac\x7b\xd7\xfc\xd2\x3d\xd2\xec\x28\x4d\xd8\x23\xf6\x29\x91\xa6\x57\x22\xf4\xe5\x92\x12\xb5\xf8\xd6\x1f\x9c\x4f\x93\xf7\x31\x63\x10\x2a\x8b\x37\xce\xcf\xeb\x7c\xf1\xee\xf0\xf7\xef\xe3\x24\x9f\x86\x07\xbe\x80\x4f\x43\xbb\xe5\xec\xa9\x0b\xd9\x9d\xe6\xed\xe9\xde\x57\x80\xe3\x9d\x9c\x37\xfb\x6b\xc0\x8a\x80\xc8\xaf\x95\x48\x19\xed\x54\xf4\xf9\x06\xef\x7d\x0d\x6c\xd4\x0a\x8a\x34\x58\xd2\xe0\x6b\x94\x7f\x88\x24\x51\xa0\x83\xdf\xef\x48\xab\x9a\x35\x7b\x12\x14\x54\x59\x7f\xa4\x39\x8a\x34\x0f\x18\x4f\x44\x99\xd3\xbb\x21\xd5\x9e\xfc\x93\x9b\xe2\x0a\x19\xfb\x8a\xf5\xfc\xdc\x54\x16\x2d\x59\x97\x9d\xca\xe2\x18\x94\x2a\x39\x2a\xb2\x36\x49\x24\xb3\x65\x87\x74\x86\x8b\xaf\x9a\x10\x3f\xdb\x21\xfb\x65\xd8\xbd\x6a\xcd\x26\x83\xbc\xb7\xda\x32\xff\x2c\xfe\xab\xe1\xbf\x14\x94\xa4\x28\x21\x14\xdc\xed\x86\xfd\xfd\x7d\x1e\x7a\x35\x89\xec\xc6\x35\xc3\xc1\x65\x51\xc6\x46\x19\x49\xc5\x3c\xaf\x5b\xb3\x47\xa4\x14\x68\x74\xd7\x98\xf9\xce\xbe\xc8\x3b\x41\x3c\x85\x92\xcf\xda\x46\x60\xa7\x06\x39\x65\x7f\x7d\xa8\x1b\xe2\x2f\x0c\xeb\x8e\x5b\x00\xa1\xde\x32\xa8\xe3\xeb\x3d\x2d\xc1\x61\xf5\x9e\xd3\xc5\xeb\xa6\x58\xae\xd4\xb6\x8d\xb3\x43\xab\x55\x6f\x2b\x97\xc8\x18\x11\x8a\xbd\xc5\xb7\xaf\xdb\x2b\xdf\x35\x86\x0e\x26\x66\x97\xaa\x71\xf9\x4e\x02\xba\x13\x19\x52\x99\x84\x37\x05\xe1\x03\xc8\x24\x42\xe8\xe3\xa1\x2d\x1c\xa5\x38\x4d\x7b\x41\xae\x41\xe0\x56\x62\xb7\x77\x1c\xe4\x99\xae\x16\x78\x1f\xbf\x7d\xce\xd8\x5f\xc7\xdc\xf4\xac\x5a\xbd\x6e\xad\x5e\xdb\x58\x63\x79\xa9\xb3\xdd\x22\xa5\x89\x66\xf1\xdb\xf7\x7f\x7b\x87\x46\xee\xfb\x97\x9f\xde\x21\x1c\x9a\xd3\x42\x24\x88\xa4\x21\x51\x0a\xb4\x0a\xd7\xc0\xa9\x90\x2a\x34\xc2\xdb\xe0\xa1\x26\x1c\x74\x10\xa9\x30\x56\xae\xf5\xbd\x6b\x8d\x84\xd0\x4a\x4b\xb2\x9a\x2c\x19\x9f\xc4\x66\x27\xb2\x05\xa2\xe3\x07\xe4\x9a\xb0\x5b\xa0\xce\xb1\x0a\x09\x6c\xd3\x5b\xdb\xb4\x5b\x04\x3f\x2e\x37\xea\x01\x51\x09\x6f\x54\x78\xf3\x6b\x06\xf2\x6e\x52\x03\xc6\xc8\x72\x73\x0e\x34\x22\x65\x18\xf6\x2e\xc1\x59\x78\xd6\x56\xa0\xc5\xbc\xbe\x10\x47\xb3\x77\x5b\x4c\xc1\xfb\x46\x85\x4b\xc2\xc9\x1c\xa4\xa5\xe2\x12\xa4\xfa\x12\xfe\x2f\x00\x00\xff\xff\x24\x0c\x7d\x86\x89\x30\x00\x00")

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

	info := bindataFileInfo{name: "templates/views/manager.html", size: 12425, mode: os.FileMode(420), modTime: time.Unix(1521893612, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsManagerMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\xdf\x8f\xe2\x36\x10\xfe\x57\xc2\xdc\x8a\xb3\x45\xce\xb0\x0f\x95\x2a\xc0\xec\x03\xbd\x87\x4a\x55\x5b\xe9\xae\x7d\x41\xa8\x72\xe2\x81\xb8\xeb\xd8\xa9\xed\xc0\xad\x56\xf9\xdf\x2b\x9b\x2c\x17\xe0\x76\xd5\x56\x7d\x41\x83\xc7\xdf\x37\xfe\xe6\x57\xee\x88\xb4\x65\x5b\xa3\x09\x94\x39\x14\xf2\x89\xec\x5a\x53\x06\x65\x0d\xa1\xcf\x77\x04\xde\x95\xd6\xec\xd4\xde\x67\xca\x34\x6d\xd8\x28\xb9\xcd\xb3\xf3\x99\x47\x8d\x65\x00\xca\xca\x4a\x98\x3d\x0e\xa1\x07\xe1\x32\xe4\x77\x24\x54\xca\xd3\xdc\xd9\x23\x47\xd6\x08\x87\x26\xf8\xdf\x4c\x50\x9a\x40\x28\xac\x7c\x82\x1c\x82\x03\x9a\x97\xad\x8b\x3e\x8e\xec\x20\x34\xa1\xb9\xc4\x1d\x8f\x50\x26\x71\x27\x5a\x1d\x7e\x17\xba\xc5\x05\x94\x15\x96\x8f\x85\xfd\x02\x3c\xd2\x39\xdb\x10\x08\x4f\x0d\x02\x7d\x20\x2f\x0c\x00\xa3\xb3\x2f\x5d\x47\x09\xb7\x7c\xeb\x93\x87\xce\xe1\xd3\xc7\x9f\x3e\xae\x3f\x03\x3f\xb9\x83\xd8\xff\x2c\x6a\x1c\x8f\x49\x44\x20\xdb\x29\x23\x09\xd8\x26\xaa\x02\xca\x76\x4a\x07\x74\x43\x9d\x0e\x43\xeb\x4c\xd6\x0b\xed\xe3\xf6\x41\x3e\xa5\xf4\xc4\xf8\x1d\x3d\xe9\xfa\x2a\x94\x4b\xdc\x3d\x38\x7b\x64\x0e\x6b\x7b\xc0\xb5\x16\xde\x13\xa8\x84\xff\x80\xce\x59\x07\x74\x1e\x9d\x42\xca\x5b\x4f\x3e\x2c\x4b\xd1\x86\x60\xcd\x26\x66\x81\xfb\xb6\xa8\x55\xd8\xc2\xf9\x19\xca\x8b\x42\xa3\x84\x7c\xc6\xf9\x10\x14\x1c\x1b\xf0\x31\x8d\x66\x1f\x2a\xda\xbd\xce\xec\xd0\x63\x22\x2e\xb5\x2a\x1f\x5f\x6b\x91\x2b\xda\xd7\x94\xfd\x87\xf7\x8f\x66\x6f\xbd\xee\x2b\xee\xe6\x79\xb1\x0d\xeb\xa4\xbd\xb6\x52\xe8\x75\x42\xc7\x7e\xe3\x00\x8b\x73\xed\xbe\x2d\x21\x0b\xb2\xef\x7a\x23\x6a\xdc\xa6\x58\x23\x5e\x29\x29\xd1\x0c\xa7\xe0\x1a\x73\x9e\x0a\x14\x65\xf5\xea\x4c\xbc\xd5\xcc\x87\xdb\x0e\x9e\x1f\xce\xb3\x51\x4e\x38\x2c\xb5\x5a\x2d\x7d\x70\xd6\xec\x57\x30\x79\xb9\x1d\xdf\x09\x74\x02\xcb\x69\xef\x9a\x67\x30\x39\xc4\xff\x5a\xad\xa0\xa3\x79\xdd\x37\x34\x4b\xd9\xf8\x90\x26\x90\xb2\x2a\xd4\x9a\xc0\x3a\xcd\xb0\xcf\x4e\xcd\xee\xe7\xcb\xc2\x65\xd3\xd5\xb2\xd5\x2b\x98\x94\x91\x23\x5a\x91\x22\x61\x09\xcd\x47\xf7\x7d\x4d\x06\xa9\xfd\x37\x75\x91\x22\x08\xfe\xdc\xfd\xf3\x32\x28\xf9\xbf\x17\x21\xbe\x61\x73\x99\xbe\x2d\x7f\xa3\x32\xb7\x75\xe9\xab\x12\x53\xc1\x1a\xeb\x03\x81\x77\x90\x47\xda\x7c\x10\xf4\xa8\x8c\xb4\x47\xa6\x6d\x29\xe2\x09\x73\xa8\xad\x90\x09\x75\xce\xe2\xe5\xa2\x4d\x09\x6c\x84\xf7\x47\xeb\xe4\x96\xbd\x58\x1f\x7c\x65\x8f\x40\x99\x35\x04\xa2\xc9\x0a\x7f\xf6\x65\x95\x92\x38\x3c\x80\xfc\x1b\xba\xe1\x1d\x4c\x2e\xb7\x95\x92\xb1\x67\xfe\x38\xc4\x15\x0b\x74\xa1\x76\x04\x5f\x56\x42\x42\xc5\xad\x1c\xf0\x4b\x20\x74\xd1\x1b\xc8\xa2\x40\x02\x3d\x84\xe6\x97\x07\x79\xa0\x5d\x47\x17\x09\x1b\x07\xf8\x72\xf7\xc4\x13\xa0\xec\x07\x11\xc4\xe7\x68\x93\x67\x2d\xcc\xbe\x15\x7b\x9c\x3f\xb7\x4e\xcf\x61\x2a\x85\xaf\x0a\x2b\x9c\x9c\x46\xd6\x04\xf0\x53\x75\xff\xbd\x61\x7f\x7a\x6b\xa0\xcb\x8b\x5f\xc5\x5e\x19\x11\x70\x3e\xba\xcf\x8b\x1f\xcd\xce\x46\x43\x3a\x71\x5c\x0b\xad\x0b\x51\x3e\xce\xaf\x94\x8b\x46\x9d\x16\xbc\x68\x14\x49\xdf\x23\xcf\x45\xa3\x58\x34\xc8\x73\x13\x83\x43\xbf\x9a\xa1\xa3\xcc\x58\x89\x9e\xd0\x5c\x0b\x1f\xb8\x69\xb5\x5e\xc4\xcb\xa5\xd5\x6d\x6d\xc8\x2c\xbf\x05\x24\xf9\xd7\xcd\xe6\xec\x31\x57\xa7\xf8\xb1\xb5\xf8\x5d\x3c\xe9\xef\xc2\xde\xd9\xb6\x01\xba\x88\x21\x46\x9c\x9b\xf4\xd5\x89\xbf\x7d\xee\xc7\xe3\x74\xdd\x53\x86\x7f\x11\x45\x59\x81\x3b\xeb\x90\xbc\x5f\x06\x97\x95\x71\xa7\xf2\x9e\x62\xb5\x0c\x32\x2b\xad\xf6\x8d\x30\x1c\xbe\x83\xd5\xfb\x49\xa4\x89\x13\x1b\xe4\x6a\x39\x0d\x2e\x8e\xed\x49\x89\xa8\xb1\x4b\xb5\xb9\xa8\x48\xbf\x06\x62\x4f\xa5\x41\x4d\x1f\x65\x76\x62\xbf\x6e\xa1\x5e\xf3\x2f\x4e\xa2\xe3\xa9\x34\xcc\x46\x9b\xd0\xcd\x6c\xbb\x98\x71\xce\x87\x37\x36\xb3\xed\x78\x0c\xc2\x97\x70\xed\xb8\xdf\x3e\x0c\xd1\x9b\x59\x0e\x12\x7d\x09\x5b\xca\x62\x1d\x09\x9d\x5f\xbb\xc5\xd0\xdb\xd1\x8e\x2e\xfe\x0e\x00\x00\xff\xff\x73\xa6\x1f\xe9\xbc\x08\x00\x00")

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

	info := bindataFileInfo{name: "assets/js/manager.min.js", size: 2236, mode: os.FileMode(420), modTime: time.Unix(1521894164, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4f\x6b\xd4\x40\x18\xc6\x9f\xae\xf5\x12\x50\x50\x2f\x0a\x82\xaf\x07\xab\x22\x53\x27\x59\x0b\x65\x76\xb3\x55\xb7\x2d\x88\x5d\x2c\x25\x7a\xf2\x32\xee\x4e\xb3\xc1\xcd\x4c\x9c\x24\x45\xa1\xa0\xa8\x78\x15\x0f\x5e\x3c\x79\xf1\x5e\xf1\xdf\xe2\x9f\x7e\x86\xc9\x17\xf0\xb3\x48\xb2\x6a\xf5\x3d\x0c\xcf\xf3\x3e\xcf\xfc\x78\x7f\x1e\x9f\x7f\x0d\x00\xf3\x00\x4e\x03\xe8\x02\x38\x0c\xe0\x2e\x66\x93\x01\x38\x02\xe0\x01\x80\x63\x00\x1e\x03\x38\x01\xe0\x15\x80\xf1\x1c\xf0\x0e\xc0\x49\x00\x47\x5b\xc0\x29\x00\xe7\x5b\xc0\x19\x00\xd7\x5a\xc0\xdc\x6f\x6e\x3d\x2d\x00\x87\x6a\xd1\x37\x7a\x3b\x89\x4b\x2b\x8b\xc4\x68\x0c\x1b\x37\xbf\xaa\xee\x95\x31\xa5\x66\xa4\x0e\x36\x3b\x6a\x62\xb2\xd9\x6e\x4b\x65\xc6\x16\x6c\x90\xc7\xc9\x88\x5d\x2f\xe3\x9c\x45\x46\xd0\x48\xed\x5c\xbd\x9f\x8c\x65\x6a\x16\x6d\xe9\x6d\xde\x8a\x58\xdf\xaa\x06\xcb\x56\x65\xa1\x04\x05\xdc\x5f\x66\xbc\xcd\x82\x36\x05\x6d\xb1\xb4\x74\x89\xb7\x39\xf7\x36\x64\x5e\xb0\xc8\x4a\x9d\x4f\x64\x61\xac\xa0\x9b\x0d\x83\x06\xa5\x95\xa9\x19\x19\xea\xfe\x07\xee\x79\x1b\x52\xc7\xa5\x8c\x15\x8b\x94\x4c\x05\xfd\xf5\x82\xb6\xca\x3c\x4f\xa4\xf6\x06\x37\x06\x6b\xec\x8e\xb2\x79\x62\xb4\x20\x7f\x91\x7b\x7d\xa3\x0b\xa5\x0b\x16\x3d\xca\x94\xa0\x42\x3d\x2c\x2e\x67\x13\x99\xe8\x0e\x0d\xc7\xd2\xe6\xaa\x08\x6f\x47\xeb\x6c\xf9\xa0\x57\xdf\xb3\xad\x2c\x5b\xd3\x43\x33\x4a\x74\x2c\xc8\xdb\x9c\x94\x56\x4e\xd8\xba\xb1\x69\x2e\x48\x67\x8d\xcd\xc3\x76\x87\x66\x32\xd4\xe7\x7c\x1e\x86\x3e\x2d\x2c\x50\x2d\xf9\xd9\xd0\xf7\x69\x85\x38\x89\xc6\xf7\xc2\xe0\x4f\xd4\x0d\xaf\xd4\xf2\x42\x53\xeb\xfa\x9c\x76\x77\x67\x5f\x7a\x61\xc0\x2f\xd2\x0a\xf9\x24\x28\xe8\xc0\xbd\x71\xfb\xee\x47\xf5\xdc\x4d\xdd\x87\xea\x59\xf5\xc4\xed\x55\x2f\xdc\xb4\x7a\x09\xf7\xd6\x7d\x72\x9f\xdd\xd4\x7d\x27\xb7\x5f\x3d\x75\xdf\xdc\x9e\xfb\xe8\xbe\xba\xe9\xbf\x49\xdd\x77\x5f\x9a\xf7\x7d\xd3\xaa\xf3\x5f\x01\x00\x00\xff\xff\xa7\x17\xf2\x1f\x61\x02\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 609, mode: os.FileMode(420), modTime: time.Unix(1521894175, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesManageMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x4f\x4b\x6f\x1b\x55\x14\xfe\xa6\x2e\x2f\x97\x67\x85\x84\x84\x58\x9c\x0a\x11\x81\xd0\x94\x19\x9b\x4a\xd1\x24\x13\x53\x92\x14\x10\x0d\x54\xc1\x94\x2d\x57\xf6\x8d\x33\x62\x3c\x63\xdd\x3b\x13\x88\x14\x89\x38\x2e\x54\x60\x44\x11\x0f\x75\x05\xa8\x3b\x96\xa6\xc4\x74\x68\x6b\x77\xcf\xea\xdc\x35\x88\x9f\xc0\x2f\x40\x02\xcd\xa3\x0f\xee\xe2\x9e\xf3\x7d\xe7\x3b\xe7\x7c\xe7\xaf\xe3\x47\xbf\x03\x80\x87\x01\x3c\x03\x60\x0f\xc0\x13\x00\xfe\x44\xf9\xc8\x02\x1e\x03\x70\xc2\x02\x1e\x01\xe0\x5a\xc0\x7d\x00\x5a\x16\xf0\x34\x80\x75\x0b\x78\x00\xc0\xfb\x16\x70\x0c\xc0\xb6\x05\x1c\x05\x90\x54\xba\xdd\x0a\x7f\x5c\xe1\x0b\x16\xf0\x20\x80\xcf\xad\x52\xfb\x8d\x05\x2c\x00\xf8\xe7\x08\xf0\x24\x80\x67\x6b\xc0\xa3\x00\x16\x6b\xc0\xeb\x00\x5e\xab\x01\x4f\x01\xf8\xaa\x06\x3c\x0e\xe0\xa7\x2a\xfe\x56\xc5\xdf\x6b\xc0\x71\x00\x7f\x54\xf8\xef\x2a\xfe\x5b\x03\x2c\x94\xde\x6e\xbf\x7c\x7f\xad\xca\x8f\x55\xf1\xa1\xea\xf6\xdc\xd3\xfd\x28\xbd\xd6\xab\xda\x91\xfc\x3b\x1d\x86\xf1\x87\xb2\x4b\x3b\x22\x4c\xa5\xf6\x70\x5a\x49\xda\x8d\x53\xd2\xa9\x92\x2d\xac\x86\xb1\x96\x58\x8d\xa3\xad\x40\xf5\x49\x8b\x1d\x49\x9d\x1c\xf4\x52\x25\x92\x20\x8e\xb0\x26\xb7\x44\x1a\x26\x58\x93\xba\xa3\x82\x41\xc1\xbd\x25\xfa\x12\x9b\x52\xcb\x04\xef\x88\x1d\x89\xf3\xf9\x68\xbc\x27\x92\xce\xb6\x54\x1a\x9b\x72\x10\xab\xc4\xde\xd0\xbd\xa0\x6b\xbf\x9a\xf6\xb4\xdd\x8e\x3d\xea\xca\x9d\x57\x3e\x08\xb6\x45\x3f\x3e\xa9\xd2\xfa\xb9\xb7\xdb\xf6\xaa\x92\xc5\x0e\x7b\x4d\x24\xd2\xa3\x86\xe3\x2e\xda\x4e\xd3\x6e\x34\xa9\xd1\xf4\x4e\x9d\x7a\xd1\x69\x3a\x4e\xfd\xac\xd0\x89\xdd\x56\x22\xd2\xa1\x48\x62\xe5\xd1\x9b\xc5\x0c\xda\x48\x95\xe8\xc7\xdd\x98\x96\xff\x37\x78\xa5\x7e\x56\x44\xbd\x54\xf4\xa4\xdd\x96\xa2\xef\xd1\x1d\xec\xd1\x66\xaa\x75\x20\xa2\xfa\xc6\x1b\x1b\xeb\xf6\x79\xa9\x74\x10\x47\x1e\xb9\x27\x9d\xfa\x6a\x1c\x25\x32\x4a\xec\xf6\xee\x40\x7a\x94\xc8\x8f\x92\x97\x06\xa1\x08\xa2\x25\xea\x6c\x0b\xa5\x65\xe2\xbf\xdb\x3e\x63\x2f\xde\xd5\xe5\x7e\xb6\xa4\xb2\xd7\xa3\x4e\xdc\x0d\xa2\x9e\x47\xf5\x73\x61\xaa\x44\x68\x9f\x89\x55\x5f\x7b\x14\x0d\x0a\xa8\xfd\xe6\x12\x95\xa9\x1f\x3d\xe7\x3a\xbe\xef\xd2\xc2\x02\xe5\xa9\x73\xc2\x77\x5d\x6a\x91\x43\x5e\x81\x57\xfc\xc6\xed\xd2\xb2\xff\x72\x9e\x3e\x5f\xc8\x96\x5d\x87\xf6\xf6\xca\x96\x15\xbf\xe1\xbc\x40\x2d\x72\xc9\xa3\xc6\x12\xf8\x5b\x9e\xf3\x2d\x33\x32\x43\x73\xc0\x19\xdf\x34\x63\x9e\x12\x5f\xe3\x19\x4f\xcc\x45\x9e\xf2\x8c\x33\x73\xc9\x03\x7f\x6d\xc6\x64\x46\x7c\x95\xa7\x66\x3f\xa7\xcd\xb8\x05\xbe\xcc\x13\xbe\x6e\xf6\xcd\xd8\x1c\x98\x2f\xc0\x3f\xf2\x9c\x0f\xcd\x41\x25\xfa\x95\x0f\xcb\x7e\x9e\x92\x19\xf2\xdc\x7c\x62\xf6\x79\xc2\xb3\x3b\x24\x5f\xe7\x39\xcf\xcc\x05\xce\xf8\x17\x33\xca\x8b\xe6\x53\xce\x38\x2b\x06\xd9\x66\xc4\x37\x79\xce\x37\xcc\xc5\xa2\x29\x33\x5f\x82\x7f\xe0\x5b\x9c\x99\x61\x49\xf0\x14\xfc\x3d\x4f\xf8\x1a\x5f\xbd\x4b\x5c\xe1\x9f\xcd\x3e\xcf\xcd\x90\xb3\xd2\xd3\x95\x7b\x16\x57\xd4\xe5\x7b\x8f\x2b\x9b\x6e\xf0\x94\x0f\xcd\x25\xf3\x59\x41\xfc\x17\x00\x00\xff\xff\x0b\x24\x32\x8e\x10\x04\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/manage.mo", size: 1040, mode: os.FileMode(420), modTime: time.Unix(1521894175, 0)}
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
