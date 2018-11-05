/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// HistoryCmd describes and implements the `sampo history` command
var HistoryCmd = &cobra.Command{
	Use:   "history [class]",
	Short: "Show history",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("history has not been implemented yet") // TODO
	},
}

func init() {
	RootCmd.AddCommand(HistoryCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// historyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
