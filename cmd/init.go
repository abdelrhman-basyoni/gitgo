/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// metadata directory
	gMetadataDir = ".gitgo"
	/**
	Chmod 0744 (chmod a+rwx,g-wx,o-wx,ug-s,-t) sets permissions so that, (U)ser / owner can read, can write and can execute. (G)roup can read, can't write and can't execute. (O)thers can read, can't write and can't execute.
	*/
	gDefaultPermissions = os.FileMode(0744)
	gMetadataDirContent = []string{
		"objects",
		"refs",
	}
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")

		// fetch the current working directory
		dir, err := os.Getwd()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		// Create the metadata directory  inside the current directory
		path := strings.Join([]string{dir, gMetadataDir}, string(os.PathSeparator))
		err = os.Mkdir(path, gDefaultPermissions)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		for _, content := range gMetadataDirContent {
			path := strings.Join([]string{dir, gMetadataDir, content}, string(os.PathSeparator))
			err = os.Mkdir(path, gDefaultPermissions)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

		}
		fmt.Fprint(os.Stdout, " Initialized gitgo repository in ", path, "\n")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
