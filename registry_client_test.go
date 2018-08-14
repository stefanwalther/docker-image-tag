package main

import (
	"strings"
	"testing"
	"reflect"
)

func TestParseBearer(t *testing.T) {

	input := []string {
		"Bearer realm=\"https://auth.docker.io/token\",service=\"registry.docker.io\",scope=\"repository:foo/bar:pull\"",
	}

	expected := map[string]string {
		"realm": "https://auth.docker.io/token",
		"service": "registry.docker.io",
		"scope": "repository:foo/bar:pull",
	}

	r := parseBearer(input)
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Parsing the bearer does not return the expected result.\nExpected: %v\nResult:%v", expected, r)
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
		}{
			{repoUrl, "foo/bar", "tags/list"},
			{repoUrl, "foo/bar", "/tags/list"},
			{repoUrl, "foo/bar", "tags/list"},
			{repoUrl, "foo/bar", "tags/list/"},
			{repoUrl, "foo/bar/", "/tags/list"},
			{repoUrl, "/foo/bar", "/tags/list"},
		}

	for _, table := range tables {
		rq := RepositoryRequest{
			RepositoryUrl: &repoUrl,
			Repo:          &table.repo,
			Endpoint:      table.endpoint,
		}
		r := getRepoUrl(&rq)
		if strings.Compare(result, r) != 0 {
			t.Errorf("Resulting Url (%v) does not match expected result (%v)", r, result)
		}
	}

}
