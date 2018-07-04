// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/grpc.mo
// locales/ru/LC_MESSAGES/manager.mo
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

var _templatesViewsManagerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1b\x5d\x73\xdb\xb8\xf1\x3d\xbf\x02\x83\xf3\x75\xe2\x36\x94\x1c\x3b\xe9\xb4\x8a\xe4\x4c\xe7\xda\x4e\xd3\x49\x72\x9d\xdc\xf5\x29\x4d\x3d\x10\xb9\x14\xe1\x42\x04\x0f\x00\x25\xbb\x1a\xfd\xf7\x0e\xc0\x2f\x50\x22\x29\x89\x84\x12\xb7\x53\x3e\x24\x32\x17\xfb\x81\xdd\x05\xb0\xbb\x58\x6e\x36\x28\x80\x90\xc6\x80\xb0\xcf\x63\x05\xb1\xc2\x68\xbb\x7d\x36\x0d\xe8\x0a\xf9\x8c\x48\x39\xc3\x09\x59\x80\xa7\xa8\x62\x80\x6f\x9f\x21\x84\x90\x0d\x34\xef\xef\x18\x84\x2a\x07\x9a\x01\xd1\xcd\xed\xe2\xd3\xdf\x7e\x98\x8e\xa3\x9b\x1c\x65\x1c\xd0\xd5\xed\xb3\xfc\xbf\x1a\x79\x9f\x01\x11\x21\x7d\xc0\xb7\x05\x74\xb3\x41\x34\x44\x23\x10\x82\x8b\x5d\x61\x08\x03\xa1\x90\xf9\xd7\x0b\x48\xbc\x00\x51\xfc\x41\xe5\x92\x4a\x49\xe6\x95\x98\xf3\x54\x29\x1e\x23\xf5\x98\xc0\x0c\x67\x7f\xe0\x8a\x2b\x97\x80\x51\x40\x14\x29\x50\x73\xe2\x18\x11\x41\x89\x17\xd1\x20\x80\x78\x86\x95\x48\x01\xdf\xfe\x4a\xd1\x25\xc8\x37\xd3\x71\x46\x26\x63\xb0\xd9\xd8\x42\x66\xc2\x6f\x36\x08\xe2\x40\xbf\xa8\x89\x2d\xf8\xba\x41\x79\x3e\x67\xde\x32\xf0\x5e\x5e\xe3\x92\xa0\xd0\x73\x42\x17\xf2\x05\xba\x90\x20\x56\xd4\x07\x34\x99\xa1\x51\xfe\x5b\x6a\xc2\x85\x96\x37\x9b\x72\xcc\xbb\x40\x8f\x12\x90\x30\xe2\x43\xf9\x76\xf4\x91\x2c\x01\xe1\x11\x46\xf8\xce\x58\xb5\x32\x90\x25\xc4\xc3\x5d\x42\x62\x60\x96\xf9\xf6\x47\xd8\xd6\xaf\x8d\x8a\xae\x6f\x2d\x31\x32\x86\xdb\xed\x74\x1c\x5d\x37\x0c\xee\x32\x7a\x6d\x60\xc3\xab\x9a\x38\x85\xa7\xde\x3e\xeb\x64\x42\x7c\x9f\x8b\x80\x6a\xab\xd3\x60\x86\x89\x57\xd7\xd8\x76\x8b\x91\xe0\x0c\x66\x58\x91\x39\xa3\xb2\xb0\xfc\x32\x65\x8a\x4a\x60\xe0\xeb\xf7\x50\xb8\xc0\x1e\xab\xba\xc5\x96\x2f\xd0\xc5\x12\x54\xc4\x8d\x29\x4a\x8d\x7c\x30\xaf\x6a\x76\xeb\x12\xb9\xc9\x14\x3b\xfc\x72\x2e\x3b\x26\xcf\x5e\xb6\x5b\xbc\x81\x10\x0d\x4b\xb4\x77\x71\x92\xaa\x9f\x1f\x13\xe8\x92\xd3\xc8\x4a\x6a\x92\x7a\x11\x90\x80\xc6\x0b\x4b\x91\x99\xae\xa3\x3d\x5d\x7b\x35\xd1\xb5\xee\xcd\xda\x53\x7c\xb1\xd0\x98\x3e\x67\x8c\x24\xe5\x92\x4c\x88\x80\x58\xcd\xf0\x77\x4d\x46\x8b\x04\x84\x33\xfc\x9d\x7f\x04\x0f\x63\x50\x78\x48\x48\x1c\x40\x30\xc3\x21\x61\x9a\x85\x79\xab\x9d\x48\x70\xa6\x9d\xf1\x30\xa1\x76\x8b\x94\x9a\x89\x5e\xd5\x55\x93\xaf\x99\x8a\x92\xb5\x3a\x5e\x75\xd3\x9b\x8e\xc9\x81\x01\xda\x65\xb4\x9e\x8f\xd1\x41\x4d\xaa\x42\xcf\xa8\x52\x78\x69\xba\xcc\xf9\x32\xed\x30\x32\x07\xc6\x20\x98\x3f\x1e\x65\xcc\x23\xf4\xb3\xeb\xe5\xde\x9c\x07\x8f\x47\x20\xee\x22\x37\xef\x57\xc7\x63\xb7\xed\x65\x9d\x14\x0e\x59\xac\x11\x89\x16\x4c\x43\x82\x42\xe2\x11\x21\xf8\xda\x13\x74\x11\x29\xfd\x67\xb8\xd6\x9b\x1f\x3d\x9d\xae\xe5\x51\xe5\xba\x2d\x7c\xab\x0f\xb1\xdd\x7d\xe0\x27\x25\x80\x2c\xfb\x10\xd3\xcf\x54\x26\x24\x2e\x26\x6e\xbc\x08\x99\x7f\xbd\x35\x11\xb1\xde\x2a\xf4\x8a\xa0\x2f\x7f\x17\x23\x2c\x0d\x23\x8c\x2e\xcc\xa2\xd0\x78\xbd\x94\x91\x9f\xb7\x27\x09\x69\x96\xe0\x69\x28\xc7\x9e\x5e\xdd\x7c\x4f\x1b\xde\x72\xea\x9d\x24\x77\xc8\xc5\x12\x65\xf6\x9d\xe1\x84\x9b\x63\xce\x57\x94\xc7\x33\xfc\x36\xff\xe1\x13\xc6\xaa\xc8\x88\x30\x86\x34\x92\x17\x71\x41\xff\xcd\x63\x45\x58\xbe\x2b\x0b\x90\x29\x53\x33\x2c\x8e\xd8\x75\xb2\x5d\x45\xd3\xd9\xd9\xe9\x57\x84\xd1\x80\x28\x2e\x4e\x9c\x88\x99\x0c\xd5\x0e\x9a\x47\x74\x59\x84\x86\x51\x4c\x96\x30\xc3\x77\xb9\x38\x18\xad\x08\x4b\x61\x86\x1b\xe2\x12\xc7\x1c\xb3\xf9\xd6\x18\xd6\x77\xfa\xa6\x00\xe5\xd0\xd3\x72\x30\x8f\xfe\x4c\x81\x75\xc7\x11\xad\x53\xb0\xc3\x75\x1d\xd1\x68\x3b\x26\x3c\x96\x74\x75\xea\x3e\x58\x92\x34\x74\x6a\x44\x51\x46\x3a\xe2\x2b\x10\xf9\x6f\xa9\x04\x4d\x20\xe8\xc9\x23\xe3\xa3\x03\x8c\x21\xf8\xa2\x3f\x72\x2e\x00\x5a\xd3\x40\x45\x33\xfc\xf2\xea\x7b\x6b\xef\x32\xe6\x28\xb6\x2e\x15\x39\xe3\xf2\xba\xc6\x45\x5b\xfe\x0c\x4c\xea\x53\x79\xaf\xb7\xe7\x33\x70\xb9\xa9\x71\xf9\x83\xd9\x6a\x5c\xb0\x99\x8e\x87\x58\x55\xf3\x1e\xe8\x53\x3a\x70\xe9\x8f\x5f\xa5\x0c\xe1\x0b\x74\x11\x6a\x3f\x32\x19\x83\xc3\x15\x6f\xb1\xba\x58\x11\xb1\x93\x2a\x18\x96\x7b\x99\x42\x5f\x1e\x0e\x96\x58\x60\x82\x64\x4b\x2c\xe3\x21\x03\x4c\x54\x90\x1d\x44\x00\x55\x9b\x71\x26\xda\x3b\xf9\x01\xa4\x24\x8b\x5e\x71\x56\xa3\x88\x24\x4f\x65\xee\xc9\x8a\x48\x5f\xd0\x44\x4d\x56\x9c\x06\xcf\xaf\x2e\xbb\xb3\x23\x45\xc4\x02\x74\x76\x74\xf8\x20\xf6\x2a\x1f\xd8\x6e\xbd\x7b\xc9\x63\x6c\x69\x3b\x4f\xf9\x0e\xe6\x1b\x47\xcf\x28\x11\x60\xf2\x92\x5e\x82\x59\x65\x91\x2c\x47\xe1\xa9\x32\xd2\x0a\xb2\x2e\x24\xfe\xeb\x4f\x3f\x7e\x34\x12\x27\x02\x9c\xd8\x17\x98\x74\x66\xd0\x3d\xc5\x3a\x11\xf0\xf4\xf0\x76\xf7\x71\xb3\x9a\x5a\x23\x7b\x99\xfa\x3e\x48\x69\x3b\x96\x39\x53\xaa\xb0\xde\x89\x00\x55\x59\xc7\x9c\x24\xc3\xad\x6f\x07\x46\x26\xcc\xf3\x16\x82\xa7\x49\x19\xec\xfe\x92\x52\x01\x81\x09\xec\xec\x5d\xe0\x53\x0e\x40\xdb\xad\x12\x29\x54\x3e\x64\xaa\x0b\xa5\xc5\x4a\x32\x09\x10\xd5\x48\x26\x03\x1c\x22\xe3\x68\x65\xd6\xea\x62\x5a\x81\x9e\x99\x72\x26\x54\xcc\x55\xe3\xfc\x50\x44\x03\xe7\xa2\x64\x2c\xb9\x40\xcf\xe1\x97\xda\x7a\xc1\x01\x4f\xe7\x0c\xf0\xe5\x3e\xe4\xf3\x97\x76\x58\xc8\x38\x51\xcd\x48\x05\xc8\xd5\x86\x6d\x27\x03\x71\xba\x9c\x83\x28\x92\x81\xbd\x13\xac\xdc\xcd\x4c\x2a\x95\x17\x9b\x30\x4a\x88\x52\x20\xe2\x19\xfe\xe7\x67\xef\x37\x5f\xde\x7e\xbe\xf2\x7e\xff\xe5\xd7\xcf\xff\x31\xca\x7e\x5c\xbe\xbd\xa8\xa5\x13\x19\xc5\x3f\x42\x48\x52\xa6\x0c\xd1\xb1\x33\x1b\x18\x6f\x6b\x31\x04\x8d\xd5\xcd\x75\xb3\x4a\x5b\x41\x34\x56\xbf\x7d\xd5\x8a\xd3\x0c\x92\x1d\x8c\xda\x61\xb2\x83\x55\x3b\x4c\x86\xf4\x01\x82\x56\x6e\x1d\xd0\x0c\xd6\xca\xb1\x82\x3e\x71\x37\x7b\x1a\xbe\x95\x76\xd8\xbc\x1d\x96\x76\xd8\xbc\x1d\xd6\x69\xf2\x0e\x60\xa7\xc1\x9f\xb8\xbd\x9f\x98\xb9\xe7\x9c\xb3\x66\x3d\xe6\x90\x73\x28\xd1\x8f\xc0\xff\xd7\x9c\x3f\x1c\x56\xe3\xbd\xf4\xe4\x9a\x2a\x3f\x42\x75\x85\xd6\x0e\xeb\x4a\x6b\x86\x30\x04\x55\x5c\xf6\x75\x94\x28\x95\xa0\xf1\xa2\x65\xff\x69\x85\xcd\x1f\x15\xc8\x16\xdd\xe7\xa0\x73\x28\x5f\xc1\x83\x3a\xd5\x7f\xbf\xa6\xa3\x3a\x9a\x33\xda\xcf\x0f\xff\x14\xa7\xbd\xeb\xe6\x4d\xcf\x34\xbb\x85\x3c\x49\x97\x28\xc3\xb9\x76\x14\xac\x15\x4f\x55\xb5\x80\x17\xe8\x02\xf4\x44\x27\xb3\x42\x1e\xd7\xf3\xd6\xcf\x94\x27\x3a\x4a\xb5\x3d\x43\x73\x1d\x7d\x34\xfb\xa3\x9e\x7d\xae\x7c\xf3\xf6\x9d\xac\x1c\x26\xd7\x80\x0e\xb9\x8b\x5f\xb8\x5c\xaf\xb7\x15\xa1\xa2\xd2\x90\x31\x72\xae\x2e\x07\x79\x9b\xfd\x4c\xc7\xd9\x64\xdc\xc9\x79\x86\xf5\x30\xd5\x6b\x9f\x08\x20\x27\x2e\x7f\xc1\xd7\x72\x86\x6f\xec\x0c\xb2\x4c\xf4\x0b\x92\x6e\x27\xee\xce\x38\x6e\xa9\x15\x3d\x30\x2e\x68\x55\x89\x56\x63\x7e\x77\xd9\x98\x8b\x3a\xcf\x36\xad\xcc\xda\x9b\x2b\x17\x79\x7b\xc9\xa5\xab\xc1\x67\xae\x62\x34\x57\x71\xd9\x24\x94\xe5\xbc\x02\x96\x7c\x05\xdf\x20\xe9\x2d\x65\xb6\x8b\x28\xd9\xbd\xb0\x12\x44\x46\x1e\xc7\xb7\x7d\xef\x40\x5b\x79\xd5\xba\x95\x9c\x50\x3c\x46\xe3\x34\x0e\x79\xa1\x6f\x12\x0c\xb9\xfe\x69\x14\x61\x5f\x81\x09\x4b\xe5\x53\xd7\xde\x89\x37\xbf\x5d\x8f\xbb\x0d\xc7\x81\x50\xc3\x2a\x7b\xc3\xae\x72\x86\x2b\x62\x3a\x1e\x70\x99\x33\x1d\x9b\x6b\xce\x1e\x77\xca\xfd\xd4\x6e\xb5\x16\x9e\xcc\xb1\xe8\x17\x3a\xe6\xe6\xde\xea\x03\xc8\x6f\xfc\x31\x92\xea\x91\xc1\x0c\x07\x54\x26\x8c\x3c\x4e\x62\x1e\xf7\xbe\x3a\x8e\x44\xdf\x80\xfe\x88\x86\xce\xe6\x46\xcd\x93\x19\xd9\x9d\x9b\xd9\x5d\x39\x64\x6d\xa6\x27\x77\x7c\x94\x24\x7b\x9e\xeb\xb6\x28\x2c\xbe\x93\x9c\xd1\xa0\xea\x94\x1d\x42\xcd\x44\x60\x59\xc1\x7b\xb8\x96\x74\x90\xe6\xf9\x10\x2b\xe8\xd3\xd6\x51\x52\xb4\xd3\x48\x99\xce\x97\x54\xed\x9d\x2e\xc5\x65\x83\x95\x18\x64\x17\xcc\x3f\x98\x26\x96\x8b\x21\x19\xe3\x30\xdb\x9e\x86\xa2\xd5\x7f\x96\xc6\xa1\x53\x7c\xe3\x7f\xa8\xaf\x8e\x41\xe8\xac\xad\xee\xc7\x54\xb9\xed\xab\xab\x08\xfe\xbf\xad\x6e\x07\xe5\xbf\xb5\xad\xce\x6d\x53\xd5\xd7\x68\xa8\x1a\xd2\x4c\x35\xa8\xcb\xc3\xee\x3c\x7a\x7d\x96\x26\x2a\xbb\xeb\xe8\xf5\x39\x1a\xa8\x5a\x3b\xb4\x9c\x34\x4f\xf5\x0f\x83\x07\x75\x33\x0d\xe9\x64\x3a\xd0\xc5\x64\xed\xa0\xc3\xda\x98\xce\xdd\xc2\x34\xd0\xb1\x9d\xb7\x2e\x0d\x6e\x5b\x3a\x5f\xcb\xd2\xf0\x76\xa5\xf5\x13\x6a\x57\x2a\x5b\x95\x7a\x09\xf5\x75\x5b\x95\x1c\x56\x6a\xdd\xb6\x28\x39\xc9\xc0\xbf\xe5\x6a\x19\xd0\xd3\x34\xac\x46\xbb\xbb\x4c\x49\xe2\x64\x89\x1e\x9c\xce\x92\x24\x2e\x0a\x65\xdf\xd2\xf0\xfd\x4f\xcb\x61\x52\xf7\x2e\x16\xf5\x2a\x14\x9d\x2f\x04\x3e\x72\xe8\x11\xc3\x0e\x0c\x39\x72\xdb\x3a\xe6\x9b\xc2\x96\xe3\xa5\xf5\x3c\x7a\x42\xdf\xeb\x1d\x76\xba\x0e\x35\xb6\x23\x1f\xfe\x6e\xd6\x2e\x04\xd4\xe9\xd4\xbf\xc6\xb6\xea\x8a\xd6\xd7\xe0\xda\x0a\xd9\xa7\xe0\xa6\xf6\x67\xdd\xed\x8f\x7d\xbd\x91\x18\x42\xa3\xb2\x3a\xb6\x29\xd9\xae\x23\xaa\xc0\x93\x09\xf1\x61\x82\x12\x01\xde\x5a\x90\xe4\x4d\x09\xd6\xe9\x4c\xc8\xf8\xda\x7b\x9c\x20\x92\x2a\x5e\x41\x96\xe4\xc1\x8b\x80\x2e\x22\x35\x41\x37\x57\xb0\xac\x20\x26\xf2\x9e\xa0\x97\x57\x57\xdf\x67\x2f\xf3\x12\xe8\x68\xe7\x6e\x09\xe5\x05\xc2\x8d\x45\x53\x2c\x68\x9c\x7d\x74\x37\x41\x57\x6f\x76\x01\x73\xae\x14\x5f\x1a\x48\x8d\xb0\x55\xfc\xb4\xa8\x25\x5c\x52\x45\x79\x3c\x41\x02\x18\x51\x74\x05\x75\x71\x4c\x0d\xb2\x71\x3c\x99\x4b\xce\x52\x05\x95\x00\xb9\x48\xaf\x93\x87\xea\x9d\xe2\xc9\x04\x5d\x97\xaf\x0a\xb2\xd9\x5d\x8a\x45\xd7\x54\xdb\x08\xa3\x0b\x2d\x89\xa6\xd3\x38\x7e\x64\xf7\x79\xb6\x60\x33\x08\x5b\x90\x2d\xd5\xda\xc6\x6d\xb3\x44\x03\x9a\xb1\x48\x85\xba\x02\xa1\xa8\x4f\x58\xc1\x5a\xf1\xa4\x20\x31\x1d\x1b\x1f\x6b\x73\xc5\x7b\x69\x1c\x71\xb3\x41\x52\x11\x45\xfd\xbf\xfc\xfc\xe1\x3d\x7a\x9e\xfd\xfe\xfb\xa7\xf7\x08\x8f\x03\x22\xa3\x39\x27\x22\x18\x13\x29\x41\xc9\xf1\x0a\xe2\x80\x0b\x39\x2e\x3f\x36\x1b\xdf\x5b\x7f\x8c\x96\x34\x1e\x69\xaa\xa6\xd5\xf6\xf2\x00\xf1\x85\x48\xfc\x82\xee\xbd\x1c\x2f\x49\x4c\x16\x20\x0c\xbe\x12\x69\x89\x9e\x09\xfe\x9f\x00\x00\x00\xff\xff\xa5\x69\x8d\xb4\x50\x41\x00\x00")

func templatesViewsManagerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsManagerHtml,
		"templates/views/manager.html",
	)
}

func templatesViewsManagerHtml() (*asset, error) {
	bytes, err := templatesViewsManagerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/manager.html", size: 16720, mode: os.FileMode(420), modTime: time.Unix(1530694973, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsManagerMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4d\x6f\xdb\x38\x10\xfd\x2b\x5a\xae\x61\x90\xb0\x4c\xec\x2e\xf6\x24\x2d\x11\x2c\xd2\x53\x81\xb6\x87\xa2\xa7\x24\x07\x46\x1a\x59\x8c\x69\x92\x1d\x52\x4a\x5c\x47\xff\xbd\xa0\x3e\xfc\x15\xbb\x0d\x74\x10\x66\x38\xf3\x38\xf3\xde\x70\x66\xb4\x6a\x4c\x11\x94\x35\x94\xed\x66\x94\x70\xd9\x1b\xc9\x63\x13\x82\x35\xa3\xb5\x44\xd8\xd8\x16\x08\xe3\x85\x56\xc5\xfa\x38\xa5\x95\x98\x80\x98\xd1\x50\x2b\xcf\x52\x14\xc0\x0b\x6d\x3d\xf8\x40\x09\x57\xc6\x35\x61\xb9\x42\xdb\x38\xc2\x52\x27\x90\x3b\x89\x60\x02\x65\xb9\xaa\x28\xf2\x52\x06\x49\x09\x82\x03\x19\xa0\x24\x6c\x40\x43\x2f\x1c\xaf\x94\x29\xcf\x11\xf2\xbf\xff\x43\xcf\x35\x98\x55\xa8\x6f\x28\xf2\xa1\x28\xca\xd2\x7f\x84\x10\xfb\x93\xf9\xfc\x00\xfc\xbd\x51\x18\x81\xa3\x6f\x44\x3c\x6f\x48\x96\xe5\xad\x96\xde\x53\x52\xab\x12\x08\x63\x19\x7d\x77\x6c\xfa\xf6\xa6\xd7\xd7\xf3\xec\xbe\x85\x4b\x17\x75\xa0\x3d\x24\x70\x09\xf5\x14\x41\x96\x25\x61\x63\xb7\xbf\x8e\xbc\x76\x57\xc7\xd2\xab\xd2\x0e\xe8\xef\xd0\x75\xd2\xee\xaa\xc0\xad\xb8\x52\xcf\xa0\xf6\x15\x52\x2f\xf7\xf5\x66\x34\x54\x45\x5b\x5e\x4b\x7f\xca\x61\x7b\x29\x3d\x8f\xc4\xf6\x0d\x14\x02\x63\xb5\x06\xe8\x1f\x7f\xb1\xfc\x30\x7e\x5c\x3a\x07\xa6\xa4\xc5\x28\xc2\x45\x98\xf4\xad\x34\x23\x8f\x85\xd4\x7a\x89\xe0\x1b\x1d\x26\x32\x7b\x4e\x2e\xf1\x38\x32\x78\xb8\x3a\x22\xd1\x11\xa9\xb2\xb8\xe9\xd1\x08\xe3\xbe\x79\xdc\xa8\x70\x9c\x0a\x2d\x98\xc0\x1d\xf6\xff\x0f\x50\xc9\x46\xc7\xa7\x73\xaa\x4c\x7d\xa4\x4c\x3a\xd4\x24\x66\x94\xfc\x49\x16\xb0\x27\x31\x3a\x09\xeb\x8f\x9d\x35\x1e\xc4\xe0\x9a\x14\x99\xdc\x84\xa5\xc1\xae\x56\x1a\xc4\xbe\x08\x40\xb4\xc8\x76\x53\xc4\x09\x4d\x47\x95\x22\x84\x06\x4d\xd2\x47\xdf\x10\xa9\x01\xc3\xd2\x37\x45\x01\xde\x93\x6c\xb4\x4b\x69\x56\x80\xa4\x3b\x9a\xce\xdf\x20\x8c\x19\xd9\x19\x60\xc7\xba\xbc\xe6\xcf\x52\x85\x4f\x40\x77\x50\x55\x50\x84\x8c\x54\xb2\x80\x47\x6b\xd7\x24\x32\xcb\xe5\x93\x7c\xa1\xbb\xb0\x75\x90\x01\x97\x21\x20\x25\x1b\x08\xb5\x2d\x09\x4b\x1b\xd4\x7b\xe7\x30\x8c\x84\xa5\x91\xaa\x0c\xb8\x07\x54\x52\xab\x1f\x71\xab\x14\x76\xe3\x34\x04\xc8\x4e\xca\xec\x89\xf3\xb5\x7d\xa6\x2c\xdd\x57\x31\x8d\x47\xda\xd7\x7f\x96\x30\x30\x57\x87\x8d\xa6\xe4\xff\x27\xf9\x92\xc4\x6d\x01\x3e\x24\x95\x54\x3a\xce\xf6\xc8\x7a\x9c\xd1\x2e\x1d\xbb\x3c\x60\xe0\xf0\x16\x9d\x20\x64\x78\x47\x35\xc8\x12\xd0\x0f\xee\xd1\x10\x77\x0f\x79\x65\x91\x46\xd7\x1a\xb6\x89\x32\xc9\x21\x70\x3a\x68\xa5\x3e\x39\xb8\x5b\xc3\xf6\x81\x4d\xc6\xf8\x1f\xb7\xe8\x83\x58\xc3\x96\x07\xfb\xcd\x39\xc0\x5b\xe9\x81\xb2\x05\xc9\x12\xb2\x68\xa5\xce\x9d\x98\x82\x9f\xac\x32\x94\xdc\x1b\xc2\x16\xe4\xde\xdc\x1b\xd2\xc5\xb5\x1c\x39\xba\x71\x0b\xf1\xf1\xeb\x97\xcf\xdc\x07\x54\x66\xa5\xaa\x2d\xed\x4d\x27\xd1\x03\x9d\xa2\x58\x6a\x1a\xad\xd3\x7f\x59\x86\xbc\xe7\x6e\x3e\xa7\x6e\x21\x46\xe3\x30\xb1\x03\x7d\x6e\x4f\xd5\x14\xd0\x75\x2c\x7e\xf9\xcf\x00\x00\x00\xff\xff\x8a\x1d\x3b\x9a\xc6\x06\x00\x00")

func assetsJsManagerMinJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsManagerMinJs,
		"assets/js/manager.min.js",
	)
}

func assetsJsManagerMinJs() (*asset, error) {
	bytes, err := assetsJsManagerMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/manager.min.js", size: 1734, mode: os.FileMode(420), modTime: time.Unix(1530705018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xcb\x6b\x1b\x57\x14\xc6\x3f\xc9\xee\x03\xb5\xdd\x18\x0a\x5d\x74\x71\x5c\xa8\x69\x29\xe3\xce\x48\x35\x75\xc7\x1e\xbb\xd4\x0f\x5a\x6a\x61\x63\xd4\xee\xaf\xa5\xab\xd1\x90\xd1\x1d\x71\x67\x64\x6c\xf0\x42\x96\x43\xb2\x88\xf0\x22\x24\xbb\x64\x91\x40\xf6\x22\x46\x20\xbf\xe4\x7f\xe1\xdc\x7d\x92\x65\xc8\x3e\x90\x75\x18\x59\xb2\xc9\x5d\x9c\xf9\x7e\x67\x7e\xe7\x70\xde\x4c\x4d\x3e\x06\x80\x2f\x00\x7c\x0f\xa0\x04\xe0\x2b\x00\x1d\x5c\xbf\x3e\x80\x1f\x00\x9c\x8e\xfa\xaf\x00\x7c\x03\xe0\x1d\x80\x69\x00\x1f\x46\xfc\x5d\x06\x98\x02\x30\x93\x01\x6a\x19\xe0\x8f\x0c\xe0\x02\x78\x9e\x05\xbe\x04\xf0\x3a\x0b\x7c\x0d\xe0\x6d\x16\x58\x01\xf0\x7e\xc4\xbf\x4f\x00\xdf\x02\x58\x9d\x00\x32\xb8\x7d\x9f\x8f\xbe\x13\xa3\xdb\xb2\x00\x26\x47\xbd\xcf\xd2\x52\x8e\x54\x35\xf0\x27\xd7\x94\xd8\x09\x65\x85\xb4\xf4\x83\x38\x91\x9a\xb4\xac\x86\xb2\x9c\x04\x91\x1a\x2b\x7f\x47\x71\x32\xce\x1b\x43\x69\x4c\x45\xb1\x47\xa1\xdc\x95\x21\x45\x55\x6a\x08\x1d\x07\xca\xa7\x64\xbf\x21\xe3\xb1\xb1\x99\xd4\xa4\xbe\xa1\xad\x48\x27\xa4\x9a\xf5\x1d\xa9\xb1\x2d\x1b\x91\x4e\xac\x62\xec\x07\x15\xeb\xaf\xa6\x1f\x5b\xa5\xc8\xa5\x8a\xdc\xfd\xf3\x4e\x50\x13\xf5\x68\x56\x37\x73\x5b\x9b\x25\x6b\x45\x4b\x91\x5e\x63\xad\x8a\x44\xba\x94\xb7\x9d\x79\xcb\x2e\x58\xf9\x02\xe5\x0b\xee\xdc\xdc\x2f\x76\xc1\xb6\x73\x1b\x22\x4e\xac\x92\x16\x2a\x0e\x45\x12\x69\x97\xfe\x1d\xee\xa0\x62\x53\x8b\x7a\x54\x89\x68\xf1\x93\xc5\x4b\xb9\x0d\xa1\xfc\xa6\xf0\xa5\x55\x92\xa2\xee\xd2\x0d\xbb\xb4\xdd\x8c\xe3\x40\xa8\x5c\xf1\x9f\xe2\x9a\xf5\xbf\xd4\x71\x10\x29\x97\x9c\x59\x3b\xb7\x12\xa9\x44\xaa\xc4\x2a\xed\x37\xa4\x4b\x89\xdc\x4b\x7e\x6d\x84\x22\x50\x0b\x54\xae\x09\x1d\xcb\xc4\xfb\xaf\xb4\x6e\xcd\xdf\x7a\xe9\x3d\x55\xa9\xad\x35\x55\x8e\x2a\x81\xf2\x5d\xca\x6d\x85\x4d\x2d\x42\x6b\x3d\xd2\xf5\xd8\x25\xd5\x18\x62\xec\x15\x16\xe8\x3a\x7a\xea\x47\xc7\xf6\x3c\x87\x66\x66\x28\x8d\xf6\xb4\xe7\x38\xb4\x4c\x36\xb9\x43\x5e\xf2\xf2\xe3\x5f\x8b\xde\x6f\x69\xfc\x69\xa8\x2d\x3a\x36\x1d\x1c\x5c\x8f\x2c\x79\x79\xfb\x67\x5a\x26\x87\x5c\xca\x2f\x80\x1f\xf2\x19\x9f\x9b\x63\x73\x9f\xfb\xa6\x6d\x3a\x64\x5a\xdc\xe3\x13\xee\x9b\x43\xd3\x36\x2d\xee\x9a\x7b\xdc\x37\xc7\xc3\xb6\xb9\xcb\xe7\xdc\xe3\x33\x73\xc8\x7d\xee\x83\x5f\xf0\x20\xb5\xc0\xcf\xb8\x67\x5a\xfc\x32\xad\xe0\x27\xdc\x1d\x29\x17\xdc\xe5\x73\xd3\xe1\x4b\xf3\x80\x4f\xc9\x1c\x99\x16\x0f\x52\x8b\x2f\x4d\x87\xf8\x8a\xbb\xa6\x35\xf4\x2e\xf9\x84\xbb\x64\xda\xdc\xe7\xab\xd4\x00\x3f\x32\x2d\x73\xc4\x27\x3c\xe0\x1e\xf8\x29\x0f\xf8\x22\xdd\x9d\xce\x0c\x4c\xcb\xb4\xb9\x8b\x8f\x01\x00\x00\xff\xff\x6e\xd3\x27\x73\x5a\x03\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 858, mode: os.FileMode(420), modTime: time.Unix(1530705024, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesGrpcMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xcd\x6e\xd4\x30\x14\x85\xcf\x54\x65\x93\x25\x6b\x16\x97\x05\x15\x08\x5c\xec\x84\x4a\x95\x27\x9e\x22\x86\x56\x42\x34\x22\x8c\x42\xf7\xd6\xc4\x64\x22\x12\x3b\xb2\x1d\x04\xd2\xbc\x06\xcf\xc3\x9b\xf0\x2c\x68\x92\xe1\xa7\x67\x75\x3e\xdd\x73\x8f\xae\xee\xaf\x87\xa7\x3f\x00\xe0\x04\xc0\x23\x00\x2f\x00\x3c\x00\x90\x63\x56\x09\xe0\x14\xc0\x47\x00\xbb\x05\x70\x77\xe4\x9f\x0b\x60\x71\xcc\x9c\xe0\x3f\x35\x9b\x72\x8d\x8d\x19\x9c\x8f\xac\x08\x4d\x5b\xb3\x37\x63\x13\x58\xe5\x24\xd5\xe6\xeb\xeb\x2f\xed\x4e\xf7\xee\xdc\x8f\x49\xf9\xa1\x62\x6b\x6f\x74\x6c\x9d\x65\x6f\x75\x34\x92\x52\x2e\x2e\x19\xcf\x58\x9a\x51\x9a\xc9\x8b\x8b\xe7\x3c\xe3\x3c\xb9\xd5\x21\xb2\xca\x6b\x1b\x3a\x1d\x9d\x97\xf4\x7e\xea\xa0\x62\xf4\xba\x77\xb5\xa3\xfc\x5e\xf1\x2a\xb9\xd5\xb6\x19\x75\x63\x58\x65\x74\x2f\xe9\x2f\x4b\xda\x8c\x21\xb4\xda\x26\xc5\xbb\xe2\x9a\xdd\x19\x1f\x5a\x67\x25\x89\x73\x9e\xac\x9d\x8d\xc6\x46\x56\x7d\x1f\x8c\xa4\x68\xbe\xc5\x97\x43\xa7\x5b\xbb\xa4\xed\x4e\xfb\x60\xa2\xfa\x54\xdd\xb0\xcb\x7f\xb9\xc3\x3d\x9f\x8d\x67\xd7\x76\xeb\xea\xd6\x36\x92\x92\xb2\x1b\xbd\xee\xd8\x8d\xf3\x7d\x90\x64\x87\x09\x83\xca\x96\x34\x5b\x65\x9f\x08\xae\x94\xa0\xb3\x33\x3a\x58\xfe\x58\x09\x41\x57\xc4\x49\x4e\xbc\x52\xe9\x9f\x51\xae\x5e\x1d\xec\xd3\x29\x96\x0b\x4e\xfb\xfd\xbc\xb2\x52\x29\x7f\x46\x57\x24\x48\x52\xba\x9c\xbf\xfd\x3b\x00\x00\xff\xff\x73\x83\x3e\x08\xc4\x01\x00\x00")

