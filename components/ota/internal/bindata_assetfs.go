// Code generated by go-bindata.
// sources:
// templates/views/releases.html
// templates/views/update.html
// locales/ru/LC_MESSAGES/ota.mo
// locales/ru/LC_MESSAGES/releases.mo
// locales/ru/LC_MESSAGES/update.mo
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

var _templatesViewsReleasesHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5f\x6f\xdb\x36\x10\x7f\xef\xa7\x38\x68\x29\xe4\xa0\x91\x64\x67\xeb\x36\xb8\x56\x86\xa2\xdd\xb0\x02\xdd\x56\xb4\xe9\x1e\xf6\x52\xd0\xe2\xc9\x62\x4a\x93\x2a\x49\xd9\xf1\x0c\x7f\xf7\x81\xfa\x63\xcb\xb2\x6c\x2b\xa9\x13\x0c\xc3\x04\xb4\xa1\xc9\xfb\x7f\xbf\x3b\x1d\xb5\x5c\x02\xc5\x98\x09\x04\x27\x92\xc2\xa0\x30\x0e\xac\x56\x4f\x46\x94\xcd\x20\xe2\x44\xeb\xd0\x51\x72\xee\x5c\x3d\x01\x00\xa8\xef\x46\x92\x7b\x53\xea\x0d\x2e\xc1\xae\xf4\xb4\x5a\xdd\x6a\x6f\x70\x59\xd2\x37\x79\x6e\x3f\xa5\x44\x20\xaf\x9d\xee\x52\x18\x66\x38\x36\x28\x72\xaa\xe4\xf2\x6a\xb9\x04\x36\xf8\x51\x80\xf3\x1e\x39\x12\x8d\xda\x01\x1f\x56\xab\x51\x90\x5c\xb6\x30\xd4\x8d\xe5\x48\x54\xcc\x6e\x9d\xab\x51\x40\xd9\xac\xa1\xbf\x65\x6b\xcb\xa4\x2a\x30\x87\x75\x18\x32\xe6\xe8\x29\xd4\xa9\x14\x9a\xcd\xda\x7c\xc8\x59\x72\xba\x2d\x26\x28\x58\xb5\x51\x2c\x45\x0a\x94\x18\x52\xec\x53\x53\x93\x07\x42\xce\x15\x49\x1d\xd0\x66\xc1\x31\x74\xe6\x8c\x9a\x64\x38\xe8\xf7\x9f\xee\xd1\x54\x68\x4b\x90\xd0\x43\xe7\x6a\xff\x61\x29\x60\x13\xf6\x3f\x51\x69\x26\x45\x15\x75\x93\xdc\x81\xf7\x65\x64\x98\x14\xfa\x5e\xbc\x1f\x53\x2e\x09\x45\x0a\xc4\xdc\x8b\xff\x03\xfb\x1b\xef\xc5\xf8\x2a\xc1\xe8\xb3\xce\xa6\xf7\xf3\x58\x45\x09\x33\x18\x99\x4c\xdd\x4f\xfb\x3b\x62\x92\x2e\x8c\xa3\xe0\x50\x16\x2d\xef\x11\x0c\x8c\x25\x5d\xec\x3f\x5f\x2e\x41\x11\x31\x41\x38\x63\x17\x70\xa6\x8a\xd2\x83\x61\x08\x7e\xb9\xd6\xb6\x63\xec\x97\x7e\x14\x61\xd4\xba\x5c\x09\xf6\x4b\x90\xc1\x6a\x65\x35\xb3\x78\x73\xf2\x46\xbf\xca\x94\x42\x61\x6c\x44\x74\x4a\x44\x55\x45\x9c\x8c\x91\x43\xfe\xbf\xa7\xb3\x28\x42\xad\x9d\x4d\x18\xa3\x82\xc9\x81\xb3\x3c\x92\x96\xd1\x1e\xa2\xa0\x45\x64\x0f\x84\xa6\xb2\xef\x20\x01\x34\x1a\xc1\xd8\x08\x6f\xa2\x64\x96\x3a\xa0\xa4\xad\xd4\xe2\xc7\x71\x21\xb9\x20\x02\x89\xc2\x38\x74\xea\x21\x79\x2d\xe7\xc2\x56\xc0\xc7\xf7\x6f\x61\xb5\x72\xc0\x10\x35\x41\x13\x3a\x9f\xc6\x9c\x88\xcf\x4e\x4d\x31\x58\xe5\x65\x08\xf2\x35\x8b\x64\xb1\x79\xab\x3b\x5a\x90\x5b\xc1\x2a\x99\x31\xd1\x10\x13\x2f\x66\x1c\x3d\x5a\x9a\xe1\x40\xde\xa4\x73\x1b\x8b\x10\xbf\x5e\x9f\xd8\x18\xdb\x2e\xcb\x3a\xba\x1b\x90\xab\x27\x9d\x28\x0b\x2c\x08\x6c\x83\x83\x51\x19\x1e\xc2\x60\xbb\x2c\xfc\x02\x67\x7e\x09\x0e\x5b\xab\x1b\xc9\xf5\xca\xbd\x8b\xdc\xd1\x38\x33\x46\x0a\x30\x8b\x14\x43\xa7\xf8\xb1\x93\x1d\x6a\x6b\x49\xed\x24\x27\x6f\xfb\x9e\x91\x93\x89\x8d\xec\x54\x52\xc2\xab\xbd\x32\xdb\xdf\xd4\x37\xf3\xb5\xd7\xcc\xc3\x2b\x29\x62\xa6\xa6\x90\xa5\x13\x45\x28\x82\x91\x50\xd5\xeb\x53\x6d\x93\x23\x18\x5f\xff\x6b\xa9\xb8\x2d\xe1\x11\xe1\x7c\x4c\xa2\xcf\xa1\x53\x12\x7e\x2c\xa4\xf6\xdc\x3a\x36\xdf\xbc\x86\xd5\xca\xbd\x80\x98\x70\x8d\xe7\x2f\xee\x00\x32\x68\x03\x1a\x51\x4a\xce\xbd\x88\xa9\x88\xa3\x67\x4b\xa8\xe9\x62\x69\x04\x90\x34\xe5\x2c\x22\x26\x7f\x1f\xdd\x11\x75\xb9\xe6\xa0\xc8\xcf\x1d\x38\xfe\xd5\xd9\x05\x22\x28\x28\xd4\x86\x28\xd3\x0c\xce\xc3\x65\xdd\x56\xde\xd7\x27\x3d\xe2\x32\xa3\x5e\x96\xbf\xe5\x3d\xc2\x4d\xa7\xac\xc3\x9c\x99\xa4\xf2\xf8\x91\x20\xb0\x7e\x6d\x74\xeb\x59\xb0\xee\x35\xb5\xa6\xf5\x1e\xa7\x72\x66\xc7\xbb\x53\x37\x96\x39\x51\x82\x89\xc9\xc3\x62\x4f\x59\xeb\xf1\x94\x5d\x25\x8f\x47\x2b\xbc\xbe\x1e\x58\x46\x11\x9d\xec\xa2\xa9\x50\xf9\xe8\x98\x39\x2d\x75\xcb\xb5\x65\x97\xe4\x24\xc3\x4d\x03\xc2\xd5\x34\xfe\xd2\x74\x75\x6a\xa4\x23\xc5\x52\x53\xc2\xb7\x56\xc3\xc1\x0d\x99\x91\xe2\xd0\xb9\xa2\x32\xca\xa6\x28\x8c\x3f\x57\xcc\x60\x8f\x12\x83\xd7\xf2\x83\x51\x4c\x4c\xb6\xd1\xb1\x31\xc0\xff\x45\xaa\x29\x31\xe0\x5c\xf6\xfb\xdf\x7b\xfd\x81\xd7\xbf\xbc\x1e\x3c\x1f\xf6\xbf\x1b\xf6\x9f\xff\xd5\xff\x61\xd8\xef\x3b\x39\x90\xce\x47\x41\xa1\xa5\x93\xb3\x1d\x12\xd0\x2d\xb2\x75\xa3\xed\x1d\xa4\xf3\xbc\x59\x67\xac\xee\x20\xf7\x62\x6e\x4c\x32\x77\x17\x60\xef\x20\x47\x19\x0f\x5f\x40\x8e\x47\x74\x14\x1c\xb8\x82\x8c\x82\xfc\x32\xdc\x72\xf9\x3e\x7a\x95\xaf\xfd\x2c\x97\xe5\x9f\x5a\x1f\xaf\x7d\x00\xb1\xd7\x24\xa7\xb2\x73\xb9\x04\x6d\x88\x61\xd1\xaf\xd7\xbf\xbd\x85\x5e\xb1\xb6\x03\xb8\x13\x50\xa2\x93\xb1\x24\x8a\x06\x44\x6b\x34\x3a\x98\xa1\xa0\x52\xe9\x60\x7d\x73\xd7\xbe\x40\xe3\x8d\x75\x10\xe9\x62\xf7\xba\xd8\x1d\x4b\x69\xb4\x51\x24\xf5\xa7\x4c\xf8\x91\xd6\x4e\x39\x37\x9d\x50\xeb\xe6\x8b\x41\x65\xc0\x66\xe7\xb0\x01\xed\x51\xb9\xd1\x27\x8c\x49\x70\xa3\x83\x9b\x2f\x19\xaa\x85\x5f\x0b\x8b\xb5\xe5\xe6\x21\x62\x31\xd6\x56\xe1\xde\x04\x3c\x88\xce\x4d\xb4\x1b\xba\x6b\x69\x78\x04\xe5\xa5\xef\x7b\x73\xbf\xad\xbe\xa8\x91\x4e\x6d\x7a\x5d\x5d\x67\xbd\xaa\x63\x9f\xfb\x0a\x09\x5d\xf4\xe2\x4c\xe4\x5f\x79\xa0\x77\x0e\xcb\xad\xc2\x9c\x33\x41\xe5\xdc\xdf\x1e\x29\x21\x84\x0d\x07\xa3\x17\xd5\x30\xd7\x64\xb6\xcf\x8c\x28\xc8\x20\x04\x37\x90\x86\x04\xa5\x9c\xc0\x85\x67\xc0\x28\x3c\x03\x37\x28\x87\x62\xf7\xc5\xee\x68\xc6\x62\xe8\x1d\x10\x6d\x9f\x0c\x9e\x85\xe0\xfe\x54\x52\x85\x03\x77\x87\xaa\x65\xe4\x3b\xf3\xc9\x0d\xb9\xed\xb5\x4b\xb4\x41\x1c\x82\xfb\xee\x8f\x0f\xd7\xee\x45\xbb\x4e\xc5\x87\x90\xb5\x9f\x95\x57\xf8\x61\x2d\x40\x6a\x9f\xed\x6b\x17\x2d\xc2\x32\x6e\x20\x0c\x43\x70\x63\xc2\x38\x52\xf7\x10\x93\x7d\x04\xce\xe1\xdd\xef\xd2\xb0\x78\xb1\xc7\x91\x2d\xa7\xec\x20\x35\x04\xf7\x67\xa5\xa4\xda\xe3\xd6\x16\x3d\xde\x9a\x21\x28\x7f\x8a\x5a\x93\x09\x76\x60\x28\xa2\x86\x1d\xe5\x27\x8c\xe2\xb0\x80\xf1\x71\x62\x6d\x16\x9c\x89\xc9\x10\xdc\x75\x25\x7c\xbb\x9b\xe7\xfa\xb3\x3a\x7f\xb1\xf7\xbc\xfd\x6d\xb6\xbb\xdb\x94\xb1\x6a\x00\x74\xbb\x32\x8a\xc9\xb4\x51\x18\x6d\x39\x3c\x05\xf6\xf6\xd6\x52\x31\xe4\xef\x61\xfd\x1f\x9a\xff\x61\x68\x36\x8f\x36\x73\xf3\x7a\x38\xf8\x27\x00\x00\xff\xff\x30\x60\xeb\xc9\x32\x1a\x00\x00"

func templatesViewsReleasesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsReleasesHtml,
		"templates/views/releases.html",
	)
}

func templatesViewsReleasesHtml() (*asset, error) {
	bytes, err := templatesViewsReleasesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/releases.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsUpdateHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x17\x5b\x6f\xdb\xbc\xf5\x3d\xbf\xe2\x4c\xed\x20\x1b\x8d\x25\xd8\xfb\x1e\x06\x7f\x96\x87\x61\xdb\x87\x15\xe8\xba\x2c\x97\xee\x31\xa0\xc5\x63\x8b\x2e\x45\xaa\x24\xe5\xc4\x31\xfc\xdf\x07\x8a\x92\x2d\x2b\x92\x2f\x0d\x56\xf4\xe1\x23\x90\x58\x3a\x3a\xf7\x1b\xcf\xd9\x6c\x80\xe2\x9c\x09\x04\x2f\x96\xc2\xa0\x30\x1e\x6c\xb7\x57\x13\xca\x56\x10\x73\xa2\x75\xe4\x29\xf9\xe4\x4d\xaf\x00\x00\xea\xd0\x58\xf2\x41\x4a\x07\xc3\x11\xd8\x27\x9d\x56\x4f\xcf\x7a\x30\x1c\x95\xf8\x4d\x9a\xe7\xc7\x8c\x08\xe4\xb5\xaf\xaf\x31\x2a\x2d\x0e\x71\x0a\xbc\x64\x34\xdd\x6c\x80\x0d\xff\x2c\xc0\x7b\xc8\x16\x8a\x50\x04\x85\x1c\x89\x46\x0f\x02\xd8\x6e\x27\x61\x32\x6a\xa1\xb3\xfc\x19\x8d\xbc\x27\xf6\x42\x14\xf5\x2a\x59\x73\xa9\xd2\x47\x07\x03\xf7\xf3\xb8\x42\x65\x58\xcc\xb1\x45\x7a\xc1\x29\xe7\x15\x31\x67\xda\x0c\x72\xa1\xcd\x9a\xe3\x8e\x5c\x1b\xcc\x74\x07\x6d\x41\xcf\xd9\x74\x42\x20\x51\x38\x8f\xbc\x77\x16\x7b\x30\xf4\xa6\x13\x9d\x11\x51\xf1\xb5\xc0\x47\x21\xbd\xe9\x70\x12\x5a\xf8\x74\x12\x92\xe9\x24\xe4\xec\x22\xae\xa3\x2e\xae\xa3\xb3\xb8\x4e\xc2\x9c\x77\x7c\xa9\x7c\x59\x29\xdf\xad\x55\xf2\x4b\x25\xfd\xce\x60\x76\xcf\x8c\x75\xeb\x2e\x7e\x16\x06\x43\x78\xc8\xb8\x24\x14\x66\x4c\x10\xb5\x86\x39\xe3\xfb\x48\xfe\x72\x84\xb7\x0d\x5d\x3d\x8c\x83\x44\x2a\xf6\x22\x85\x21\x1c\x8a\x77\x4e\x66\xc8\x07\x1c\xe7\x06\xa8\x92\xd9\x8b\x14\xe8\x81\x92\x1c\x1d\xbe\x07\x29\x9a\x44\xd2\xc8\xcb\xa4\x36\x5e\x61\x91\x34\xc4\xeb\x94\xe8\x0e\x89\x0d\x93\x22\xf2\x36\x1b\x08\x6e\xf1\x5b\x8e\xda\x04\x0f\xb7\x9f\x82\x1b\x62\x12\xd8\x6e\x3d\x10\x72\x45\x38\xa3\xc4\xe0\x74\x12\x5a\x41\x5d\xfe\xa5\x6c\x75\x8e\x83\x47\x6f\x74\xf0\x08\xbe\xa0\x62\xf3\x35\xe4\x85\x9f\x91\x9e\xef\x62\xba\x4b\x75\xca\x6b\xee\x3d\xa2\x90\x23\x33\x7b\x15\xbe\xa0\xd2\x4c\x0a\x27\x6e\x3c\x09\xa9\x39\x45\x4c\x2b\x99\x2b\x47\x3a\x58\x11\x9e\xa3\x37\x9d\x84\x94\x4e\xaf\xce\x97\xfc\xb7\x04\xe3\xaf\x3a\x4f\xbf\x47\x74\x5c\xd2\x7e\xaf\xec\xbf\xaa\x38\x61\x06\x63\x93\x2b\xbc\x44\xfe\x61\xc5\x12\x15\x27\x7b\x0d\x8a\x9a\x85\x02\xe1\x44\x8a\xba\x53\x67\x42\x89\x58\xa0\x82\xa2\x22\xdc\xff\x12\x54\x4b\x16\xa6\x41\x48\x63\x4b\x07\xe2\x5c\x29\x14\x06\x48\x96\x71\x16\x13\x9b\xf0\x40\x6a\x26\xc1\x1f\xb5\xb5\x4a\x30\xbe\xfb\x0b\x16\xd2\x62\x14\x39\x55\x68\x7a\xdc\xd4\x0b\xfd\x79\xc7\x5e\xde\xe0\x47\xcd\x5e\xb0\xe9\xc7\x1d\xeb\xd9\xda\xa0\xae\xaa\xc1\x6a\x75\x44\xe7\x8e\x86\x08\x8d\x0b\xac\xe8\x3e\x0b\x25\xf3\xec\x54\xa5\xcc\x72\x63\xa4\x00\xb3\xce\x30\xf2\xdc\xcb\xee\x6a\x9a\x19\x01\x33\x23\x06\x3a\x8f\x63\xd4\xda\x03\x4a\x0c\x19\x18\xb9\x58\xd8\x0e\x96\x4a\x4a\xf8\xa9\x5e\xe5\x8e\xa3\x23\x6a\x81\x26\xf2\xde\x5d\x4a\x58\xe0\x0f\x8c\xed\x2c\x45\xd3\x2b\x8b\x4b\x8a\x39\x53\x29\xe4\xe5\x15\x5c\xcb\x14\xe7\xcb\x8b\x25\xc4\x84\xf3\x19\x89\xbf\x46\x5e\x79\x9d\x97\xb7\x7b\x6f\x4e\xb8\xc6\xfe\xaf\xde\xab\x7b\xdf\x09\x3a\x91\x67\xce\xa9\xa7\x72\xed\x9c\x38\x3c\x11\x25\x98\x58\xfc\xe4\x71\x10\x14\x14\x6a\x43\x94\xf9\x7f\xc6\xc4\xa8\xbc\x35\x24\x75\xf9\x3f\x38\x3c\x65\x3f\xfb\x29\xa3\xa3\x30\x95\x2b\x6c\xbb\x80\xdf\x1e\x92\xdb\x82\x77\xef\x20\x1a\x0e\x76\x51\x00\xba\x31\xba\xc7\x94\xf6\x4f\x2d\xe0\x06\xa8\xf6\x5a\x3e\x96\x3f\x9b\x0d\xa0\xa0\x56\xe9\xab\xda\x4e\x92\x20\xa1\x5e\x65\xc9\x66\x03\xda\x10\xc3\xe2\x7f\xde\xff\xeb\x13\xf4\xdc\xf3\xc3\xed\x27\xf0\x42\x4a\x74\x32\x93\x44\xd1\x90\x68\x8d\x46\x87\x2b\x14\x54\x2a\x1d\x56\xd3\xdf\x52\x87\xb1\xde\xbf\x06\x29\x13\x41\x6c\x3b\xab\x6b\x31\x56\x44\xbb\x06\x4b\xed\x95\x1f\xdf\x22\x7b\xd9\x10\xbd\x6c\x4a\xbe\x98\xf9\xf2\x3f\x39\xaa\xf5\xe0\x2e\x25\xca\x0c\xfe\x5b\xec\x1f\x56\xca\xf2\x9b\x05\x07\xda\x82\x1d\xb4\x45\xde\xd5\x44\xc7\x8a\x65\xa6\xac\xaa\x5a\xb3\x08\x97\x64\x45\xdc\xc7\xf2\xfe\x7a\xdf\xa3\x32\xce\x53\x14\xa6\x1f\x28\x24\x74\xdd\x9b\xe7\xa2\x98\x83\xa1\xd7\x87\xcd\x2e\xb0\xef\x7b\xfe\x3b\xb7\x06\xf9\xfd\xba\xf8\x5e\xff\xd7\xab\x3a\x52\xe0\xb2\xee\x33\x3e\x1b\xbf\x1f\x24\x8c\xda\xfc\x6d\x41\xb8\x51\xb8\x62\x32\xd7\x47\x91\x7e\x63\x82\xe9\xa4\x86\xb2\xc3\xe1\x68\xaa\xdd\xf0\x23\x85\x08\x7c\xbf\xf6\xf1\xef\x55\x28\x64\x66\x0d\xd1\x81\x34\x04\xa2\x9a\x31\xf6\x64\x44\x11\xce\x91\xbb\x15\x45\x8f\x61\x78\xfd\xea\x7b\xfa\x99\xa4\x38\x06\xbf\x94\xe4\x1f\x62\xc4\x0a\x89\xc1\x8f\x29\x59\xe0\x7d\x92\xa7\x33\x41\x18\xd7\x63\x17\x87\x43\x4c\x12\xc7\x98\x19\xa4\xbf\x31\x8e\x7a\x0c\x7e\x3d\x22\x29\x89\xdd\x7a\x74\x5d\x87\x3e\x0f\x5a\x80\x2f\x2c\xbb\x0e\x66\x4c\x5c\x07\x2f\x2c\x6b\x28\x53\x8e\x12\x63\xd8\x87\x6f\x7e\x0d\xaa\xdf\xb0\xfa\x68\x28\xfd\x85\xbc\x97\x76\xb1\xf0\xaf\x61\x54\x0b\xc8\x69\x42\xca\x34\x99\x71\x2c\x69\x87\xf5\x48\x55\xa7\x1e\x2d\x15\x30\xda\x82\x62\x03\x7f\xb0\x19\xf8\xfd\xc0\xe0\xb3\xe9\xa9\x0a\xdc\xae\x54\x70\x38\xd3\xef\xa9\x2a\x78\x07\xd9\x7e\x08\xdf\x93\xd4\x67\xe1\x0e\xb2\xfd\xcc\xb9\x27\xb3\xb0\x36\xab\xd9\x1c\x1a\x4c\xe1\x0f\x11\xf8\x76\xc3\xdc\x4d\xd5\x7e\x5b\x90\x2a\x7f\xbb\x3d\x11\x82\xda\xb8\xd8\x52\x32\xc7\xa8\xca\xe1\xe6\x34\x55\x50\xdb\x27\x6c\x80\x13\xf9\xd4\x86\xbd\x05\xe4\x1a\x2f\x56\xb9\x8b\xdd\x71\x95\x4f\x50\x35\x54\xee\x32\xf0\xf0\xa6\xdc\xbf\x6d\x6b\x11\x7b\x62\x82\xca\xa7\xe0\x70\x1a\x82\xa8\x56\x4f\xe5\xf0\xd3\x8c\xd6\x8a\x28\xc8\x6d\x0b\x0a\xa5\x21\x61\x49\x1f\xfa\xf0\xa1\x96\xf2\x1f\xc0\x0f\xcb\x39\xce\x6f\x64\x49\x91\x21\xed\x9c\xed\xc9\xe1\x43\x04\xfe\x5f\x4a\x8c\x68\xe8\x37\x4c\x39\x78\x7d\x1f\x90\x25\x79\xee\xbd\xe6\x62\x2f\x82\x31\xf8\x37\xff\xbe\xbb\x6f\xf4\x8d\x42\x86\xe2\x63\xc8\x5f\xc3\x5b\x7a\x4a\x6b\x43\xd9\x99\x11\x28\xd4\x39\x37\x10\x45\x11\xf8\x73\xc2\x38\xd2\xce\xe4\xb6\x47\xe0\x13\xdc\x7c\x96\x86\xcd\xd7\x2d\x4a\x1f\x18\x60\xc7\xaf\x31\xf8\xff\x50\x4a\xaa\x16\x13\x0e\x70\xf1\xd9\x8c\x41\x05\x29\x6a\x4d\x16\x78\x02\xd9\x79\x06\xcf\xe0\x6b\xb3\xab\xb5\xbd\x37\x8f\x36\x6b\xce\xc4\x62\x0c\xfe\x4c\x4a\xa3\x8d\x22\xd9\x9f\xfc\x4e\x8a\x6d\x47\x7a\x1f\xad\x33\xf8\x2e\xef\xdd\x95\xd5\x78\x8e\xff\x2a\xe4\x6a\xff\x38\x45\xe4\xfc\xa8\xcf\x92\xf0\xa3\x3d\x79\xaa\x21\xd4\xe8\xba\x5b\x82\x9b\xbc\x0f\x3a\x42\x33\xb7\xdf\x52\x7f\xa7\xda\x87\x5b\x34\x5a\xa8\x7f\xaf\xd2\x83\xf3\xb3\x55\xe9\x59\xd3\xd6\xb0\x83\xfb\x71\x06\x87\x53\x57\xdb\xc4\x06\x97\x67\x7f\x05\x9a\x84\x6e\x4f\xa8\xad\x6e\xff\x0b\x00\x00\xff\xff\x58\x36\x7a\x9d\x4d\x1a\x00\x00"

func templatesViewsUpdateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsUpdateHtml,
		"templates/views/update.html",
	)
}

