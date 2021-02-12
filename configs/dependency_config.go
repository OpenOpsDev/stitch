package configs

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/openopsdev/go-cli-commons/logger"
)

type LanguageCommands struct {
	Init    string `yaml:"init"`
	Install string `yaml:"install"`
}

type Dependency struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type DependencyConfig struct {
	Language     string            `yaml:"language"`
	Version      string            `yaml:"version"`
	Commands     LanguageCommands  `yaml:"commands"`
	Dependencies []Dependency 		`yaml:"dependencies"`
	Vars         map[string]string
}

func replaceVariable(s, newValue string) string {
	re := regexp.MustCompile(`{{([^}]*)}}`)
	str := re.ReplaceAll([]byte(s), []byte("github.com/roger-king/test"))
	return string(str)
}

func (d *DependencyConfig) Init() error {
	initCmd := replaceVariable(d.Commands.Init, d.Vars["initVar"])
	separateInitCmd := strings.Split(initCmd, " ")
	cmd := exec.Command(separateInitCmd[0])
	args := separateInitCmd[0:]
	cmd.Args = args
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info(out.String())
	return nil
}

func (d *DependencyConfig) Install(dep string) error {
	splitInstallCmd := strings.Split(d.Commands.Install, " ")
	cmd := exec.Command(splitInstallCmd[0])
	args := splitInstallCmd[0:]
	args = append(args, dep)
	cmd.Args = args
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info(out.String())
	return nil
}

func (d *DependencyConfig) InstallDependencies() error {
	logger.Info("Installing dependencies")
	for _, dep := range d.Dependencies {
		depString := dep.Name

		if len(dep.Version) > 0 {
			depString = fmt.Sprintf("%s@%s", depString, dep.Version)
		}

		err := d.Install(depString)

		if err != nil {
			return err
		}
	}

	return nil
}
