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
	"github.com/openopsdev/go-cli-commons/logger"
	"github.com/openopsdev/stitch/pkg/configs"
	"github.com/spf13/cobra"
)

// initAppCmd represents the initApp command
var initAppCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var answers map[string]string
		var a string

		if len(args) == 0 {
			logger.Fatal("missing applate")
		} else if len(args) > 1 {
			logger.Fatal("too many arguments")
		}

		a = args[0]
		appplate := configs.NewApplate(a, answers)
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
	appCmd.AddCommand(initAppCmd)
}
