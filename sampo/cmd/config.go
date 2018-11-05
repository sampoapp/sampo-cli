/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigCmd describes and implements the `sampo config` command
var ConfigCmd = &cobra.Command{
	Use:   "config [var]",
	Short: "Show configuration variables",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			for _, key := range viper.AllKeys() {
				fmt.Printf("%s=%s\n", key, viper.GetString(key))
			}
		} else {
			for _, key := range args {
				fmt.Printf("%s=%s\n", key, viper.GetString(key))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(ConfigCmd)
}
