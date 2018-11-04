/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"

	"github.com/sampoapp/sampo-cli/sampo/schema"
	"github.com/sampoapp/sampo-cli/sampo/store"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [class]",
	Short: "List data",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := store.Open("./app.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		rows, err := schema.QueryEntities(db)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			entity, err := schema.ScanEntity(rows)
			if err != nil {
				panic(err)
			}
			fmt.Println(entity.UUID)
		}

		err = rows.Err()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