func templatesViewsUpdateHtml() (*asset, error) {
	bytes, err := templatesViewsUpdateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/update.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesOtaMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4d\x6b\xd4\x40\x1c\xc6\x9f\x6e\x57\x0a\xc1\x37\x44\x3c\x79\xf8\x7b\xb0\x28\x32\x35\xa9\x1e\x64\x76\xb3\xeb\x5b\x0b\x62\x17\xa5\xa4\xde\xc7\xcd\x34\x0d\x26\x33\x61\x66\x22\x16\x7a\x10\x4f\x8a\x5e\x05\x4f\xe2\xc9\xbb\x82\x76\x45\x5c\x0f\x7e\x81\xd9\x2f\xe0\x67\x91\x6c\xb6\x8a\x03\x03\xcf\xef\xf7\x7f\xe6\xe5\xf7\x99\xee\x5b\x00\x38\x06\xe0\x3c\x80\xbb\x00\x56\x00\x14\x68\xd7\x4b\x00\x27\x01\xbc\x02\x70\x02\xc0\x3b\x00\xa7\x00\x7c\x04\x70\x1c\xc0\x57\x00\xc3\x25\xe0\x17\x80\xe5\x66\x77\x5a\xbf\xd2\x01\xce\x01\x38\xdb\x01\x4e\x03\xb8\xd2\x01\x96\x16\x77\x76\x01\x74\x16\x6f\x2e\x2f\x1c\xc6\x5a\xed\xe6\x19\x73\xe2\x71\x57\x3b\x81\x52\xaa\xba\xbb\x2d\x0b\x29\xac\xb4\x47\x54\x69\x9b\x3b\x6d\xf6\x5b\xde\xa9\x32\x23\x52\x89\xc6\x1b\xc7\x46\x36\xcb\x53\x76\xbb\xce\x2c\x4b\x34\xa7\x54\x3e\xbd\xf9\x24\xdf\x13\xa5\x5e\x33\x75\xb0\x25\xac\x63\x89\x11\xca\x16\xc2\x69\xc3\xe9\xfe\x7c\x44\xa3\xda\x88\x52\xa7\x9a\xfa\xff\xf5\x07\xc1\x96\x50\x59\x2d\x32\xc9\x12\x29\x4a\x4e\x7f\x99\xd3\x76\x6d\x6d\x2e\x54\x30\xba\x37\xda\x60\x8f\xa4\xb1\xb9\x56\x9c\xa2\xb5\x30\xb8\xa3\x95\x93\xca\xb1\x64\xbf\x92\x9c\x9c\x7c\xe6\xae\x56\x85\xc8\x55\x8f\xc6\x7b\xc2\x58\xe9\xe2\x9d\x64\x93\xdd\xf8\xd7\x6b\xfe\xb3\x2b\x0d\xdb\x50\x63\x9d\xe6\x2a\xe3\x14\x3c\x2c\x6a\x23\x0a\xb6\xa9\x4d\x69\x39\xa9\x6a\x8e\x36\xbe\xd6\xa3\x36\xc6\xea\x62\x14\xc6\x71\x44\xab\xab\xd4\xc4\xf0\x42\x1c\x45\x34\xa4\x90\xf8\x9c\x07\xf1\xfa\xd1\xa8\x1f\x5f\x6f\xe2\xa5\x79\xad\x1f\x85\x74\x70\xd0\x1e\x19\xc4\xeb\xe1\x65\x1a\x52\x44\x9c\xd6\x7b\x78\x90\xdc\x82\xff\xe0\xbf\xf8\xef\x7e\xe2\x0f\x67\xaf\x5b\xf8\xe9\xa7\xfe\xd0\x4f\x66\x2f\xfc\x74\xf6\xdc\x4f\xfc\x37\xf8\xf7\xfe\x93\xff\xe1\xa7\xfe\x73\xa3\x67\x6f\xf0\x27\x00\x00\xff\xff\xb2\xdb\xf7\x31\x3c\x02\x00\x00"

func localesRuLc_messagesOtaMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesOtaMo,
		"locales/ru/LC_MESSAGES/ota.mo",
	)
}

