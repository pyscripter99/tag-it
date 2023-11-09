package tagger

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

	// Config
	VPrefix bool
}

func (s Semver) String() string {
	return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}

func ParseSemverString(verString string) (Semver, error) {
	var ver Semver

	reg, err := regexp.Compile(`^v?\d*\.\d*\.\d*$`)
	if err != nil {
		return ver, err
	}

	if !reg.MatchString(verString) {
		return ver, fmt.Errorf("%s not a valid/supported semver version", verString)
	}

	ver.VPrefix = strings.HasPrefix(verString, "v")

	verString = strings.TrimPrefix(verString, "v")

	segments := strings.Split(verString, ".")

	if len(segments) != 3 {
		return ver, errors.New("length of segments not equal to 3")
	}

	segment64, err := strconv.ParseUint(segments[0], 10, 32)
	if err != nil {
		return ver, err
	}
	ver.Major = uint32(segment64)

	segment64, err = strconv.ParseUint(segments[1], 10, 32)
	if err != nil {
		return ver, err
	}
	ver.Minor = uint32(segment64)

	segment64, err = strconv.ParseUint(segments[2], 10, 32)
	if err != nil {
		return ver, err
	}
	ver.Patch = uint32(segment64)

	return ver, nil
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
