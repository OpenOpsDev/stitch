package configs

import "github.com/openopsdev/go-cli-commons/prompts"

type PromptConfig struct {
	Prompts []map[string]string `yaml:"prompts"`
}

func (p PromptConfig) Build() prompts.Prompts {
	pmpts := make(prompts.Prompts, len(p.Prompts))

	for _, v := range p.Prompts {
		for key, val := range v {
			if key == "key" {
				pmpts[val] = &prompts.UserInput{
					Label: "name of service",
				}
			}
		}
	}

	return pmpts
}
