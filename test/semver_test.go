package test

import (
	"fmt"
	"tag-it/internal/tagger"
	"testing"
)

func TestBump(t *testing.T) {
	var tests = []struct {
		ver  tagger.Semver
		bump tagger.SemverBump
		want tagger.Semver
	}{
		{tagger.Semver{Patch: 0, Minor: 0, Major: 1}, tagger.Patch, tagger.Semver{Patch: 1, Minor: 0, Major: 1}},
		{tagger.Semver{Patch: 0, Minor: 0, Major: 1}, tagger.Minor, tagger.Semver{Patch: 0, Minor: 1, Major: 1}},
		{tagger.Semver{Patch: 0, Minor: 0, Major: 1}, tagger.Major, tagger.Semver{Patch: 0, Minor: 0, Major: 2}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Bump: base=%s, want=%s", tt.ver, tt.want)
		t.Run(testname, func(t *testing.T) {
			result := tagger.Bump(tt.ver, tt.bump)
			if result != tt.want {
				t.Errorf("got=%s, want=%s", result, tt.want)
			}
		})
	}
}