func localesRuLc_messagesOtaMo() (*asset, error) {
	bytes, err := localesRuLc_messagesOtaMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/ota.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesReleasesMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\xdf\x6b\x1c\x55\x14\xfe\xf6\x97\xd6\x55\xa3\x54\x50\x14\x85\x23\xd2\x52\x1f\xa6\xee\x46\x05\x99\x64\x13\x63\xda\x42\xb1\x81\x12\x37\xbe\x89\x5c\x76\x6f\x77\x87\xee\xce\xac\x77\xee\xf4\x87\xf6\x61\x7f\x20\x22\x8a\x15\x51\x7c\x11\xab\x88\xf8\xba\xad\xd9\x64\x9a\xa4\x9b\xbf\x40\x3c\xf7\x1f\x28\xbe\xfa\x0f\xf8\x2c\x77\x66\x36\x6d\x21\x88\x2f\xbd\x2f\x73\xce\xf9\xce\xf7\x9d\xef\xdc\x3b\x77\x8f\x16\xbf\x03\x80\xa7\x01\xbc\x08\xe0\x07\x00\xcf\xd9\x38\x87\xe4\x5c\xc9\x01\x8f\x02\xb8\x9a\x03\x9e\x00\x30\xc8\x01\x47\x00\x7c\x95\x03\x9e\x07\xf0\x7d\x0e\x78\x09\xc0\xcd\x1c\xf0\x26\x80\x3f\x33\xbc\x94\x07\x8a\x00\xe6\xf2\x69\xfe\x4c\x1e\x78\xc4\xea\x66\xf5\x57\xf2\x00\x01\x38\x91\x07\x1e\x07\x70\x36\x9f\xce\xd9\xc8\xbe\x1f\xe4\x81\xe5\x1c\x20\xf3\xa9\xb7\x5f\x0b\xc0\xb3\x00\xb6\x0a\x69\xfe\x47\x01\x58\x00\x70\xb7\x00\x9c\x02\xf0\x42\x11\xf8\x08\xc0\x87\x45\x60\x0e\x80\x29\xa6\x73\xff\x2a\xa6\xbe\xff\xce\xea\xff\x64\x79\xa9\x04\xac\x00\x38\x5a\x4a\xf7\xa8\x97\xd2\xba\x5f\x4a\xfb\x3e\x29\x01\xf6\x0a\x9e\x02\x50\x4e\xaf\x22\xf1\x5d\x42\xda\x67\x77\x29\x64\xf5\x39\xdc\x3b\x47\xf0\xe0\xc9\x03\x78\x12\xc0\x63\x48\xf7\xb2\xde\xed\xbe\x58\x69\x68\x2f\xf0\x43\xac\xa8\x46\xdb\xd3\xb2\xa1\x23\x25\xb1\xda\x96\x8d\x8b\x61\xd4\xc5\x6a\xe0\x5f\xf0\x54\x97\x94\xec\x06\x97\x24\x29\xd9\x91\x22\x94\x74\x2c\x3c\x40\xa2\x5e\x4b\x89\xa6\x24\x1d\xfc\x6f\x94\x84\xdf\x24\x25\x43\x2d\x94\x26\xd1\xeb\x75\xbc\x86\xb0\x26\x70\x2a\xb8\xec\x77\x02\xd1\xc4\x79\xa1\xdb\x58\x4f\x09\x21\xd6\x93\xe1\x78\xcf\xfb\x58\x62\x23\x13\xbc\x8f\x46\x97\x3d\xdd\x9e\xe9\x61\xa3\x67\x15\x64\x93\x84\xc6\xfb\x52\x85\x56\xb7\x11\x29\x25\x7d\x8d\x75\xd9\x0b\x94\x76\xd6\xc2\x96\xd7\x74\xde\x89\x5a\xa1\x53\x0f\x5c\x6a\xca\x4b\x6f\x5f\xf4\xda\xa2\x1b\x9c\x54\x51\xf9\x9c\x08\xb5\x53\x57\xc2\x0f\x3b\x42\x07\xca\xa5\x77\x13\x88\xd6\x22\x25\xba\x41\x33\xa0\xc5\x07\xfa\x97\xca\xe7\x84\xdf\x8a\x44\x4b\x3a\x75\x29\xba\x2e\x1d\xe4\x2e\xad\x47\x61\xe8\x09\xbf\xbc\x76\x76\xed\xb4\x93\x79\x71\xa9\x7a\xb2\x52\x5e\x0d\x7c\x2d\x7d\xed\xd4\xaf\xf6\xa4\x4b\x5a\x5e\xd1\xaf\xf5\x3a\xc2\xf3\x17\xa8\xd1\x16\x2a\x94\xba\xb6\x51\x3f\xe3\xbc\x75\xaf\xcf\xfa\xb9\x20\x95\x73\xda\x6f\x04\x4d\xcf\x6f\xb9\x54\x3e\xdf\x89\x94\xe8\x38\x67\x02\xd5\x0d\x5d\xf2\x7b\x49\x1a\xd6\x5e\x5f\xa0\x34\xac\xf9\xc7\xaa\x95\x5a\xad\x4a\xc7\x8f\x93\x0d\x2b\x2f\xd7\xaa\x55\x5a\xa6\x0a\xb9\x49\xbe\x54\x9b\x9f\x41\x8b\xb5\x37\x6c\x78\x22\x69\x5b\xac\x56\xe8\xda\xb5\x94\xb2\x54\x9b\xaf\xbc\x4a\xcb\x54\x25\x97\xe6\x17\xc0\xdf\xf2\x84\x6f\x9b\x81\x19\xf2\x2d\x8e\xcd\x75\xf0\xd7\xa6\x6f\x3e\xe5\xd8\x0c\x79\xc2\x3b\x66\x68\x46\xa6\xcf\x63\xf0\x6f\x49\x3a\x30\x23\xde\xe3\x3d\x5b\xb8\xc1\x53\xde\x4c\x68\x13\xd3\xe7\x2d\xde\xe4\x09\xdf\xe1\x98\x27\x64\x46\xbc\xc9\x63\xde\x4d\x0b\xe6\x3a\x99\x3e\x4f\x78\x97\x63\xde\xe6\xb1\xfd\x95\xfe\x83\xcb\x53\xbe\xc9\x77\x78\xca\xb7\xee\xe3\xf3\x26\x4f\x1f\x8a\x08\x71\x4c\xbc\x9f\xf0\x27\xb6\xc4\xfb\x66\x64\x06\xbc\x43\xbc\x6f\xfa\x1c\xf3\x2e\x4f\x79\x6b\xa6\x00\xfe\x99\x77\x78\x6c\x3e\xe3\xb1\x19\x9a\x2f\xc1\x37\xcc\x28\x0d\x7e\x9a\x89\x9a\x2f\xc0\xbf\x64\xcb\xc7\x33\x6c\xcc\xdb\xbc\x67\x67\x80\x7f\x3c\xf0\x95\xa0\x87\x8c\xb1\xd7\x37\x38\xd4\x13\x4f\x79\xcf\x3e\xd7\xd8\x0c\x79\x4c\x09\xf2\xbb\xe9\x9b\x11\x6f\xf3\x0e\xc7\xe0\x6f\x2c\xc5\x0c\x12\xa7\xd9\xe3\x8d\xcc\xe7\x1c\xf3\x6d\xfc\x1b\x00\x00\xff\xff\xea\xd6\x9e\x47\x8a\x05\x00\x00"

func localesRuLc_messagesReleasesMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesReleasesMo,
		"locales/ru/LC_MESSAGES/releases.mo",
	)
}

