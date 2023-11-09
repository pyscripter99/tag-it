package tagger

import "github.com/Masterminds/semver"

type SemverBump int

const (
	Patch SemverBump = 0
	Minor SemverBump = 1
	Major SemverBump = 2
)

func BumpVersion(ver semver.Version, bump SemverBump) semver.Version {
	bumped := ver

	switch bump {
	case Patch:
		bumped = ver.IncPatch()
	case Minor:
		bumped = ver.IncMinor()
	case Major:
		bumped = ver.IncMajor()
	}

	return bumped
}
