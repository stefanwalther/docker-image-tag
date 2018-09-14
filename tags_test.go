package main

import (
	"testing"
	"reflect"
)

func TestTags2SemVer(t *testing.T) {

	tables := []struct {
		desc           string
		tagList        []string
		expectedErrors []string
	}{
		{
			"Null test",
			[]string{},
			[]string(nil),
		},
		{
			"One value",
			[]string{"1.0"},
			[]string(nil),
		},
		{
			"Simple two values",
			[]string{"1.0", "2.0"},
			[]string(nil),
		},
		{
			"One err value",
			[]string{"a"},
			[]string{"a"},
		},
		{
			"Two err values",
			[]string{"a", "b"},
			[]string{"a", "b"},
		},
		{
			"Three values, mixed",
			[]string{"1.0", "a", "2.0"},
			[]string{"a"},
		},
		{
			"Not that clear, but still semver compatible",
			[]string{"1.0-alpine", "2.0-alpine"},
			[]string(nil),
		},
	}

	for _, fixture := range tables {
		_, et, err := tags2SemVer(fixture.tagList)
		if err != nil {
			t.Errorf("An error has occurred: %v", err)
		}
		if !reflect.DeepEqual(et, fixture.expectedErrors) {
			t.Errorf("Test \"%v\"\nExpected error list does not match the returned one.\nExpected:\t%v\nCurrent:\t%v", fixture.desc, fixture.expectedErrors, et)
		}
	}
}
