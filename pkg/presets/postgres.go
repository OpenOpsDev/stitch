package presets

import (
	"fmt"

	"github.com/openopsdev/go-cli-commons/prompts"
)

type PostgresServiceSpec struct {
	Tag           string
	LocalPort     string
	User          string
	Password      string
	DB            string
	ContainerName string
	HostName      string
	VolumeName    string
}

var (
	commonPostgresPrompts = prompts.Prompts{
		"dbUser": &prompts.UserInput{
			Label: "name of the user to create inside of the DB",
		},
		"dbPassword": &prompts.UserInput{
			Label: "password to create for the user",
		},
		"dbName": &prompts.UserInput{
			Label: "name of the DB to connect to",
		},
	}
)

func buildPostgresServicePrompts() prompts.Answers {
	ps := prompts.Prompts{
		"localPort": &prompts.UserInput{
			Label:   "local port to bind to (e.g. 5432)",
			Default: "5432",
		},
		"volumeName": &prompts.UserInput{
			Label:   "name of local volume to create (e.g. db)",
			Default: "db",
		},
	}

	for k, v := range commonPostgresPrompts {
		ps[k] = v
	}

	return ps.Run()
}

func NewPostgresServiceSpec(serviceName, version string) (*PostgresServiceSpec, []string, []string) {
	// Run Prompts here
	answers := buildPostgresServicePrompts()
	networks := make([]string, 0)
	volumes := []string{answers["volumeName"]}
	return &PostgresServiceSpec{
		Tag:           version,
		LocalPort:     answers["localPort"],
		VolumeName:    answers["volumeName"],
		ContainerName: serviceName,
		HostName:      serviceName,
		User:          answers["dbUser"],
		Password:      answers["dbPassword"],
		DB:            answers["dbName"],
	}, volumes, networks
}

func NewPostgresServiceSpecWithDefaults(serviceName, version string) (*PostgresServiceSpec, []string, []string) {
	// Run Prompts here
	answers := commonPostgresPrompts.Run()
	networks := make([]string, 0)
	volumes := []string{"db"}
	return &PostgresServiceSpec{
		Tag:           version,
		LocalPort:     "5432",
		VolumeName:    "db",
		ContainerName: serviceName,
		HostName:      serviceName,
		User:          answers["dbUser"],
		Password:      answers["dbPassword"],
		DB:            answers["dbName"],
	}, volumes, networks
}

func (p *PostgresServiceSpec) Generate() *Service {
	return &Service{
		Image:         fmt.Sprintf("postgres:%s", p.Tag),
		Ports:         []string{fmt.Sprintf("%s:5432", p.LocalPort)},
		Restart:       "always",
		ContainerName: p.ContainerName,
		HostName:      p.HostName,
		Volumes:       []string{fmt.Sprintf("%s:/var/lib/postgresql/data", p.VolumeName)},
		Environment: map[string]string{
			"POSTGRES_USER":     p.User,
			"POSTGRES_PASSWORD": p.Password,
			"POSTGRES_DB":       p.DB,
		},
	}
}
