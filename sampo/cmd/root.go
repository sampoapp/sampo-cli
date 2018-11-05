/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var debug bool
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "sampo",
	Short: "Sampo command-line interface (CLI)",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Version: "0.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "C", "", "Set config file (default: $HOME/.sampo/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debugging")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Be verbose")
	rootCmd.SetVersionTemplate(`Sampo CLI {{printf "%s" .Version}}
`)
}

// initConfig reads the configuration file and environment variables.
func initConfig() {
	if configFile != "" {
		// Use the specified config file:
		viper.SetConfigFile(configFile)
	} else {
		// Search for config file in the current directory and under the home directory:
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.sampo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in:
	if err := viper.ReadInConfig(); err == nil {
		if debug {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
