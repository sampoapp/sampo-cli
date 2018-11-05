/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// TodayCmd describes and implements the `sampo today` command
var TodayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today's agenda",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("today has not been implemented yet") // TODO
		os.Exit(1)
	},
}

func init() {
	RootCmd.AddCommand(TodayCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
