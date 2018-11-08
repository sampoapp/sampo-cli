/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// InitCmd describes and implements the `sampo init` command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init has not been implemented yet") // TODO
		os.Exit(1)
		// TODO: Create the $HOME/.sampo directory.
		// TODO: Download the schema.sql script to $HOME/.sampo/schema.sql.
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
