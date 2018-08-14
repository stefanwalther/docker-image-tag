package main

import (
	"strings"
	"testing"
)

func TestTagFetch(t *testing.T) {

	tables := []struct {
		desc     string
		tags     []string
		strategy string
		result   string
	}{
		{
			"Just a basic test",
			[]string{"1.0"},
			"latest",
			"1.0",
		},
		{
			"A basic test with two versions, using strategy 'latest'",
			[]string{"1.0", "2.0"},
			"latest",
			"2.0",
		},
		{
			"A basic test with two versions, but reversed order, using strategy 'latest'",
			[]string{"2.0", "1.0"},
			"latest",
			"2.0",
		},
		{
			"Using a semver-invalid version",
			[]string{"1.0", "2.0", "renovate-xx"},
			"latest",
			"2.0",
		},
		{
			"A basic test with two versions, using strategy 'oldest'",
			[]string{"1.0", "2.0"},
			"oldest",
			"1.0",
		},
		{
			"Dealing with no results",
			[]string{},
			"latest",
			"",
		},
	}

	for _, versionSet := range tables {
		r, err := tagFetch(versionSet.tags, versionSet.strategy)
		if err != nil {
			t.Errorf("An error has occorred: %v\n", err)
		}
		if strings.Compare(r, versionSet.result) != 0 {
			t.Errorf("Error in test `%v`:\n\tThe result (%v) does not match the expected result (%v), using strategy '%v'", versionSet.desc, r, versionSet.result, versionSet.strategy)
		}
	}
}
