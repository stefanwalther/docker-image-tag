package main

import (
	"github.com/Masterminds/semver"
	"sort"
)

// Fetch a single tag based on the given strategy.
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
