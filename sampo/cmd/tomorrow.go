/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TomorrowCmd describes and implements the `sampo tomorrow` command
var TomorrowCmd = &cobra.Command{
	Use:   "tomorrow",
	Short: "Show tomorrow's agenda",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tomorrow has not been implemented yet") // TODO
	},
}

func init() {
	RootCmd.AddCommand(TomorrowCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tomorrowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
