package presets

import (
	"fmt"

	"github.com/openopsdev/go-cli-commons/prompts"
)

type RabbitMQServiceSpec struct {
	Image            string
	Tag              string
	LocalPort        string
	ContainerName    string
	HostName         string
	WithManagementUI bool
}

func buildRabbitMQServicePrompts() prompts.Answers {
	ps := prompts.Prompts{
		"localPort": &prompts.UserInput{
			Label:   "local port to bind to (e.g. 5672)",
			Default: "5672",
		},
		"withManagementUI": &prompts.ConfirmInput{
			Label: "would you like complementary management UI?",
		},
	}

	return ps.Run()
}

func NewRabbitMQServiceSpec(serviceName, version string) (*RabbitMQServiceSpec, []string, []string) {
	// Run Prompts here
	answers := buildRabbitMQServicePrompts()
	networks := make([]string, 0)
	volumes := []string{answers["volumeName"]}
	return &RabbitMQServiceSpec{
		Tag:              version,
		LocalPort:        answers["localPort"],
		ContainerName:    serviceName,
		HostName:         serviceName,
		WithManagementUI: answers["withManagementUI"] == "yes",
	}, volumes, networks
}

func NewRabbitMQServiceSpecWithDefaults(serviceName, version string) (*RabbitMQServiceSpec, []string, []string) {
	networks := make([]string, 0)
	volumes := make([]string, 0)
	return &RabbitMQServiceSpec{
		Tag:              version,
		LocalPort:        "5672",
		ContainerName:    serviceName,
		HostName:         serviceName,
		WithManagementUI: true,
	}, volumes, networks
}

func (r *RabbitMQServiceSpec) Generate() *Service {
	// TODO: optional for redislabs images
	// TODO: Redis UI Management
	ports := []string{fmt.Sprintf("%s:5672", r.LocalPort)}

	if r.WithManagementUI {
		ports = append(ports, "15672:15672")
	}

	return &Service{
		Image:         fmt.Sprintf("rabbitmq:%s", r.Tag),
		Ports:         ports,
		Restart:       "always",
		ContainerName: r.ContainerName,
		HostName:      r.HostName,
	}
}
