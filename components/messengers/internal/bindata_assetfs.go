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

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x52\x5d\x8b\x1c\x55\x10\x3d\x6d\x8f\x1a\xc7\xaf\xa8\xa0\x3e\x28\x94\x4a\x42\x44\x7a\xd3\x3d\x63\x20\xf4\x6e\xef\x4a\xf6\x03\x16\x77\x4d\x58\x26\xbe\xdf\x6c\xdf\xed\x69\x76\xe6\xde\xa1\xef\x1d\x3f\x20\x4a\xd8\x10\x83\x18\x12\x14\x7d\xf3\xc9\x5f\x30\xae\xbb\x38\xd9\x25\x83\x3f\x40\xa1\xae\xef\xea\xab\x3f\x43\xba\xe7\xc3\x64\x53\x2f\x5d\xa7\xce\xa9\x53\xd5\xdc\xfa\xfb\xe5\xda\x0f\x00\xf0\x2c\x80\x37\x00\x58\x00\x2f\x01\xf8\x1d\xe3\x38\xed\x01\x6f\x96\x35\x0f\x78\x11\xc0\x79\x0f\x78\x01\xc0\x92\x07\xbc\x0a\xe0\x23\x0f\x78\x1d\x80\xf4\x80\xe7\x01\x7c\x39\xd1\xdd\xf1\x80\x77\x00\xfc\xe8\x01\xcf\x01\xf8\x6d\x52\xff\xc3\x03\xda\x1e\xf0\x97\x07\x9c\x03\xb0\xe6\x8f\xfb\xb4\x0f\x9c\x06\xf0\x85\x0f\x9c\x01\x70\xd7\x1f\xcf\x3d\xf0\xc7\xfd\x7f\xfa\xc0\x29\x00\xff\xf8\xc0\x1c\x80\x7f\x7d\xa0\x0e\xe0\xb5\x1a\xf0\x0a\x80\xb7\x6b\x80\x37\xd9\xf9\x29\x00\xcf\x60\xcc\x97\x71\x6a\xf2\x7f\xd3\x78\x62\xf2\x7d\x12\xc0\xd3\x00\x7c\x00\xb5\x29\xb9\xad\xd5\x4e\x9e\xd5\x2e\x09\x23\xe9\xea\xd6\x06\xed\xe8\x82\x3e\x95\xd7\xa8\xad\xf5\xae\x99\xb2\xcb\x6d\x61\x0d\xad\xaf\x4c\xf1\xaa\x12\xd7\x3a\x32\x3d\x01\xa9\xdf\x4b\x85\x95\xe6\x64\xf9\x31\xbb\xcb\xb6\x2d\x8b\x19\x6a\xc9\x8e\xcc\x0a\xd1\x3d\x89\x49\x28\xa5\xad\xb0\xb9\x56\x86\x8c\xd5\x85\xc8\xe4\x4c\xa3\x77\xa5\x9a\x02\x91\xa6\x24\xca\xed\xb6\x64\x4f\x17\x36\xd8\x34\x59\x9e\x06\x97\xfa\x99\x09\x5a\x3a\xa6\x54\x7e\xf2\xc1\x6e\xde\x16\x5d\x3d\x57\xf4\xeb\x57\x2e\xb7\x82\xe5\x42\x56\xb6\xc1\x8a\xb0\x32\xa6\x46\x18\x5d\x0c\xc2\x66\xd0\x68\x52\xa3\x19\x5f\xb8\xf0\x5e\xd8\x0c\xc3\xfa\x86\x30\x36\x68\x15\x42\x99\x8e\xb0\xba\x88\xe9\xc3\xca\x83\x36\xfb\x85\xe8\xea\x54\xd3\xc2\x23\xc6\x8b\xf5\x0d\xa1\xb2\xbe\xc8\x64\xd0\x92\xa2\x1b\xd3\x0c\xc7\xb4\xd5\x37\x26\x17\xaa\xbe\xb9\xbe\xb9\x1a\x7c\x2c\x0b\x93\x6b\x15\x53\x34\x17\xd6\x97\xb5\xb2\x52\xd9\xa0\xf5\x79\x4f\xc6\x64\xe5\x67\xf6\x7c\xaf\x23\x72\x35\x4f\xdb\x6d\x51\x18\x69\x93\xab\xad\xb5\xe0\xe2\xff\xba\x72\x9f\x1d\x59\x04\xab\x6a\x5b\xa7\xb9\xca\x62\xaa\x5f\xe9\xf4\x0b\xd1\x09\xd6\x74\xd1\x35\x31\xa9\x5e\x05\x4d\xd2\x9c\xa7\x71\x9a\xa8\x33\x51\x98\x24\x11\x9d\x3d\x4b\x65\x1a\xbe\x95\x44\x11\x2d\x51\x48\x71\x85\x17\x93\xc6\x94\x5a\x48\xde\x2f\xd3\x73\x95\x6c\x21\x0a\xe9\xfa\xf5\x71\xcb\x62\xd2\x08\xdf\xa5\x25\x8a\x28\xa6\xc6\x3c\xf8\x5b\x1e\xf0\xaf\x3c\xe2\x7d\xf7\x0d\xdf\xaf\xee\x86\x0f\xf8\xd8\xdd\xab\xde\xda\xdd\x72\x37\xf9\xa8\x64\xb1\xbe\x42\xee\x36\x0f\xdc\x5e\x85\xf8\x3b\x3e\xe2\x63\x77\xd7\xdd\xe6\x43\x7e\xc0\xa3\x87\x0b\x43\xb7\xe7\xee\x10\x8f\xf8\xe7\x92\xe1\x7d\x3e\xae\x34\x43\x77\xef\x71\xd5\x43\x43\x86\xe0\xef\xdd\x0d\x77\x93\x7f\xe1\x11\x1f\x62\x76\x4b\xb3\x23\x72\xb7\xdc\x0d\x1e\x94\x4e\x7c\xcc\x43\xf7\x35\x1f\x52\x05\x1f\xf0\xc8\xed\xf1\xc0\x7d\xc5\x43\xbe\x0f\xfe\x89\x47\x7c\x54\x4e\x04\x1f\x54\x4b\x0c\x78\x7f\x32\x6d\x7d\x05\xff\x05\x00\x00\xff\xff\x8e\x86\x73\x10\x36\x04\x00\x00"

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
