package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/openopsdev/go-cli-commons/logger"
	"github.com/roger-king/stitch/templates"
	"github.com/roger-king/stitch/utils"
	"gopkg.in/yaml.v2"
)

type Templater interface {
	Run()
}

type Applate struct {
	Source     string
	Answer     map[string]string
	Dependency DependencyConfig
}

func findCacheDir() string {
	homedir, _ := os.UserHomeDir()
	return path.Join(homedir, ".stitch/applates")
}

var cacheDir = findCacheDir()

func NewApplate(source string) Applate {
	var dependencyConfig DependencyConfig
	applateDir := path.Join(cacheDir, source)
	applateConfig := path.Join(applateDir, ".stitch/applate.yaml")
	contents, err := ioutil.ReadFile(applateConfig)

	if err != nil {
		logger.Fatal(fmt.Errorf("failed to read existing config: %v", err).Error())
	}

	err = yaml.Unmarshal(contents, &dependencyConfig)

	if err != nil {
		logger.Fatal(fmt.Errorf("failed to build applate config: %v", err).Error())
	}

	err = dependencyConfig.Init()

	if err != nil {
		logger.Fatal(fmt.Errorf("failed to init project", err).Error())
	}

	return Applate{
		Source:     applateDir,
		Dependency: dependencyConfig,
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
		isConfig := strings.Contains(path, "applate.yaml")
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
