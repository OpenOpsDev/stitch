package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/openopsdev/go-cli-commons/logger"
	"gopkg.in/yaml.v2"
)

// DockerComposeConfig -
type DockerComposeConfig struct {
	Version string `yaml:"version"`
}

// DockerConfig -
type DockerConfig struct {
	Registry string               `yaml:"registry"`
	Compose  *DockerComposeConfig `yaml:"compose"`
}

// StitchConfig -
type StitchConfig struct {
	Docker *DockerConfig `yaml:"docker"`
}

// NewStichConfig -
func NewStichConfig() (*StitchConfig, bool) {
	var c StitchConfig
	cwd, err := os.Getwd()

	if err != nil {
		logger.Fatal(fmt.Errorf("Failed to get current working directory: %v", err).Error())
	}

	existingConfigPath := path.Join(cwd, ".stitch/config.yaml")
	if _, err := os.Stat(existingConfigPath); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(cwd, ".stitch"), os.ModePerm)

		if err != nil {
			logger.Fatal(fmt.Errorf("Failed to get current working directory: %v", err).Error())
		}
		return &c, false
	}

	contents, err := ioutil.ReadFile(existingConfigPath)

	if err != nil {
		logger.Fatal(fmt.Errorf("Failed to read existing config: %v", err).Error())
	}

	err = yaml.Unmarshal(contents, &c)

	if err != nil {
		logger.Fatal(fmt.Errorf("Failed to unmarshal config %v", err).Error())
	}

	return &c, true
}
