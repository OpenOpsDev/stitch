package configs

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// DockerComposeConfig -
type DockerComposeConfig struct {
	Version string `yaml:"version"`
}

// DockerConfig -
type DockerConfig struct {
	Registry string `yaml:"registry"`
	Compose *DockerComposeConfig `yaml:"compose"`
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
		os.Exit(1)
	}

	existingConfigPath := path.Join(cwd, ".stitch/config.yaml")
	if _, err := os.Stat(existingConfigPath); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(cwd, ".stitch"), os.ModePerm)

		if err != nil {
			os.Exit(1)
		}
		return &c, false
	}

	contents, err := ioutil.ReadFile(existingConfigPath)

	if err != nil {
		os.Exit(1)
	}

	err = yaml.Unmarshal(contents, &c)

	if err != nil {
		os.Exit(1)
	}

	return &c, true
}
