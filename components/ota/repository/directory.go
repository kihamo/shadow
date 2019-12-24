package repository

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/kihamo/shadow/components/ota"
	"github.com/kihamo/shadow/components/ota/release"
)

type Directory struct {
	*Memory

	lock     sync.RWMutex
	releases []ota.Release
}

func NewDirectory() *Directory {
	return &Directory{
		Memory: NewMemory(),
	}
}

func (r *Directory) Load(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.HasPrefix(info.Name(), "release-file-") {
			return nil
		}

		rl, err := release.NewLocalFile(path, "")
		if err == nil {
			r.Add(release.NewCompress(rl))
		}

		return err
	})
}

func (r *Directory) Remove(release ota.Release) (err error) {
	err = os.Remove(release.Path())

	if err == nil {
		err = r.Memory.Remove(release)
	}

	return err
}
