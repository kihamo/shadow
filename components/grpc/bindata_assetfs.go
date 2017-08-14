// Code generated by go-bindata.
// sources:
// templates/views/index.html
// assets/js/index.js
// assets/js/index.min.js
// DO NOT EDIT!

package grpc

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

var _templatesViewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\x5b\x8f\xe3\x34\x14\x7e\x9f\x5f\x61\x05\x84\x40\x22\x0d\x2c\xfb\x34\xb4\x5d\xad\x56\x42\xda\x07\x2e\x02\xde\x47\x6e\x7c\xd2\x78\xe4\x5c\x64\xbb\xed\x8c\x56\xfd\xef\x1c\xdb\xb9\xb6\xe9\x25\x71\x16\x01\xa2\xd2\x74\x12\xd7\xe7\x92\xf3\x1d\xfb\x7c\x27\xfe\xf4\x89\x30\x48\x78\x0e\x24\x88\x8b\x5c\x43\xae\x03\x72\x3c\x3e\x2c\x19\xdf\x93\x58\x50\xa5\x56\x41\x49\xb7\x10\x6a\xae\x05\x04\xeb\x07\x82\x9f\xee\x8f\x76\xfc\x49\x40\xa2\xab\x1f\xed\x84\xf4\x87\xf5\xf6\xf7\xdf\x3e\x2c\x23\xbc\x70\x22\x11\xca\xac\x1f\xaa\x7f\x3d\xf5\xb1\x00\x2a\x13\xfe\x12\xac\xeb\x5f\xd1\x27\x9e\x90\x05\x48\x59\xc8\x53\x67\xa8\x00\xa9\x89\xfd\x0e\x19\xcd\xb7\x20\xeb\x1b\xae\x32\xae\x14\xdd\xb4\x6e\x6e\x76\x5a\x17\x39\xd1\xaf\x25\xac\x02\x77\x13\xb4\x56\x0b\x05\x01\x61\x54\xd3\x5a\xb4\x52\x1e\x10\x2a\x39\x0d\x53\xce\x18\xe4\xf8\x80\x72\x87\x0a\xbf\xd2\x3c\x03\xf5\xe3\x32\x72\x6a\x9c\x01\xf4\xb3\xe3\xa4\x73\x1e\xc7\x20\x67\x66\xa0\xe7\xb6\x2c\x0e\x03\xc1\x8b\x0b\x11\x66\x2c\xfc\xfe\x4d\xd0\x28\x94\xe6\x99\xc8\x97\xea\x5b\xfc\x03\xb9\xe7\x31\x90\xc7\x15\x59\x54\xd7\xca\x28\xae\xa3\x8c\xb3\xeb\x39\x1f\x99\x99\x25\xa1\x14\x34\x86\x66\x74\xf1\x0b\xcd\x10\xd7\x45\x40\x82\x27\x8b\x6a\x0b\x50\xc7\x89\x97\xa7\x92\xe6\x20\x3a\xf0\x9d\xcf\xe8\xa2\xdf\x9b\x95\xbe\x59\x77\xdc\x70\x06\x8f\x47\xc4\xfd\xcd\xc0\xe4\x6b\xa0\xf7\x26\x0e\x0c\xf5\xdc\xa9\x33\x75\xfd\x70\xd5\x08\x8d\xe3\x42\x32\x6e\x50\xe7\x0c\x6f\xc3\x7e\xc4\x8e\xc7\x80\xc8\x42\x60\x6e\x68\xcc\x1a\xae\x6a\xe4\xb3\x9d\xd0\x5c\x81\x80\xd8\x8c\x43\x9d\x02\x67\xa6\xfa\x88\x65\x88\x58\x06\x3a\x2d\x2c\x14\x4d\x44\x7e\xb6\x43\x3d\xdc\xae\xb9\x3c\x04\xc5\x89\xbd\xca\xca\x09\xe4\x6e\xf0\x32\xe2\x03\x8a\x70\x8d\xd5\x62\x1f\xf3\x72\xa7\xff\xc4\x75\x72\xcd\x4f\xeb\x2b\xed\x79\x1a\xa6\x40\x19\xcf\xb7\x9d\x40\xba\x58\xa7\x67\xb1\x0e\x7b\xae\x9b\xd8\xdb\xb5\xa7\x8b\xed\xd6\x48\xe2\x52\x10\xb4\x6c\x96\x64\x49\x25\x02\xbc\x0a\xbe\x18\x02\x2d\x95\x90\xe0\x4f\xf1\x1d\x36\x2c\xa0\xf0\x82\xce\x32\x40\xb7\x12\x2a\x8c\x09\x3b\x6a\x92\x08\xbd\x36\xc9\x78\x5b\xd1\x65\x44\x9a\xc8\xa4\x6f\xfb\xa1\xa9\xd6\x4c\xab\xa9\xb3\x3a\xde\x5e\xd7\xb7\x8c\xe8\x8d\x09\x26\x65\x4c\x9c\xef\x89\x41\xcf\xab\x3a\xce\xa4\x0d\x78\x03\x9d\x4b\x3e\x17\x1d\x41\x37\x20\x04\xb0\xcd\xeb\x5d\x60\xde\x11\x9f\xd3\x2c\x0f\x37\x05\x7b\xbd\x43\xf0\x54\x78\x78\xbf\xba\x5f\xfa\xd2\x5e\x76\x55\xc3\x2d\xc4\x06\x85\x78\x6d\x34\xa1\x24\xa1\x21\xc5\x52\x71\x08\x25\xdf\xa6\xda\xdc\x26\x07\xb3\xf9\xf1\xf1\x7a\x3b\x19\xd5\xac\xdb\x3a\xb7\xa6\x28\x3b\xdd\x07\xfe\xd0\x12\x68\x36\x45\x99\xf9\x2c\x15\xc2\x53\x3f\xb8\xcd\x22\x62\xbf\xc3\x03\x95\xb9\xd9\x2a\xd6\xca\xea\x5f\x46\x66\xe2\xa4\xa7\xaf\x0a\xec\x28\xaf\xec\x9a\x1b\x27\x72\x6f\xb9\xba\x6e\x77\xdc\xf4\x0b\x65\x6e\x94\xdf\x49\x21\x33\xe2\x00\xc5\xa5\x56\xd8\xba\x16\x6b\x2c\x83\xab\xe0\x5d\x75\x11\x53\x21\x5a\x2a\x64\x6f\xec\xb6\x2b\x41\x61\xf1\x43\xba\x32\xc7\x8a\x3f\x73\x8c\x9b\xec\xaa\xe8\x98\xa3\x57\x01\xc9\x31\x6f\x57\x41\x65\x29\x20\x7b\x2a\x76\x78\x3f\xc0\x29\xe6\x35\xe8\x9e\xa4\x67\xaf\xbf\x49\x0f\x71\x8b\x5b\x9f\x0b\x35\x75\xf1\x13\x07\x71\x9d\x02\x5c\x7c\x82\x2e\xd3\x36\x64\xc4\x20\x54\x16\xb9\xe2\xfb\xb1\x5b\x58\xa3\xd2\xea\xe9\x29\x25\x4e\x75\x5a\xec\x91\x4d\xbb\x6b\x5c\xa3\xbc\x04\x36\xd1\x86\xb3\x63\xb8\x81\x8f\xbc\x9c\x2e\x5c\x39\xb0\xb6\x91\x5f\x46\x78\xe5\xad\xca\x20\x39\x8f\xa6\xf7\x76\x0d\xfa\xe9\x42\x69\x8f\xf0\x18\xdb\x9e\xe0\x98\xe2\x3d\x5d\xbe\xa5\xcd\x09\xd2\xe6\xc4\xa0\x64\x59\xf3\x8c\x4b\xa7\x75\xd5\x3b\x8f\x98\x25\x71\xd6\xcb\x96\xc3\x69\x8f\xf0\x9d\xab\xad\xb8\xf7\x3c\x6a\xfd\xb4\xf8\xa5\xd6\xb4\xf2\x7c\x62\x7f\x7a\x72\xa1\xb0\xd9\xbf\x26\xd4\x8a\x71\x85\xba\xfe\x74\xda\xfd\xd1\x16\x6b\x0e\x7f\x47\xb1\xed\x96\xea\xaa\x48\x07\x44\xe9\x57\x43\xdc\x19\x57\xd8\x03\xbe\x3e\xe6\x45\x3e\xb9\x26\xa4\x92\x44\x13\x45\x6f\xbf\x64\x19\x7e\x79\x32\xda\x50\xf7\x6d\x8a\x2b\x82\xd0\xc4\xe0\x90\x72\x8d\x45\xab\xc4\x56\xf8\xb1\x94\x10\x1e\x24\x2d\x47\x73\xb5\xc6\x52\xf5\x1a\xca\xc7\x43\x91\x3f\xa9\x42\x70\xd6\xbe\xd4\xf2\xd1\x66\x18\x5d\xb8\x95\xc5\xae\x9c\x0a\x70\x97\x47\xc0\x8b\x0e\x63\xe4\x95\x20\x7d\xca\x7b\x97\x5b\xa9\xdd\x26\xe3\xba\x81\x7d\xa3\x73\x82\x7f\xa1\xda\xc5\x31\x28\xd5\xb0\xac\x0f\x96\x68\x4e\x4d\x34\x2f\x30\xc7\x89\x98\x78\x7f\x16\x8e\x3f\x26\x19\xfe\x43\x3d\xaf\x79\x3d\x3c\x57\xcb\xfb\xeb\x4e\xcf\xdb\xf3\xb6\x0a\xff\x6f\x79\xff\xad\x2d\xef\xbc\x6d\xd2\xdf\xd1\x22\xf9\xb4\x47\x5e\x94\x76\xa6\xb6\xc8\xbf\x25\x9a\xce\x33\xbd\xda\x17\x9f\xd6\xe5\x46\xdb\xd2\xd9\x9a\xfc\xfa\x16\x4f\x80\x67\xef\x57\x66\xee\x55\xa6\x23\xef\xd7\x5d\x4c\xee\x2c\x26\x75\x15\x9f\x6f\xdb\xbc\x73\xea\x1d\xd3\x6e\x4c\x31\xe1\x16\x6a\x96\x43\xa1\x0b\x87\x3d\xee\x28\xe7\x99\xee\xa9\x8a\x71\x3b\xd5\x8f\xfb\x82\xb3\xaf\xbf\xfb\xe6\x9f\x74\xe0\x72\x3b\xe9\xae\x84\xf1\xb2\xf0\xed\x83\xcf\x2e\x5b\xec\xeb\xe9\x1f\xa7\x77\x9a\xd0\xce\x71\xfe\xb3\x72\x27\xf9\x2e\xb0\x44\xc9\x78\x15\x44\x5b\x59\xc6\x11\xc6\x0c\xb4\x8a\x9e\x55\xc4\x73\x06\x2f\x8b\x8c\xe7\x8b\x67\xf5\x6e\xbf\x32\x07\xda\xef\xcb\x52\xf0\x98\x9a\x57\x54\x8b\xcd\x8e\x0b\xf7\xb2\x17\xa9\x8b\x55\xd3\x31\xf6\x57\x00\x00\x00\xff\xff\x59\xa2\x1a\x1f\x3a\x20\x00\x00")

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

	info := bindataFileInfo{name: "templates/views/index.html", size: 8250, mode: os.FileMode(420), modTime: time.Unix(1502718322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsIndexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xc1\x8e\xd3\x30\x10\xbd\xef\x57\x58\xa2\x92\x1d\x51\x7c\xe2\xd4\x15\x37\x4e\x48\xc0\x81\x2f\x70\x9d\x31\x31\xeb\xd8\xd1\x78\xb2\x4b\x41\xfd\x77\x6c\xa7\x71\x92\x36\x07\x46\x8a\x46\x89\xdf\x9b\x79\x7e\x33\x39\x08\x33\x7a\x4d\x36\x78\xd1\xb0\xbf\x4f\x2c\xc5\x41\x70\xa9\x95\x73\x1f\x10\xe2\xe8\x88\x9d\x47\xa2\xe0\xa5\x76\x21\x02\x6f\x52\xb6\xfa\x65\xc5\x9a\x48\x13\x91\x3a\x1b\x1b\x39\x28\x04\x4f\xa2\x91\x9d\x6d\x41\x34\xcf\x05\x70\xbd\xe5\xb9\x85\x09\xd8\x97\x36\xa9\x64\x1c\xcf\xbd\xa5\x5a\x93\x55\x29\x39\xe0\x35\x15\x93\x03\x96\xfc\x19\x8c\x4a\x9a\xe6\xa2\x39\x5e\x15\x32\x60\x9f\xe6\xee\xdb\x83\x6e\x39\xa8\xb2\x9e\x9f\x2a\xa4\x93\x6f\xca\xd2\x57\x10\x4b\xbb\xd2\xd2\x18\xd0\xc4\x4e\x8c\x1b\xa5\xe1\x1c\xc2\x0b\xaf\xe7\xd7\x35\xff\x20\xd5\x2f\xf5\xfb\x8e\x4d\x97\x01\x4e\x0c\xa4\x22\x42\xc1\x7b\xa0\x2e\xb4\xbc\x39\x6e\x30\x23\xba\x05\xa2\xca\xad\xef\x21\xad\x22\x95\x31\x11\xd0\x2a\x67\xff\x24\x27\xb7\x80\x38\x6a\x0d\x31\x9e\x58\x9d\x05\xae\x7d\x5b\xdb\x30\x24\x1b\x38\x5f\x09\x9f\xc3\x1a\x26\x50\x4e\x83\xde\x23\xe7\xc8\xe4\x2f\x3f\xbe\x7f\x93\x91\xd0\xfa\x9f\xd6\x5c\x44\x79\x4d\x7e\x46\x58\xd8\x47\xe6\x47\xe7\x8e\xec\xe3\x6a\x04\xd5\x34\x06\x2e\xc2\xad\x1b\x20\x86\x5d\xa5\x73\xb3\x1b\xe4\xb1\xca\xa3\xfe\x7c\x39\x2c\x33\xe6\xef\x38\x7b\x9f\xec\xca\xb6\x09\x3e\x69\xe2\xcd\x8e\x96\x87\x0f\x28\x8d\xf5\x6d\x5a\xfa\x44\x1a\x82\x2f\x4b\xde\x51\xef\xc4\xb0\xc3\x46\x19\xbb\xf0\x26\xfe\xa7\x6e\x5d\x2e\x9e\xff\x03\x7e\x47\xb9\x6e\x36\x6a\xce\xe9\xf9\x17\x00\x00\xff\xff\xd6\x9e\xa0\x29\x92\x03\x00\x00")

func assetsJsIndexJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsIndexJs,
		"assets/js/index.js",
	)
}

