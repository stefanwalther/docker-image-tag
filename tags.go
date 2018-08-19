package main

import "github.com/Masterminds/semver"

func getFirst(s []*semver.Version) *semver.Version {
	return s[0]
}

// Get the last non-empty version ...
func getLast(s []*semver.Version) *semver.Version {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != nil && s[i].Original() != "" {
			return s[i]
		}
	}
	return nil
}

// Parse a list of tags and return a slice of SemVer versions.
func tags2SemVer(tags []string) (semVerTags []*semver.Version, errTags []string, err error) {

	for _, t := range tags {
		v, err := semver.NewVersion(t)
		if err == nil {
			semVerTags = append(semVerTags, v)
		} else {
			errTags = append(errTags, t)
		}
	}
	return semVerTags, errTags, nil
}