func localesRuLc_messagesGrpcMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesGrpcMo,
		"locales/ru/LC_MESSAGES/grpc.mo",
	)
}

func localesRuLc_messagesGrpcMo() (*asset, error) {
	bytes, err := localesRuLc_messagesGrpcMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/grpc.mo", size: 452, mode: os.FileMode(420), modTime: time.Unix(1530705024, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesManagerMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x4f\x6b\x13\x51\x14\xc5\x4f\xc7\xfa\x87\x11\x44\x5c\xbb\xb8\x2e\x2c\x8a\xbc\x3a\x33\xb1\x50\x5e\x3b\xa9\x1a\x1b\x10\x13\x2c\x65\xd4\xf5\x33\x79\x4d\x06\x27\xef\x85\x37\x33\xa2\xd2\x4d\xb3\xb5\x2b\xd1\x9d\x20\xf8\x09\x4a\x21\x98\x56\x48\xdd\xb8\x72\xf3\x66\xe3\x4e\x3f\x8b\x4c\x32\x5a\xbc\x9b\x7b\xce\x9d\xdf\x9c\x7b\xdf\xef\x2b\x8b\x1f\x00\xe0\x3c\x80\xab\x00\x22\x00\x17\x01\xec\x63\x5e\x13\x00\xe7\x00\x1c\x01\x58\x04\xf0\x0d\xc0\x59\x00\xdf\xab\xfe\xa3\x9a\x17\x15\xf7\x13\x40\x7f\x01\xf8\x05\xe0\x32\x80\xa6\x03\x5c\x02\xf0\xcc\x01\x2e\x00\xd8\x71\x00\x17\x80\x76\xe6\xfc\x9b\xca\xef\x39\xc0\x42\xb5\xd3\xa9\xfa\x22\x4e\xab\x64\xcf\x54\x77\x96\x7b\x71\xaf\x93\xc5\x5a\xa1\x21\x92\x04\xcd\x58\x26\x5d\xb4\xc4\x73\x99\x20\x7a\x3d\x94\x48\x33\x23\xc5\x00\xdb\x72\xa8\x4d\xc6\xda\x69\x2f\xee\xb2\xfb\x79\x2f\x65\x91\xe6\xd4\x95\x2f\xef\xbe\x88\xfb\x62\xa0\x97\x4d\xee\x6e\x3d\x8e\x58\xc3\x48\x51\xa6\xb1\x07\x22\x93\x9c\x02\xcf\x5f\x65\x5e\x8d\x05\x35\x0a\x6a\x7c\x65\xe5\x96\x57\xf3\x3c\xb7\x25\xd2\x8c\x45\x46\xa8\x34\x11\x99\x36\x9c\x1e\xcd\x32\xa8\x9d\x1b\x31\xd0\x5d\x4d\xeb\xff\x05\xd7\xdd\x96\x50\xbd\x5c\xf4\x24\x8b\xa4\x18\x70\xfa\xe7\x39\x6d\xe7\x69\x1a\x0b\xe5\xb6\x1f\xb6\x37\xd9\x53\x69\xd2\x58\x2b\x4e\xfe\xb2\xe7\x36\xb4\xca\xa4\xca\x58\xf9\x0a\x4e\x99\x7c\x95\xdd\x1e\x26\x22\x56\x6b\xd4\xe9\x0b\x93\xca\x2c\x7c\x12\x35\xd9\xea\x29\x57\xde\xb3\x23\x0d\xdb\x54\x1d\xdd\x8d\x55\x8f\x93\xbb\x95\xe4\x46\x24\xac\xa9\xcd\x20\xe5\xa4\x86\x33\x9b\x86\xb5\x35\x9a\xcb\x50\x5d\xf7\xbd\x30\xf4\x69\x69\x89\x4a\xe9\x5d\x0b\x7d\x9f\x36\xc8\x23\x3e\xf3\xf5\x30\xf8\xfb\x69\x3d\xbc\x53\xca\x1b\x33\x6c\xdd\xf7\x68\x77\x77\xfe\x4b\x3d\x0c\xbc\x9b\xb4\x41\x3e\x71\x0a\xd6\x60\xdf\xdb\xb1\x3d\x2a\xf6\x8a\x91\x3d\xb4\x13\x3b\x86\x7d\x57\xbc\xb5\x5f\xec\xa1\x3d\x28\x46\xc5\x3e\xec\x27\x3b\xb5\x5f\xcb\xf9\x47\x3b\x2e\x46\xf6\xd8\x1e\xc0\x7e\xb6\x13\x7b\x02\x7b\x62\xa7\xc5\xc8\x4e\xed\x31\xfe\x04\x00\x00\xff\xff\x12\x91\x5c\x79\x8c\x02\x00\x00")