func assetsJsIndexJs() (*asset, error) {
	bytes, err := assetsJsIndexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/index.js", size: 914, mode: os.FileMode(420), modTime: time.Unix(1502482867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsIndexMinJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x91\xc1\x6e\xe3\x20\x10\x86\x5f\x25\x62\xa3\x88\xd1\xb2\x9c\xf6\x64\xcb\xda\xcb\x9e\x56\xda\xf6\xd0\x27\x20\x78\xa8\x69\x30\xa0\x61\x9c\x34\xb5\xf2\xee\xc5\x4e\x1a\xb5\x95\x0f\x63\x7e\xcd\x3f\xf3\xf1\xb3\x95\x6e\x8a\x96\x7d\x8a\x12\xe6\xad\x14\xda\x9a\x10\x7e\x11\x96\x29\xf0\x66\x3f\x31\xa7\xa8\x6d\x48\x05\x05\xd4\xea\xed\xe1\x6b\x3f\x0f\xbe\x80\xce\x86\x30\xb2\x04\x3d\xf8\x1e\x25\x5c\x40\xd5\x49\x2e\xd1\xb8\x4e\xab\xce\x32\xed\x47\xcf\x9f\xad\x78\xac\x0e\x9d\x69\xad\x7f\xd1\x99\xba\x4f\x42\x7b\x34\xb4\xc1\xee\x36\x57\x0d\xdd\xf7\x0d\xed\xa0\x4f\xc6\xf3\x7f\x94\x33\x3a\x87\x96\x1b\xe1\x8c\xc5\x7d\x4a\x07\xb1\xac\xd5\xe6\xc5\xbc\xca\x99\xcf\x19\x1b\xd4\x86\x99\xa4\x18\x91\x87\xd4\x0b\x50\x13\x85\xbb\x68\x56\x90\x2a\xf6\x86\x4d\x55\x0b\x92\x37\xc1\xbf\x55\x7c\x55\x26\x6b\xb1\x94\xe6\x8e\x4b\x30\x2f\x60\xb9\x13\xa2\x25\x7d\x0d\xe7\x4f\xee\xfe\x3d\x3d\x3e\xe8\xc2\xe4\xe3\xb3\x77\x67\xb9\x1e\x2b\x68\x41\xf9\xd1\x04\x2a\x4e\x21\xa8\xdf\xd0\x90\x46\xa2\x44\xbb\x9d\xcc\xdd\xed\x1f\x94\xa4\x7a\x41\xf1\x43\xfc\x44\xbd\x60\x48\x71\x75\x09\x00\xd0\xce\xc7\xbe\x3e\x47\x55\x72\x8a\x6b\xfc\x03\x8f\x41\x66\x50\xa4\xcb\x90\x4e\x95\xf3\x9e\x85\x58\x72\x17\x70\xb9\xc0\xf2\xb5\xef\x01\x00\x00\xff\xff\x66\xf4\x41\x35\xd4\x01\x00\x00")

func assetsJsIndexMinJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsIndexMinJs,
		"assets/js/index.min.js",
	)
}

func assetsJsIndexMinJs() (*asset, error) {
	bytes, err := assetsJsIndexMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/index.min.js", size: 468, mode: os.FileMode(420), modTime: time.Unix(1502719057, 0)}
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
	"assets/js/index.js": assetsJsIndexJs,
	"assets/js/index.min.js": assetsJsIndexMinJs,
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
			"index.js": &bintree{assetsJsIndexJs, map[string]*bintree{}},
			"index.min.js": &bintree{assetsJsIndexMinJs, map[string]*bintree{}},
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
