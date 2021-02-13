package configs

import (
	"github.com/openopsdev/go-cli-commons/prompts"
)

type PromptMapping struct {
	Key string `yaml:"key"`
	Type string `yaml:"type"`
	Title string `yaml:"title"`
}

type PromptConfig struct {
	Prompts []PromptMapping `yaml:"prompts"`
}

func (p PromptConfig) Build() prompts.Prompts {

	pmpts := make(prompts.Prompts, len(p.Prompts))

	for _, v := range p.Prompts {
		pmpts[v.Key] = &prompts.UserInput{
			Label: v.Title,
		}
	}

	return pmpts
}
