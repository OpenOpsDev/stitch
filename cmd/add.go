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
	"os"
	"strings"

	"github.com/openopsdev/go-cli-commons/prompts"
	"github.com/roger-king/stitch/configs"
	"github.com/roger-king/stitch/services"
	"github.com/spf13/cobra"
)

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
			"restart": &prompts.SelectInput{
				Label:   "restart",
				Options: []string{"no", "always", "on-failure", "unless-stopped"},
			},
			"ports": &prompts.MultiInput{
				Label:       "Ports to bind (e.g. 9000:9000)",
				ValidString: `()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5]):()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])`,
			},
			"volumes": &prompts.MultiInput{
				Label: "Volumes to bind (e.g. db:/var/internal)",
			},
		}

		dc, _ := configs.NewDockerCompose()
		answers := ps.Run()

		dc.Services[answers["serviceName"]] = &configs.Service{
			Image:         imageName,
			ContainerName: answers["serviceName"],
			HostName:      answers["serviceName"],
			Restart:       answers["restart"],
			Volumes:       strings.Split(answers["volumes"], ","),
			Ports:         strings.Split(answers["ports"], ","),
		}

		configs.Render("docker-compose.yml", dc)
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
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
