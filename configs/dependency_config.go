package configs

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"

	"github.com/openopsdev/go-cli-commons/logger"
)

type LanguageCommands struct {
	Init    string `yaml:"init"`
	Install string `yaml:"install"`
}

type DependencyConfig struct {
	Language     string            `yaml:"language"`
	Version      string            `yaml:"version"`
	Commands     LanguageCommands  `yaml:"commands"`
	Dependencies map[string]string `yaml:"dependencies"`
	Vars         map[string]string
}

func replaceVariable(s string) string {
	re := regexp.MustCompile(`{{([^}]*)}}`)
	str := re.ReplaceAll([]byte(s), []byte("github.com/roger-king/test"))
	return string(str)
}

func (d *DependencyConfig) Init() error {
	initCmd := replaceVariable(d.Commands.Init)
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

func (d *DependencyConfig) Install() (string, error) {
	cmd := exec.Command(d.Commands.Install)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
