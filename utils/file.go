package utils

import (
	"io/ioutil"
	"os"
	"path"
)

type File struct {
	Name   string
	Source string
	Data   string
	Dest   string
}

func (f File) Read() (string, error) {
	var data string

	b, err := ioutil.ReadFile(f.Source)

	if err != nil {
		return data, err
	} else {
		data = string(b)
	}

	return data, nil
}

func (f File) Write() error {
	directory := Directory{
		Path: f.Dest,
	}
	err := directory.Mkdir()

	if err != nil {
		return err
	}

	file, err := os.Create(path.Join(f.Dest, f.Name))

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(f.Data)

	if err != nil {
		return err
	}

	return nil
}
