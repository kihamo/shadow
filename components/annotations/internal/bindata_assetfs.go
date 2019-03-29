// Code generated by go-bindata.
// sources:
// locales/ru/LC_MESSAGES/config.mo
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

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x5d\x6b\x5c\x45\x18\x7e\x36\xdd\xf8\xb1\xe0\x57\x15\x05\xf1\xe2\xf5\xc2\xd2\x22\x27\x3d\xbb\xb1\x18\x4e\x72\x52\x6d\x93\x6a\xb0\xb1\x4b\xd9\x7a\x29\x4c\x72\x66\x77\x0f\xdd\x9d\x59\x66\xce\xaa\x91\x5c\xc4\x54\x14\x69\xd1\x20\x7a\x23\x88\x20\x78\xbf\x86\xac\xdd\x36\xdd\xf5\x2f\xbc\xf3\x07\xbc\xf3\x5e\xff\x81\xcc\xd9\x9c\x86\x56\x04\x11\x9c\x9b\x9d\x67\xde\x79\x9f\x8f\xd9\xf7\xfc\x76\xb2\xfc\x2d\x00\x54\x00\xbc\x04\xa0\x03\xe0\x09\x00\xf7\x30\x5d\x7f\x00\x78\x1d\xc0\x9f\x00\x9e\x03\x70\xa6\x04\x3c\x09\x60\xb9\x04\xc4\x00\xd6\x4b\xc0\xf3\x00\x6e\x96\x80\x0b\x00\x7e\x3c\xfa\xfd\xbd\x04\x3c\x05\xe0\xf4\xcc\xb4\xbe\x30\x03\xb4\x4b\xc0\x95\x19\xe0\x7d\x00\xa3\x13\xc0\x33\x00\x5e\x28\x03\x4f\x7b\xde\x32\xb0\x0c\x20\x2a\x03\x2f\x02\xd8\x2a\x03\x5d\x00\xdf\x94\x81\x2d\x00\x33\xb3\xc0\xb3\x00\xb6\x67\x81\x93\x00\xf6\x66\x81\x12\x80\xc7\x01\xcc\x02\x78\xec\xc8\xef\x23\x38\x5e\x33\x00\xca\x00\x1e\x3d\xca\x77\xc2\x1f\x6e\x6a\xd5\x4c\x5b\xe5\x37\xeb\x6b\x74\x5d\x6e\xcd\xd1\xbb\x9a\x94\x94\x09\xa5\x4d\xea\x5b\x69\x94\xe8\x4a\x12\x2a\xa1\x9e\xb0\xf6\x43\x6d\x12\x4a\x2d\x59\x99\x15\x7d\x2b\xc2\xb6\x37\xb4\x30\x89\xa5\xb5\x95\xe2\x70\x55\x89\x8d\x8e\x4c\x0a\xf8\x96\x11\x4d\xa1\x04\x89\x24\x31\xd2\x5a\xd2\x4d\x7a\xbb\xd1\xa8\x93\xd7\x4c\x15\x35\xb5\xe9\x8a\x8c\xda\x59\xd6\x8b\xce\x9e\x6d\x6b\x9b\x45\x3d\x6d\xb2\x87\xbb\x6d\xa6\x8d\x68\xc9\xe2\xb8\x5e\xf8\x69\x6a\x43\x1b\xc2\xa6\x9b\x24\xfa\x59\x5b\x9b\xf4\x63\x91\xa5\x5a\x3d\x90\xe4\x28\xdd\x43\xde\xaf\x15\xf9\xfe\x3b\x85\x48\x12\x12\xc7\xc9\x83\x4c\x6c\x94\x85\x52\x3a\xcb\x09\x2c\xae\x4a\x1f\x25\x58\xb7\xad\x34\x09\x2e\xf4\x5b\x36\x68\xe8\x88\x12\xf9\xc1\x1b\xd7\xd3\xb6\xe8\xea\x39\xd3\xaf\xd4\xaf\x34\x82\x8b\x46\xe6\x1d\xc1\x8a\xc8\x64\x44\xb5\xb0\xba\x10\x84\xf3\x41\x6d\x9e\x6a\xf3\xd1\xb9\x73\xaf\x86\xf3\x61\x58\xb9\x2c\x6c\x16\x34\x8c\x50\xb6\x23\x32\x6d\x22\x7a\x27\xe7\xa0\xf5\xbe\x11\x5d\x9d\x68\x5a\x7a\x80\x78\xb9\x72\x59\xa8\x56\x5f\xb4\x64\xd0\x90\xa2\x1b\xd1\x7d\x1c\xd1\xd5\xbe\xb5\xa9\x50\x95\xf5\xb5\xf5\xd5\xe0\x3d\x69\x6c\xaa\x55\x44\xd5\xb9\xb0\x72\x51\xab\x4c\xaa\x2c\x68\x6c\xf5\x64\x44\x99\xfc\x28\x3b\xdb\xeb\x88\x54\x2d\xd2\x66\x5b\x18\x2b\xb3\xf8\x5a\xe3\x52\xb0\x70\x7c\xcf\xfb\x69\x4a\x13\xac\xaa\x4d\x9d\xa4\xaa\x15\x51\xa5\xde\xe9\x1b\xd1\x09\x2e\x69\xd3\xb5\x11\xa9\x5e\x0e\x6d\x3c\xbf\x48\xd3\x6d\xac\x5e\xa9\x86\x71\x5c\xa5\x53\xa7\xc8\x6f\xc3\x97\xe3\x6a\x95\xce\x53\x48\x51\x8e\x97\xe3\x5a\x51\x5a\x8a\x5f\xf3\xdb\xd3\xf9\xb5\xa5\x6a\x48\xdb\xdb\xd3\x96\xe5\xb8\x16\x9e\xa1\xf3\x54\xa5\x88\x6a\x8b\xe0\xef\xf8\xd0\x7d\xe9\x3e\xf7\xff\xd3\x1c\xf1\xf7\x3c\x24\x1e\xbb\x1b\xfc\x0b\x0f\x79\x4c\x3c\x74\x9f\xf0\x21\x8f\xc8\xdd\xe0\xbb\x3c\xe0\xdb\x3c\xe0\xb1\xbb\x49\xfc\x2b\x4f\xf8\xd0\xdd\xe2\xdb\x3c\xe1\x7d\x1e\xb8\x5d\x1e\x7a\x4c\x3c\xf2\xb5\x81\xdb\x99\xd6\xb1\xb6\x92\x63\x1e\xfb\x3a\x0f\xf9\x0e\xf8\x6b\xbe\x3b\x95\xf4\x0a\x3c\x01\xef\xf1\x81\xdb\xf1\x4a\x54\x8c\xec\xfd\x29\xe7\x7d\x72\x9f\xf2\xc4\xed\xf0\xbd\xa9\xc8\xdf\xa7\x9d\x7f\x72\x3b\xb9\xc0\xc8\x1b\x75\x5f\xf0\xb0\xa0\x01\xff\x70\xec\x84\xf8\x80\x0f\xdd\x57\xc4\x3f\xe7\x31\xbc\xeb\x09\xdf\x21\x1e\xf0\xbe\xdb\xcd\x15\x46\x3e\x9d\xfb\x8c\x47\x3c\xfa\x97\x0f\x41\x45\x12\xef\xd5\xab\xfd\xc3\x9b\xfc\xef\xca\x07\x3c\xc9\xd9\xf7\x79\xe4\x76\xdd\x2d\xff\x5d\xf1\x1e\x8f\xfd\xf3\xba\xdd\x82\x1a\x7f\x05\x00\x00\xff\xff\xd8\x97\x38\xdc\xa3\x05\x00\x00"

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
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
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
				"config.mo": &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
			}},
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
