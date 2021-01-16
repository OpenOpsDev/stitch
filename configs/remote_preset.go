package configs

import (
	"strings"

	"github.com/openopsdev/go-cli-commons/constants"
	"github.com/openopsdev/go-cli-commons/prompts"
)

type RemoteServiceSpec struct {
	Image        string
	ServiceName  string
	Restart      string
	Ports        []string
	Volumes      []string
	Environments []string
}

func buildRemoteSpecPrompts() prompts.Answers {
	ps := prompts.Prompts{
		"restart": &prompts.SelectInput{
			Label:   "restart",
			Options: []string{"no", "always", "on-failure", "unless-stopped"},
		},
		"ports": &prompts.MultiInput{
			Label:       "Ports to bind (e.g. 9000:9000)",
			ValidString: constants.PortRegex,
		},
		"volumes": &prompts.MultiInput{
			Label:       "Volumes to bind (e.g. db:/var/internal)",
			ValidString: constants.DockerVolumeRegex,
		},
		"environments": &prompts.MultiInput{
			Label:       "Environment variables to set (e.g. key=value)",
			ValidString: constants.KeyValueRegex,
		},
	}
	return ps.Run()
}

func NewRemoteServiceSpec(serviceName, imageName string) (*RemoteServiceSpec, []string, []string) {
	answers := buildRemoteSpecPrompts()
	networks := make([]string, 0)
	v := strings.Split(answers["volumes"], ",")
	volumes := make([]string, len(v))

	for i, vol := range v {
		key := strings.Split(vol, ":")[0]
		volumes[i] = key
	}

	return &RemoteServiceSpec{
		Image:        imageName,
		ServiceName:  serviceName,
		Restart:      answers["restart"],
		Ports:        strings.Split(answers["ports"], ","),
		Volumes:      v,
		Environments: strings.Split(answers["environments"], ","),
	}, volumes, networks
}

func (r *RemoteServiceSpec) Generate() *Service {
	envmap := map[string]string{}
	if len(r.Environments) > 0 {
		for _, e := range r.Environments {
			splitenvs := strings.Split(e, "=")
			key := splitenvs[0]
			value := splitenvs[1]
			envmap[key] = value
		}
	}

	return &Service{
		Image:         r.Image,
		Restart:       r.Restart,
		ContainerName: r.ServiceName,
		HostName:      r.ServiceName,
		Ports:         r.Ports,
		Volumes:       r.Volumes,
		Environment:   envmap,
	}
}
