// Code generated by go-bindata.
// sources:
// templates/config.tpl.html
// templates/logs.tpl.html
// templates/mail.tpl.html
// templates/workers.tpl.html
// DO NOT EDIT!

package system

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

var _templatesConfigTplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x3f\x6f\xf3\x20\x10\xc6\x77\x3e\xc5\x09\xe5\xdd\xde\x3a\xed\xd0\xa5\x72\x98\xaa\x8e\xdd\xda\x9d\x98\xb3\x8d\x44\xc0\x82\xb3\xab\x28\xca\x77\xef\x61\x5b\x72\xdc\x3f\x69\x3c\x71\xc7\xef\x31\xf7\xc0\x73\x3a\x81\xf5\x95\xeb\x0d\x82\xac\x63\xf0\x84\xde\x48\x90\x4e\x1f\x43\x4f\xdb\x16\xb5\xc1\x58\x50\xe7\x8a\x96\x0e\x4e\xc2\xf9\x2c\x44\x69\xec\x00\x95\xd3\x29\xed\x24\xe9\xbd\xc3\xbb\x88\xa9\x0b\x3e\xd9\x01\xa5\x12\xc0\x5f\x39\xf6\x57\x10\x4c\x68\x1b\x06\x8c\xf3\x3a\x51\xb4\x1d\x9a\x59\x33\xe9\xf2\x89\x4b\x3d\xf5\xe2\xba\x31\x83\xf0\x61\x0d\xb5\x3b\xf9\x70\xff\x4f\xaa\x57\x7d\xc0\x72\x4b\xed\x55\xf2\x31\x93\xef\xda\xf5\xbf\xa3\xea\x19\x53\xc5\x53\x91\x0d\xfe\x3b\xc4\x9d\x8b\x59\xf2\xfe\x6a\xda\x92\xf6\xc1\x1c\x97\x9a\xef\x36\x6a\xdf\x20\x6c\xec\x7f\xd8\xd4\x4e\x37\xf0\xb4\x83\xe2\x85\x17\x29\xdf\xe4\x0d\x2e\x8d\xe2\x9f\x8c\xd2\x22\x5b\x64\x15\x9f\x6a\xfe\x20\x47\x8b\xb7\xa1\x6f\x49\x37\x3f\xa3\x6b\xaf\x2c\xe0\x60\x5c\x0e\xcd\xfb\x8b\x5b\x2e\xf2\x8b\x2a\x51\x6e\x39\x1c\x4a\x88\xeb\xb1\xaa\x43\xa0\x2f\xb1\xfa\x0c\x00\x00\xff\xff\x5f\x33\xbe\x76\x88\x02\x00\x00")

func templatesConfigTplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfigTplHtml,
		"templates/config.tpl.html",
	)
}

