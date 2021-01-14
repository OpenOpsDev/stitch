package configs

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// Service - representation of docker-compose service options
type Service struct {
	Image string `yaml:"image,omitempty"`
	ContainerName string `yaml:"container_name,omitempty"`
	HostName string `yaml:"hostname,omitempty"`
	Restart string `yaml:"restart,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Volumes []string `yaml:"volumes,omitempty"`
	Ports []string `yaml:"ports,omitempty"`
}

// DockerCompose -
type DockerCompose struct {
	Version string `yaml:"version"`
	Services map[string]*Service `yaml:"services,omitempty"`
	Volumes map[string]string `yaml:"volumes,omitempty"`
	Networks map[string]string `yaml:"networks,omitempty"`
}

func NewDockerCompose() (*DockerCompose, bool) {
	var c DockerCompose
	cwd, err := os.Getwd()

	if err != nil {
		os.Exit(1)
	}

	existingConfigPath := path.Join(cwd, "docker-compose.yml")
	if _, err := os.Stat(existingConfigPath); os.IsNotExist(err) {
		return &DockerCompose{
			// TODO: get from config
			Version: "3",
			Services: map[string]*Service{},
		}, false
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