package main

import (
	"fmt"
	"github.com/Masterminds/semver"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"sort"
)

var (
	buildVersion    string
)

var (
	app = kingpin.New("docker-image-tag", "Docker registry V2 search tool to list and search for image tags.")

	registryUrl = app.Flag("registry", "Registry url").Short('r').Default("https://index.docker.io/v2").String() //.URL()
	username    = app.Flag("username", "Username").Short('u').Default(os.Getenv("DOCKER_USER")).String()
	password    = app.Flag("password", "Password").Short('p').Default(os.Getenv("DOCKER_PASS")).String()
	debug       = app.Flag("debug", "Debug mode").Bool()

	version = app.Command("version", "Get the version of docker-image-tag.")

	get      = app.Command("get", "Get a specific tag version, based on the strategy.").Default()
	getImage = get.Arg("image", "The Docker image to use.").Required().String()
	strategy = get.Flag("strategy", "Strategy to use, defaults to `latest`.").Default("latest").String()

	list          = app.Command("list", "List all tags")
	listImage     = list.Arg("image", "The Docker image to use.").Required().String()
	listSortOrder = list.Flag("order", "The sort order, `asc` or `desc`, defaults to `desc`.").Default("desc").String()
)

func main() {

	repository := RepositoryRequest{
		RepositoryUrl: registryUrl,
		User:          username,
		Password:      password,
		Endpoint:      "/tags/list",
	}

	app.Version(buildVersion)
	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		app.Fatalf("An error has occurred:\n\n- %v\n\nUse `docker-image-tag help` to get further usage information.", err)
	}

	switch cmd {
	case version.FullCommand():
		fmt.Printf("Version: %v\n", buildVersion)
	case get.FullCommand():
		repository.Repo = getImage
		tags := getTags(&repository)
		result, err := tagFetch(tags, *strategy)
		failIfErrorIsNotNil(err)
		fmt.Println(result)
	case list.FullCommand():
		repository.Repo = listImage
		tags := getTags(&repository)
		semverTags, _, _ := tags2SemVer(tags)
		if *listSortOrder == "desc" {
			sort.Sort(sort.Reverse(semver.Collection(semverTags)))
		}
		for _, t := range semverTags {
			fmt.Printf("%v\n", t)
		}
	}
}