func templatesConfigTplHtml() (*asset, error) {
	bytes, err := templatesConfigTplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/config.tpl.html", size: 648, mode: os.FileMode(420), modTime: time.Unix(1441722307, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLogsTplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x55\x4d\x6f\xe3\x36\x10\xbd\xeb\x57\x4c\xb5\x06\x24\x35\xb1\x94\x1e\x7a\x71\x64\xf7\xb0\x2d\x8a\x14\xdb\xee\x21\xed\x29\xdd\x03\x2d\x8d\x64\x66\x69\x52\x20\x29\x6f\x0c\xc3\xff\xbd\x43\xea\x23\xb2\x9c\xa6\x40\x0d\xc4\x31\x47\xf3\xde\xcc\x9b\x0f\xea\x74\x02\x2e\x0b\xd1\x96\x08\x61\xa5\x95\xb4\x28\xcb\x10\x42\xc1\x8e\xaa\xb5\xd9\x0e\x59\x89\x3a\xb5\x8d\x48\x77\x76\x2f\x42\x38\x9f\x83\x20\x2f\xf9\x01\x0a\xc1\x8c\x59\x87\x0d\x93\x28\xc0\x7f\x2f\x4b\xac\x58\x2b\x6c\xb8\x09\x80\x3e\x57\x5e\x4b\x47\xc6\x65\x1d\x6e\x1e\x77\xea\x1b\x50\xe0\xf4\x77\xf6\xf2\x60\x71\x6f\x88\x16\x04\xb3\x68\x2c\x68\x2c\x94\x2e\x8d\xa7\x98\xd3\x6c\xad\x5c\xd6\x5a\xb5\x0d\x34\xad\x10\x4b\xcd\xeb\xdd\x10\x6d\x74\xdf\xb6\xd6\x2a\x09\xbc\x5c\x87\x1a\x2b\x8d\x66\x17\x82\x3d\x36\x48\x68\xff\x24\x9c\x90\x81\x23\xec\xb3\xf6\xbf\x5f\xcc\x8c\xce\x53\xf2\x01\x52\x31\xa8\xd8\x72\xa0\xdd\xe4\x19\x9f\x05\xcf\xba\x18\xaf\xd6\x3c\xa3\xf4\xfb\x7a\x4c\x7e\x5e\x95\x66\xab\xca\xe3\x24\x74\xde\x8a\xe1\xb9\x64\x07\xa0\xbf\xa5\x65\x5b\x13\x82\x56\x82\x94\xd0\x6f\xc1\xcd\x5c\x3a\x15\x54\x33\x59\x23\x2c\x24\xdb\xe3\x2d\x2c\x50\x5a\xcd\xd1\xc0\x6a\x0d\xe9\x27\x55\xd7\xa8\x5d\xa1\x2f\x33\x16\x7c\x93\x33\xd8\x91\xa6\x75\xf8\x81\x18\x3c\x96\xbc\x26\xa1\x42\x28\x99\x65\x4b\x4b\x0c\x83\x65\x33\xf1\xcc\x33\x46\x95\x10\xfc\x2a\x19\x1a\xa4\x69\xb8\x3c\x6b\xc5\x26\x78\xb3\xaf\x44\xb9\x2c\xfc\xe8\xfd\xa7\x26\xe1\x75\xbc\x2b\xc9\x11\x8f\xc9\xfb\xf2\x86\xd3\x40\xce\x12\xfa\xf9\x98\xca\x7d\xa3\xed\xcd\x26\x6f\x34\x92\xb6\xfe\x7b\xde\xeb\xb1\x9d\xef\x69\x9e\xb7\xbf\xff\x17\xe4\xa6\xd0\xbc\xb1\x9b\x60\x11\x97\xaa\x68\xf7\x24\x3d\x49\x35\x2d\xc8\x31\xae\x5a\x59\x58\xae\x64\x9c\xc0\x89\xa0\xc3\x11\xda\x86\xfa\x80\xbd\x15\xe0\xc0\x34\x90\x1e\x58\xc3\x22\x8e\xd2\x41\xd9\xea\xc0\x0d\xdf\x0a\x5c\x55\x5c\x1b\x1b\x25\xf7\x5d\xc5\x79\x05\xf1\x77\xe4\x33\x80\x81\xf6\xcc\xb6\x5a\xde\xfb\xd3\x39\x18\x19\xa9\xbe\x0f\x25\x71\x92\x6f\xca\xac\xd5\x71\xc4\x4b\xc7\x32\x3c\xa7\x65\xef\x1e\x56\x5c\x96\x71\x44\xa5\x89\x12\x7f\x35\xc4\xd1\x18\x6c\x91\xb2\x67\xf6\x12\x0f\x91\xdc\xf6\xad\x20\xfa\xf5\x97\x3f\xa3\xdb\xde\xd4\x6a\x41\x96\xcc\x1c\x0d\x6d\x7f\x46\x31\xcd\x4f\xf4\xb5\x8e\xe0\xa6\x4b\x60\xf0\x33\x6d\x51\xa0\x31\xab\xb1\x08\x31\x6d\x5e\xa3\xa4\xc1\xe4\x34\x16\x99\xfb\x1b\x64\x0d\x4f\x5f\xee\x47\x5b\xa5\x34\xc4\x2e\x5f\x4e\xb7\x1b\x8c\x20\x38\x5d\x34\xcc\x3b\x10\x9a\xc0\x83\xcb\x13\x9f\xb0\x8c\xec\x64\x25\x1f\x89\xdf\xe0\x67\xd7\x02\x67\x4b\x2d\xdf\x63\x92\x5a\xf5\x49\x15\x4c\xe0\x23\xad\x9a\xac\xe3\xe4\x6a\x8c\xfc\xe7\x06\x22\x78\x72\xea\x3c\x52\xe0\x01\x05\x21\xff\x6a\x1a\xd4\x1f\x99\x71\x3d\x25\x8f\x2f\x10\xfd\x1b\xda\xc3\xf6\x54\x09\x56\xe3\x7d\x70\x99\x1f\x75\xf6\xf3\xf6\x19\x0b\x9b\x7e\xc5\xa3\xe9\x72\xab\x38\x8a\xd2\x24\x14\x49\xd6\x76\x37\x97\x7d\x21\xeb\x66\x4d\xb9\xb9\xd4\x7e\x7b\xfc\xfc\x47\x6a\xbc\x0c\x5e\x1d\x2f\x78\x2e\x2b\xf2\x3a\xde\xe7\xd7\x54\xb0\x7b\x41\x78\x98\x49\x9f\x15\x97\x71\xf8\xb7\x0c\x93\x11\xdb\xa1\xce\xfe\xec\x71\x6e\x6a\xdd\xe5\xc6\x68\x80\x0a\xc1\x8b\xaf\x71\x37\xe0\xc3\xa9\x9b\x26\xd7\xa2\xfe\xc2\xed\x26\xfd\x43\x7f\xea\x86\xb2\x3f\xf4\x90\xd9\xe6\x74\xc5\x19\x5c\x76\xcc\x7c\x74\xb7\x40\x1c\x31\x72\x3a\xd0\xdc\x4e\x97\xa1\xf3\xd1\xb8\x57\x07\x9c\xb9\xf5\x3b\x42\x12\x0d\x5e\x21\x58\x59\xbe\xed\xee\x64\x26\x4e\x81\x41\xfb\x40\x57\x9b\x3e\x30\xf1\xff\x13\x1c\x76\x7f\x42\x7e\x0b\x3f\xdc\xdd\xdd\xc1\xf7\xf0\x23\x59\x5d\x5d\xf3\x6c\xb8\x53\x82\xf7\x5f\xea\x95\x52\x76\xf6\x52\xff\x27\x00\x00\xff\xff\x90\x25\x9c\x18\x06\x08\x00\x00")

func templatesLogsTplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLogsTplHtml,
		"templates/logs.tpl.html",
	)
}

func templatesLogsTplHtml() (*asset, error) {
	bytes, err := templatesLogsTplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/logs.tpl.html", size: 2054, mode: os.FileMode(420), modTime: time.Unix(1441722307, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMailTplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x51\x6f\xdb\x36\x10\x7e\xf7\xaf\xb8\x71\x43\x64\x63\x95\xb4\x06\x7b\x8b\x24\x0c\xd8\xcb\x1e\x56\xb4\x40\xfb\x36\xec\x81\x96\x4e\x31\x0b\x8a\xd4\x48\x2a\xad\x1b\xf8\xbf\xef\x48\x51\xaa\x92\xa8\xf1\xd6\x2c\x43\x09\xd8\x22\x8f\xc7\xbb\xef\x3e\x1e\x8f\xbc\xbd\x05\xa1\x6a\x39\x34\x08\xac\x35\x5a\x39\x54\x0d\x03\x26\xf9\x51\x0f\x2e\x3f\x20\x6f\xd0\x64\xae\x97\xd9\xc1\x75\x92\xc1\xe9\xb4\xd9\x14\xb6\x36\xa2\x77\x20\xb9\xba\x1e\xf8\x35\x96\xec\x3d\xbf\xe1\xa3\x90\x55\x1b\xa0\xf6\xc3\xb6\x1d\x54\xed\x84\x56\xb0\xdd\xc1\x6d\x90\xf9\x76\xc3\x0d\xb4\xda\x74\x50\x92\x4a\xf2\xbd\x25\x67\x69\xc7\x85\x4c\x76\x57\x9b\x59\xc9\x2b\x64\x76\xd8\x77\xc2\x7d\xc1\x4c\x70\x91\xf5\xda\x92\x82\x57\xe6\xce\x99\x6d\xc2\x83\x66\xb2\x7b\x11\x2d\xa0\x11\x5c\x8a\x4f\xb8\xdd\xed\xb2\x46\x2b\x9c\x8d\x6d\x1b\xee\xf8\x7d\x7b\xbe\x89\x16\xc2\x5c\x76\xe0\xf6\xf5\x07\xf5\xc6\xe8\x1e\x8d\x3b\x6e\x13\x34\x46\x9b\x64\x07\x17\x17\xe0\xe7\xff\x88\x82\x3f\xe1\xbb\xb2\x04\x35\x48\xf9\x60\x26\x93\xa8\xae\xdd\x61\xcd\xcb\x1c\x63\x6f\xb0\x27\x06\xb6\x49\xd1\x88\x1b\xa8\x25\xb7\xb6\x64\x5c\x92\x4b\x08\xff\x69\x43\x0c\xa3\x99\x06\xc2\x76\xc2\x5a\xbe\x97\xc8\xaa\x04\x7e\x5c\xb5\x3b\xb5\xa4\xd8\x0f\xce\x11\x6f\xee\xd8\xd3\x06\x8d\x03\x36\xf9\xa8\xa5\xb6\xc8\x02\xe0\xc9\x6c\x74\xcc\xaa\x0b\x27\x3a\xb4\x57\x45\x3e\xae\x39\xeb\xe9\x2e\x1f\xe7\x50\xe5\x14\x69\x95\xec\x1e\x68\x9d\x00\xa5\xc5\x27\x91\x65\x87\xba\x46\x6b\xbf\x6d\xb6\x92\x57\x04\x91\xce\x0c\xf8\xd4\x87\x08\xf9\x3c\xbc\x2f\xd2\x76\x47\x72\xda\x65\x2d\x1d\xa6\xcf\x79\xbe\x96\x7d\xcf\x99\x79\xff\x0f\x8f\x33\x87\xbc\x69\xc0\x69\xf8\x6b\xc0\x01\xc1\x47\x8e\xcd\xe3\xe8\xd6\x68\x3c\x2d\x4b\x8f\x6f\x06\xdd\x60\x14\xd9\xa3\x84\xbc\xda\x2c\xd5\xa6\x6f\x91\x8f\xc5\xae\xa2\x5a\x18\xaa\x59\x0c\xd0\xf7\xd3\x83\x36\xe2\x13\xd5\x51\x4e\xd5\xd2\x68\x89\xa3\x98\x41\x87\xee\xa0\x9b\x92\xf9\xaa\xc5\x40\x50\x6f\xae\x7e\x0c\xc6\xd2\x55\xb2\xdc\x1e\xad\xc3\x2e\x0f\xd2\xb1\x94\x2e\x77\x28\x38\xb8\x36\x7a\xe8\xe3\x64\x50\x90\x7c\x8f\xd2\xef\x6b\xc9\x9c\xfe\xcc\xb6\x96\xa9\xed\xd2\x4b\xa8\x09\x0c\x01\x49\x83\x1a\xab\xde\xbd\x2e\xf2\xd0\x5d\x58\x58\xb8\x88\xcb\x5e\xfe\xb4\xf0\x10\x74\x84\xea\x07\x17\x37\xd6\xe1\x47\xc7\xee\x80\x8a\x4e\xc6\xc0\x3c\x0a\xc5\x3b\x1c\x7b\xbd\xe4\x35\x1e\xb4\xa4\x8b\xa4\x64\xf8\x91\x77\xbd\xc4\x97\xbf\x34\x9a\x62\x54\x59\xad\xbb\x17\x10\x85\x97\x0b\x21\xa3\xab\x42\x0e\x64\x61\x19\x68\xd8\xbe\xcd\xfd\xee\xbf\xe1\x87\x6e\x95\xf7\x58\xbb\xb3\x24\xbd\x1d\xf5\x9e\x9d\xa9\x19\xcf\x48\xd7\x3c\x7c\x96\xe0\x3d\xa2\xb3\x91\xff\x1a\x9e\x00\x23\xfc\xaf\x0b\xdf\xa2\xa4\x18\xd6\x63\x36\xfa\x03\xc9\x7e\x9e\xd3\xc3\x23\xaa\x1e\x1c\xd8\x42\xf7\xe1\xce\x8f\x2c\x50\x02\x09\xc5\xaa\x37\xfe\x03\x9e\xd0\x22\x1f\x15\xce\xae\x0c\x6f\x96\xea\xb7\x77\xaf\x7e\x5f\x5f\x42\x07\x39\x80\xfd\x4f\x69\xee\xc6\xea\x74\x96\xe9\x58\xc5\xbe\x8e\x64\x4f\x03\x37\xc8\xff\x11\xcd\x13\xa2\xaa\xc8\xa7\x75\x4f\x8f\x78\x0d\x25\xc4\x9e\x6e\x5b\x8b\x2e\xbd\x5c\x3f\x19\x31\xf1\xe9\x79\xb7\xc8\xfb\x30\x1a\xcf\xcc\x34\x8a\x9b\xf8\x36\xbc\x48\xa3\xab\xbd\x53\x40\xbf\xb4\x37\xa2\xe3\xe6\xf8\xe8\x01\x29\x72\x8f\x9c\x8a\xf4\xe3\x6f\xdc\x56\x6b\x77\xef\x8d\xfb\x77\x00\x00\x00\xff\xff\xc8\xda\xe0\x41\x15\x0b\x00\x00")

func templatesMailTplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMailTplHtml,
		"templates/mail.tpl.html",
	)
}

func templatesMailTplHtml() (*asset, error) {
	bytes, err := templatesMailTplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/mail.tpl.html", size: 2837, mode: os.FileMode(420), modTime: time.Unix(1463942024, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesWorkersTplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\xdf\x6f\xdb\x36\x10\x7e\xf7\x5f\x41\x78\x05\x24\x6f\x95\xec\xac\xed\x80\x25\x71\x86\x61\x19\x8a\x02\xc3\x5e\x1a\x2c\x0f\xc3\x50\x30\xd2\xc9\xe6\xc2\x50\x02\x49\xc5\x31\xd2\xfc\xef\x23\x29\xea\x37\x65\xd9\x49\x50\x24\xc0\xfc\x60\x49\xc7\xef\x8e\x47\xde\xf7\x91\x94\x7d\x7f\x8f\x08\x8b\x68\x1e\x03\x9a\x26\x3c\x65\x12\x58\x3c\x45\x53\x8a\xb7\x69\x2e\xe7\x6b\xc0\x31\xf0\x50\x66\x34\x5c\xcb\x1b\x3a\x45\x0f\x0f\x93\xc9\x69\x4c\x6e\x51\x44\xb1\x10\xcb\x29\x4f\x37\xd3\xb3\x09\x52\x9f\xa6\x35\x4a\x69\x70\x27\x82\x9f\x6c\xd3\x1e\xcd\x5d\x48\x86\x19\x50\x64\xbe\x83\x15\x07\x60\x1d\xac\x13\x1f\xe8\x74\x09\x5b\x39\xb0\x5d\x7c\x9d\xf7\x18\xd2\x26\xfb\x6e\x07\xdc\xb8\x90\xd2\x21\xc1\x28\xc1\x41\x2e\x80\x0b\x7d\xf3\xe1\x6e\x7a\x76\x3a\x27\x3b\x3a\x9b\xab\xde\x0e\xca\xe5\x67\x24\xe1\x4e\x06\x9c\xac\xd6\x72\x2c\xad\x86\xf7\x3a\x5f\xc1\x14\x91\x78\x39\xdd\xa4\xfc\x5a\xa5\x17\x44\x69\xce\x54\x84\xc5\x48\x0a\x65\xa0\xb3\xcb\xc2\x6f\x2c\xe3\xe1\xe6\x81\xa6\x21\x33\x46\x6b\x0e\xc9\x72\xfa\x9d\x4d\x78\x8f\xc2\x16\x44\x48\xd2\x54\x02\xdf\x55\x61\xa1\x80\x95\x4f\x4e\x69\x40\x21\x51\x53\xf1\x17\x81\x0d\x3a\x07\x89\x09\x55\xc3\xd4\xa0\x43\x62\xd8\x8a\x74\xb9\x80\xb9\xa2\x5b\x10\x11\x1e\x51\xa8\x30\x8a\x12\xa3\x1d\x34\x4b\x4f\x01\xf3\x84\x18\x32\x1d\x3e\xbd\xb8\xa3\xb4\x36\xae\xfb\xf8\x78\xa9\x72\x88\x5f\x9b\x50\x25\x16\xd7\x2f\x56\xa8\x26\xb9\x60\x83\x89\x3c\x5c\xab\xca\x09\x19\xff\x6f\x2f\x57\xd3\xed\xff\x62\x7d\xe1\x62\xdd\x02\xa5\x4e\x75\xbd\x68\xbd\x0a\x22\xe1\x06\x67\x2f\x55\xb1\xab\x94\xab\x83\x13\x61\x20\x0e\xd0\xea\xc7\xca\xe9\x9b\x69\x75\x48\x85\x48\xc8\x2d\x05\x35\x22\xd0\x53\x70\xfc\xfe\x28\x73\xb3\x78\x98\x95\xf6\xd6\x5e\x26\x43\xec\x8b\x21\xc1\x39\x95\xad\x03\x89\xe3\x28\xd9\xe1\x5f\xff\x04\xd2\x07\x5f\xa5\xf1\x76\xe0\xe4\x29\xf1\x95\x56\x35\x88\x2c\x65\x82\xdc\x42\x57\x29\xa6\xbd\x05\x46\x85\xcb\x3a\xbd\x75\x2e\x51\xa7\x52\xa7\xe6\xb2\xf3\x81\x32\xc9\xf5\xd9\xa7\xf3\xd3\xb9\xba\x0c\xb6\xff\xc6\x01\x4b\x88\x77\x83\x2e\xd4\x1a\x8b\xc6\x22\x19\x10\xc3\x37\xb0\x07\x4c\x48\x2c\x73\xb1\x07\x30\xda\x95\x9e\xb2\x3a\x46\xae\xb1\x03\xf3\xa4\xab\xe5\x74\xe8\x37\x28\xa3\x2e\xc6\xb3\xb1\xae\xb9\x4d\xed\xe2\x9c\x6b\x27\x7d\x85\xb4\x7b\x76\xc6\xbc\x0a\x22\x88\x88\x93\x4c\x16\x6d\xb7\x98\x1b\x9a\x8b\xf3\x94\xe5\xf2\x64\x62\x8c\x49\xce\x22\x49\x52\x86\x56\x20\xf5\xb0\x3e\x1b\x1d\xfc\xa9\xa6\xc0\x2f\x24\x31\x43\xf7\xb6\x23\xb1\x21\x32\x5a\xf7\xcc\x08\x45\x58\x00\x5a\x1c\x37\x92\xe4\x20\x73\xce\x90\xa7\xcf\x6e\xde\x49\xa3\xe1\x4a\x4d\xda\xb5\xed\xba\x72\x3d\x72\xb9\x66\x3c\x8d\x40\x88\x71\xef\x1f\x5d\xde\x22\x8f\xf6\xf3\x7e\xe7\xf2\x4e\xd4\x99\x6a\xdc\xf5\xfd\xd0\x88\xd5\x43\xa6\xb8\xb1\x33\x82\x95\xa2\x2b\x44\xce\xae\x59\xba\x61\x2e\xf7\xe2\xe1\x61\x52\x7c\xb7\x0b\x98\x67\xb1\xe2\xa3\x5f\x17\xe6\x4d\x88\xff\xc5\x77\x7e\x5d\x27\xb9\xcd\xe0\x18\x79\x1f\x7f\xbf\xf0\xde\x56\xc6\x9c\x53\x65\x9b\x8b\xad\x50\x27\x8b\xb9\xdd\x8a\x1a\xed\x76\x26\x8f\xab\x8e\x7c\x3e\xbb\x6f\x64\xa6\x59\x25\x2f\xd1\x12\xbd\xf1\xbd\xf2\x4d\x15\x19\xd6\x7a\xb3\x10\x6e\x32\xb9\xf5\x67\x27\x5d\xfc\x85\xc5\x17\xaf\x1f\x3d\x74\x07\x5e\x46\x5d\x22\x1e\xda\xfb\x2f\xe6\x85\x00\x7d\xfd\x8a\x16\xbd\xe0\x3a\xe6\x17\x53\x87\x65\xbf\xb5\xf4\xb7\xed\x75\x40\x63\x70\xc5\x2b\x01\x57\xb9\xd8\xb6\x3c\x8c\xc1\xe5\x51\x9f\x82\x0c\xbe\xf1\xa8\xd1\xcd\xd1\x91\x04\xf9\xbf\x72\x8e\xb7\x21\x11\xe6\xea\x57\xe1\x67\x4d\x85\x15\x9f\x24\xe5\xc8\xd7\x1d\x10\x44\x58\x9d\x48\x1f\xd8\x4e\xbd\x99\xf4\xdf\xe4\x9f\x93\x01\xb0\x9e\x37\x05\x2d\x80\xa1\x7e\x6a\x15\xa2\xfc\xc8\xcb\x10\x67\x19\xb0\xd8\xf7\xf4\x7a\xeb\xed\x38\xdf\xfd\x80\x14\x24\x3e\xf3\xd4\x8d\x8d\x4a\x62\x63\x9c\x6b\xeb\x7e\x8e\x4c\xbf\xe4\x68\x5e\xdb\x08\x76\xd5\x9d\x85\x32\xfd\x23\x8d\x30\x85\xcf\x92\xab\x7d\x4a\xd1\xfe\xc0\xc0\xbe\x19\xef\x2f\x66\xd8\x3a\x2f\x25\x02\xef\x69\x41\xf4\xbe\xf1\xb4\x30\xfd\x35\xd8\x04\x2e\x57\xdc\x27\xc5\xae\x26\xd2\x84\x1c\x9e\xc6\xc3\x3a\xd1\xdb\x9c\x37\xeb\x52\xea\xa1\x4b\x1c\x87\x82\x43\x0a\x6c\x25\xd7\x93\x01\x2f\x97\x32\x6a\x65\xef\x21\x8e\x06\x78\x50\x1f\x96\xf2\x4d\xb0\x11\x88\x8b\xf6\x17\x8f\xa1\x7d\x49\xae\x03\x6b\x56\xd3\xe9\xb1\x6a\x19\x29\xf2\x33\x97\xb7\xb5\xe4\x36\x67\x73\xa4\xc8\xd5\x16\x10\x48\x22\x29\xa8\x2d\x40\xbf\x7f\xfa\x8d\x70\x2a\x09\xb4\xa9\x4e\xa0\xed\x54\xb4\x77\xbd\xb6\x96\xce\xb5\xa5\x07\x6e\xfd\xf0\x5b\xe2\xcb\x45\xb4\x0b\xee\xfe\xf8\xd4\x4f\xae\xe1\xf2\x50\xee\xcc\xd6\x66\x47\xf9\xc6\xaf\x36\xcd\x9a\x84\xe5\x3e\x5d\x7a\x0b\x90\x9f\x98\x7a\xf9\xbc\xc5\xd4\x2f\xda\xde\xa2\xa3\xc5\x62\x81\xbe\x47\x1f\xca\x70\xea\x7a\x3a\x2f\xcf\x72\x93\xdd\xff\x5a\x14\xaf\xb2\xad\x7f\x2d\xfe\x0b\x00\x00\xff\xff\x9b\x5b\xf8\xba\xe7\x18\x00\x00")

func templatesWorkersTplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesWorkersTplHtml,
		"templates/workers.tpl.html",
	)
}

func templatesWorkersTplHtml() (*asset, error) {
	bytes, err := templatesWorkersTplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/workers.tpl.html", size: 6375, mode: os.FileMode(420), modTime: time.Unix(1466119265, 0)}
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
	"templates/config.tpl.html": templatesConfigTplHtml,
	"templates/logs.tpl.html": templatesLogsTplHtml,
	"templates/mail.tpl.html": templatesMailTplHtml,
	"templates/workers.tpl.html": templatesWorkersTplHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"config.tpl.html": &bintree{templatesConfigTplHtml, map[string]*bintree{}},
		"logs.tpl.html": &bintree{templatesLogsTplHtml, map[string]*bintree{}},
		"mail.tpl.html": &bintree{templatesMailTplHtml, map[string]*bintree{}},
		"workers.tpl.html": &bintree{templatesWorkersTplHtml, map[string]*bintree{}},
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
