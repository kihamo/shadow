// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/grpc.mo
// locales/ru/LC_MESSAGES/manager.mo
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

var _templatesViewsManagerHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1b\x5d\x93\xdb\xb8\xed\x3d\xbf\x82\xc3\xdb\xeb\x24\x6d\x64\x6f\x76\x93\x4e\xeb\xd8\x9b\xe9\x5c\xdb\x69\x3a\x49\xae\x93\xbb\x3e\xa5\x69\x86\x96\x60\x8b\x5b\x9a\xd4\x91\x94\xbd\x5b\x8f\xff\x7b\x87\xd4\x17\xe5\x95\x64\x5b\xa2\x93\x6d\xa7\x7a\x48\xbc\x02\x09\x80\x00\x08\x02\x20\xb4\xdd\xa2\x08\x16\x94\x03\xc2\xa1\xe0\x1a\xb8\xc6\x68\xb7\x7b\x32\x8d\xe8\x1a\x85\x8c\x28\x35\xc3\x09\x59\x42\xa0\xa9\x66\x80\x6f\x9e\x20\x84\x90\x0b\xb4\xef\xbf\x30\x58\xe8\x1c\x68\x07\xc4\xd7\x37\xcb\x8f\x7f\xfb\x61\x3a\x8e\xaf\xf3\x29\xe3\x88\xae\x6f\x9e\xe4\xff\xd5\xd0\x87\x0c\x88\x5c\xd0\x3b\x7c\xd3\x04\x95\x62\xd3\x40\x35\x14\x2c\x58\x45\xc1\x8b\xab\x1c\xb6\xdd\x22\x49\xf8\x12\xd0\x85\x7a\x8e\x2e\x14\xc8\x35\x0d\x01\x4d\x66\x68\x94\xff\x56\x66\x51\x05\x7b\xdb\x6d\x39\xe6\x6d\x64\x46\x49\x48\x18\x09\xa1\x7c\x3b\xfa\x40\x56\x80\xf0\x08\x23\xfc\xc5\x8a\xa3\x5a\x99\xc3\xc4\xdd\x97\x84\x70\x60\xce\xba\x1f\x8e\x70\xc5\x56\x1b\x15\x5f\xdd\x38\x6c\x64\x04\x77\xbb\xe9\x38\xbe\x6a\x18\xdc\x25\xad\xda\xc0\x86\x57\x35\x76\x0a\x15\xdf\x3c\xe9\x24\x42\xc2\x50\xc8\x88\x0a\x8e\x11\x8d\x66\x98\x04\x75\x89\xed\x76\x18\x49\xc1\x60\x86\x35\x99\x33\xaa\x34\x46\x44\x52\x12\xac\x52\xa6\xa9\x02\x06\xa1\x79\x6f\xc0\x32\x6d\x5a\x7c\x5d\x63\xab\xe7\xe8\x62\x05\x3a\x16\x56\x15\xa5\x44\xde\xdb\x57\x35\xbd\x75\xb1\xdc\xa4\x8a\x3d\x7a\x39\x95\x3d\x95\x67\x2f\xdb\x35\xde\x80\x88\x2e\xca\x69\x6f\x79\x92\xea\x9f\xef\x13\xe8\xe2\xd3\xf2\x4a\x6a\x9c\x06\x31\x90\x88\xf2\xa5\x23\xc8\x4c\xd6\xf1\x03\x59\x07\x35\xd6\x8d\xec\x23\xa2\x49\xa0\xc5\x72\x69\x66\x86\x82\x31\x92\x28\xc8\x5f\x27\x44\x02\xd7\x33\xfc\x5d\x93\xd2\x62\x09\x8b\x19\xfe\x2e\x3c\x82\x86\x55\x28\xdc\x25\x84\x47\x10\xcd\xf0\x82\x30\x43\xc2\xbe\x35\x46\x24\x05\x33\xc6\x78\x18\x51\xbb\x46\x4a\xc9\xc4\x2f\xeb\xa2\xc9\xf7\x4c\x85\xc9\xd9\x1d\x2f\xbb\xf1\x4d\xc7\xe4\xc0\x00\x63\x32\x46\xce\xc7\xc8\xa0\xc6\x55\x21\x67\x54\x09\xbc\x54\x5d\x66\x7c\x99\x74\x18\x99\x03\x63\x10\xcd\xef\x8f\x52\xe6\x11\xf2\xd9\xb7\xf2\x60\x2e\xa2\xfb\x23\x26\xee\x4f\x6e\xf6\x57\xc7\xcf\x6e\xf3\x65\x9d\x18\x0e\x69\xac\x71\x12\x2d\x88\x2e\x08\x5a\x90\x80\x48\x29\x36\x81\xa4\xcb\x58\x9b\x3f\x17\x1b\xe3\xfc\xe8\xe9\x78\x1d\x8b\x2a\xf7\x6d\x61\x5b\x7d\x90\xed\xfb\x81\x9f\xb4\x04\xb2\xea\x83\xcc\x3c\x53\x95\x10\x5e\x2c\xdc\x5a\x11\xb2\xff\x06\x1b\x22\xb9\x71\x15\x66\x47\xd0\x17\xbf\xe3\x08\x2b\x4b\x08\xa3\x0b\xbb\x29\xcc\xbc\x5e\xc2\x00\x1e\x9d\xca\x6c\xb6\x05\x4f\x9b\x72\xec\xe9\xd5\x4d\xf7\xb4\xe1\x2d\xa7\xde\x49\x7c\x2f\x84\x5c\xa1\x4c\xbf\x33\x9c\x08\x7b\xcc\x85\x9a\x0a\x3e\xc3\x6f\xf2\x1f\x21\x61\xac\xf4\x13\xe6\x0f\x64\x26\x05\xb1\x90\xf4\xdf\x82\x6b\xc2\x72\xaf\x2c\x41\xa5\x4c\xcf\xb0\x3c\xc2\xeb\x64\x5e\xc5\xe0\xd9\xf3\xf4\x6b\xc2\x68\x44\xb4\x90\x27\x2e\xc4\x2e\x86\x1a\x03\x45\xfa\x3e\x81\x19\x8e\x69\x14\x01\xc7\x88\x93\x15\xcc\xf0\x97\x9c\x1d\x8c\xd6\x84\xa5\x30\xc3\x0d\x71\x89\x67\x8a\xd9\x7a\x6b\x04\xeb\x9e\xbe\x29\x40\x39\xf4\xb4\x1c\xcc\xa3\x3f\x53\x60\xdd\x71\x44\xeb\x12\xdc\x38\xd7\x44\x34\x46\x8f\x89\xe0\x8a\xae\x4f\xf5\x83\x25\x4a\x8b\xa7\x86\x14\x65\xa8\x63\xb1\x06\x99\xff\x56\x5a\xd2\x04\xa2\x9e\x34\x32\x3a\x26\xc0\x18\x32\x5f\xf6\x9f\x9c\x33\x80\x36\x34\xd2\xf1\x0c\xbf\xb8\xfc\xde\xf1\x5d\x56\x1d\x85\xeb\xd2\xb1\x37\x2a\xaf\x6a\x54\x8c\xe6\xcf\x40\xa4\xbe\x94\x77\xc6\x3d\x9f\x81\xca\x75\x8d\xca\x1f\xac\xab\xf1\x41\x66\x3a\x1e\xa2\x55\x43\x7b\xa0\x4d\x99\xc0\xa5\xff\xfc\x2a\x65\x58\x3c\x47\x17\x0b\x63\x47\x36\x63\xf0\xb8\xe3\x1d\x52\x17\x6b\x22\xf7\x52\x05\x4b\xf2\x41\xa6\xd0\x97\x86\x87\x2d\x16\xd9\x20\xd9\x61\xcb\x5a\xc8\x00\x15\x15\x68\x07\x21\x40\x95\x33\xce\x58\x7b\xab\xde\x83\x52\x64\xd9\x2b\xce\x6a\x64\x91\xe4\xa9\xcc\x2d\x59\x13\x15\x4a\x9a\xe8\xc9\x5a\xd0\xe8\xe9\xe5\xb3\xee\xec\x48\x13\xb9\x04\x93\x1d\x1d\x3e\x88\x83\xca\x06\x76\xbb\xe0\x56\x09\x8e\x1d\x69\xe7\x29\xdf\xc1\x7c\xe3\xe8\x15\x25\x12\x6c\x5e\xd2\x8b\x31\xa7\x2c\x92\xe5\x28\x22\xd5\x96\x5b\x49\x36\x05\xc7\x7f\xfd\xe9\xc7\x0f\x96\xe3\x44\x82\x17\xfd\x02\x53\xde\x14\xfa\x40\xb0\x5e\x18\x3c\x3d\xbc\xdd\x7f\xfc\xec\xa6\xd6\xc8\x5e\xa5\x61\x08\x4a\xb9\x86\x65\xcf\x94\x2a\xac\xf7\xc2\x40\x55\xd6\xb1\x27\xc9\x70\xed\xbb\x81\x91\x0d\xf3\x82\xa5\x14\x69\x52\x06\xbb\xbf\xa4\x54\x42\x64\x03\x3b\xd7\x0b\x7c\xcc\x01\x68\xb7\xd3\x32\x85\xca\x86\x6c\x75\xa1\xd4\x58\x89\x26\x01\xa2\x1b\xd1\x64\x80\x43\x68\x3c\xed\xcc\x5a\x5d\xcc\x08\x30\xb0\x4b\xce\x98\xe2\x42\x37\xae\x0f\xc5\x34\xf2\xce\x4a\x46\x52\x48\xf4\x14\x7e\xa9\xed\x17\x1c\x89\x74\xce\x00\x3f\x7b\x08\xf9\xf4\xb9\x1d\xb6\x60\x82\xe8\xe6\x49\x05\xc8\x97\xc3\x76\x93\x01\x9e\xae\xe6\x20\x8b\x64\xe0\xc1\x09\x56\x7a\x33\x9b\x4a\xe5\xc5\x26\x8c\x12\xa2\x35\x48\x3e\xc3\xff\xfc\x14\xfc\xe6\xf3\x9b\x4f\x97\xc1\xef\x3f\xff\xfa\xe9\x3f\x46\xd9\x8f\x67\x6f\x2e\x6a\xe9\x44\x86\xf1\x8f\xb0\x20\x29\xd3\x16\xe9\xd8\x9b\x0e\xac\xb5\xb5\x28\x82\x72\x7d\x7d\xd5\x2c\xd2\x56\x10\xe5\xfa\xb7\x2f\x5b\xe7\x34\x83\x54\x07\xa1\x76\x98\xea\x20\xd5\x0e\x53\x0b\x7a\x07\x51\x2b\xb5\x0e\x68\x06\x6b\xa5\x58\x41\x1f\xb9\x99\x3d\x0e\xdb\x4a\x3b\x74\xde\x0e\x4b\x3b\x74\xde\x0e\xeb\x54\x79\x07\xb0\x53\xe1\x8f\x5c\xdf\x8f\x4c\xdd\x73\x21\x58\xb3\x1c\x73\xc8\x39\x84\x18\xc6\x10\xfe\x6b\x2e\xee\x0e\x8b\xf1\x56\x05\x6a\x43\x75\x18\xa3\xba\x40\x6b\x87\x75\x25\x35\x8b\x18\xa2\x2a\x2e\xfb\x3a\x42\x54\x5a\x52\xbe\x6c\xf1\x3f\xad\xb0\xf9\xbd\x06\xd5\x22\xfb\x1c\x74\x0e\xe1\x6b\xb8\xd3\xa7\xda\xef\xd7\x34\x54\x4f\x6b\x46\x0f\xf3\xc3\x3f\xf1\xb4\x77\xdd\xbc\xe9\x99\x66\xb7\x90\x27\xc9\x12\x65\x73\xae\x3c\x05\x6b\xc5\x53\x55\x2d\xe0\x39\xba\x00\xb3\xd0\xc9\xac\xe0\xc7\xf7\xba\xcd\x33\x15\x89\x89\x52\x5d\xcb\x30\x54\x47\x1f\xac\x7f\x34\xab\xcf\x85\x6f\xdf\xbe\x55\x95\xc1\xe4\x12\x30\x21\x77\xf1\x0b\x97\xfb\xf5\xa6\x42\x54\x54\x1a\x32\x42\xde\xc5\xe5\x21\x6f\x73\x9f\xe9\x38\x5b\x8c\x3f\x3e\xcf\xb0\x1f\xa6\x66\xef\x13\x09\xe4\xc4\xed\x2f\xc5\x46\xcd\xf0\xb5\x9b\x41\x96\x89\x7e\x81\xd2\xef\xc2\xfd\x29\xc7\x2f\xb6\xa2\x79\xc4\x07\xae\x2a\xd1\x6a\xcc\xef\x9e\x35\xe6\xa2\xde\xb3\x4d\x27\xb3\x0e\xe6\xda\x47\xde\x5e\x52\x99\xa7\x5a\x0b\x9e\x1f\x3b\xd9\x1f\xa5\x81\xcd\x35\x47\x73\xcd\x83\xc8\x78\x2d\x99\x5f\x79\x05\x12\x56\x62\x0d\xdf\x20\xe9\x2d\x79\x76\x8b\x28\xd9\xbd\xb0\x96\x44\xc5\x01\x61\x1a\xdf\xf4\xbd\x05\x6d\xa5\x36\xce\xa4\xf2\x95\x65\x4e\xf9\x42\x14\x12\x27\xd1\x90\x0b\xa0\x46\x16\x1e\x8a\x30\x61\xa9\x7a\xec\xd2\x3b\xf1\xee\xb7\xeb\xf1\xe7\x72\x3c\x30\x35\xac\xb6\x37\xec\x32\x67\xb8\x20\xa6\xe3\x01\xd7\x39\xd3\xb1\xbd\xe8\xec\x71\xab\xdc\x4f\xec\xd5\x72\xfb\x5d\x02\x1f\x57\x99\xaf\x75\x02\xe4\x77\xfe\x18\x29\x7d\xcf\x60\x86\x23\xaa\x12\x46\xee\x27\x5c\xf0\xde\x97\xc7\xb1\xec\x1b\xd2\x77\x7a\x9f\x90\x09\x05\xf8\xe6\x57\x9a\xae\x40\xbd\x1e\xb6\x79\x6b\xbd\x9b\xd9\x6d\x39\x20\xc2\x40\xea\x93\x7b\x3e\x4a\x94\x3d\x4f\x76\x97\x15\xc6\xbf\x28\xc1\x68\x54\x35\x99\x0e\xc1\x66\x63\xb0\xac\xe4\x3d\x5c\x4a\x26\x4c\x0b\x42\xe0\x1a\xfa\x34\x76\x94\x18\xdd\x44\x52\xa5\xf3\x15\xd5\x0f\x4e\x97\xe2\xba\xc1\x49\x0d\xb2\x2b\xe6\x1f\x6c\x1b\xcb\xc5\x90\x9c\x71\x98\x6e\x4f\x9b\x62\xc4\x7f\x96\xd6\xa1\x53\x6c\xe3\x7f\xa8\xb3\x8e\xc1\xc2\x5b\x63\xdd\x8f\xa9\xf6\xdb\x59\x57\x21\xfc\x7f\x63\xdd\xde\x94\xff\xd6\xc6\x3a\xbf\x6d\x55\x5f\xa3\xa5\x6a\x48\x3b\xd5\xa0\x3e\x0f\xb7\xf7\xe8\xd5\x59\xda\xa8\xdc\xbe\xa3\x57\xe7\x68\xa1\x6a\xed\xd1\xf2\xd2\x3e\xd5\x3f\x0c\x1e\xd4\xcf\x34\xa4\x97\xe9\x40\x1f\x93\xe3\x41\x87\x35\x32\x9d\xbb\x89\x69\xa0\x61\x7b\x6f\x5e\x1a\xdc\xb8\x74\xbe\xa6\xa5\xe1\x0d\x4b\x9b\x47\xd4\xb0\x54\x36\x2b\xf5\x62\xea\xeb\x36\x2b\x79\xac\xd5\xfa\x6d\x52\xf2\x92\x81\x7f\xcb\xdd\x32\xa0\xab\x69\x58\x95\x76\x7f\x9b\x92\xc4\xcb\x16\x3d\xb8\x9c\x15\x49\x7c\x14\xca\xbe\xa5\xe2\xfb\x9f\x96\xc3\xb8\xee\x5d\x2c\xea\x55\x28\x3a\x5f\x08\x7c\xe4\xd0\x23\x86\x1d\x18\x72\xa4\xdb\x3a\xe6\xab\xc2\x96\xe3\xa5\xf5\x3c\x7a\x44\x5f\xec\x1d\x36\xba\x0e\x31\xb6\x4f\x3e\xfc\xe5\xac\x5b\x08\xa8\xe3\xa9\x7f\xc8\xec\xd4\x15\x9d\x0f\xa9\x8d\x16\xb2\xaf\xa8\x6d\xed\xcf\xb9\xdd\x1f\x87\xc6\x91\x58\x44\xa3\xb2\x3a\xb6\x2d\xc9\x6e\x62\xaa\x21\x50\x09\x09\x61\x82\x12\x09\xc1\x46\x92\xe4\x75\x09\x36\xe9\xcc\x82\x89\x4d\x70\x3f\x41\x24\xd5\xa2\x82\xac\xc8\x5d\x10\x03\x5d\xc6\x7a\x82\xae\x2f\x61\x55\x41\x6c\xe4\x3d\x41\x2f\x2e\x2f\xbf\xcf\x5e\xe6\x25\xd0\xd1\xde\xed\x12\xca\x0b\x84\x5b\x07\xa7\x5c\x52\x9e\x7d\x76\x37\x41\x97\xaf\xf7\x01\x73\xa1\xb5\x58\x59\x48\x0d\xb1\x53\xfc\x74\xb0\x25\x42\x51\x4d\x05\x9f\x20\x09\x8c\x68\xba\x86\x3a\x3b\xb6\x06\xd9\x38\x9e\xcc\x95\x60\xa9\x86\x8a\x81\x9c\xa5\x57\xc9\x5d\xf5\x4e\x8b\x64\x82\xae\xca\x57\x05\xda\xec\x2e\xc5\xc1\x6b\xab\x6d\x84\xd1\xa5\xe1\xc4\xe0\x69\x1c\x3f\x72\x3b\x3d\x5b\x66\x33\x58\xb4\x4c\x76\x44\xeb\x2a\xb7\x4d\x13\x0d\xd3\xac\x46\xaa\xa9\x6b\x90\x9a\x86\x84\x15\xa4\xb5\x48\x0a\x14\xd3\xb1\xb5\xb1\x36\x53\xbc\x55\xd6\x10\xb7\x5b\xa4\x34\xd1\x34\xfc\xcb\xcf\xef\xdf\xa1\xa7\xd9\xef\xbf\x7f\x7c\x87\xf0\x38\x22\x2a\x9e\x0b\x22\xa3\x31\x51\x0a\xb4\x1a\xaf\x81\x47\x42\xaa\x71\xf9\xb9\xd9\xf8\xd6\xf9\x63\xb4\xa2\x7c\x64\xb0\xda\x66\xdb\x67\x07\x90\x2f\x65\x12\x16\x78\x6f\xd5\x78\x45\x38\x59\x82\xb4\xf3\xb5\x4c\xcb\xe9\x19\xe3\xff\x09\x00\x00\xff\xff\x68\xc6\xc7\x5c\x8b\x40\x00\x00"

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

