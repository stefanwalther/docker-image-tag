package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app = kingpin.New("docker-image-tag", "Docker registry V2 search tool to list and search for image tags.")

	registryUrl = app.Flag("registry", "Registy url").Short('r').Default("https://index.docker.io/v2").String() //.URL()
	username    = app.Flag("username", "Username").Short('u').Default(os.Getenv("DOCKER_USER")).String()
	password    = app.Flag("password", "Password").Short('p').Default(os.Getenv("DOCKER_PASS")).String()
	debug       = app.Flag("debug", "Debug mode").Bool()

	get      = app.Command("get", "Get a specific tag version, based on the strategy.").Default()
	image    = get.Arg("image", "The Docker image to use").Required().String()
	strategy = get.Flag("strategy", "Strategy to use, defaults to `latest`.").Default("latest").String()
)

//func handleResult(items []string, err error) {
//	failIfErrorIsNotNil(err)
//	for _, item := range items {
//		fmt.Printf("%s\n", item)
//	}
//}

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
		Repo:          image,
		Endpoint:      "/tags/list",
	}

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		app.Fatalf("An error has occurred:\n\n- %v\n\nUse `docker-image-tag help` to get further usage information.", err)
	}
	switch cmd {
	case get.FullCommand():
		tags := getTags(&repository)
		result, err := tagFetch(tags, *strategy)
		failIfErrorIsNotNil(err)
		fmt.Println(result)
	}
}
