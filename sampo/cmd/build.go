/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sampoapp/sampo-cli/sampo/store"
	"github.com/spf13/cobra"
)

// BuildCmd describes and implements the `sampo build` command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build database snapshot",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Read the schema file:
		sqlSchemaData, err := ioutil.ReadFile(fmt.Sprintf("%s/.sampo/schema.sql", os.Getenv("HOME")))
		if err != nil {
			panic(err)
		}
		sqlSchema := string(sqlSchemaData)

		// Validate all input arguments:
		for _, arg := range args {
			info, err := os.Stat(arg)
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Printf("%s does not exist\n", arg)
					os.Exit(66) // EX_NOINPUT
				}
				panic(err)
			}
			if !info.IsDir() {
				fmt.Printf("%s is not a directory\n", arg)
				os.Exit(66) // EX_NOINPUT
			}
		}

		// Process the input arguments:
		for _, arg := range args {
			db, err := store.Create(fmt.Sprintf("%s.db", arg))
			if err != nil {
				panic(err)
			}
			defer db.Close()

			if err := db.Init(sqlSchema); err != nil {
				panic(err)
			}

			if _, err := db.CreateUser("owner", "Owner"); err != nil {
				panic(err)
			}

			if err := db.Compact(); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(BuildCmd)
}
