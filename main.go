package main

import (
	"brunch/internal"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var rootCmd = &cobra.Command{
	Use:   "internal",
	Short: "internal displays all recent branches for a given git repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}

		var displayObjects []internal.DisplayObject
		displayObjects, err = internal.Brunch(count)
		if err != nil {
			return err
		}

		selected, err := internal.Prompt(displayObjects)
		if err != nil {
			return err
		}

		err = exec.Command("git", "checkout", selected).Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func main() {
	rootCmd.Flags().IntP("count", "c", 8, "how many branches to display")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
