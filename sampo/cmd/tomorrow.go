/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tomorrowCmd represents the tomorrow command
var tomorrowCmd = &cobra.Command{
	Use:   "tomorrow",
	Short: "Show tomorrow's agenda",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tomorrow called") // TODO
	},
}

func init() {
	rootCmd.AddCommand(tomorrowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tomorrowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tomorrowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
