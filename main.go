package main

import (
	"brunch/brunch"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "brunch",
	Short: "brunch displays all recent branches for a given git repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		repositoryName, err := cmd.Flags().GetString("repo")
		if err != nil {
			return err
		}

		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}

		var displayObjects []brunch.DisplayObject
		displayObjects, err = brunch.Brunch(repositoryName, count)
		if err != nil {
			return err
		}

		brunch.Prompt(displayObjects)

		return nil
	},
}

func main() {
	rootCmd.Flags().StringP("repo", "r", ".", "repository to explore (default \".\")")
	rootCmd.Flags().IntP("count", "c", 8, "how many branches to display")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
