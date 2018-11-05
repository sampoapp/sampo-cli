/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"

	"github.com/sampoapp/sampo-cli/sampo/schema"
	"github.com/sampoapp/sampo-cli/sampo/store"
	"github.com/spf13/cobra"
)

// ExportCmd describes and implements the `sampo export` command
var ExportCmd = &cobra.Command{
	Use:   "export [uuid]...",
	Short: "Export data",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := store.OpenDefault()
		if err != nil {
			panic(err)
		}
		defer db.Close()

		for _, arg := range args {
			entity, err := schema.LookupEntity(db, arg)
			if err != nil {
				panic(err)
			}

			fmt.Println(entity)
		}
	},
}

func init() {
	RootCmd.AddCommand(ExportCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
