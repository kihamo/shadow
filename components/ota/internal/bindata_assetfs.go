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

var _templatesViewsReleasesHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x6d\x8f\xdb\x36\x12\xfe\x9e\x5f\x31\xd0\x6d\x60\x2f\xb2\xb2\xec\xbd\xcb\xdd\xc1\xb1\x36\x08\x92\x16\x0d\x90\xb6\x41\xb2\xdb\x0f\xfd\x12\x8c\x45\xca\xe2\x86\x26\x55\x72\x64\xaf\x6b\xf8\xbf\x17\xd4\x8b\x2d\xdb\xf2\xcb\xba\xde\xa0\x28\x2a\x20\x59\x9a\xe4\xcc\x70\x9e\x79\x66\x34\xd4\x7c\x0e\x8c\xc7\x42\x71\xf0\x22\xad\x88\x2b\xf2\x60\xb1\x78\x36\x60\x62\x02\x91\x44\x6b\x43\xcf\xe8\xa9\x77\xf3\x0c\x00\xa0\x3e\x1b\x69\xe9\x8f\x99\xdf\xbb\x06\x37\xb2\xe3\x6a\xf4\x60\xfd\xde\x75\xb9\x7f\x53\xe6\xe1\x4b\x8a\x8a\xcb\xda\xea\xf6\x0e\x12\x24\xf9\xc6\x8e\x7c\x57\x72\x7d\x33\x9f\x83\xe8\xfd\x5f\x81\xf7\x89\x4b\x8e\x96\x5b\x0f\x3a\xb0\x58\x0c\x82\xe4\xba\x41\x20\x93\x95\x56\x85\x13\x50\x38\x19\xa2\xf1\x8d\x18\x25\x04\xf9\x31\xbe\x90\xd6\x72\xa8\x1f\x1a\x8c\xe5\xf2\x52\x34\x2f\xe4\x8b\x08\x89\xe1\x71\xe8\x05\x9a\x30\x78\x9d\xa5\x0c\x89\x87\xbd\x1d\xaa\x96\x52\xa2\x3a\xd1\x48\xce\xd2\x44\x44\x5a\xc1\x72\xe4\x1b\x1e\x1b\x6e\x13\x0f\x72\x08\x42\x6f\xe9\xed\x5d\xae\x1d\x0c\x4f\xb5\x15\xa4\x8d\xa8\x1c\xf7\x80\x21\xa1\x4f\x7a\x34\x72\x02\xce\x1f\x12\x69\x39\x9b\x4a\x8c\xf8\x98\x2b\x0a\xbd\xa1\x26\xd2\x63\xef\x66\x10\xec\x73\x29\xc0\x1d\x40\x04\x4d\x48\x0c\x82\x4c\x36\xcc\xd6\x19\x22\x39\x9a\x58\x3c\x38\xbb\x4c\x4c\x36\x82\xde\x30\xb5\xc6\x83\x8a\x8d\xfb\x6d\x10\x0e\x25\xf7\x0d\xb7\xa9\x56\x56\x4c\x9a\x88\x93\x8b\xe4\xfb\xd6\x84\xa0\x10\xb5\x64\x44\xca\x59\x0e\x59\x31\xcf\xa8\xa6\x0f\x94\x9e\x1a\x4c\x3d\xb0\x34\x73\x10\x4f\x05\xa3\xa4\xdf\xeb\x76\x9f\xef\x09\xf5\x80\x12\x8e\x6c\xdf\xba\x39\xc0\x13\x4a\x56\x5c\xff\x85\x1b\x2b\xb4\xaa\xa8\x4e\xc9\x23\x64\xdf\x44\x24\xb4\xb2\x27\xc9\xbe\x35\x1c\x89\x33\x40\x3a\x49\xfc\xb3\xf8\x9d\x9f\x66\x37\xe1\xd1\x57\x9b\x8d\x4f\x73\xd8\x44\x89\x20\x1e\x51\x66\x4e\xb3\xfe\x11\x29\x39\x46\x70\x10\xec\x0b\xa2\x93\x3d\x40\x81\xa1\x66\xb3\xdd\xeb\xf3\x39\x18\x54\x23\x0e\x17\xe2\x0a\x2e\x4c\x51\xee\xa0\x1f\x42\xa7\x1c\x5b\x57\xa5\x77\x6b\x3f\x48\x30\xe6\x5c\xae\x14\x77\x4a\x8e\xc1\x62\xe1\x2c\x8b\x78\xb5\xf2\xde\xbe\xcd\x8c\xe1\x8a\x1c\x22\x36\x45\x55\x25\x91\xc4\x21\x97\x90\xff\xef\xdb\x2c\x8a\xb8\xb5\xde\x0a\xc6\xa8\x10\xf2\xe0\x22\x47\xd2\x09\xba\x45\xae\x58\x81\xec\x1e\x68\xaa\xf3\xed\xdd\x00\x1b\x75\x60\x48\xca\x1f\x19\x9d\xa5\x1e\x18\xed\x12\xb5\xf8\x71\x58\x09\xd4\x8b\x79\x1d\x92\x77\x7a\xaa\xa4\x46\x76\xf7\xe9\x43\x5e\x69\x09\xcd\x88\x53\xe8\x7d\x19\x4a\x54\x5f\xbd\x9a\x61\x70\xc6\x4b\x08\xf2\x71\x5e\xd7\xdd\xe0\xc1\x1e\x79\x02\x58\x7b\x39\xc4\x68\x21\x46\x3f\x16\x92\xfb\xac\x3c\xc6\xf6\x5b\xe1\xdd\x72\xc5\x61\x7c\xa0\xb8\xaf\x19\x72\x85\xfe\xa8\x9d\x5b\x5c\xb8\x4b\x47\x06\x19\xcf\xeb\xe4\x1e\xfe\x6d\x59\x1c\x66\x44\x5a\x01\xcd\x52\x1e\x7a\xc5\x8f\x2d\x04\x99\xe3\xbb\xd9\x02\x70\xfd\x15\x37\xd6\x0c\x65\x35\x57\x46\xe4\x5f\xf5\xc9\x7c\xec\x6f\x62\xf5\x56\xab\x58\x98\x31\x64\xc5\xf9\x81\x34\x54\x39\xf5\xdc\x3a\x00\x95\x90\xcb\x7f\x0d\x59\xb1\xa6\x3c\x42\x29\x87\x18\x7d\x0d\xbd\x72\x63\x89\x4a\xbb\x55\xe7\xcf\xfb\x77\xb0\x58\xb4\xae\x20\x46\x69\xf9\xe5\xab\x47\x10\x01\x9a\xc8\x80\xc6\xe8\xa9\x1f\x09\x13\x49\xee\x3b\x9a\x6f\x37\x09\x85\x6b\x98\xa6\x52\x44\x48\xf9\x2b\xe3\x91\xcc\xc8\x2d\x07\x45\x7c\x1e\x21\xf1\x97\x8e\x2e\xa0\x62\x60\xb8\x25\x34\xb4\x09\xce\xd3\x45\x9d\x4c\x76\x86\xa0\x47\x52\x67\xcc\xcf\x52\x97\xe6\x3e\x4a\x3a\x2a\xea\x30\x15\x94\x54\x1e\x7f\x03\x0a\x2c\xcb\xfa\x89\x35\xe5\x13\x1f\xeb\x89\xab\x28\xe7\x2e\x28\x53\x34\x4a\xa8\xd1\xd3\x72\xce\xb8\xd3\xf3\x73\x56\x93\x1c\x8f\x46\x5a\xfd\x79\x42\x91\xc1\xa6\x0b\x46\x61\xf2\x9b\x72\xe5\xb0\xde\xad\xfb\xc1\xf6\x96\xb3\xb4\x11\x1b\x84\xbc\xcb\xb3\x8d\xb3\x37\x74\xec\x1b\x6e\x60\x23\x23\x52\x2a\xc9\x58\xcb\xc4\xe0\x1e\x27\x58\x2c\x7a\x37\x4c\x47\x99\xbb\x8d\x75\xa6\x46\x10\x6f\xbb\x1b\xdd\xad\xfe\x4c\x46\xa8\xd1\x7a\xac\x57\x07\xe8\x7c\xaf\xcd\x18\x09\xbc\xeb\x6e\xf7\xbf\x7e\xb7\xe7\x77\xaf\x6f\x7b\x2f\xfb\xdd\xff\xf4\xbb\x2f\x7f\xed\xfe\xaf\xdf\xed\x7a\x39\x2d\x2e\x07\x41\x61\xe5\x28\x67\x8f\x08\xc0\x71\xc8\xd6\x0f\xed\xba\xfd\xa3\x3b\xbb\xba\x60\xd5\xed\x9f\x24\x5c\xef\xf6\x4f\x52\xe0\xba\xfd\x83\x82\xfb\x5b\xfd\xc3\x88\x0e\x82\x3d\xcd\xfe\x20\xc8\x6f\x9d\x4d\xf7\xeb\x43\x77\xe6\xda\xcf\x72\x58\xfe\xa9\x55\xe4\xda\xe7\x1d\x77\x21\xf1\xaa\x73\xce\xe7\x60\x09\x49\x44\x3f\xdc\xfe\xf8\x01\xda\xc5\xd8\xb5\xba\x5e\xc0\xd0\x26\x43\x8d\x86\x05\x68\x2d\x27\x1b\x4c\xb8\x62\xda\xd8\x60\x79\x45\xb6\x1d\xc5\xc9\x1f\xda\x20\xb2\xc5\xec\x6d\x31\x3b\xd4\x9a\x2c\x19\x4c\x3b\x63\xa1\x3a\x91\xb5\x5e\xd9\xfd\x9c\xd1\xea\xea\x6a\x5e\x1d\x60\x35\xb3\xff\x00\xcd\xa8\xdc\xdb\x33\x62\x12\xdc\xdb\xe0\xfe\xb7\x8c\x9b\x59\xa7\x06\x8b\x3b\xcb\xfd\x53\x60\x31\xb4\xce\xe0\xce\x00\x3c\x89\xcd\x15\xda\x1b\xb6\x6b\x61\xf8\x06\xc6\x4b\xdf\x77\xc6\x7e\xdd\x7c\x91\x23\x47\x95\xe9\x65\x76\x5d\xb4\xab\x8a\x7d\xd9\x31\x1c\xd9\xac\x1d\x67\x2a\xff\x9c\x02\xed\x4b\x98\xaf\x25\xe6\x54\x28\xa6\xa7\x9d\xf5\xc6\x10\x42\x58\x49\x08\x76\x55\xb5\x64\x9b\xc2\xee\x99\xa0\x81\x0c\x42\x68\xe5\x9f\x13\x4b\x3d\x41\x0b\x5e\x80\x60\xf0\x02\x5a\x41\xd9\xda\xb6\x5e\x6d\x37\x59\x22\x86\xf6\x1e\xd5\xee\xc9\xe0\x45\x08\xad\xd7\xe5\xae\xb0\xd7\xda\xda\xd5\xd0\xbc\x5d\x74\xf0\x1e\x1f\xda\xcd\x1a\x1d\x88\x7d\x68\x7d\xfc\xf9\xf3\x6d\xeb\xaa\xd9\xa6\x91\x7d\xc8\x9a\xd7\xca\xcb\x72\xbf\x06\x90\xd9\x75\xf6\xa5\x8b\x8e\x61\x99\x24\x08\xc3\x10\x5a\x31\x0a\xc9\x59\x6b\x9f\x90\x7b\x14\x9f\xc2\xc7\x9f\x34\x89\x78\xb6\xc3\x91\x35\xa7\x5c\x5b\xd4\x87\xd6\x77\xc6\x68\xb3\xc3\xad\xb5\xfd\xfc\x81\xfa\x60\x3a\x63\x6e\x2d\x8e\xf8\x11\x02\x05\x6a\xfc\x48\xfd\x89\x60\xbc\x5f\xd0\xf8\xf0\x66\x4b\x33\x29\xd4\xa8\x0f\xad\x65\x26\xfc\x7b\x3b\xce\xf5\x67\x71\xf9\x6a\xe7\x7a\xf3\xdb\x6c\x7b\x76\x53\xc7\x62\x83\xa0\xeb\x99\x51\xf4\x99\x1b\x89\xd1\x14\xc3\x73\x70\x6f\x67\x2e\x15\x2d\xfb\x0e\xd1\x7f\xa8\xf9\x37\xa6\xe6\xe6\xd2\xaa\x6f\x5e\x36\x07\x7f\x04\x00\x00\xff\xff\xb0\x3b\x89\x69\x10\x1b\x00\x00"

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

