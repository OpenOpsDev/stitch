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
	"path"

	"github.com/openopsdev/go-cli-commons/logger"
	"github.com/openopsdev/stitch/cmd/app"
	"github.com/openopsdev/stitch/cmd/compose"
	configSub "github.com/openopsdev/stitch/cmd/config"
	"github.com/openopsdev/stitch/cmd/inits"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var stitch = &cobra.Command{
	Use:   "stitch",
	Short: "A tool to help streamline the development process!",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := stitch.Execute(); err != nil {
		logger.Fatal(fmt.Errorf("Failed to start CLI: %v", err).Error())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	stitch.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $ROOT_APP/.stitch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	stitch.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setupSubCommands()
}

func setupSubCommands() {
	stitch.AddCommand(app.Cmd)
	stitch.AddCommand(compose.Cmd)
	stitch.AddCommand(inits.Cmd)
	stitch.AddCommand(configSub.Cmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find current working directory
		pwd, err := os.Getwd()
		if err != nil {
			logger.Fatal(fmt.Errorf("Failed to get current working directory: %v", err).Error())
		}

		// Search config in home directory with name ".stitch" (without extension).
		viper.AddConfigPath(path.Join(pwd, ".stitch"))
		viper.SetConfigName("config.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
