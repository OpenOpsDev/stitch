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
package app

import (
	"github.com/openopsdev/go-cli-commons/logger"
	"github.com/openopsdev/stitch/pkg/configs"
	"github.com/openopsdev/stitch/pkg/utils"
	"github.com/spf13/cobra"
)

// initAppCmd represents the initApp command
var initAppCmd = &cobra.Command{
	Use:   "init",
	Short: "streamlines the creation of templates (applates)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := utils.CheckRequiredNumOfArgs(args, 1); err != nil {
			logger.Error(err.Error())
		}

		var answers map[string]string

		source := args[0]
		appplate := configs.NewApplate(source, answers)
		applateErrors := appplate.Run()

		if len(applateErrors) > 0 {
			for _, e := range applateErrors {
				logger.Error(e.Error())
			}
			return
		}
	},
}

func init() {
	Cmd.AddCommand(initAppCmd)
}
