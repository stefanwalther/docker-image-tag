package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseBearer(t *testing.T) {

	input := []string{
		"Bearer realm=\"https://auth.docker.io/token\",service=\"registry.docker.io\",scope=\"repository:foo/bar:pull\"",
	}

	expected := map[string]string{
		"realm":   "https://auth.docker.io/token",
		"service": "registry.docker.io",
		"scope":   "repository:foo/bar:pull",
	}

	r := parseBearer(input)
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Parsing the bearer does not return the expected result.\nExpected: %v\nResult:%v", expected, r)
	}

}

func TestFixOfficialRepos(t *testing.T) {
	input := "foo"
	expected := "library/foo"
	r := fixOfficialRepos(input)
	if r != expected {
		t.Errorf("The expected result for official repos does not match.\nExpected: %v\nResult: %v", expected, r)
	}
}

func TestGetRepoUrl(t *testing.T) {

	repoUrl := "https://index.docker.io/v2"
	result := "https://index.docker.io/v2/foo/bar/tags/list/"

	tables :=
		[]struct {
			repositoryUrl string
			repo          string
			endpoint      string
			expected      string
		}{
			{repoUrl, "foo/bar", "tags/list", result},
			{repoUrl, "foo/bar", "/tags/list", result},
			{repoUrl, "foo/bar", "tags/list", result},
			{repoUrl, "foo/bar", "tags/list/", result},
			{repoUrl, "foo/bar/", "/tags/list", result},
			{repoUrl, "/foo/bar", "/tags/list", result},
			{repoUrl, "foo", "/tags/list", "https://index.docker.io/v2/library/foo/tags/list/"},
		}

	for _, table := range tables {
		rq := RepositoryRequest{
			RepositoryUrl: &repoUrl,
			Repo:          &table.repo,
			Endpoint:      table.endpoint,
		}
		r := getRepoUrl(&rq)
		if strings.Compare(r, table.expected) != 0 {
			t.Errorf("Resulting Url does not match expected result:\nInput:\t\t\t\t%v\nExpected result:\t%v\nResult:\t\t\t\t%v", table.repo, table.expected, r)
		}
	}
}
