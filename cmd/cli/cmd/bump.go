package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"tag-it/internal/tagger"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:   "bump <patch|minor|major> [version] [flags]",
	Short: "Bump the version by specified amount",
	Long:  "Bump the version by specified amount.\nIf authentication is required set: GIT_USERNAME and GIT_PASSWORD environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		useGit, err := cmd.Flags().GetBool("git")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		usePrefix, err := cmd.Flags().GetBool("prefix")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		usePush, err := cmd.Flags().GetBool("push")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		useCommit := usePush
		if !usePush {
			useCommit, err = cmd.Flags().GetBool("commit")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		tagMessage, err := cmd.Flags().GetString("tag-message")
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
		var repo *git.Repository

		if useGit || useCommit || usePush {
			repositoryPath, err := cmd.Flags().GetString("repository")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			repo, err = git.PlainOpen(repositoryPath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if useGit {
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

		var bump tagger.SemverBump

		switch args[0] {
		case "patch":
			bump = tagger.Patch
		case "minor":
			bump = tagger.Minor
		case "major":
			bump = tagger.Major
		default:
			fmt.Println("Invalid bump factor, please use: patch, minor or major")
			os.Exit(1)
		}

		absVer := *ver
		absVer = tagger.BumpVersion(absVer, bump)
		ver = &absVer

		var verString string
		if usePrefix {
			verString = strings.Join([]string{"v", ver.String()}, "")
		} else {
			verString = ver.String()
		}

		if useCommit {
			ref, err := repo.Head()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			_, err = repo.CreateTag(verString, ref.Hash(), &git.CreateTagOptions{
				Message: tagMessage,
			})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if usePush {
			var username, password string

			for _, env := range os.Environ() {
				envSplit := strings.Split(env, "=")
				if envSplit[0] == "GIT_USERNAME" {
					username = envSplit[1]
				}
				if envSplit[0] == "GIT_PASSWORD" {
					password = envSplit[1]
				}
			}

			var auth *http.BasicAuth

			if username != "" && password != "" {
				auth = &http.BasicAuth{
					Username: username,
					Password: password,
				}
			}

			if err := tagger.PushTags(repo, "origin", auth); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		fmt.Println(verString)
	},
}

func init() {
	rootCmd.AddCommand(bumpCmd)

	bumpCmd.Flags().BoolP("git", "g", false, "Specifies if tag-it should get the current version from the current git repository")
	bumpCmd.Flags().BoolP("push", "p", false, "Create git commit and push to specified remote")
	bumpCmd.Flags().BoolP("commit", "c", false, "Specifies if the new tag should be committed to the current git repository.")
	bumpCmd.Flags().String("tag-message", "Bump version to '$version$'", "String to use in commit message.")
	bumpCmd.Flags().StringP("repository", "r", ".", "Folder containing .git/")
	bumpCmd.Flags().Bool("prefix", false, "Adds 'v' to beginning of version")
}
