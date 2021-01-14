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

		image, err := dockerAPI.FindImage(imageName)

		if err != nil {
			fmt.Print("Failed to find image name")
		}

		ps := prompts.Prompts{
			"serviceName": &prompts.UserInput{
				Label: "name of service",
			},
		}
		fmt.Print(image.IsDatabase())
		if image.IsDatabase() {
			fmt.Print("is a db")
			ps["volume"] = &prompts.ConfirmInput{
				Label: "Persist data?",
			}
		}

		answers := ps.Run()

		dc, _ := configs.NewDockerCompose()

		dc.Services[answers["serviceName"]] = &configs.Service{
			Image: imageName,
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
