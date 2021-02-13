/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package inits

import (
	"fmt"

	"github.com/openopsdev/go-cli-commons/logger"
	"github.com/openopsdev/stitch/configs"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Interactively create or update a .stitch/config.yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		// globalConfig, hasExisting := configs.NewGlobalStitchConfig()
		config, hasExisting := configs.NewStichConfig()
		if !hasExisting {
			config.Docker = &configs.DockerConfig{
				Compose: &configs.DockerComposeConfig{
					Version: "3",
				},
			}
		}
		err := configs.Render("./.stitch/config.yaml", config)

		if err != nil {
			logger.Fatal(fmt.Errorf("found an error rendering %v", err).Error())
		}
	},
}
