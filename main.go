package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
)

var (
	app = kingpin.New("docker-getImage-tag", "Docker registry V2 search tool to list and search for getImage tags.")

	registryUrl = app.Flag("registry", "Registy url").Short('r').Default("https://index.docker.io/v2").String() //.URL()
	username    = app.Flag("username", "Username").Short('u').Default(os.Getenv("DOCKER_USER")).String()
	password    = app.Flag("password", "Password").Short('p').Default(os.Getenv("DOCKER_PASS")).String()
	debug       = app.Flag("debug", "Debug mode").Bool()

	get      = app.Command("get", "Get a specific tag version, based on the strategy.").Default()
	getImage = get.Arg("image", "The Docker image to use.").Required().String()
	strategy = get.Flag("strategy", "Strategy to use, defaults to `latest`.").Default("latest").String()

	list		= app.Command("list", "List all tags")
	listImage = list.Arg("image", "The Docker image to use.").Required().String()
	listSortOrder = list.Flag("order", "The sort order, `asc` or `desc`, defaults to `asc`.").String()
)

func failIfErrorIsNotNil(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func main() {

	repository := RepositoryRequest{
		RepositoryUrl: registryUrl,
		User:          username,
		Password:      password,
		Endpoint:      "/tags/list",
	}

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		app.Fatalf("An error has occurred:\n\n- %v\n\nUse `docker-getImage-tag help` to get further usage information.", err)
	}
	switch cmd {
	case get.FullCommand():
		repository.Repo = getImage
		tags := getTags(&repository)
		result, err := tagFetch(tags, *strategy)
		failIfErrorIsNotNil(err)
		fmt.Println(result)
	case list.FullCommand():
		repository.Repo = listImage
		tags := getTags(&repository)
		//if *listSortOrder == "desc" {
		//	sort.Sort(sort.Reverse(tags))
		//}
		fmt.Printf(strings.Join(tags, "\n") + "\n")
	}
}
