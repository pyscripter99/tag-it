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

func TestParse(t *testing.T) {
	var tests = []struct {
		str  string
		want tagger.Semver
	}{
		{"1.0.0", tagger.Semver{Major: 1, Minor: 0, Patch: 0}},
		{"1.1.0", tagger.Semver{Major: 1, Minor: 1, Patch: 0}},
		{"1.0.1", tagger.Semver{Major: 1, Minor: 0, Patch: 1}},
		{"1.9.3", tagger.Semver{Major: 1, Minor: 9, Patch: 3}},
		{"999.999.999", tagger.Semver{Major: 999, Minor: 999, Patch: 999}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Parse: str=%s, want=%s", tt.str, tt.want)
		t.Run(testname, func(t *testing.T) {
			result, err := tagger.ParseSemverString(tt.str)
			if err != nil {
				t.Fatal(err)
			}

			if result != tt.want {
				t.Errorf("got=%s, want=%s", result.String(), tt.want.String())
			}
		})
	}
}
