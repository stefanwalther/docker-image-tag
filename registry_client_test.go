package main

import (
	"strings"
	"testing"
)

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