var _assetsJsManagerMinJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4d\x6f\xdb\x38\x10\xfd\x2b\x5a\xae\x61\x90\xb0\x4c\xec\x2e\xf6\x24\x2d\x11\x2c\xd2\x53\x81\xb6\x87\xa2\xa7\x24\x07\x46\x1a\x59\x8c\x69\x92\x1d\x52\x4a\x5c\x47\xff\xbd\xa0\x3e\xfc\x15\xbb\x0d\x74\x10\x66\x38\xf3\x38\xf3\xde\x70\x66\xb4\x6a\x4c\x11\x94\x35\x94\xed\x66\x94\x70\xd9\x1b\xc9\x63\x13\x82\x35\xa3\xb5\x44\xd8\xd8\x16\x08\xe3\x85\x56\xc5\xfa\x38\xa5\x95\x98\x80\x98\xd1\x50\x2b\xcf\x52\x14\xc0\x0b\x6d\x3d\xf8\x40\x09\x57\xc6\x35\x61\xb9\x42\xdb\x38\xc2\x52\x27\x90\x3b\x89\x60\x02\x65\xb9\xaa\x28\xf2\x52\x06\x49\x09\x82\x03\x19\xa0\x24\x6c\x40\x43\x2f\x1c\xaf\x94\x29\xcf\x11\xf2\xbf\xff\x43\xcf\x35\x98\x55\xa8\x6f\x28\xf2\xa1\x28\xca\xd2\x7f\x84\x10\xfb\x93\xf9\xfc\x00\xfc\xbd\x51\x18\x81\xa3\x6f\x44\x3c\x6f\x48\x96\xe5\xad\x96\xde\x53\x52\xab\x12\x08\x63\x19\x7d\x77\x6c\xfa\xf6\xa6\xd7\xd7\xf3\xec\xbe\x85\x4b\x17\x75\xa0\x3d\x24\x70\x09\xf5\x14\x41\x96\x25\x61\x63\xb7\xbf\x8e\xbc\x76\x57\xc7\xd2\xab\xd2\x0e\xe8\xef\xd0\x75\xd2\xee\xaa\xc0\xad\xb8\x52\xcf\xa0\xf6\x15\x52\x2f\xf7\xf5\x66\x34\x54\x45\x5b\x5e\x4b\x7f\xca\x61\x7b\x29\x3d\x8f\xc4\xf6\x0d\x14\x02\x63\xb5\x06\xe8\x1f\x7f\xb1\xfc\x30\x7e\x5c\x3a\x07\xa6\xa4\xc5\x28\xc2\x45\x98\xf4\xad\x34\x23\x8f\x85\xd4\x7a\x89\xe0\x1b\x1d\x26\x32\x7b\x4e\x2e\xf1\x38\x32\x78\xb8\x3a\x22\xd1\x11\xa9\xb2\xb8\xe9\xd1\x08\xe3\xbe\x79\xdc\xa8\x70\x9c\x0a\x2d\x98\xc0\x1d\xf6\xff\x0f\x50\xc9\x46\xc7\xa7\x73\xaa\x4c\x7d\xa4\x4c\x3a\xd4\x24\x66\x94\xfc\x49\x16\xb0\x27\x31\x3a\x09\xeb\x8f\x9d\x35\x1e\xc4\xe0\x9a\x14\x99\xdc\x84\xa5\xc1\xae\x56\x1a\xc4\xbe\x08\x40\xb4\xc8\x76\x53\xc4\x09\x4d\x47\x95\x22\x84\x06\x4d\xd2\x47\xdf\x10\xa9\x01\xc3\xd2\x37\x45\x01\xde\x93\x6c\xb4\x4b\x69\x56\x80\xa4\x3b\x9a\xce\xdf\x20\x8c\x19\xd9\x19\x60\xc7\xba\xbc\xe6\xcf\x52\x85\x4f\x40\x77\x50\x55\x50\x84\x8c\x54\xb2\x80\x47\x6b\xd7\x24\x32\xcb\xe5\x93\x7c\xa1\xbb\xb0\x75\x90\x01\x97\x21\x20\x25\x1b\x08\xb5\x2d\x09\x4b\x1b\xd4\x7b\xe7\x30\x8c\x84\xa5\x91\xaa\x0c\xb8\x07\x54\x52\xab\x1f\x71\xab\x14\x76\xe3\x34\x04\xc8\x4e\xca\xec\x89\xf3\xb5\x7d\xa6\x2c\xdd\x57\x31\x8d\x47\xda\xd7\x7f\x96\x30\x30\x57\x87\x8d\xa6\xe4\xff\x27\xf9\x92\xc4\x6d\x01\x3e\x24\x95\x54\x3a\xce\xf6\xc8\x7a\x9c\xd1\x2e\x1d\xbb\x3c\x60\xe0\xf0\x16\x9d\x20\x64\x78\x47\x35\xc8\x12\xd0\x0f\xee\xd1\x10\x77\x0f\x79\x65\x91\x46\xd7\x1a\xb6\x89\x32\xc9\x21\x70\x3a\x68\xa5\x3e\x39\xb8\x5b\xc3\xf6\x81\x4d\xc6\xf8\x1f\xb7\xe8\x83\x58\xc3\x96\x07\xfb\xcd\x39\xc0\x5b\xe9\x81\xb2\x05\xc9\x12\xb2\x68\xa5\xce\x9d\x98\x82\x9f\xac\x32\x94\xdc\x1b\xc2\x16\xe4\xde\xdc\x1b\xd2\xc5\xb5\x1c\x39\xba\x71\x0b\xf1\xf1\xeb\x97\xcf\xdc\x07\x54\x66\xa5\xaa\x2d\xed\x4d\x27\xd1\x03\x9d\xa2\x58\x6a\x1a\xad\xd3\x7f\x59\x86\xbc\xe7\x6e\x3e\xa7\x6e\x21\x46\xe3\x30\xb1\x03\x7d\x6e\x4f\xd5\x14\xd0\x75\x2c\x7e\xf9\xcf\x00\x00\x00\xff\xff\x8a\x1d\x3b\x9a\xc6\x06\x00\x00"

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

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xcf\x6b\x13\x5b\x14\xc7\xbf\xf9\xf1\xfa\x1e\x79\xef\x21\x14\xdc\xb9\x38\x15\x2c\x75\x31\x35\xa9\x22\x32\xe9\xb4\x62\x69\xb1\xd8\x60\x2d\x51\x10\xdc\xdc\x26\x37\xd3\xc1\xc9\x9d\x70\x67\x52\x5a\xe8\x22\x4d\x45\x17\x96\x2e\x44\x17\x42\x5d\xb8\x70\x1f\x2c\x81\xd8\x1f\xe9\x56\x77\xe7\x82\x5b\xfd\x03\xc4\x3f\xc0\xa5\x4c\x7e\xb4\x78\x16\x33\xdf\xcf\x39\x9f\x73\x66\xbe\x8f\xa6\xdf\x00\xc0\x3f\x00\x2e\x01\x78\x02\xe0\x5f\x00\xfb\xe8\xd7\x67\x00\x97\x01\x7c\x19\xf4\x7f\x02\xf8\x1f\xc0\x2f\x00\x63\x00\x46\x12\x7d\x9e\x48\x00\xa3\x00\x6e\x26\x80\x0b\x00\x16\x13\xc0\x6c\x02\x78\x9c\x00\x6c\x00\x6f\x93\xfd\x6f\x7c\x4d\x02\xff\x01\xf8\x96\x04\xe6\x00\xfc\x18\xf0\x64\x0a\xb8\x08\x20\x9f\x02\xd2\x00\x1e\xa4\x80\x04\xce\x6b\x64\xf0\x4e\x01\xf8\x1b\x40\x12\x7d\x2f\xbe\xf9\x57\x3c\x28\x05\xaa\xe2\xb9\xe9\x79\x25\x56\x7d\x59\x26\x2d\x5d\x2f\x8c\xa4\x26\x2d\x2b\xbe\x2c\x45\x5e\xa0\x86\xca\xdd\x20\x8c\x86\x79\xa9\x27\x0d\xa9\x20\x36\xc8\x97\xeb\xd2\xa7\xa0\x42\x35\xa1\x43\x4f\xb9\x14\x6d\xd6\x64\x38\x34\xee\x47\x6b\x52\x9f\xd1\x72\xa0\x23\x52\xf5\xea\xea\xd9\x09\x2b\x12\xab\x69\x57\xd7\x4a\x58\x91\xb5\x40\x47\x56\x21\x74\xbd\xb2\x75\xa7\xee\x86\x56\x31\xb0\xa9\x2c\xd7\x6f\x3f\xf5\xd6\x44\x35\x98\xd4\xf5\xcc\x92\x08\x23\xab\xa8\x85\x0a\x7d\x11\x05\xda\xa6\x7b\xbd\x11\x15\xea\x5a\x54\x83\x72\x40\xd3\x7f\xf8\x33\x99\x25\xa1\xdc\xba\x70\xa5\x55\x94\xa2\x6a\xd3\x19\xdb\xb4\x52\x0f\x43\x4f\xa8\x4c\x61\xb1\x30\x6f\x3d\x92\x3a\xf4\x02\x65\x53\x6e\x32\x9b\x99\x0b\x54\x24\x55\x64\x15\x37\x6b\xd2\xa6\x48\x6e\x44\xd7\x6a\xbe\xf0\x54\x9e\x4a\x6b\x42\x87\x32\x72\x1e\x16\x17\xac\x5b\xe7\x5e\xfc\x3f\x15\xa9\xad\x79\x55\x0a\xca\x9e\x72\x6d\xca\x2c\xfb\x75\x2d\x7c\x6b\x21\xd0\xd5\xd0\x26\x55\xeb\x61\xe8\x5c\xcf\x53\x3f\x3a\xea\x4a\x2e\xeb\x38\x39\x1a\x1f\xa7\x38\x66\xc7\x9c\x5c\x8e\x66\x29\x4b\x76\x8f\x67\x9c\xa9\xe1\x68\xda\xb9\x11\xc7\x89\x9e\x36\x9d\xcb\xd2\xd6\x56\x7f\x65\xc6\x99\xca\x5e\xa5\x59\xca\x91\x4d\x53\x79\xf0\x2b\x3e\xe4\x23\xb3\x67\x5e\x70\xc7\x34\xcd\x2e\x99\x06\xb7\xf9\x80\x3b\x66\xdb\x34\x4d\x83\x5b\xe6\x39\x77\xcc\x5e\xaf\x6d\x9e\xf1\x11\xb7\xf9\xd0\x6c\x73\x87\x3b\xe0\x0f\xdc\x8d\x2d\xf0\x7b\x6e\x9b\x06\x7f\x8c\x9f\xe0\x7d\x6e\x0d\x94\x63\x6e\xf1\x91\xd9\xe5\x13\xf3\x92\x3f\x91\xd9\x31\x0d\xee\xc6\x16\x9f\x98\x5d\xe2\x53\x6e\x99\x46\xcf\x3b\xe1\x03\x6e\x91\x69\x72\x87\x4f\x63\x03\xfc\xda\x34\xcc\x0e\x1f\x70\x97\xdb\xe0\x77\xdc\xe5\xe3\xf8\x76\xbc\xd3\x35\x0d\xd3\xe4\x16\xdc\x95\xe5\x39\xfc\x0e\x00\x00\xff\xff\x4c\xb9\xe1\xf1\x56\x03\x00\x00"

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

