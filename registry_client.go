package main

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"strings"
)

// This function parses the Www-Authenticate header provided in the challenge
// It has the following format
// Bearer realm="https://gitlab.com/jwt/auth",service="container_registry",scope="repository:andrew18/container-test:pull"
func parseBearer(bearer []string) map[string]string {
	out := make(map[string]string)
	for _, b := range bearer {
		for _, s := range strings.Split(b, " ") {
			if s == "Bearer" {
				continue
			}
			for _, params := range strings.Split(s, ",") {
				fields := strings.Split(params, "=")
				key := fields[0]
				val := strings.Replace(fields[1], "\"", "", -1)
				out[key] = val
			}
		}
	}
	return out
}

type TagResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type RepositoryRequest struct {
	RepositoryUrl *string
	User          *string
	Password      *string
	Repo          *string
	Endpoint      string
}

func fixSuffixPrefix(s string) string {

	sep := "/"

	s = strings.TrimPrefix(s, sep)
	if !strings.HasSuffix(s, sep) {
		return s + sep
	}
	return s
}

func getRepoUrl(repository *RepositoryRequest) string {
	var out string
	out += fixSuffixPrefix(*repository.RepositoryUrl)
	out += fixSuffixPrefix(*repository.Repo)
	out += fixSuffixPrefix(repository.Endpoint)
	return out
}

func getTags(repository *RepositoryRequest) []string {
	request := gorequest.New()

	url := getRepoUrl(repository)

	// First step is to tagFetch the endpoint where we'll be authenticating
	resp, _, _ := request.Get(url).End()

	// This has the various things we'll need to parse and use in the request
	params := parseBearer(resp.Header["Www-Authenticate"])
	paramsJSON, _ := json.Marshal(&params)

	// Get the token
	challenge := gorequest.New()
	resp, body, _ := challenge.Get(params["realm"]).
		SetBasicAuth(*username, *password).
		Query(string(paramsJSON)).
		End()

	token := make(map[string]string)
	json.Unmarshal([]byte(body), &token)

	// Now reissue the challenge with the token in the Header
	// curl -IL https://index.docker.io/v2/odewahn/image/tags/list
	authenticatedRequest := gorequest.New()

	resp, body, _ = authenticatedRequest.Get(url).
		Set("Authorization", "Bearer "+token["token"]).
		End()

	var tagResponse TagResponse
	json.Unmarshal([]byte(body), &tagResponse)

	return tagResponse.Tags
}
