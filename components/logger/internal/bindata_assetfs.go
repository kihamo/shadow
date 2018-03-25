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

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xcf\x6b\x24\x45\x18\x7d\x9d\x19\x5d\x1d\x7f\x2f\x22\x28\x22\xdf\x0a\x2e\x8a\x76\xec\xee\x49\x60\xe9\xa4\x13\x35\x9b\xe0\x92\x44\x97\x30\xea\x45\x58\xca\x99\x4a\xa7\xb1\xa7\x3a\x54\x77\x07\x17\x73\x88\x59\x71\x05\xc5\x15\x51\x14\x0f\x2e\xe8\x49\xbc\xcc\x8e\x19\x19\x37\x3b\xb3\x27\xf1\xfa\xb5\x67\xf1\x0f\xf0\x2f\xf0\x28\xdd\x3d\x9d\x44\xf0\xb2\x75\xa9\xef\x7d\xef\xd5\xab\xf7\x15\xf5\xd7\xe9\xfa\x57\x00\x70\x3f\x80\x27\x01\xec\x02\x78\x04\xc0\x9f\x28\x17\x19\x25\x77\xc6\x00\x1e\x02\x30\x3d\xc1\x0b\x06\xf0\x30\x80\xd5\x09\x7e\xdb\x00\x5e\x05\x10\x18\xc0\xa3\x00\xfa\x06\xf0\x34\x80\xdf\x0d\xe0\x01\x00\x7f\x1b\xc0\x83\x00\xfe\x31\x80\xd3\x00\x4e\x4d\x01\x5b\x06\xf0\xc4\x54\xd9\xbf\x52\x03\x1e\x03\xf0\x45\xad\xc4\x3f\xd6\x4a\xdf\xe1\x64\xff\xad\x06\xac\x02\xf8\xa3\x56\xe6\x7c\xaa\x0e\x2c\x02\x98\xab\x97\x7e\xef\xd7\xcb\xfe\x67\x75\xe0\x71\x00\x3f\xd5\x01\x03\xc0\x5d\x00\x1a\x93\xf9\x6a\x93\x99\xea\x93\xfd\x14\x80\xfb\x00\x4c\x01\xb8\x1b\xc7\xeb\x1e\x00\xf7\x56\xa0\x1d\xa9\xcd\xc0\xaf\xbf\x1c\x4a\x9d\x54\x60\x49\x07\x49\xd0\x16\x61\x85\xcf\xcb\x77\x52\xbf\x02\xcb\x5d\xa9\x7d\xa9\xda\x97\x8f\x1a\x5a\x47\xba\x02\x2b\x81\x0c\x3b\x31\x05\x8a\x36\x23\xdd\x15\x09\x6d\xe6\x8d\x4b\x4a\x74\xa5\x57\x94\xf6\xa5\x1d\x11\xa6\xf2\x85\x02\x38\x27\x08\xa7\x24\x2a\xa3\x0b\xaa\x34\x08\x22\x75\x1c\x64\x2d\xf2\x29\x94\x3b\x32\x3c\x71\xc1\xc6\xca\xd2\xec\x8c\x33\x53\x49\x5e\x8b\x92\xa0\x7d\xe4\xf2\x96\xd0\x2a\x50\x47\xd9\x45\xa7\x43\xa2\x8c\x84\x0d\xb9\x1d\xe9\xc4\x5c\x8f\xfd\xa0\x63\xbe\x92\xfa\xb1\xd9\x8a\x5c\xea\xc8\x9d\x97\xde\x0d\xb6\x44\x37\x9a\xd6\x69\xe3\xe2\xeb\x2d\x73\x49\xcb\x22\x85\x79\x5e\x24\xd2\x25\xc7\xb2\xcf\x99\x56\xd3\x74\x9a\xe4\x34\xdd\xd9\xd9\xe7\xad\xa6\x65\x35\xd6\x44\x9c\x98\x2d\x2d\x54\x1c\x8a\x24\xd2\x2e\xad\x16\x1e\xb4\x9e\x6a\xd1\x8d\x3a\x11\xcd\xff\xc7\x78\xa1\xb1\x26\x94\x9f\x0a\x5f\x9a\x2d\x29\xba\x2e\x1d\x61\x97\x36\xd2\x38\x0e\x84\x6a\xac\x5f\x58\x5f\x36\xdf\x94\x3a\x0e\x22\xe5\x92\x3d\x6d\x35\x96\x22\x95\x48\x95\x98\xad\xcb\xdb\xd2\xa5\x44\xbe\x97\xbc\xb8\x1d\x8a\x40\xcd\x51\x7b\x4b\xe8\x58\x26\xde\x1b\xad\x15\xf3\xdc\xb1\x2e\xcf\xb3\x29\xb5\xb9\xac\xda\x51\x27\x50\xbe\x4b\x8d\x8b\x61\xaa\x45\x68\xae\x44\xba\x1b\xbb\xa4\xb6\x0b\x18\x7b\xcd\x39\x2a\x4b\x4f\x3d\x63\x5b\x9e\x67\xd3\xd9\xb3\x94\x97\xd6\x19\xcf\xb6\x69\x91\x2c\x72\x0b\xbc\xe0\x39\x15\x35\xef\xcd\xe4\xe5\xb3\x85\x6c\xde\xb6\x68\x77\xb7\x3c\xb2\xe0\x39\xd6\x73\xb4\x48\x36\xb9\xe4\xcc\x81\xbf\xcf\xf6\x78\xc0\x7d\x1e\xf3\xcf\xdc\x03\x7f\x9b\xed\xf1\x30\xdb\xe7\x61\x76\x95\x07\xd9\x07\x7c\x93\x87\xfc\x2b\xf8\xbb\x6c\x9f\x0f\xb9\xc7\x07\x7c\x33\x57\x7d\xce\x7d\xee\x15\xca\x6b\x39\xf7\x31\x0f\xf9\x46\xc9\x5c\xe7\x31\x1f\x66\xd7\x88\xfb\x94\x7d\xc8\xe3\x6c\x8f\x6f\x71\x2f\xdb\xe7\xc1\x9d\xff\x35\xfe\x86\x47\x27\x3c\x3e\xe2\x21\x8f\x79\xc4\xa3\xec\x93\x3c\xd1\x0f\xd9\x1e\x8f\xb9\xcf\x03\x1e\x65\x9f\x12\x1f\x16\x03\x0c\x27\xcd\x1e\x8f\xf2\x6c\xff\x93\xa2\xfa\x90\xfc\x35\xf7\xf8\x16\x0f\xb2\xab\xa5\x98\x07\xe0\xeb\xc5\x53\x1c\x64\x57\xf8\x76\x51\xfd\xc2\x07\xb9\x7d\x49\x7e\xc9\x63\xbe\xc1\x3d\xee\xe7\xcf\x93\xdf\x78\x3b\x9f\x94\x07\xf8\x37\x00\x00\xff\xff\x49\x6d\xd6\x53\xc5\x04\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 1221, mode: os.FileMode(420), modTime: time.Unix(1521991982, 0)}
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
