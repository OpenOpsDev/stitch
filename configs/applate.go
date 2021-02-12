package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/openopsdev/go-cli-commons/logger"
	"github.com/openopsdev/stitch/services"
	"github.com/openopsdev/stitch/templates"
	"github.com/openopsdev/stitch/utils"
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

func NewApplate(source string, answers map[string]string) Applate {
	var promptConfig PromptConfig
	var dependencyConfig DependencyConfig
	applateDir := path.Join(cacheDir, source)

	if err := configExists(applateDir, ".stitch/applate.yaml"); os.IsNotExist(err) {
		err = services.GitClone(source, applateDir)

		if err != nil {
			logger.Fatal(fmt.Errorf("failed to pull %s applate: %v", source, err.Error()).Error())
		}
	}

	// TODO: handle no prompt config file
	applatePromptConfigPath := path.Join(applateDir, ".stitch/prompts.yaml")
	applatePromptContent, err := ioutil.ReadFile(applatePromptConfigPath)

	if err != nil {
		logger.Fatal(fmt.Errorf("failed to read prompt config: %v", err).Error())
	}

	err = yaml.Unmarshal(applatePromptContent, &promptConfig)

	if err != nil {
		logger.Fatal(fmt.Errorf("failed to build prompt config: %v", err).Error())
	}
	prompts := promptConfig.Build()
	promptAnswers := prompts.Run()

	// TODO: split into a Go Routine
	applateConfigFilePath := path.Join(applateDir, ".stitch/applate.yaml")
	applateConfigContents, err := ioutil.ReadFile(applateConfigFilePath)

	if err != nil {
		logger.Fatal(fmt.Errorf("failed to read existing config: %v", err).Error())
	}

	err = yaml.Unmarshal(applateConfigContents, &dependencyConfig)

	if err != nil {
		logger.Fatal(fmt.Errorf("failed to build applate config: %v", err).Error())
	}

	dependencyConfig.Vars = promptAnswers

	err = dependencyConfig.Init()

	if err != nil {
		logger.Fatal(fmt.Errorf("failed to init project: %v", err).Error())
	}

	return Applate{
		Source:     applateDir,
		Dependency: dependencyConfig,
		Answer:     promptAnswers,
	}
}

func (a Applate) Run() []error {
	var runErrors []error
	files, err := a.findTemplates()

	if err != nil {
		runErrors = append(runErrors, fmt.Errorf("cannot find tempalte: %v", err.Error()))
		return runErrors
	}

	for _, f := range files {
		file := utils.File{
			Source: f,
		}
		rawData, err := file.Read()

		if err != nil {
			runErrors = append(runErrors, fmt.Errorf("failed to read %s: %v", f, err.Error()))
			continue
		}

		template := templates.HandlebarTemplate{
			Raw:         rawData,
			Destination: a.buildDestinationPath(f),
		}

		err = template.Save(a.Answer)

		if err != nil {
			runErrors = append(runErrors, fmt.Errorf("failed to generate %s template: %v", f, err.Error()))
		}
	}

	if a.Dependency.Dependencies != nil {
		err = a.Dependency.InstallDependencies()

		if err != nil {
			runErrors = append(runErrors, fmt.Errorf("failed to install: %v", err.Error()))
		}
	}

	return runErrors
}

func (a Applate) findTemplates() ([]string, error) {
	var files []string

	err := filepath.Walk(a.Source, func(path string, info os.FileInfo, err error) error {
		isConfig := strings.Contains(path, "applate.yaml")
		gitDir := strings.Contains(path, ".git")
		isFileToAdd := !isConfig && !info.IsDir() && !gitDir
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
