/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/sampoapp/sampo-cli/sampo/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type record = map[interface{}]interface{}

var ownerNick, ownerName, ownerUUID string

// BuildCmd describes and implements the `sampo build` command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build database snapshot",
	Long: `Sampo is a personal information manager (PIM) app.
This is the command-line interface (CLI) for Sampo.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Read the schema file:
		sqlSchema := readSchemaFile()

		// Validate all input arguments:
		for _, arg := range args {
			validateInputDirectory(arg)
		}

		// Process the input arguments:
		for _, arg := range args {
			createSnapshot(arg, sqlSchema)
		}
	},
}

func init() {
	RootCmd.AddCommand(BuildCmd)
	BuildCmd.Flags().StringVarP(&ownerName, "owner-name", "", "Owner", "Set owner name")
	BuildCmd.Flags().StringVarP(&ownerNick, "owner-nick", "", "owner", "Set owner nick")
	BuildCmd.Flags().StringVarP(&ownerUUID, "owner-uuid", "", "", "Set owner UUID (default: random)")
	viper.BindPFlag("owner.name", BuildCmd.Flags().Lookup("owner-name"))
	viper.BindPFlag("owner.nick", BuildCmd.Flags().Lookup("owner-nick"))
	viper.BindPFlag("owner.uuid", BuildCmd.Flags().Lookup("owner-uuid"))
}

func readSchemaFile() string {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/.sampo/schema.sql", os.Getenv("HOME")))
	if err != nil {
		panic(err)
	}
	return string(data)
}

func validateInputDirectory(arg string) {
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

func createSnapshot(inputDirPath string, sqlSchema string) {
	// Create the database file:
	db, err := store.Create(fmt.Sprintf("%s.db", inputDirPath))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Load the database schema:
	if err := db.Init(sqlSchema); err != nil {
		panic(err)
	}

	// Parse the given user UUID, if any:
	userUUID := uuid.Nil
	if viper.GetString("owner.uuid") != "" {
		userUUID, err = uuid.FromString(viper.GetString("owner.uuid"))
		if err != nil {
			panic(err)
		}
	}

	// Create the row for user ID #1:
	if _, err := db.CreateUser(&userUUID, viper.GetString("owner.nick"), viper.GetString("owner.name")); err != nil {
		panic(err)
	}

	// Find all relevant YAML input files:
	var yamlFiles []string
	err = filepath.Walk(inputDirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Base(path)[0] != '.' && filepath.Ext(path) == ".yaml" {
			yamlFiles = append(yamlFiles, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	var readerLatch, writerLatch sync.WaitGroup
	records := make(chan record, 256)

	// Spawn the writer process:
	writerLatch.Add(1)
	go writeRecords(db, records, &writerLatch)

	// Spawn reader processes for the set of YAML input files:
	for _, yamlFile := range yamlFiles {
		readerLatch.Add(1)
		go processInputFile(yamlFile, records, &readerLatch)
	}

	// Wait for the reader processes to finish:
	readerLatch.Wait()
	close(records)

	// Wait for the writer process to finish:
	writerLatch.Wait()

	// Attempt to reduce the size of the final database:
	if err := db.Compact(); err != nil {
		panic(err)
	}
}

func processInputFile(path string, output chan record, latch *sync.WaitGroup) {
	defer latch.Done()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	yaml := yaml.NewDecoder(bufio.NewReader(file))

	record := make(map[interface{}]interface{})
	for yaml.Decode(record) == nil {
		output <- record
	}
}

func writeRecords(db *store.Store, input chan record, latch *sync.WaitGroup) {
	defer latch.Done()

	for record := range input {
		fmt.Println("Processing record:", record) // TODO
	}
}
