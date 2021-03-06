// Code generated by go-bindata.
// sources:
// templates/views/send.html
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/mail.mo
// locales/ru/LC_MESSAGES/send.mo
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

var _templatesViewsSendHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4f\x6f\xdb\x3e\x0c\xbd\xf7\x53\x10\x3a\xfd\x7e\xc3\x1c\x21\xe9\x50\xec\x60\x07\x03\x76\xd9\xa1\xc5\x86\xb5\x3b\x17\xb2\xc5\xd4\x2a\xf4\xc7\x93\xe4\x34\xab\x91\xef\x3e\xc8\xb2\x5d\x37\xf1\xb2\x16\x5b\x80\x2c\x97\x90\x12\x25\x52\xef\x3d\xc9\x6c\x1a\xe0\xb8\x12\x1a\x81\x14\x46\x7b\xd4\x9e\xc0\x76\x7b\x96\x72\xb1\x86\x42\x32\xe7\x32\x62\xcd\x03\x59\x9e\x01\x00\x8c\x47\x0b\x23\x13\xc5\x93\xf9\x02\x82\xe5\x54\x6f\x6d\x5c\x32\x5f\x74\xf1\xbb\x6b\x36\xb7\x15\xd3\x28\x47\xb3\xfb\x11\x5e\x78\x89\x3b\x11\x6d\x54\xb9\x58\x36\x0d\x88\xf9\x7b\x0d\xe4\x8a\x09\x49\x60\x06\xdb\x6d\x4a\xcb\xc5\x44\xf0\xb8\x50\x89\xcc\xae\xc4\x86\x2c\x53\xca\xc5\x7a\x27\xf7\xc4\xd0\xb3\x72\x7a\x50\x26\x72\xac\x8c\x55\x7d\x60\xb0\x93\xd2\x58\xf1\x68\xb4\x67\x12\x5a\x5f\xb2\x1c\x65\x22\x71\xe5\x09\x58\x23\x31\x86\x11\x50\xe8\x4b\xc3\x33\x52\x19\xe7\x09\x08\x9e\x11\x87\x9a\x27\xaa\x3d\x14\x2b\xbc\x30\x3a\x23\x34\xb8\x34\x4c\x50\x02\xda\xac\x99\x14\x9c\x79\xdc\xaf\x63\xb7\x66\xe1\x51\xc5\xfc\x77\xd6\xd4\xd5\x44\xe5\xc3\xaa\xb6\xc0\x10\x9b\x11\x6f\xc8\x13\xb3\xda\x5b\x23\x63\xf9\xd0\xf1\x7c\xde\xd3\x7c\x3e\xc9\xf2\xd4\x6f\x60\xeb\xe6\x73\xe4\x0a\x52\x57\x31\x3d\xc8\x0a\xbf\xd7\xc2\x22\x27\xcb\x37\x29\x0d\x13\x07\x0a\xa5\x6d\x2d\x07\x02\xf6\x85\x79\xd1\x17\x7c\xf1\xe2\x82\x53\xa1\xab\xda\x83\xff\x51\x61\x46\x3c\x6e\x3c\x79\x46\x6f\x87\x4b\x64\x2c\xe0\xa5\x99\xc2\x68\x55\x92\x15\x58\x1a\xc9\xd1\x66\x04\x37\x4c\x55\x12\xe7\x1f\xb8\x51\x4c\xe8\x59\x61\xd4\x5b\xe8\x06\x17\xa3\x41\x02\x6b\x26\x6b\xcc\x08\x81\x1e\x8b\x31\x2a\x07\xd0\xd8\xd3\xec\x4b\xa6\xfe\x58\x22\xae\xce\xef\xb1\xf0\xc7\xd3\xc9\x75\x9f\x20\x88\xe5\x5f\x12\xc3\x80\x4c\x54\xc4\xe0\xf6\xfc\x9e\x1e\x97\xe1\x58\xc7\x23\xf2\x63\x7c\x34\x21\x66\x39\x09\x36\x1d\x4a\x2c\xfc\x14\x85\x10\xa7\x16\xe1\x8d\x7e\x70\x19\x79\x37\x5c\xec\xb6\xfa\xf6\xb2\x07\xeb\x70\x82\x36\x89\xa9\xc2\xdb\xdd\xd3\x5e\x49\x26\x34\x79\xfa\x68\x7d\x09\x3e\x44\x29\xc5\x4f\x57\x8c\x7f\xf5\xc6\xa5\x57\x72\xb4\xef\xa7\x9b\xab\xcb\xd7\xec\x98\xd2\x78\xe2\xd3\x53\xa5\x42\xe7\xd8\xdd\x11\x85\x79\xd5\x27\x38\x09\x4d\x06\x29\x30\x8b\x6c\x52\x95\x16\x9d\x78\x64\xb9\xc4\xdb\x3e\x6c\x4f\xa0\x03\x5c\x41\xa3\xbd\xb3\x4c\x69\xbf\xe0\x98\x04\x4b\x7d\xeb\x8c\x14\x7c\xb2\xab\x9a\x5a\xf0\x32\x31\xfc\x0a\x58\xc5\x13\xb3\x5a\x39\xf4\xc9\xf9\xef\x50\xcd\x6b\xef\x8d\x1e\xda\x2a\xd2\x3d\xe1\xae\xce\x95\x78\x7a\xc4\x73\xaf\x21\xf7\x3a\x71\x75\x51\xa0\x73\xa3\xeb\x74\xdd\x2e\x8a\xd7\x29\xee\xf5\x37\x71\x4c\x69\x00\xe2\x60\x13\x3a\x72\x3b\xb3\xfb\x6b\x1a\x40\xcd\x83\x72\xcf\x46\x6d\xfb\xbd\x6b\x3b\xf6\xa6\x01\xe7\x99\x17\x45\x78\x0e\xe0\xbf\x68\x7f\xfb\x7a\x09\x84\x72\xe6\xca\xdc\x30\xcb\x29\x73\x0e\xbd\xa3\x6b\xd4\xdc\x58\x47\xbb\xa6\xd2\x58\x7a\x3f\x72\x66\x4a\xe8\x59\xd8\x75\xc5\xa4\xc3\xff\xbb\xcd\x63\xe6\x9f\x01\x00\x00\xff\xff\x8e\x6a\xf5\xf6\x2f\x0c\x00\x00"

func templatesViewsSendHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsSendHtml,
		"templates/views/send.html",
	)
}

func templatesViewsSendHtml() (*asset, error) {
	bytes, err := templatesViewsSendHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/send.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x51\x4b\x4f\x13\x5f\x14\xff\x0d\xd3\xff\x5f\x2d\x1a\x95\x44\xe3\x42\xc3\xc5\x44\xa2\x8b\xc1\x16\x5d\x98\x81\x01\x1f\x81\x48\xa4\x09\xc1\xe2\x96\x5c\xe9\xa5\x4c\x6c\xe7\x36\xf7\xce\xf8\x64\x51\xba\x61\x43\x82\x1b\xa3\x71\xc3\xca\xb8\x32\x69\x88\x44\x22\x30\x7c\x01\x17\xe7\x7e\x01\xb7\x7e\x0d\x33\x9d\x29\x8f\x18\xe3\x59\xcc\xf9\xbd\xce\x99\x93\xdc\x9f\x7d\xb9\x77\x00\xd0\x0b\xe0\x32\x80\x10\xc0\x79\x00\x3f\x90\xd6\x39\x0b\xc8\x25\x9a\x95\x66\x2e\x5a\xc0\x19\x00\x57\x2d\xe0\x12\x80\x82\x05\x5c\x00\x30\x65\x01\x67\x01\xcc\x67\x39\x99\xf5\x37\x99\xbe\x9a\xf5\xf7\x16\x30\x6e\x01\x9f\x2d\x60\x00\xc0\xaf\x1e\xe0\x24\x80\x53\x36\x70\x1a\x40\x9f\x9d\xea\xfd\x36\x70\x05\xc0\xbd\x4c\x9f\xb7\xd3\x5c\xcd\x4e\xef\x89\x6c\xa0\x1f\xc0\xeb\x2c\xff\xd1\x06\x2c\x00\xff\x23\xcd\xe5\x01\xfc\x97\xec\x05\xd0\x03\xc0\xc6\xf1\xca\x1d\xc1\xbd\x59\x3f\x91\x7c\x4a\xdc\xaf\x61\x41\x06\x8b\x7e\x35\xf7\x50\xea\xb0\x8b\xa7\x45\x18\x0a\xd5\x65\x49\x8a\x2d\x2a\x59\x67\xbc\x52\x51\x42\xeb\x3f\x8d\x80\xd7\x45\x57\x9d\xe1\x5a\xbf\x90\xaa\x72\xc0\xa5\x3a\x58\xfc\xb8\x54\x9e\xe9\xe2\x39\x2d\xd4\x91\x39\x27\xe4\x4f\x73\xf5\xe4\xa0\x59\xd1\x90\x2a\x74\x4a\xba\xea\x57\x9c\xfb\x51\x55\x3b\x65\xe9\xb2\x8a\x78\x7e\xf7\x99\xbf\xc4\xeb\x72\x48\x45\xf9\x69\xae\x43\xa7\xac\x78\xa0\x6b\x3c\x94\xca\x65\x8f\x3a\x16\x2b\x45\x8a\xd7\x65\x45\xb2\xd1\x63\xf9\xb1\xfc\x34\x0f\xaa\x11\xaf\x0a\xa7\x2c\x78\xdd\x65\x07\xdc\x65\xb3\x91\xd6\x3e\x0f\xf2\xa5\xa9\xd2\x84\xf3\x44\x28\xed\xcb\xc0\x65\xc5\xa1\x42\xfe\x81\x0c\x42\x11\x84\x4e\xf9\x55\x43\xb8\x2c\x14\x2f\xc3\x9b\x8d\x1a\xf7\x83\x11\xb6\xb0\xc4\x95\x16\xa1\x37\x57\x9e\x74\xee\x1c\xe6\x92\x7b\x16\x85\x72\x26\x82\x05\x59\xf1\x83\xaa\xcb\xf2\x33\xb5\x48\xf1\x9a\x33\x29\x55\x5d\xbb\x2c\x68\x74\xa8\xf6\x6e\x8d\xb0\x14\x7a\xc1\xb5\x62\xc1\xf3\x8a\x6c\x70\x90\x25\xb0\x30\xe0\x15\x8b\x6c\x9c\x15\x98\xdb\xe1\x63\xde\x70\xd7\x1a\xf5\x6e\x27\xf0\x7a\x27\x36\x5a\x2c\xb0\xe5\xe5\x74\x64\xcc\x1b\x2e\xdc\x60\xe3\xac\xc8\x5c\x36\x3c\x02\xfa\x42\x3b\xb4\x45\xdf\x4d\xcb\x34\x29\xa6\x3d\xda\xa3\xb6\x59\x67\xb4\x4f\xb1\x59\x35\x2d\x6a\x83\x3e\x51\x6c\x56\x4c\x0b\xb4\x41\xdb\x66\xc5\xac\xd1\x2e\xc5\xa0\xb7\xf4\xd5\x34\x69\xcb\xac\x30\x8a\x4d\x8b\xf6\x4d\x93\xda\xb4\x49\xdb\xa6\x45\x5b\xb4\x63\xd6\x41\x1f\x68\x37\xd9\xf4\x17\x77\x83\xda\x9d\x3f\xee\x98\xb5\x84\xc4\xa6\x69\x5a\xe8\x3c\x79\x77\x70\x3f\x75\xe9\x1b\xc5\xb4\x49\xed\xc3\xd1\x7f\x9e\xfc\x3b\x00\x00\xff\xff\x99\x0a\xb1\xb2\xbc\x03\x00\x00"

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

var _localesRuLc_messagesMailMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbd\x8e\xd3\x40\x14\x85\x4f\x56\x4b\xe3\x92\x9a\xe2\x6e\xc1\x0a\x24\x66\xb1\x03\x05\x9a\x78\x12\x04\xda\x95\x10\x6b\x69\x59\x99\xed\x47\xf1\xe0\x58\xd8\x33\xd6\xcc\x18\x81\x94\x02\xa5\xa1\xa4\x82\xb7\x40\x4a\x43\x03\x4a\x78\x85\xf1\x0b\xf0\x2c\x28\x0e\x7f\x7b\xaa\xf3\xe9\x7c\x57\xba\x3f\x6f\x1e\x7e\x02\x80\x03\x00\xb7\x00\xdc\x03\x70\x03\x40\x8a\x7d\x2e\x00\x1c\x02\x78\x01\x60\x36\x02\xae\x00\x1c\x01\xf8\x3c\x02\x46\xbf\x9d\x03\xfc\x97\x4c\x56\x35\x2e\x55\x6b\xac\x67\x99\x2b\xab\x82\x3d\xe9\x4a\xc7\x72\xc3\xa9\x50\x6f\x1e\xbf\xae\x16\xb2\x31\x27\xb6\x8b\xce\xa5\xf3\x2c\xb7\x52\xbb\x5a\x7a\x63\x39\x3d\x1f\x26\xca\x3a\x2b\x1b\x53\x18\x4a\xaf\xf9\xd3\xe8\x5c\xea\xb2\x93\xa5\x62\xb9\x92\x0d\xa7\xbf\xcc\xe9\xb2\x73\xae\x92\x3a\xca\x9e\x65\xa7\xec\x4a\x59\x57\x19\xcd\x29\x39\x89\xa3\xa7\x46\x7b\xa5\x3d\xcb\xdf\xb5\x8a\x93\x57\x6f\xfd\xfd\xb6\x96\x95\x9e\xd0\x7c\x21\xad\x53\x5e\xbc\xcc\xcf\xd8\xa3\x7f\xde\xee\x9f\x57\xca\xb2\x53\x3d\x37\x45\xa5\x4b\x4e\xd1\x45\xdd\x59\x59\xb3\x33\x63\x1b\xc7\x49\xb7\x03\x3a\xf1\x60\x42\xfb\x2a\xf4\xed\x24\x16\x22\xa1\xe3\x63\xda\xd5\xf8\x48\x24\x09\xcd\x28\x26\x3e\xf0\x54\x8c\xff\x4c\xa9\x78\xb8\xab\x77\x06\x2d\x4d\x62\x5a\x2e\xf7\x27\x53\x31\x8e\xef\xd2\x8c\x12\xe2\x34\x9e\x20\x7c\x09\xdf\xc3\xd7\xf0\xad\x5f\xf5\xef\xc3\x36\x6c\xc2\x26\xac\xfb\x8f\x14\x7e\x84\x6d\xff\xa1\x5f\x85\x35\x7e\x05\x00\x00\xff\xff\xb1\x6b\x5b\x81\xb8\x01\x00\x00"

func localesRuLc_messagesMailMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesMailMo,
		"locales/ru/LC_MESSAGES/mail.mo",
	)
}

