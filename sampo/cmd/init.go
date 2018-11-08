/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"io"
	"net/http"
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
		configDirPath := fmt.Sprintf("%s/.sampo", os.Getenv("HOME"))
		schemaFilePath := fmt.Sprintf("%s/schema.sql", configDirPath)

		// Create the $HOME/.sampo directory, if it doesn't exist:
		if err := os.Mkdir(configDirPath, 0700); err != nil && !os.IsExist(err) {
			panic(err)
		}

		// Set the correct permissions on the $HOME/.sampo directory:
		if err := os.Chmod(configDirPath, 0700); err != nil {
			panic(err)
		}

		// Create the $HOME/.sampo/schema.sql file, overwriting any existing file:
		schemaFile, err := os.Create(schemaFilePath)
		if err != nil {
			panic(err)
		}
		defer schemaFile.Close()

		// Set the correct permissions on the $HOME/.sampo/schema.sql file:
		if err := schemaFile.Chmod(0600); err != nil {
			panic(err)
		}

		// Download the current schema.sql script from GitHub:
		response, err := http.Get("https://raw.githubusercontent.com/sampoapp/sampo/master/etc/schema.sql")
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		// Copy the HTTP response stream to $HOME/.sampo/schema.sql:
		if _, err := io.Copy(schemaFile, response.Body); err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
