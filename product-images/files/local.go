package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Local struct {
	basePath    string
	maxFileSize int
}

func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{basePath: p, maxFileSize: maxSize}, nil
}

func (l *Local) Save(p string, contents io.Reader) error {
	fp := l.fullPath(p)
	d := filepath.Dir(fp)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		fmt.Println("Unable to create directory")
		return err
	}

	_, err = os.Stat(fp)
	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			fmt.Println("Unable to delete file")
			return err
		}
	} else if !os.IsNotExist(err) {
		fmt.Println("Unable to get file info")
		return err
	}

	f, err := os.Create(fp)
	if err != nil {
		fmt.Println("Unable to create file")
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, contents)
	if err != nil {
		fmt.Println("Unable to write file")
		return err
	}

	return nil
}

func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
