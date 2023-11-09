package cmd

import (
	"fmt"
	"os"
	"slices"
	"tag-it/internal/tagger"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:   "bump <patch|minor|major> [version] [flags]",
	Short: "Bump the version by specified amount",
	Long:  "Bump the version by specified amount",
	Run: func(cmd *cobra.Command, args []string) {
		useGit, err := cmd.Flags().GetBool("git")
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

		if len(args) < 2 && !useGit {
			fmt.Println("Please specify version or use --git")
			os.Exit(1)
		}

		bumpFactors := []string{"patch", "minor", "major"}
		if !slices.Contains(bumpFactors, args[0]) {
			fmt.Println("Invalid bump factor, please use: patch, minor or major")
			os.Exit(1)
		}

		var ver *semver.Version

		if useGit {
			repositoryPath, err := cmd.Flags().GetString("repository")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			repo, err := git.PlainOpen(repositoryPath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			ver, err = tagger.LoadGitLatestVersion(repo)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if !useGit {
			ver, err = semver.NewVersion(args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		fmt.Println(ver)
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
