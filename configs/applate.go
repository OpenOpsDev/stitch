package configs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/roger-king/stitch/templates"
	"github.com/roger-king/stitch/utils"
)

type Templater interface {
	Run()
}

type Applate struct {
	Source     string
	Answer     map[string]string
	Dependency DependencyConfig
	cacheDir   string
}

func cacheDir() string {
	homedir, _ := os.UserHomeDir()
	return path.Join(homedir, ".stitch/applates")
}

func NewApplate(source string) Applate {
	return Applate{
		Source: path.Join(cacheDir(), source),
	}
}

func (a Applate) Run() []error {
	var readErrors []error
	files, err := a.findTemplates()

	if err != nil {
		readErrors = append(readErrors, fmt.Errorf("cannot find tempalte: %v", err.Error()))
		return readErrors
	}

	for _, f := range files {
		file := utils.File{
			Source: f,
		}
		rawData, err := file.Read()

		if err != nil {
			readErrors = append(readErrors, fmt.Errorf("failed to read %s: %v", f, err.Error()))
			continue
		}

		template := templates.HandlebarTemplate{
			Raw:         rawData,
			Destination: a.buildDestinationPath(f),
		}

		err = template.Save(a.Answer)

		if err != nil {
			readErrors = append(readErrors, fmt.Errorf("failed to generate %s template: %v", f, err.Error()))
		}
	}

	return readErrors
}

func (a Applate) findTemplates() ([]string, error) {
	var files []string

	err := filepath.Walk(a.Source, func(path string, info os.FileInfo, err error) error {
		isConfig := strings.Contains(path, "config.yaml")
		isFileToAdd := !isConfig && !info.IsDir()
		if isFileToAdd {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

func (a Applate) buildDestinationPath(p string) string {
	s := strings.Split(p, a.Source)
	_, s = s[0], s[1:]
	updatedPath := strings.Join(s, "/")
	cwd, _ := os.Getwd()
	return path.Join(cwd, updatedPath)
}