func localesRuLc_messagesMailMo() (*asset, error) {
	bytes, err := localesRuLc_messagesMailMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/mail.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesSendMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x51\x41\x6b\xd4\x40\x14\xfe\xb2\xad\xda\x2e\x0a\x52\x04\x11\x3d\x3c\x0f\x96\x8a\xa4\x6e\xaa\x88\xa4\x4d\x2b\x4a\x8b\x62\x17\x4b\x1b\xbd\x8f\x9b\x71\xbb\x76\x77\xb2\x64\x12\xb1\xd0\x43\xbb\xbd\x54\xf0\xea\x45\x41\xd1\xb3\x87\x2a\x2c\xd4\xda\xa6\x7f\xe1\x8d\x3f\x40\xf4\x2e\x5e\x3c\x7a\x50\x92\x6c\xad\x82\xef\x30\xef\xfb\xde\xf7\x0d\xef\x1b\xe6\xf3\x50\xff\x33\x00\x18\x04\x70\x06\x40\x00\xe0\x18\x80\xb7\x28\xea\x0b\x80\xa3\x00\xbe\x02\xe8\x07\xf0\x1d\xc0\x11\x00\x3f\x00\x9c\x00\xf0\x13\x40\x19\xc0\x71\xab\xd0\x4f\x5b\x85\x4e\x16\x50\x02\x30\x62\x01\x53\x16\x70\xc1\x02\x4e\x02\x58\x2a\x15\xbe\xf5\x12\x30\x04\x60\xa3\x04\x5c\x01\xf0\xa2\x04\x9c\x02\xf0\xa9\x37\xff\x56\x02\x06\x00\xfc\xea\xf5\x81\x3e\xc0\x02\x70\xb8\xb7\xff\x10\x80\x3e\x14\x3b\x06\x7b\x59\xfb\x71\x50\x03\xfb\xe0\x46\xa8\x62\xa9\x62\x8a\x97\xdb\x12\x37\xfd\xea\x2c\xaa\x52\x6b\x51\x97\xfb\x9d\xb4\x54\x01\xe9\xa4\x56\x93\x5a\x63\xae\x29\x1a\x8a\x62\xf9\x38\xc6\x82\x54\x01\x16\x92\xfb\x0f\x65\x2d\x86\x7f\x07\xf3\xb2\x1d\x46\xb1\x5d\xd5\xf5\x46\x60\x5f\x4f\xea\xda\xf6\x43\x97\x02\xf9\xe8\xda\x52\x63\x51\xb4\xc2\xd1\x28\x29\xcf\x0a\x1d\xdb\x7e\x24\x94\x6e\x8a\x38\x8c\x5c\xba\x9d\x4b\x54\x4d\x22\xd1\x0a\x83\x90\x26\xfe\xf1\x4f\x96\x67\x85\xaa\x27\xa2\x2e\x6d\x5f\x8a\x96\x4b\x7f\xb8\x4b\xf3\x89\xd6\x0d\xa1\xca\xd5\x5b\xd5\x69\xfb\x9e\x8c\x74\x23\x54\x2e\x39\xa3\x95\x72\xef\x49\xb6\xbf\xdc\x96\x6e\x9e\xf5\x62\x3b\x8b\x3d\x4e\xb5\x45\x11\x69\x19\x7b\x77\xfd\x19\xfb\xea\x81\x2f\xcb\xf3\x40\x46\xf6\xb4\xaa\x85\x41\x43\xd5\x5d\x2a\xcf\x35\x93\x48\x34\xed\x99\x30\x6a\x69\x97\x54\x3b\xa7\xda\xbb\x34\x4e\x05\xf4\xd4\x39\xa7\xe2\x79\x0e\x0d\x0f\x53\x06\x2b\x67\x3d\xc7\xa1\x29\xaa\x90\x9b\xf3\x49\x6f\x6c\x5f\x9a\xf0\x2e\x67\x70\x24\xb7\x4d\x38\x15\x5a\x59\x29\xae\x4c\x7a\x63\x95\xf3\x34\x45\x0e\xb9\x34\x36\x0e\x7e\xc3\x5b\xbc\x47\xbc\xcd\x29\xef\x9a\x0e\x77\xf3\x73\xb3\xf8\x14\x7e\xcd\x29\xa7\xfc\xce\x3c\xc9\xe6\xbc\xc5\xdd\xff\x8c\xc8\xac\x9b\x35\xde\xe3\xae\xd9\xe0\x5d\x4e\x89\x53\xd3\xe1\x3d\xb3\xca\x9b\xfc\x9e\x3f\xe6\xae\x14\xfc\xca\xac\x72\x6a\xd6\x4c\x87\x53\xfe\x40\xf9\xa2\xed\x8c\x82\x5f\xfe\x65\xdf\x32\x1d\xf3\x34\xcb\xd4\xe5\x1d\xde\x04\x3f\xe7\x94\x77\xcc\x3a\x7e\x07\x00\x00\xff\xff\xba\x72\x79\x44\x11\x03\x00\x00"

func localesRuLc_messagesSendMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesSendMo,
		"locales/ru/LC_MESSAGES/send.mo",
	)
}

func localesRuLc_messagesSendMo() (*asset, error) {
	bytes, err := localesRuLc_messagesSendMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/send.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/views/send.html": templatesViewsSendHtml,
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
	"locales/ru/LC_MESSAGES/mail.mo": localesRuLc_messagesMailMo,
	"locales/ru/LC_MESSAGES/send.mo": localesRuLc_messagesSendMo,
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
				"mail.mo": &bintree{localesRuLc_messagesMailMo, map[string]*bintree{}},
				"send.mo": &bintree{localesRuLc_messagesSendMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"send.html": &bintree{templatesViewsSendHtml, map[string]*bintree{}},
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
