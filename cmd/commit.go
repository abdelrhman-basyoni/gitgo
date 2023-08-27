/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdelrhman-basyoni/gitgo/core"
	"github.com/spf13/cobra"
)

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
		blob := &core.Blob{}
		if err := db.Store(blob); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to create new Database -  %v\n", err)
			os.Exit(1)
		}
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
}
