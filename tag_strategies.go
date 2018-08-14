package main

import (
	"github.com/Masterminds/semver"
	"sort"
)

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

func tagFetch(tags []string, strategy string) (s string, err error) {

	raw := tags
	var vs []*semver.Version
	for _, r := range raw {
		v, err := semver.NewVersion(r)
		if err == nil {
			vs = append(vs, v)
		}
	}
	sort.Sort(semver.Collection(vs))

	out := ""

	switch strategy {
	case "latest":
		x := getLast(vs)
		if x != nil {
			out = x.Original()
		}
	case "oldest":
		out = getFirst(vs).Original()
	}

	return out, nil
}
