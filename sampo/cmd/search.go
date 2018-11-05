/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// SearchCmd describes and implements the `sampo search` command
var SearchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search data",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search has not been implemented yet") // TODO
		os.Exit(1)
	},
}

func init() {
	RootCmd.AddCommand(SearchCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