var _localesRuLc_messagesReleasesMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\x4d\x6f\x1b\x55\x14\x3d\xe3\xd8\xb4\x18\x28\x5f\x02\x84\x44\xa5\xcb\xa2\x81\x4a\x4c\xb0\x53\x10\x68\x92\x49\x48\xdd\x56\x42\x34\xa2\x0a\x0e\xac\x9f\xec\x57\x7b\x54\x7b\xc6\xbc\x79\xd3\x14\xd4\x85\x3f\x84\x68\x05\x52\x01\xc1\xb6\x05\x81\xd4\xad\xa9\xe2\xc6\xf9\xa8\x23\x36\x88\xe5\x7d\x4b\x36\xf0\x17\x58\xf0\x03\xd0\x7c\x38\x6d\x55\x84\xd8\xf4\x6d\xe6\xde\x7b\xce\x7d\xe7\xdc\xfb\xe6\x8f\x67\xf2\xdf\x01\xc0\xd3\x00\x5e\x02\xf0\x13\x80\x17\x00\xbc\x66\x21\x39\x57\x2c\xe0\x10\x80\xab\x16\xf0\x38\x80\xaf\x2d\xe0\x30\x80\xeb\x16\xf0\x22\x80\x9b\x16\x70\x14\xc0\x2f\x16\xf0\x26\x80\x3f\x2d\xa0\x08\xe0\xb9\x5c\xca\xa3\x1c\x90\x07\x70\x3c\xcb\xe7\x72\xc0\x23\x00\xde\xca\xea\x4b\x39\xe0\x59\x00\x95\x1c\x40\x00\x3e\xca\xa5\x7a\x97\xb2\x6f\x37\x07\x2c\x5b\xc0\x95\x1c\xf0\x14\x80\xdf\x66\x80\xe7\x01\xfc\x3e\x93\xe6\x7f\xcd\x00\x0b\x00\xac\x3c\x70\x0a\x80\x9b\x07\x3e\x06\xd0\xcf\xa7\xfe\xfe\xce\x03\x47\x00\x1c\x29\xa4\xfa\x47\x0b\xe9\x1c\xb3\x85\xb4\x7e\x22\xcb\x4f\x16\x80\x57\x00\xbc\x5f\x00\x56\x00\x6c\x64\xf5\xed\x8c\xf7\x6b\x01\x88\x57\x12\x6b\x3e\x96\xae\x26\xf1\x5f\x00\xf0\x04\xd2\x99\x66\x90\x72\x0f\xe3\xee\x79\x14\xf7\x9f\x1c\x80\x27\x91\xee\xe8\x50\xb6\xf7\x58\x07\x2b\x35\xed\x05\x7e\x88\x15\x55\x6b\x7a\x5a\xd6\x74\xa4\x24\x2a\x4d\x59\xbb\x10\x46\x6d\x54\x02\xff\xbc\xa7\xda\xa4\x64\x3b\xb8\x28\x49\xc9\x96\x14\xa1\xa4\x63\xe1\x01\x12\x75\x1a\x4a\xd4\x25\xe9\xe0\x7f\xa3\x24\xfc\x3a\x29\x19\x6a\xa1\x34\x89\x4e\xa7\xe5\xd5\x44\x6c\x02\x15\x25\x85\x96\x75\x12\x1a\xa7\x82\x0d\xbf\x15\x88\x3a\xce\x09\xdd\xc4\x5a\xda\x1b\x62\x2d\xf1\x81\x0f\xbc\x4f\x25\xd6\x3b\x75\xa1\x63\x4f\x9d\x20\xf4\x74\xa0\x3c\x19\x62\x3d\xd3\xbb\xe7\x56\xda\xf0\x74\x73\x2a\x87\x0f\xa5\x0a\x63\xa9\x5a\xa4\x94\xf4\x35\xd6\x64\x27\x50\xda\x5e\x0d\x1b\x5e\xdd\x3e\x19\x35\x42\xbb\x1a\x38\x54\x97\x17\xdf\xb9\xe0\x35\x45\x3b\x98\x53\x51\xf1\xac\x08\xb5\x5d\x55\xc2\x0f\x5b\x42\x07\xca\xa1\xf7\x12\x88\x56\x23\x25\xda\x41\x3d\xa0\xc5\xfb\xf8\x4b\xc5\xb3\xc2\x6f\x44\xa2\x21\xed\xaa\x14\x6d\x87\x0e\x72\x87\xd6\xa2\x30\xf4\x84\x5f\x5c\x7d\x77\xf5\xb4\x9d\x79\x71\xa8\x3c\x57\x2a\x56\x02\x5f\x4b\x5f\xdb\xd5\x4f\x3a\xd2\x21\x2d\x2f\xe9\xd7\x3b\x2d\xe1\xf9\x0b\x54\x6b\x0a\x15\x4a\xed\xae\x57\xcf\xd8\x6f\xdf\xe5\xc5\x7e\xce\x4b\x65\x9f\xf6\x6b\x41\xdd\xf3\x1b\x0e\x15\xcf\xb5\x22\x25\x5a\xf6\x99\x40\xb5\x43\x87\xfc\x4e\x92\x86\xee\x89\x05\x4a\x43\xd7\x3f\x56\x2e\xb9\x6e\x99\x66\x67\x29\x0e\x4b\x2f\xbb\xe5\x32\x2d\x53\x89\x9c\x24\x5f\x72\xe7\xa7\xd0\xa2\xfb\x46\x1c\xbe\x9a\xd0\x16\xcb\x25\xba\x7c\x39\x6d\x59\x72\xe7\x4b\xc7\x69\x99\xca\xe4\xd0\xfc\x02\xf8\x5b\x1e\xf1\xb6\xe9\x99\x3e\xdf\xe2\xb1\xb9\x06\xfe\xca\x74\xcd\x67\x3c\x36\x7d\x1e\xf1\x8e\xe9\x9b\x81\xe9\xf2\x10\x7c\x33\x49\x7b\x66\xc0\x7b\xbc\x17\x17\x6e\xf0\x84\x37\x93\xb6\x91\xe9\xf2\x6d\xde\xe4\x11\xdf\xe1\x31\x8f\xc8\x0c\x78\x93\x87\xbc\x9b\x16\xcc\x35\x32\x5d\x1e\xf1\x2e\x8f\x79\x8b\x87\xf1\xdf\xf5\x1f\xbd\x3c\xe1\x9f\xf9\x0e\x4f\xf8\xd6\x3d\xfd\xbc\xc9\x93\x87\x72\x09\xf1\x98\x78\x3f\xe9\x1f\xc5\x25\xde\x37\x03\xd3\xe3\x1d\xe2\x7d\xd3\xe5\x31\xef\xf2\x84\x6f\x4f\x6f\x88\x77\x35\x34\x7d\x1e\x92\xe9\xf1\x84\xb7\x92\x21\x33\xe4\x07\xde\xe1\xa1\xf9\x3c\xc6\xcd\x97\xe0\x1b\x66\x90\x06\xdf\x4f\xe5\xcc\x17\xe0\x1f\xb3\xb5\x8c\xa7\xd8\x90\xb7\x78\x2f\x56\x07\x5f\x3f\x70\x9c\xa0\xa9\xd1\xfd\x44\x26\x7e\x8a\x49\x62\x67\xfc\x20\xef\x41\xa3\xf1\x03\xf4\xfe\x75\x2a\x9e\xf0\x1e\xf8\x9b\x18\x30\xbd\xc4\x77\xf6\xc8\x03\x73\x95\xc7\xbc\x8d\x7f\x02\x00\x00\xff\xff\xeb\x18\x60\xea\xd5\x05\x00\x00"

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
	"templates/views/releases.html": templatesViewsReleasesHtml,
	"templates/views/update.html": templatesViewsUpdateHtml,
	"locales/ru/LC_MESSAGES/ota.mo": localesRuLc_messagesOtaMo,
	"locales/ru/LC_MESSAGES/releases.mo": localesRuLc_messagesReleasesMo,
	"locales/ru/LC_MESSAGES/update.mo": localesRuLc_messagesUpdateMo,
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
				"ota.mo": &bintree{localesRuLc_messagesOtaMo, map[string]*bintree{}},
				"releases.mo": &bintree{localesRuLc_messagesReleasesMo, map[string]*bintree{}},
				"update.mo": &bintree{localesRuLc_messagesUpdateMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"releases.html": &bintree{templatesViewsReleasesHtml, map[string]*bintree{}},
			"update.html": &bintree{templatesViewsUpdateHtml, map[string]*bintree{}},
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