func localesRuLc_messagesManagerMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesManagerMo,
		"locales/ru/LC_MESSAGES/manager.mo",
	)
}

func localesRuLc_messagesManagerMo() (*asset, error) {
	bytes, err := localesRuLc_messagesManagerMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/manager.mo", size: 652, mode: os.FileMode(420), modTime: time.Unix(1530705024, 0)}
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
	"templates/views/manager.html": templatesViewsManagerHtml,
	"assets/js/manager.min.js": assetsJsManagerMinJs,
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
	"locales/ru/LC_MESSAGES/grpc.mo": localesRuLc_messagesGrpcMo,
	"locales/ru/LC_MESSAGES/manager.mo": localesRuLc_messagesManagerMo,
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
			"manager.min.js": &bintree{assetsJsManagerMinJs, map[string]*bintree{}},
		}},
	}},
	"locales": &bintree{nil, map[string]*bintree{
		"ru": &bintree{nil, map[string]*bintree{
			"LC_MESSAGES": &bintree{nil, map[string]*bintree{
				"config.mo": &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
				"grpc.mo": &bintree{localesRuLc_messagesGrpcMo, map[string]*bintree{}},
				"manager.mo": &bintree{localesRuLc_messagesManagerMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"manager.html": &bintree{templatesViewsManagerHtml, map[string]*bintree{}},
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
