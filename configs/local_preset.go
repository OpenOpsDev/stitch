package configs

import (
	"strings"

	"github.com/openopsdev/go-cli-commons/constants"
	"github.com/openopsdev/go-cli-commons/prompts"
)

type LocalServiceSpec struct {
	Hostname     string
	Context      string
	Dockerfile   string
	Target       string
	Restart      string
	DependsOn    []string
	Links        []string
	Ports        []string
	Volumes      []string
	Environments []string
	Command      []string
}

func buildLocalServicePrompts() prompts.Answers {
	pre := prompts.Prompts{
		"doesLink": &prompts.SelectInput{
			Label:   "Does this service link to another?",
			Options: []string{"yes", "no"},
		},
		"doesDepend": &prompts.SelectInput{
			Label:   "Does this service depend on another?",
			Options: []string{"yes", "no"},
		},
	}

	preAnswers := pre.Run()

	ps := prompts.Prompts{
		"buildDockerfile": &prompts.UserInput{
			Label:   "Name of your Dockerfile (e.g. Dockerfile)",
			Default: "Dockerfile",
		},
		"buildTarget": &prompts.UserInput{
			Label:   "Multi-Stage build stage (leave blank if single build)",
			Default: "",
		},
		"restart": &prompts.SelectInput{
			Label:   "restart",
			Options: []string{"no", "always", "on-failure", "unless-stopped"},
		},
		"command": &prompts.UserInput{
			Label:   "Command to start application",
			Default: "",
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

	if preAnswers["doesLink"] == "yes" {
		ps["links"] = &prompts.MultiInput{
			Label: "Link to another service (e.g. db)",
		}
	}

	if preAnswers["doesDepend"] == "yes" {
		ps["dependsOn"] = &prompts.MultiInput{
			Label: "Depends on another service (e.g. api1)",
		}
	}

	return ps.Run()
}

func NewLocalServiceSpec(serviceName, buildContext string) (*LocalServiceSpec, []string, []string) {
	answers := buildLocalServicePrompts()
	networks := make([]string, 0)
	v := strings.Split(answers["volumes"], ",")
	volumes := make([]string, len(v))

	for i, vol := range v {
		key := strings.Split(vol, ":")[0]
		volumes[i] = key
	}

	return &LocalServiceSpec{
		Hostname:     serviceName,
		Context:      buildContext,
		Dockerfile:   answers["buildDockerfile"],
		Target:       answers["buildTarget"],
		Restart:      answers["restart"],
		Ports:        strings.Split(answers["ports"], ","),
		Volumes:      v,
		DependsOn:    strings.Split(answers["dependsOn"], ","),
		Links:        strings.Split(answers["links"], ","),
		Environments: strings.Split(answers["environments"], ","),
		Command:      strings.Split(answers["command"], " "),
	}, volumes, networks
}

func (l *LocalServiceSpec) Generate() *Service {
	envmap := map[string]string{}
	if len(l.Environments) > 0 {
		for _, e := range l.Environments {
			splitenvs := strings.Split(e, "=")
			key := splitenvs[0]
			value := splitenvs[1]
			envmap[key] = value
		}
	}

	serviceBuild := &ServiceBuild{
		Context:    l.Context,
		Dockerfile: l.Dockerfile,
	}

	if len(l.Target) > 0 {
		serviceBuild.Target = l.Target
	}

	return &Service{
		HostName:    l.Hostname,
		Restart:     l.Restart,
		Build:       serviceBuild,
		Command:     l.Command,
		Volumes:     l.Volumes,
		Ports:       l.Ports,
		Links:       l.Links,
		DependsOn:   l.DependsOn,
		Environment: envmap,
	}
}