var _localesRuLc_messagesGrpcMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x4f\x6b\xdb\x30\x18\xc6\x9f\x84\xec\xe2\xe3\xce\x3b\xbc\x3b\x2c\x6c\x30\x65\x76\xb6\xc3\x50\xac\xa4\x34\x24\x50\x1a\x43\x1a\xdc\xdc\x45\xac\x3a\xa6\xb6\x64\x24\xb9\xb4\x90\xaf\xd1\x7e\xbc\x7e\x96\x12\x3b\xfd\xf7\x9c\x9e\x1f\xcf\x4f\x20\xde\xe7\xaf\x83\x47\x00\xe8\x03\xf8\x06\xe0\x37\x80\x2f\x00\x62\x74\x59\x03\x18\x00\xb8\x02\x30\xeb\x01\xdb\x13\x3f\xf5\x80\xde\xc9\xe9\xe3\x43\xf2\xcd\x7a\x8e\x8d\xaa\x8d\xf5\x2c\x71\x79\x91\xb1\xf3\x26\x77\x2c\x35\x9c\x32\x75\x77\x76\x5b\xec\x65\x65\x46\xb6\x09\x56\xd2\x79\x96\x5a\xa9\x5d\x29\xbd\xb1\x9c\x2e\xdb\x89\x92\xc6\xca\xca\x64\x86\xe2\x4f\xfe\x34\x58\x49\x9d\x37\x32\x57\x2c\x55\xb2\xe2\xf4\xc6\x9c\x36\x8d\x73\x85\xd4\x41\x72\x91\x2c\xd8\x56\x59\x57\x18\xcd\x29\x1a\x85\xc1\xdc\x68\xaf\xb4\x67\xe9\x43\xad\x38\x79\x75\xef\xff\xd4\xa5\x2c\xf4\x84\x76\x7b\x69\x9d\xf2\xe2\x3a\x5d\xb2\xff\xef\xde\xf1\x3f\x37\xca\xb2\x85\xde\x99\xac\xd0\x39\xa7\x60\x5d\x36\x56\x96\x6c\x69\x6c\xe5\x38\xe9\xba\x45\x27\xfe\x4e\xa8\xab\x42\xff\x88\x42\x21\x22\x1a\x0e\xe9\x58\xc3\xef\x22\x8a\x68\x46\x21\xf1\x96\xa7\x62\xfc\x3a\xc5\xe2\xdf\xb1\xfe\x6c\xb5\x38\x0a\xe9\x70\xe8\x9e\x4c\xc5\x38\xfc\x45\x33\x8a\x88\xd3\x78\xd2\x1d\xf1\x25\x00\x00\xff\xff\x53\x05\x69\x39\x9b\x01\x00\x00"

