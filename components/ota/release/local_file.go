package release

import (
	"crypto/md5"
	"io"
	"io/ioutil"
	"os"

	"github.com/kihamo/shadow/components/ota"
)

type LocalFileRelease struct {
	file         *os.File
	path         string
	version      string
	checksum     []byte
	architecture string
	fileInfo     os.FileInfo
}

func NewLocalFile(path, version string) (*LocalFileRelease, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return NewLocalFileFromFD(fd, version)
}

func NewLocalFileFromStream(stream io.Reader, version, dir string) (*LocalFileRelease, error) {
	if dir == "" {
		dir = os.TempDir()
	}

	fd, err := ioutil.TempFile(dir, "release-")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fd, stream)
	if err != nil {
		return nil, err
	}

	return NewLocalFileFromFD(fd, version)
}

func NewLocalFileFromFD(fd *os.File, version string) (*LocalFileRelease, error) {
	stat, err := fd.Stat()
	if err != nil {
		return nil, err
	}

	h := md5.New()
	if _, err := io.Copy(h, fd); err != nil {
		return nil, err
	}

	fd.Seek(0, 0)

	if version == "" {
		version = stat.Name()
	}

	return &LocalFileRelease{
		file:         fd,
		path:         fd.Name(),
		version:      version,
		checksum:     h.Sum(nil),
		architecture: ota.GoArch(fd),
		fileInfo:     stat,
	}, nil
}

func (f *LocalFileRelease) Version() string {
	return f.version
}

func (f *LocalFileRelease) BinFile() io.ReadCloser {
	f.file.Seek(0, io.SeekStart)
	return f.file
}

func (f *LocalFileRelease) Path() string {
	return f.path
}

func (f *LocalFileRelease) Checksum() []byte {
	return f.checksum
}

func (f *LocalFileRelease) Size() int64 {
	return f.fileInfo.Size()
}

func (f *LocalFileRelease) Architecture() string {
	return f.architecture
}

func (f *LocalFileRelease) FileInfo() os.FileInfo {
	return f.fileInfo
}
