package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/openopsdev/go-cli-commons/logger"
	"gopkg.in/yaml.v2"
)

type ServiceBuild struct {
	Context    string `yaml:"context,omitempty"`
	Dockerfile string `yaml:"dockerfile,omitempty"`
	Target     string `yaml:"target,omitempty"`
}

// Service - representation of docker-compose service options
type Service struct {
	Image         string            `yaml:"image,omitempty"`
	Build         *ServiceBuild     `yaml:"build,omitempty"`
	Command       []string          `yaml:"command,omitempty,flow"`
	ContainerName string            `yaml:"container_name,omitempty"`
	HostName      string            `yaml:"hostname,omitempty"`
	Restart       string            `yaml:"restart,omitempty"`
	Environment   map[string]string `yaml:"environment,omitempty"`
	Volumes       []string          `yaml:"volumes,omitempty"`
	Ports         []string          `yaml:"ports,omitempty"`
	Links         []string          `yaml:"links,omitempty"`
	DependsOn     []string          `yaml:"depends_on,omitempty"`
}

// DockerCompose -
type DockerCompose struct {
	Version  string                 `yaml:"version"`
	Services map[string]*Service    `yaml:"services,omitempty"`
	Volumes  map[string]interface{} `yaml:"volumes,omitempty"`
	Networks map[string]string      `yaml:"networks,omitempty"`
}

func NewDockerCompose() (*DockerCompose, bool) {
	var c DockerCompose
	cwd, err := os.Getwd()

	if err != nil {
		logger.Fatal(fmt.Errorf("Failed to get current working directory: %v", err).Error())
	}

	existingConfigPath := path.Join(cwd, "docker-compose.yml")
	if _, err := os.Stat(existingConfigPath); os.IsNotExist(err) {
		return &DockerCompose{
			// TODO: get from config
			Version:  "3",
			Services: map[string]*Service{},
		}, false
	}

	contents, err := ioutil.ReadFile(existingConfigPath)

	if err != nil {
		logger.Fatal(fmt.Errorf("Failed to read existing docker-compose file: %v", err).Error())
	}

	err = yaml.Unmarshal(contents, &c)

	if err != nil {
		logger.Fatal(fmt.Errorf("Failed to unmarshal existing docker-compose file: %v", err).Error())
	}

	return &c, true
}

func (d *DockerCompose) CreateVolumes(volumes []string) {
	if len(volumes) > 0 {
		if len(d.Volumes) == 0 {
			d.Volumes = make(map[string]interface{}, len(volumes))
		}

		for _, v := range volumes {
			d.Volumes[v] = make(map[interface{}]interface{})
		}
	}
}
