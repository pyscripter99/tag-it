package tagger

import (
	"errors"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func LoadGitLatestVersion(repo *git.Repository) (*semver.Version, error) {
	tagrefs, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	tags := []*semver.Version{}

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		ver, err := semver.NewVersion(t.Name().Short())
		if err != nil {
			return err
		}

		tags = append(tags, ver)

		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Sort(semver.Collection(tags))
	if len(tags) == 0 {
		return nil, errors.New("no tags found")
	}

	return tags[len(tags)-1], err
}
