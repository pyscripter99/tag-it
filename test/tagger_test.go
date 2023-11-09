package test

import (
	"fmt"
	"tag-it/internal/tagger"
	"testing"

	"github.com/Masterminds/semver"
)

func TestBumpVersion(t *testing.T) {
	tests := []struct {
		base *semver.Version
		bump tagger.SemverBump
		want *semver.Version
	}{
		{semver.MustParse("1.0.0"), tagger.Patch, semver.MustParse("1.0.1")},
		{semver.MustParse("1.0.0"), tagger.Minor, semver.MustParse("1.1.0")},
		{semver.MustParse("1.0.0"), tagger.Major, semver.MustParse("2.0.0")},
		{semver.MustParse("v12.6.2-alpha"), tagger.Minor, semver.MustParse("12.7.0")},
		{semver.MustParse("5.0.6"), tagger.Major, semver.MustParse("6.0.0")},
		{semver.MustParse("8.9.3"), tagger.Patch, semver.MustParse("8.9.4")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Bump: %s to %s", tt.base, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := tagger.BumpVersion(*tt.base, tt.bump)
			if got.String() != tt.want.String() {
				t.Errorf("got=%s, want=%s", got.String(), tt.want.String())
			}
		})
	}
}
