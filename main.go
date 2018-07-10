package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/InVisionApp/github-url-previewer/config"
	"gopkg.in/go-playground/webhooks.v4"
	"gopkg.in/go-playground/webhooks.v4/github"
	"mvdan.cc/xurls"
)

const (
	path = "/payload"
	port = 3016
)

func main() {
	// loads envorinment.env config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failure loading ENV vars: %v", err)
	}

	fmt.Println(cfg.GithubToken)
	hook := github.New(&github.Config{Secret: cfg.GithubToken})
	hook.RegisterEvents(HandlePullRequest, github.PullRequestEvent)

	err = webhooks.Run(hook, ":"+strconv.Itoa(port), path)
	if err != nil {
		fmt.Println(err)
	}
}

// HandlePullRequest handles GitHub pull_request events
func HandlePullRequest(payload interface{}, header webhooks.Header) {

	fmt.Println("Handling Pull Request")

	pl := payload.(github.PullRequestPayload)

	urls := xurls.Relaxed().FindAllString(pl.PullRequest.Body, -1)
	for i := range urls {
		fmt.Printf("%s", urls[i])
		fmt.Println("")
	}
}