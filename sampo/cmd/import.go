/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ImportCmd describes and implements the `sampo import` command
var ImportCmd = &cobra.Command{
	Use:   "import [url]",
	Short: "Import data",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import has not been implemented yet") // TODO
	},
}

func init() {
	RootCmd.AddCommand(ImportCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
