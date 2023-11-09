package cmd

import (
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:   "bump <patch|minor|major> [version] [flags]",
	Short: "Bump the version by specified amount",
	Long:  "Bump the version by specified amount",
	Run: func(cmd *cobra.Command, args []string) {
		git, err := cmd.Flags().GetBool("git")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(args) > 2 {
			fmt.Println("Too many arguments.")
			os.Exit(1)
		}
		if len(args) < 1 {
			fmt.Println("Please specify bump factor")
			os.Exit(1)
		}

		if len(args) < 2 && !git {
			fmt.Println("Please specify version or use --git")
			os.Exit(1)
		}

		bumpFactors := []string{"patch", "minor", "major"}
		if !slices.Contains(bumpFactors, args[0]) {
			fmt.Println("Invalid bump factor, please use: patch, minor or major")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(bumpCmd)

	bumpCmd.Flags().BoolP("git", "g", false, "Specifies if tag-it should get the current version from the current git repository")
	bumpCmd.Flags().StringP("push", "p", "", "Create git commit and push to specified remote")
	bumpCmd.Flags().BoolP("commit", "c", false, "Specifies if the new tag should be committed to the current git repository.")
	bumpCmd.Flags().String("commit-message", "Bump version to '$version$'", "String to use in commit message.")
	bumpCmd.Flags().StringP("repository", "r", ".", "Folder containing .git/")
}
