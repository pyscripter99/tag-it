package test

import (
	"fmt"
	"tag-it/internal/tagger"
	"testing"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

func TestLoadLatestGitVersion(t *testing.T) {
	tests := []struct {
		initalTags []string
		want       *semver.Version
	}{
		{[]string{"v0.0.1", "0.1.5", "v2.0.1"}, semver.MustParse("v2.0.1")},
		{[]string{"v0.0.1-alpha", "v0.1.5-beta", "1.28.6"}, semver.MustParse("v1.28.6")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test semver parse: want=%s", tt.want)
		t.Run(testname, func(t *testing.T) {
			repo, err := git.Init(memory.NewStorage(), nil)
			if err != nil {
				t.Fatal(err)
			}

			for _, tag_name := range tt.initalTags {
				repo.CreateTag(tag_name, plumbing.NewHash(tag_name), nil)
			}

			latest, err := tagger.LoadGitLatestVersion(repo)
			if err != nil {
				t.Error(err)
			}

			if latest.String() != tt.want.String() {
				t.Errorf("want=%s, got=%s", tt.want, latest)
			}
		})
	}
}
