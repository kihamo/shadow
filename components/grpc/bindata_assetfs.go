// Code generated by go-bindata.
// sources:
// templates/views/index.html
// public/js/dashboard.js
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

var _templatesViewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x4b\x8f\xdb\x36\x10\xbe\xe7\x57\x10\x6c\xd1\x53\x65\x61\x8b\x9c\x76\x6d\x07\x69\x80\x16\x39\xf4\x81\xb6\xf7\x82\x96\x46\x12\x17\xb4\x28\x90\x94\xdd\x45\xe0\xff\xde\x21\x45\xbd\x6c\x29\x12\x6d\xa0\xa7\x0a\xd8\xb5\x28\xce\x0c\xe7\xc5\x99\x8f\xfc\xf2\x85\xa4\x90\xf1\x12\x08\x4d\x64\x69\xa0\x34\x94\x5c\x2e\xef\xb6\x29\x3f\x91\x44\x30\xad\x77\x54\xc9\x33\xdd\xbf\x23\xf8\x0c\xbf\x26\x52\x44\x22\x8f\x9e\x7e\xf0\x73\x6e\xbe\x78\x6a\xa7\x2b\x96\x43\x54\x00\x4b\x41\xd1\x7d\xfe\xc7\xef\x9f\xb6\x71\xf1\xe4\xa5\xc4\x28\x66\xff\xce\xff\x8c\x96\xaa\x58\x09\x22\x3a\xc8\xf4\xcd\x4b\x45\xf5\x78\x46\x36\xa0\x94\x54\x56\xaf\x6b\x2d\x98\x00\x65\x88\xfb\x1f\xa5\xac\xcc\x41\xb5\x03\xae\x8f\x5c\x6b\x76\x10\x30\x54\xf0\x50\x1b\x23\x4b\x62\xde\x2a\xd8\xd1\x66\x40\x3b\x8b\x84\xd4\x40\x49\xca\x0c\x6b\xd9\xfd\x02\x94\x30\xc5\x59\x54\xf0\x34\x85\x72\x47\x8d\xaa\x51\xe8\x77\x86\x1f\x41\xbf\x6c\xe3\x46\x4c\xbf\x08\xea\x7c\xa5\x70\x63\xa9\x9f\x83\x32\xb5\x13\xed\x50\x59\xad\xc9\xb7\xfa\x7b\xfc\x03\x75\xe2\x09\x90\xe7\x1d\xd9\xf8\x77\xdd\xca\x40\xca\x76\xfe\x73\x6a\x29\x14\x54\x82\x25\xd0\x7d\xdd\xfc\xca\x8e\x18\xc4\x0d\x25\xf4\x6f\xda\xad\x5c\xbc\xdf\x0f\x38\x1b\x9a\xcb\x05\x63\xf1\x7e\x7f\xad\xc1\x11\x35\x38\x82\x29\xa4\x13\xdf\xb1\xfc\xe2\x3e\x0d\xf5\xc0\x80\x78\xc2\xcd\xe7\xb2\xaa\xcd\x5f\xe8\xcc\xa1\x9a\xcd\xdc\x95\x96\x9e\x61\x52\xc9\x9b\x04\xc8\x95\xac\x2b\x4a\x78\xba\xa3\x0d\x9f\x8e\x46\x82\x2f\x97\x61\x4c\xaf\xd9\x49\x23\x04\xd3\x9a\xd5\xc2\x0c\x28\xa7\x17\xb3\x49\xca\xcb\xfc\x8a\xce\x7b\x6f\x4c\x6a\xb8\x19\xa5\xd3\x88\x98\x35\x99\x63\x64\x9e\x0b\x70\xfb\x43\xb0\xaa\x4b\xa8\x8a\x29\xdc\x5b\x3b\xfa\xcd\xd0\xa0\x3e\xa0\x97\xcb\x8d\x85\x93\xab\xe0\x53\x28\xc8\x3a\x39\xcb\x62\xf6\xfd\x87\x3e\xfc\x6c\xc2\x58\x97\x13\xe3\x2f\x5d\xda\x8e\xbc\xd7\x87\x65\x79\xf5\xb1\xff\x5a\x9f\x90\xce\x39\x13\x7a\x7c\xad\x1c\x2c\x12\xfb\xd8\xf3\x32\x93\x33\x2c\xd3\x6b\xcc\x27\xc1\x88\x8f\xb7\x5c\x19\x23\x19\x8b\x18\xee\xf2\x73\xa4\x78\x5e\x18\x3b\xcc\xb0\x4e\x6e\x63\xfe\x75\x19\x83\x70\x74\xbb\xa7\x0d\xcc\x12\x23\xee\x3c\xac\x2a\x23\xf6\x3f\x8d\x02\x76\x5c\xe2\x1d\x19\x5c\x0b\xd1\xe8\xbc\x60\xad\x63\xd4\xe8\x9f\x96\x53\xb0\x03\x7a\xd8\xfd\x8f\xce\x4c\x95\xce\x61\xda\x29\xb0\x8d\x2d\xe1\x82\xf7\x6e\xf3\x69\xc2\x46\x5f\x1e\xef\x14\xb2\xcd\xa4\x3a\x92\xc6\x41\x68\xaa\xd4\xb6\x76\x27\x86\x4b\x2c\xdb\x1f\xfc\x4b\xc2\x84\xe8\x8b\xbe\x1d\x2c\xb8\xc1\x6d\x60\x05\x1a\xab\x09\x76\x43\xf7\xbb\x62\xdf\x2d\x64\x92\x0d\x9e\xef\x43\x4d\x5f\xa1\xa4\xc4\x2c\xd8\x51\x2f\x95\x92\x13\x13\x35\x8e\x27\xea\xf7\xfd\xc2\x1b\x0d\x47\xb2\xc7\xb5\x61\x49\x74\xc8\xee\x1c\x3e\x33\x7d\x63\xf3\x13\x07\xd1\x77\x97\xb5\x4b\x1b\xdb\xd7\x6d\x48\x2a\x59\x6a\x7e\x9a\xab\xc9\x37\x22\x1c\xdf\x48\xc8\x4a\xce\x86\xdb\x56\x89\x10\x7a\xb5\x9e\xd8\x2f\xb0\x77\xfe\xd8\xc6\xf8\x16\xcc\x6a\xfd\x79\x1f\xe7\x47\xb7\x31\xc2\x78\x91\x3a\xc0\x3c\x2b\x3b\xd0\x79\x36\xb1\xd6\xd3\xf7\x38\x26\x43\x1c\x93\x59\x2f\x3a\x18\xf3\x40\xc2\xf5\xaa\x04\xc7\x31\x75\x5d\xd7\x69\xd1\x37\x5d\x13\x60\xfe\xad\x18\x8f\xb2\xee\x13\x13\xc6\x15\x16\xda\xe5\x9a\x3d\x21\x7f\x7d\x70\x91\xd8\xee\xd3\x15\xad\x6a\xb9\xbd\xd8\x67\x9d\xba\x2b\x84\xdd\x96\xc2\x4c\x4a\x83\x07\x10\x03\xff\x98\x28\x41\xb4\x67\x8f\x3e\xcb\x6a\x0f\x8b\xb5\xae\x0f\x47\x6e\xba\xe6\x74\x30\x25\xc1\x3f\x07\x67\xdc\x8b\x3e\x76\xa5\xfb\x93\x6d\x5c\xf1\x43\x1d\x77\x1b\xdb\x6e\x39\x83\xac\xfc\xe9\x6c\x95\xe5\xe4\x0c\xb6\xa5\x5a\x60\xb8\xb6\x3d\xce\xeb\xad\xcd\x9b\xc5\xcf\xe7\x82\x1b\x88\x10\x55\x24\xf0\x5c\x29\x88\xce\x8a\x55\x2f\x78\x2c\xc3\xf3\xc4\xdb\x73\x29\x4b\x78\xb1\x60\x2b\x44\xc9\x26\x48\xba\x4e\xf0\x58\xa5\xff\x33\x74\x28\x20\xbb\x0b\x1c\xfe\x56\x9b\x87\xd0\x61\xc3\xff\x3f\x3c\x9c\x37\x74\x2d\x7c\x79\x10\x78\xdc\x0f\x3a\x42\x00\x47\x50\x93\xba\x13\x68\x84\x83\x8c\xf5\x9d\x24\x08\x20\x84\x80\x83\x05\x60\x30\xd8\x66\x61\xc8\x20\xd0\xe1\x0f\x23\x82\x07\xd1\xc0\xfa\x48\x84\xf5\xf3\xd5\xbd\x7c\x55\x1f\x5f\x6c\x58\xb3\xd3\x33\x53\x53\x77\x18\xe3\x4f\x83\xe1\xd5\x3d\x9d\xd0\x30\xb8\x47\x9b\xbc\x09\x9a\xba\x5b\x69\xaf\x51\xc6\x7e\x9c\x1b\xf9\x35\x07\xf7\x82\x83\x1b\xd9\x57\xdd\x5c\xc6\xea\x44\xf1\xca\x10\xad\x92\x1d\x8d\x73\x55\x25\xf1\xab\x8e\x53\xa6\x8b\x83\x64\x2a\xdd\xbc\xea\x0f\xa7\x9d\xbd\x75\xfc\x58\x55\x82\x27\xcc\x82\xf9\xcd\xcf\x60\x7e\xac\xb9\x68\xce\xa2\x58\x7f\x9d\x88\xc1\x42\xff\x06\x00\x00\xff\xff\xc0\xef\x9b\x5b\xf9\x15\x00\x00")

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

	info := bindataFileInfo{name: "templates/views/index.html", size: 5625, mode: os.FileMode(420), modTime: time.Unix(1501162841, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicJsDashboardJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\xc1\x4e\xc3\x30\x0c\x86\xef\x3c\x85\x25\x26\x25\x15\x53\x4e\x9c\x86\xb8\x71\x42\x02\x0e\x3c\x81\xd7\x39\x34\x90\x36\x95\xe3\x6e\x0c\xb4\x77\x27\x2d\x5b\xda\x75\xb3\x54\x45\x95\xbf\xdf\x76\x7e\x67\xa1\x6d\xd7\x94\xe2\x42\xa3\x0b\xf8\xbd\x81\x14\x0b\xad\x6c\xe0\xda\x94\xe8\xbd\x2a\x4c\xec\xd6\xb5\x93\x8c\x41\xe6\xfa\xa0\x2d\x35\x62\x5a\x1e\xce\x27\xb2\xd8\x79\xd1\xc5\x43\xce\x6f\x91\x81\xe0\x31\xd5\x94\xca\xc5\x59\xa2\x1a\x13\xa6\x45\x4e\x15\x7a\x69\x46\x2a\xb3\x43\x27\x2f\xa4\xc7\x76\x43\x4b\x6b\xa9\x14\x58\x81\xb2\x58\xd2\x3a\x84\x2f\x95\xf3\x87\xa9\x7e\x61\xf0\x13\xbf\x67\x6a\xd9\xb7\xb4\x02\x32\x28\xc2\x5a\xd5\x24\x55\xd8\xa8\x62\x79\xc6\x74\xec\x47\x04\x87\x5b\xcf\x91\x0d\x0a\xf6\x4c\x24\x76\xe8\xdd\x0f\xe9\x19\x10\xbb\xb2\xa4\x18\x57\x90\xed\xe5\xa9\x6f\x53\x1b\xda\x64\x83\x52\x93\xc1\x4f\xe1\x2c\x68\x36\x4c\x31\x99\x7a\x4d\xdc\x47\x2f\x7e\x7e\x7f\x7b\x35\x51\xd8\x35\x1f\xce\xee\xf5\xf0\x9b\xfc\x8c\x34\xaa\x97\xd0\x74\xde\x2f\xe1\x7e\xb2\x82\x6c\x1a\x90\x8f\x74\xec\x46\xcc\xe1\xea\xa4\xa7\x66\x47\xe4\xb2\xca\xe5\xfc\xe9\x21\xdd\x2a\xb8\x4b\x36\xf5\x76\x69\xf5\x3f\x8b\x2a\x0a\x53\x49\xed\x75\x9b\xde\x56\x15\x76\xfa\xca\x48\x79\xf5\xaa\x72\x1b\x52\x33\xe2\x70\xb6\xef\xd3\x99\xbe\xbf\x00\x00\x00\xff\xff\x26\xf7\xa1\x29\xcd\x02\x00\x00")

func publicJsDashboardJsBytes() ([]byte, error) {
	return bindataRead(
		_publicJsDashboardJs,
		"public/js/dashboard.js",
	)
}

func publicJsDashboardJs() (*asset, error) {
	bytes, err := publicJsDashboardJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/js/dashboard.js", size: 717, mode: os.FileMode(420), modTime: time.Unix(1501157366, 0)}
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
	"public/js/dashboard.js": publicJsDashboardJs,
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
	"public": &bintree{nil, map[string]*bintree{
		"js": &bintree{nil, map[string]*bintree{
			"dashboard.js": &bintree{publicJsDashboardJs, map[string]*bintree{}},
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