func localesRuLc_messagesReleasesMo() (*asset, error) {
	bytes, err := localesRuLc_messagesReleasesMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/releases.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesUpdateMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x4d\x6f\x1b\x55\x14\x3d\x43\xcc\x97\x25\x3e\x04\x62\x81\x60\x71\x11\x6a\x81\xc5\x04\x8f\x41\x80\x26\x99\x04\x51\x5a\xa9\xa2\x16\x28\x38\x5d\xb1\x79\xd8\xcf\xce\x88\xf1\x7b\xa3\x37\x33\x15\x91\xba\x30\x36\x48\xa0\x22\x85\x0d\x62\x83\xe8\x86\x05\x4b\xb7\xb5\x1b\x37\x26\xce\x0f\x60\x73\xdf\x1f\x60\xc3\x82\x2d\xfc\x03\x34\x33\xae\x49\x49\x40\xf4\x6d\xe6\x7e\x9c\x73\xee\xb9\x57\x9a\x5f\x9f\xaa\x7c\x0b\x00\x55\x00\xcf\x03\x88\x00\x3c\x06\xe0\x67\x94\xef\x8f\x45\xfd\x4f\x00\x2f\x01\x78\xdc\x01\x9e\x03\xf0\x86\x03\x3c\x0b\x60\x6b\x91\x47\x0e\xf0\x30\x80\x2f\x1d\xe0\x69\x00\xdf\x38\xc0\x13\x00\x7e\x74\x80\x55\x00\x63\x07\xd8\x74\x80\xdf\x1c\xe0\x7d\x00\x2f\xae\x00\x1f\x01\x48\x56\x80\x77\x01\xf0\x0a\xf0\x26\x80\x67\x2a\x40\x00\xa0\x51\x01\x9e\x04\x70\xad\x52\xf2\xbf\xaf\x94\x73\x7e\xa9\x00\x17\x01\xfc\x5e\x01\x1c\x00\x0f\x01\xa8\x2c\xbc\xae\x00\x78\x60\xb1\xcb\xa3\x00\x1e\x44\xe9\xe9\x11\x1c\x7b\xe7\xb4\xea\x84\xa6\x47\x46\xf6\xf4\x15\x49\x59\x1c\x69\xd1\x96\x6d\xea\x84\x91\x5c\x36\xb3\xb8\x6b\x44\x5b\x92\x50\x6d\x32\x32\x49\x85\x49\x49\xc4\x71\x14\xb6\x44\x1a\x6a\x75\x12\x77\xac\xf7\x61\x2a\x63\xf2\x68\xbb\x10\xa6\x8f\x43\x25\xcc\x6e\xa9\x5e\x74\xea\x74\x59\x9a\xb0\xb3\xfb\x8f\xc9\xdb\xa5\xd2\xdd\xef\xf1\xc9\xcb\x9a\x91\x91\x14\x89\x44\x98\x90\xd2\x29\x75\xb4\xa1\x56\x66\x8c\x54\xf7\x98\x23\x61\x5a\x3b\x61\x2a\x5b\x69\x66\x24\x9d\x49\xb0\x25\x63\x6d\x52\xb7\x91\x74\xc3\xb6\xfb\x4e\xd6\x4d\xdc\xa6\xf6\xa9\x2d\xaf\xbc\xfd\x49\xb8\x23\x7a\x7a\xd5\x64\xd5\x4b\x22\x49\xdd\xa6\x11\x2a\x89\x44\xaa\x8d\x4f\xef\x15\x2d\x6a\x64\x46\xf4\x74\x5b\xd3\xfa\x3d\xf8\x8d\xea\x25\xa1\xba\x99\xe8\x4a\xb7\x29\x45\xcf\xa7\x65\xee\xd3\x56\x96\x24\xa1\x50\xd5\xc6\xc5\xc6\x79\xf7\xb2\x34\x49\xa8\x95\x4f\xde\x6a\xad\x7a\x4e\xab\x54\xaa\xd4\x6d\xee\xc6\xd2\xa7\x54\x7e\x9a\xbe\x1a\x47\x22\x54\x6b\xd4\xda\x11\x26\x91\x69\xb0\xdd\xbc\xe0\xbe\xf5\x37\x2e\xf7\xd3\x91\xc6\x3d\xaf\x5a\xba\x1d\xaa\xae\x4f\xd5\x0f\xa2\xcc\x88\xc8\xbd\xa0\x4d\x2f\xf1\x49\xc5\x45\x9a\x04\xaf\xad\x51\x19\x06\xea\x8c\x57\x0b\x02\x8f\xce\x9e\xa5\x3c\xac\xbd\x10\x78\x1e\x6d\x52\x8d\xfc\x22\xdf\x08\xea\x77\x5b\xeb\xc1\xeb\x79\xf8\x72\x01\x5b\xf7\x6a\x74\xf5\x6a\x49\xd9\x08\xea\xb5\x57\x68\x93\x3c\xf2\xa9\xbe\x06\xbe\xce\x73\x1e\xdb\x01\xdf\xe4\x89\xed\xf3\x6d\x1e\xf3\x84\x0f\x79\xca\x13\xb2\x43\x1e\xf3\x88\x67\x65\xc1\xee\x11\xef\xf3\x88\x6f\xd9\xbe\x1d\xf2\xed\xa2\x78\xc8\x73\xbe\xc5\x73\xb2\x9f\xf3\x88\xef\xf0\x8c\x47\xff\xa5\xc7\x73\xbe\x51\x50\x6e\x1e\xd7\x9c\x12\x1f\x15\xd0\x49\x21\x7f\x64\x87\xf6\x33\x3e\x20\x3e\xb2\x7d\x9e\xf2\x8c\xe7\x8b\x59\x53\xbb\x77\xff\xe2\xa7\x8a\xfc\x94\x6f\x41\x1e\xf1\x77\xcb\x75\xf6\xf9\x80\x47\xc4\x37\x78\xca\x87\x3c\xb2\xfd\x53\x37\x2b\x79\x75\xe2\xeb\xb6\x5f\x0c\xca\x3d\x14\xbc\xff\x73\x97\x1f\x96\xfe\xa6\x76\x60\xbf\x3e\x51\x38\xfd\x12\x76\xf0\x6f\xf0\x02\x37\xe3\x29\xef\x83\x0f\xf3\x03\x8c\x79\x66\xf7\xc8\x0e\x78\xc2\x07\x76\x68\xbf\xe2\x09\xdf\xa1\x7c\x19\xfb\x45\x4e\x29\xca\x03\x3b\xb4\x7d\x7b\xed\xd4\xc3\xe4\x7f\xd3\x5f\x01\x00\x00\xff\xff\xa4\xe9\x47\xb0\x31\x05\x00\x00"

func localesRuLc_messagesUpdateMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesUpdateMo,
		"locales/ru/LC_MESSAGES/update.mo",
	)
}

func localesRuLc_messagesUpdateMo() (*asset, error) {
	bytes, err := localesRuLc_messagesUpdateMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/update.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/views/releases.html":      templatesViewsReleasesHtml,
	"templates/views/update.html":        templatesViewsUpdateHtml,
	"locales/ru/LC_MESSAGES/ota.mo":      localesRuLc_messagesOtaMo,
	"locales/ru/LC_MESSAGES/releases.mo": localesRuLc_messagesReleasesMo,
	"locales/ru/LC_MESSAGES/update.mo":   localesRuLc_messagesUpdateMo,
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
	"locales": &bintree{nil, map[string]*bintree{
		"ru": &bintree{nil, map[string]*bintree{
			"LC_MESSAGES": &bintree{nil, map[string]*bintree{
				"ota.mo":      &bintree{localesRuLc_messagesOtaMo, map[string]*bintree{}},
				"releases.mo": &bintree{localesRuLc_messagesReleasesMo, map[string]*bintree{}},
				"update.mo":   &bintree{localesRuLc_messagesUpdateMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"releases.html": &bintree{templatesViewsReleasesHtml, map[string]*bintree{}},
			"update.html":   &bintree{templatesViewsUpdateHtml, map[string]*bintree{}},
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
