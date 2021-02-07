package utils

import (
	"os"
	"path"
	"strings"
)

type Directory struct {
	Path string
}

func (d Directory) DoesExist() bool {
	if _, err := os.Stat(d.Path); os.IsNotExist(err) {
		return false
	}

	return true
}

func (d Directory) extractFileFromPath() {
	s := strings.Split(d.Path, "/")
	shouldBeFile := s[len(s)-1]
	if strings.Contains(shouldBeFile, ".") {
		_, s = s[0], s[1:]
		d.Path = strings.Join(s, "/")
	}
}

func (d Directory) Mkdir() error {
	if !d.DoesExist() {
		d.extractFileFromPath()
		err := os.Mkdir(path.Join(d.Path), os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}
