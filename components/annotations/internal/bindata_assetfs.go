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

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x4d\x6f\x1b\x45\x18\x7e\x9c\xb8\x7c\x58\xe2\xb3\x08\x01\xe2\xf0\x72\xa0\x2a\x42\x9b\xac\x1d\x2a\xaa\x4d\x36\x85\x36\x29\x44\x34\x60\x45\x0e\x47\xa4\x49\x76\x6c\xaf\x6a\xcf\x58\x33\x6b\x20\xd0\x43\x48\x11\x08\x51\xf1\x51\x89\x0b\x12\xe2\xc4\x8d\x83\x89\x62\xea\x36\xc4\xfc\x85\x77\x0e\x9c\x90\x10\xfc\x03\x7e\x02\x9a\xb5\xb7\x51\x8b\x90\x10\x12\x73\xd8\x9d\x67\xde\x99\xe7\x63\xf6\xdd\xdf\x1e\x2d\x7f\x05\x00\xf7\x03\x78\x1a\x40\x02\xe0\x01\x00\xdf\x63\x32\x7e\x07\xf0\x22\x80\x3f\x00\x3c\x06\xe0\xc9\x12\xf0\x20\x80\xb9\x12\x10\x03\x88\x4b\xc0\xe3\x00\xde\x2f\x01\xe7\x01\x5c\x9f\xbe\x7f\x29\x01\x0f\x01\x78\x62\x06\x68\x97\x80\xd3\x33\xc0\x5b\x00\xbe\x9c\x05\x1e\x01\xf0\xe7\x2c\xf0\x30\x80\x99\x32\xb0\x0c\xe0\x64\x19\x78\x0a\xc0\x46\x19\xe8\x02\xe8\x96\x81\x1d\x00\xbf\x96\x81\x93\x00\x36\x4f\x00\xa5\xa9\xcf\x13\x00\xee\x9b\xfa\xbb\x07\xc7\x63\x06\x40\x19\xc0\xbd\x53\x3c\xeb\x1f\xdb\x5a\x35\xd3\x56\xf9\xe5\xfa\x1a\x5d\x96\x3b\x73\xf4\xba\x26\x25\x65\x42\x69\x93\xfa\x56\x1a\x25\xba\x92\x84\x4a\xa8\x27\xac\x7d\x47\x9b\x84\x52\x4b\x56\x66\xc5\xb9\x15\x61\xdb\x5b\x5a\x98\xc4\xd2\xda\x4a\xb1\xb8\xaa\xc4\x56\x47\x26\x05\x7c\xc5\x88\xa6\x50\x82\x44\x92\x18\x69\x2d\xe9\x26\xbd\xda\x68\xd4\xc9\x6b\xa6\x8a\x9a\xda\x74\x45\x46\xed\x2c\xeb\x45\xf3\xf3\x6d\x6d\xb3\xa8\xa7\x4d\x76\xf7\x69\x9b\x69\x23\x5a\xb2\x58\xae\x17\x7e\x9a\xda\xd0\x96\xb0\xe9\x36\x89\x7e\xd6\xd6\x26\x7d\x4f\x64\xa9\x56\x77\x24\x99\xa6\xbb\xcb\xfb\x66\x91\xef\xbf\x53\x88\x24\x21\xe1\x93\x6f\x48\xef\x39\x58\xb7\xad\x34\x09\xce\xf7\x5b\x36\x68\xe8\x88\x12\xf9\xf6\x4b\x97\xd3\xb6\xe8\xea\x39\xd3\xaf\xd4\xdf\x68\x04\x17\x8c\xcc\xb9\x83\x15\x91\xc9\x88\x6a\x61\xf5\x6c\x10\x2e\x04\xb5\x05\xaa\x2d\x44\x67\xce\x3c\x1f\x2e\x84\x61\xe5\x92\xb0\x59\xd0\x30\x42\xd9\x8e\xc8\xb4\x89\xe8\xb5\x9c\x83\xd6\xfb\x46\x74\x75\xa2\x69\xe9\x0e\xe2\xe5\xca\x25\xa1\x5a\x7d\xd1\x92\x41\x43\x8a\x6e\x44\xb7\x71\x44\x1b\x7d\x6b\x53\xa1\x2a\xeb\x6b\xeb\xab\xc1\x9b\xd2\xd8\x54\xab\x88\xaa\x73\x61\xe5\x82\x56\x99\x54\x59\xd0\xd8\xe9\xc9\x88\x32\xf9\x6e\x36\xdf\xeb\x88\x54\x2d\xd2\x76\x5b\x18\x2b\xb3\x78\xb3\x71\x31\x38\x7b\xbc\xcf\xfb\x69\x4a\x13\xac\xaa\x6d\x9d\xa4\xaa\x15\x51\xa5\xde\xe9\x1b\xd1\x09\x2e\x6a\xd3\xb5\x11\xa9\x5e\x0e\x6d\xbc\xb0\x48\x93\x69\xac\x9e\xad\x86\x71\x5c\xa5\x53\xa7\xc8\x4f\xc3\x67\xe2\x6a\x95\xce\x51\x48\x51\x8e\x97\xe3\x5a\x51\x5a\x8a\x5f\xf0\xd3\xd3\xf9\xb6\xa5\x6a\x48\x57\xae\x4c\x8e\x2c\xc7\xb5\xf0\x39\x3a\x47\x55\x8a\xa8\xb6\x08\xfe\x9a\x0f\xdd\x67\xee\x63\xff\x41\xe6\x88\xbf\xe1\x21\xf1\x91\xbb\xca\x3f\xf2\x90\x8f\x88\x87\xee\x03\x3e\xe4\x11\xb9\xab\x7c\x8b\x07\x7c\x83\x07\x7c\xe4\x3e\x25\xfe\x99\xc7\x7c\xe8\xae\xf1\x0d\x1e\xf3\x3e\x0f\xdc\x1e\x0f\x3d\x26\x1e\xf9\xda\xc0\xed\x4e\xea\x58\x5b\xc9\x31\x1f\xf9\x3a\x0f\xf9\x26\xf8\x3a\xdf\x9a\x48\x7a\x05\x1e\x83\xbf\xe0\x03\xb7\xeb\x95\xa8\xe8\xcd\xdb\xed\xcc\xfb\xe4\x3e\xe4\xb1\xdb\xe5\x9f\x26\x22\x7f\x6f\x6b\xfe\xce\xed\xe6\x02\x23\x6f\xd4\x7d\xc2\xc3\x82\x06\xfc\xed\xb1\x13\xe2\x03\x3e\x74\x9f\x13\xff\x90\xc7\xf0\xae\xc7\x7c\x93\x78\xc0\xfb\x6e\x2f\x57\x18\xf9\x74\xee\x23\x1e\xf1\xe8\x5f\x5e\x04\x15\x49\xbc\x57\xaf\xf6\x0f\x77\xf2\xbf\x2b\x1f\xf0\x38\x67\xdf\xe7\x91\xdb\x73\xd7\xfc\x0f\xf4\x57\x00\x00\x00\xff\xff\xb5\x7a\xfe\x38\x69\x05\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 1385, mode: os.FileMode(420), modTime: time.Unix(1541273623, 0)}
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
