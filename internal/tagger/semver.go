package tagger

import "fmt"

type SemverBump int

const (
	Patch SemverBump = 0
	Minor SemverBump = 1
	Major SemverBump = 2
)

type Semver struct {
	Patch uint32
	Minor uint32
	Major uint32
}

func (s Semver) String() string {
	return fmt.Sprintf("%d.%d.%d", s.Patch, s.Minor, s.Major)
}

func Bump(ver Semver, bump SemverBump) Semver {
	switch bump {
	case Patch:
		ver.Patch += 1
	case Minor:
		ver.Minor += 1
	case Major:
		ver.Major += 1
	}
	return ver
}
