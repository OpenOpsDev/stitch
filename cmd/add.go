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
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/openopsdev/go-cli-commons/prompts"
	"github.com/roger-king/stitch/configs"
	"github.com/roger-king/stitch/services"
	"github.com/spf13/cobra"
)

var Defaults bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a service to your container orchestration config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Print("Please provide image name")
			os.Exit(1)
		}
		imageName := args[0]
		dockerAPI := services.NewDockerHubAPI()

		_, err := dockerAPI.FindImage(imageName)

		if err != nil {
			fmt.Print("Failed to find image name")
		}

		ps := prompts.Prompts{
			"serviceName": &prompts.UserInput{
				Label: "name of service",
			},
		}

		answers := ps.Run()
		serviceName := answers["serviceName"]

		// Create New instance of Docker-Compose
		dc, _ := configs.NewDockerCompose()
		// Check if service with that name exists
		// prompt if the user would like to override the configurations set
		if ok := dc.Services[serviceName]; ok != nil {
			confirm := prompts.Prompts{
				"ok": prompts.ConfirmInput{
					Label: fmt.Sprintf("Are you you want to override %s configurations?", serviceName),
				},
			}
			confirmAnswers := confirm.Run()

			if confirmAnswers["ok"] == "no" {
				os.Exit(1)
			}
		}

		// Find Kind of service (May want to offer if remote and has preset do you want to use our prest)
		service := configs.NewService(imageName, serviceName, Defaults)
		dc.Services[serviceName] = service.Service
		dc.CreateVolumes(service.Volumes)

		err = configs.Render("docker-compose.yml", dc)

		if err != nil {
			log.Print(err)
		} else {
			log.Print("Done.")
		}

		// ps := prompts.Prompts{
		// 	"serviceName": &prompts.UserInput{
		// 		Label: "name of service",
		// 	},
		// 	"restart": &prompts.SelectInput{
		// 		Label:   "restart",
		// 		Options: []string{"no", "always", "on-failure", "unless-stopped"},
		// 	},
		// 	"ports": &prompts.MultiInput{
		// 		Label:       "Ports to bind (e.g. 9000:9000)",
		// 		ValidString: `()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5]):()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])`,
		// 	},
		// 	"volumes": &prompts.MultiInput{
		// 		Label:       "Volumes to bind (e.g. db:/var/internal)",
		// 		ValidString: `.*:.*`,
		// 	},
		// 	"environments": &prompts.MultiInput{
		// 		Label:       "Environment variables to set (e.g. key=value)",
		// 		ValidString: `.*=.*`,
		// 	},
		// }

		// dc, _ := configs.NewDockerCompose()
		// answers := ps.Run()

		// ports := strings.Split(answers["ports"], ",")
		// volumes := strings.Split(answers["volumes"], ",")
		// environments := strings.Split(answers["environments"], ",")

		// // make volumes
		// if len(volumes) > 0 {
		// 	dc.Volumes = make(map[string]interface{}, len(volumes))
		// 	for _, v := range volumes {
		// 		remote := strings.Split(v, ":")[0]
		// 		dc.Volumes[remote] = make(map[interface{}]interface{})
		// 	}
		// }

		// // make environments
		// envmap := map[string]string{}
		// if len(environments) > 0 {
		// 	for _, e := range environments {
		// 		splitenvs := strings.Split(e, "=")
		// 		key := splitenvs[0]
		// 		value := splitenvs[1]
		// 		envmap[key] = value
		// 	}
		// }

	},
}

func init() {
	composeCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolVarP(&Defaults, "yes", "y", false, "add with defaults")
}
