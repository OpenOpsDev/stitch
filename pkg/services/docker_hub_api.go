package services

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/openopsdev/stitch/pkg/models"
)

// DockerHubAPI -
type DockerHubAPI struct {
	Rest *resty.Request
}

// NewDockerHubAPI -
func NewDockerHubAPI() *DockerHubAPI {
	client := resty.New()
	client.SetHostURL("https://hub.docker.com/api")
	r := client.R()
	return &DockerHubAPI{
		Rest: r,
	}
}

// FindImage - finds an image from dockerhub api
func (d *DockerHubAPI) FindImage(imageName string) (*models.DockerHubImage, error) {
	var image models.DockerHubImage
	res, err := d.Rest.Get(fmt.Sprintf("/content/v1/products/images/%s", imageName))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Body(), &image)

	if err != nil {
		return nil, err
	}

	return &image, nil
}
