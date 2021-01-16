package configs

import (
	"fmt"

	"github.com/openopsdev/go-cli-commons/prompts"
)

type RedisServiceSpec struct {
	Image         string
	Tag           string
	LocalPort     string
	ContainerName string
	HostName      string
	VolumeName    string
}

func buildRedisServicePrompts() prompts.Answers {
	ps := prompts.Prompts{
		"localPort": &prompts.UserInput{
			Label:   "local port to bind to (e.g. 6379)",
			Default: "6379",
		},
		"volumeName": &prompts.UserInput{
			Label:   "name of local volume to create (e.g. redis-db)",
			Default: "redis-db",
		},
	}

	return ps.Run()
}

func NewRedisServiceSpec(serviceName, version string) (*RedisServiceSpec, []string, []string) {
	// Run Prompts here
	answers := buildRedisServicePrompts()
	networks := make([]string, 0)
	volumes := []string{answers["volumeName"]}
	return &RedisServiceSpec{
		Tag:           version,
		LocalPort:     answers["localPort"],
		VolumeName:    answers["volumeName"],
		ContainerName: serviceName,
		HostName:      serviceName,
	}, volumes, networks
}

func NewRedisServiceSpecWithDefaults(serviceName, version string) (*RedisServiceSpec, []string, []string) {
	networks := make([]string, 0)
	volumes := []string{"db"}
	return &RedisServiceSpec{
		Tag:           version,
		LocalPort:     "6379",
		VolumeName:    "redisdb",
		ContainerName: serviceName,
		HostName:      serviceName,
	}, volumes, networks
}

func (r *RedisServiceSpec) Generate() *Service {
	// TODO: optional for redislabs images
	// TODO: Redis UI Management
	return &Service{
		Image:         fmt.Sprintf("redis:%s", r.Tag),
		Ports:         []string{fmt.Sprintf("%s:6379", r.LocalPort)},
		Restart:       "always",
		ContainerName: r.ContainerName,
		HostName:      r.HostName,
		Volumes:       []string{fmt.Sprintf("%s:/data", r.VolumeName)},
	}
}
