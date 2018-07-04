// Code generated by go-bindata.
// sources:
// locales/ru/LC_MESSAGES/config.mo
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

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4f\x6b\x13\x5d\x14\xc6\x9f\xe6\xed\xbb\x99\x8d\x50\x50\x5c\xb8\x38\x5d\x58\x14\xb9\x75\x66\x62\xa1\xdc\x66\x5a\xb1\xb6\x20\x4d\xb0\x94\xd1\x95\x9b\x4b\x72\x3b\x19\x9a\xdc\x1b\xe7\xce\x14\x0b\x5d\x18\x54\xdc\xe9\x4a\x90\xee\xfc\x00\x42\x11\x02\x21\xb1\xfd\x0c\x67\xbe\x80\x9f\x45\xf2\x4f\xf1\xae\x9e\xdf\x39\xe7\xfe\x78\x7e\xad\x2c\x7f\x01\x80\x65\x00\x77\x00\xd4\x00\xfc\x0f\xe0\x15\x66\xaf\x07\x60\x05\xc0\x6b\x00\x37\x01\xbc\x9f\xf3\x05\x80\xf6\x12\xf0\x1d\xc0\x0d\x00\xb7\x2a\xc0\x6d\x00\x6b\x15\x60\x15\xc0\x4e\x05\x58\x02\xf0\xdf\xdc\x5d\x99\xfb\xd0\xb4\xe6\x38\x4d\x96\x77\xad\x3d\x49\x35\x19\xd5\xd5\x8b\xd1\x81\x3e\xa3\xd4\x90\xd3\xce\xa5\xd6\x2c\xa6\x75\xdb\x54\x1d\x4d\x4e\x9d\x6a\x1c\xe9\x9e\xcd\x72\xd1\x70\x49\xda\x12\x4f\x8a\xc4\x89\xd8\x4a\x6a\xe9\xd3\xc7\x27\x69\x5b\x75\xed\x7a\x56\x78\x87\xcf\x63\xb1\x9b\x69\x95\xa7\xd6\x88\xa7\x2a\xd7\x92\x42\x3f\xd8\x14\x7e\x55\x84\x55\x0a\xab\x72\x63\xe3\x81\x5f\xf5\x7d\xaf\xae\x5c\x2e\xe2\x4c\x19\xd7\x51\xb9\xcd\x24\x1d\x4c\x1d\xd4\x28\x32\xd5\xb5\x2d\x4b\xb5\x7f\xc4\xdb\x5e\x5d\x99\xa4\x50\x89\x16\xb1\x56\x5d\x49\x7f\x58\xd2\x51\xe1\x5c\xaa\x8c\xd7\x78\xd6\xd8\x13\x2f\x75\x36\xe9\x2f\x29\x58\xf7\xbd\x5d\x6b\x72\x6d\x72\x11\x9f\xf5\xb4\xa4\x5c\xbf\xc9\x1f\xf6\x3a\x2a\x35\x5b\xd4\x6c\xab\xcc\xe9\x3c\x7a\x11\xef\x8b\xcd\xbf\x77\x93\x3e\xc7\x3a\x13\x7b\xa6\x69\x5b\xa9\x49\x24\x79\x87\x9d\x22\x53\x1d\xb1\x6f\xb3\xae\x93\x64\x7a\x53\x74\x51\x75\x8b\x66\x31\x32\x77\x03\x3f\x8a\x02\x5a\x5b\xa3\x49\xf4\x57\xa3\x20\xa0\x1d\xf2\x49\x4e\x79\x3b\x0a\x17\xab\x5a\xf4\x68\x12\xef\x4d\xcf\x6a\x81\x4f\xe7\xe7\xb3\x2f\xdb\x51\xe8\xdf\xa7\x1d\x0a\x48\x52\xb8\x05\xfe\xca\x3f\xcb\xcf\xc4\xa3\xf2\x1d\x8f\x78\x08\xbe\xe0\x71\xf9\xa9\xfc\x48\xfc\x83\xca\x3e\x0f\xca\x7e\xd9\xe7\xe1\x64\xf1\x8d\xaf\xcb\x0f\xe5\x5b\xbe\xe4\x2b\x1e\xf0\x15\x0f\x79\x40\x3c\xe6\x6b\x1e\xf1\x25\x8f\x79\x88\xdf\x01\x00\x00\xff\xff\xea\x89\x4d\xd3\x61\x02\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 609, mode: os.FileMode(420), modTime: time.Unix(1522961961, 0)}
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