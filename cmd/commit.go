/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/abdelrhman-basyoni/gitgo/core"
	"github.com/spf13/cobra"
)

var (
	message string
)

type ConfigVars struct {
	G_AUTHOR_NAME  string
	G_AUTHOR_EMAIL string
}

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("commit called")
		config := ConfigVars{
			G_AUTHOR_NAME:  os.Getenv("G_AUTHOR_NAME"),
			G_AUTHOR_EMAIL: os.Getenv("G_AUTHOR_EMAIL"),
		}

		author := core.Author{}
		if err := author.New(config.G_AUTHOR_NAME, config.G_AUTHOR_EMAIL, time.Now()); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed while creating author- %v\n", err)
			os.Exit(1)
		}
		rootDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		gitgodir := strings.Join([]string{rootDir, core.GMetadataDir}, string(os.PathSeparator))
		dbDir := strings.Join([]string{gitgodir, "objects"}, string(os.PathSeparator))
		db := core.Database{}

		if err := db.New(dbDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to create new Database -  %v\n", err)
			os.Exit(1)
		}

		//get the list of all the files in the root directory
		files, err := os.ReadDir(rootDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to read current directory -  %v\n", err)
			os.Exit(1)
		}
		var entries []core.Entry

		for _, file := range files {
			if file.Name() == "." || file.Name() == ".." || file.IsDir() {
				continue
			}
			file, err := os.Open(file.Name())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to open file -  %v\n", err)
				os.Exit(1)
			}
			defer file.Close()

			var content bytes.Buffer

			_, err = io.Copy(&content, file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to cop file -  %v\n", err)
				os.Exit(1)
			}

			blob := core.Blob{}
			if err := blob.New(content.Bytes()); err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to create new Blob   -  %v\n", err)
				os.Exit(1)
			}

			// store the blob in the database
			if err := db.Store(&blob); err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to store the   Blob   -  %v\n", err)
				os.Exit(1)
			}
			// create a new Entry
			entry := core.Entry{}
			if err := entry.New(blob.GetOid(), file.Name()); err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to create a new entry -  %v\n", err)
				os.Exit(1)
			}
			entries = append(entries, entry)
		}
		//create a new tree
		tree := core.Tree{}
		if err := tree.New(entries); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to create a new tree -  %v\n", err)
			os.Exit(1)
		}
		//store the tree
		if err := db.Store(&tree); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to store the  tree -  %v\n", err)
			os.Exit(1)
		}
		//create new commit
		commit := core.Commit{}
		if err := commit.New(nil, tree.GetOid(), "Initial commit", author, author); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to create a new commit - %v\n", err)
			os.Exit(1)

		}

		//store the commit
		if err := db.Store(&commit); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to store  commit - %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("[(root_commit)%x] %s\n", commit.GetOid(), commit.GetMessage())
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	commitCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message")
}
