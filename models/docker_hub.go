package models

// DockerHubCategory -
type DockerHubCategory struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

// DockerHubPublisher -
type DockerHubPublisher struct {
	ID   string `json:"id"`
	Name string `json:"Name"`
}

// DockerHubImage - represents a response for finding an image from dockerhub
type DockerHubImage struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	Slug             string              `json:"slug"`
	Type             string              `json:"type"`
	Source           string              `json:"source"`
	Popularity       int                 `json:"popularity"`
	Categories       []DockerHubCategory `json:"categories"`
	Publisher        DockerHubPublisher  `json:"publisher"`
	ShortDescription string              `json:"short_description"`
	LongDescription  string              `json:"long_description"`
	CreatedAt        string              `json:"created_at"`
	UpdatedAt        string              `json:"updated_at"`
}

// IsDatabase -
func (d *DockerHubImage) IsDatabase() bool {
	for _, c := range d.Categories {
		if c.Name == "database" {
			return true
		}
	}
	return false
}
