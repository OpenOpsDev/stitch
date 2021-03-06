package presets

import (
	"strings"
)

type ServiceSpec interface {
	Generate() *Service
}

type NewServiceOutputs struct {
	Service  *Service
	Volumes  []string
	Networks []string
}

func NewService(image, serviceName string, withDefaults bool) *NewServiceOutputs {
	var service *Service
	var spec ServiceSpec
	var volumes []string
	var networks []string
	version := "latest"
	splitImage := strings.Split(image, ":")
	name := splitImage[0]

	if len(splitImage) == 2 {
		version = splitImage[1]
	}

	if image[0] == '.' {
		spec, volumes, networks = NewLocalServiceSpec(serviceName, image)
	} else {
		switch name {
		case "postgres":
			if withDefaults {
				spec, volumes, networks = NewPostgresServiceSpecWithDefaults(serviceName, version)
			} else {
				spec, volumes, networks = NewPostgresServiceSpec(serviceName, version)
			}
		case "redis":
			if withDefaults {
				spec, volumes, networks = NewRedisServiceSpecWithDefaults(serviceName, version)
			} else {
				spec, volumes, networks = NewRedisServiceSpec(serviceName, version)
			}
		case "rabbitmq":
			if withDefaults {
				spec, volumes, networks = NewRabbitMQServiceSpecWithDefaults(serviceName, version)
			} else {
				spec, volumes, networks = NewRabbitMQServiceSpec(serviceName, version)
			}
		default:
			spec, volumes, networks = NewRemoteServiceSpec(serviceName, image)
		}
	}

	service = spec.Generate()
	return &NewServiceOutputs{
		Service:  service,
		Volumes:  volumes,
		Networks: networks,
	}
}