func localesRuLc_messagesGrpcMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesGrpcMo,
		"locales/ru/LC_MESSAGES/grpc.mo",
	)
}

func localesRuLc_messagesGrpcMo() (*asset, error) {
	bytes, err := localesRuLc_messagesGrpcMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/grpc.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesManagerMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x3f\x8b\x13\x51\x14\xc5\xcf\x8e\xf1\x0f\x23\x88\x58\x2b\x5c\x0b\x17\x2d\x26\xce\x44\x0b\x99\x64\x12\x75\xd9\x80\x98\x80\x2e\xa3\xd8\x3e\x93\x67\x76\x70\xf2\x5e\x78\x6f\x46\x14\xb6\x31\xad\x5b\x89\x76\x82\xe0\x27\x58\x16\x82\xd9\x15\xb2\x36\x56\x36\x77\x1a\x3b\xfd\x2c\x32\xc9\xe8\xb2\xb7\xb9\xe7\xdc\xfb\xbb\xdc\xf3\xe7\x52\xed\x23\x00\x9c\x05\x70\x19\x40\x0c\xe0\x3c\x80\x5d\xac\x6a\x0e\xe0\x0c\x80\x03\x00\x35\x00\xdf\x01\x9c\x06\xf0\xa3\xea\x3f\xab\x79\x51\x71\xbf\x00\x74\xd6\x80\xdf\x00\x2e\x02\xb8\xe2\x00\x17\x00\xd4\x1d\xe0\x1c\x80\xc8\x01\x5c\x00\x5d\x67\xc5\x3f\xae\xfc\x33\x07\x58\xab\x7e\x3a\x55\xaf\xe1\xb8\x4a\xf6\x54\x95\xb3\xfc\x8b\x7b\x83\x2c\xd1\x0a\x1b\x22\x4d\xd1\x4d\x64\x3a\x44\x4f\x3c\x97\x29\xe2\x37\x13\x09\x9b\x19\x29\xc6\xd8\x92\x13\x6d\x32\xaf\x6f\x47\xc9\xd0\xbb\x9f\x8f\xac\x17\xeb\x90\x86\xf2\xd5\xdd\x97\xc9\xb6\x18\xeb\xba\xc9\xdd\x9e\xb0\x99\x17\x1b\xa1\x6c\x2a\x32\x6d\x42\x7a\xb8\x5c\x51\x3f\x37\x62\xac\x87\x9a\x5a\x27\xf8\xb6\xdb\x13\x6a\x94\x8b\x91\xf4\x62\x29\xc6\x21\xfd\xf7\x21\x6d\xe5\xd6\x26\x42\xb9\xfd\x07\xfd\x4d\xef\xa9\x34\x36\xd1\x2a\xa4\xa0\xee\xbb\x1b\x5a\x65\x52\x65\x5e\x19\x2e\xa4\x4c\xbe\xce\x6e\x4e\x52\x91\xa8\x26\x0d\xb6\x85\xb1\x32\x8b\x9e\xc4\x5d\xef\xce\x31\x57\xe6\x79\x21\x8d\xb7\xa9\x06\x7a\x98\xa8\x51\x48\xee\xa3\x34\x37\x22\xf5\xba\xda\x8c\x6d\x48\x6a\xb2\xb4\x36\xba\xd5\xa4\x95\x8c\xd4\xb5\xc0\x8f\xa2\x80\xd6\xd7\xa9\x94\xfe\xd5\x28\x08\xa8\x43\x3e\x85\x4b\xdf\x8e\x1a\xff\x56\xad\xe8\x76\x29\xaf\x2f\xb1\x56\xe0\xd3\xce\xce\xea\xa4\x1d\x35\xfc\x1b\xd4\xa1\x80\x42\x6a\x34\xc1\x1f\x78\xc6\x07\xc5\xdb\x62\xca\xfb\x3c\xe7\x19\xf8\x7d\xf1\x8e\xbf\xf2\x3e\xef\x15\xd3\x62\x17\xfc\x99\x17\xfc\xad\x9c\x7f\xe2\x59\x31\xe5\x43\xde\x03\x7f\xe1\x39\x1f\x81\x8f\x78\x51\x4c\x79\xc1\x87\xf8\x1b\x00\x00\xff\xff\xd2\xd4\x28\x7f\x63\x02\x00\x00"

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
	"templates/views/manager.html": templatesViewsManagerHtml,
	"assets/js/manager.min.js": assetsJsManagerMinJs,
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
	"locales/ru/LC_MESSAGES/grpc.mo": localesRuLc_messagesGrpcMo,
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
				"config.mo": &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
				"grpc.mo": &bintree{localesRuLc_messagesGrpcMo, map[string]*bintree{}},
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
