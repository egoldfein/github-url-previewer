package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/InVisionApp/github-url-previewer/config"
	"github.com/InVisionApp/github-url-previewer/services"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/go-playground/webhooks.v4"
	ghub "gopkg.in/go-playground/webhooks.v4/github"
	"mvdan.cc/xurls"
)

const (
	path = "/payload"
	port = 3016
)

var client *github.Client
var ctx = context.Background()

func main() {

	// loads envorinment.env config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failure loading ENV vars: %v", err)
	}

	hook := ghub.New(&ghub.Config{Secret: cfg.GithubToken})
	hook.RegisterEvents(HandlePullRequest, ghub.PullRequestEvent)

	err = webhooks.Run(hook, ":"+strconv.Itoa(port), path)
	if err != nil {
		fmt.Println(err)
	}
}

// HandlePullRequest handles GitHub pull_request events
func HandlePullRequest(payload interface{}, header webhooks.Header) {

	// loads envorinment.env config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failure loading ENV vars: %v", err)
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.GithubPersonalAccessToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	pl := payload.(ghub.PullRequestPayload)

	// ensures comments are only added when a PR is opened
	if pl.Action == "opened" {
		urls := xurls.Relaxed().FindAllString(pl.PullRequest.Body, -1)

		for i := range urls {

			str := services.GetPreview(cfg.LinkPreviewAccessKey, urls[i])
			newPRComment := &github.IssueComment{
				Body: &str,
			}

			_, _, err = client.Issues.CreateComment(ctx, pl.PullRequest.User.Login, pl.PullRequest.Head.Repo.Name, int(pl.PullRequest.Number), newPRComment)
			if err != nil {
				log.Fatalf("Error creating PR Comment: %v", err)
			}
		}
	}
}
