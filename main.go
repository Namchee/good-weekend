package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

const (
	closed = "closed"
)

type Event struct {
	Action string "json:action"
	Number int    "json:number"
}

type Configuration struct {
	token    string
	location *time.Location
	message  string
}

type Repository struct {
	owner string
	name  string
}

func getConfiguration() *Configuration {
	token := os.Getenv("INPUT_ACCESS_TOKEN")

	if len(token) == 0 {
		log.Fatalln("GitHub access token is required")
	}

	timezoneLocation := os.Getenv("INPUT_TIMEZONE")
	location, err := time.LoadLocation(timezoneLocation)

	if err != nil {
		log.Fatalln("Ilegal timezone location name")
	}

	message := os.Getenv("INPUT_MESSAGE")

	return &Configuration{
		token,
		location,
		message,
	}
}

func getRepository(metadata string) *Repository {
	tokens := strings.Split(metadata, "/")

	if len(tokens) != 2 {
		return nil
	}

	return &Repository{
		owner: tokens[0],
		name:  tokens[1],
	}
}

func main() {
	config := getConfiguration()

	ctx := context.Background()

	token_source := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.token},
	)
	oauthClient := oauth2.NewClient(ctx, token_source)
	client := github.NewClient(oauthClient)

	repository := getRepository(os.Getenv("GITHUB_REPOSITORY"))

	if repository == nil {
		log.Fatalln("Failed to read repository metadata")
	}

	var event Event
	event_file, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))

	if err != nil {
		log.Fatalln("Failed to read event metadata")
	}

	if err = json.NewDecoder(event_file).Decode(&event); err != nil {
		log.Fatalln("Failed to parse event metadata")
	}

	pullRequest, _, err := client.PullRequests.Get(
		ctx,
		repository.owner,
		repository.name,
		event.Number,
	)

	if err != nil {
		log.Fatalln("Failed to fetch pull request information")
	}

	day := pullRequest.GetCreatedAt().In(config.location).Weekday()

	if day == 6 || day == 0 { // if PR is submitted on weekends, close it
		err = closePullRequest(
			&ctx,
			client,
			config.message,
			repository.name,
			repository.owner,
			pullRequest.GetNumber(),
		)

		if err != nil {
			log.Fatalln("Failed to close pull request due to unexpected errors")
		}
	}
}

func closePullRequest(
	ctx *context.Context,
	client *github.Client,
	reason string,
	repository string,
	owner string,
	number int,
) error {
	_, _, err := client.Issues.CreateComment(
		*ctx,
		owner,
		repository,
		number,
		&github.IssueComment{
			Body: github.String(reason),
		},
	)

	if err != nil {
		return err
	}

	_, _, err = client.Issues.AddLabelsToIssue(
		*ctx,
		owner,
		repository,
		number,
		[]string{"good-weekend"},
	)

	if err != nil {
		return err
	}

	_, _, err = client.PullRequests.Edit(
		*ctx,
		owner,
		repository,
		number,
		&github.PullRequest{
			State: github.String(closed),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
