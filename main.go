package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

const (
	closed = "closed"
)

type Configuration struct {
	token   string
	offset  int8
	message string
}

type Repository struct {
	author string
	name   string
}

func getConfiguration() Configuration {
	token := os.Getenv("INPUT_ACCESS_TOKEN")

	if len(token) == 0 {
		log.Fatalln("GitHub access token is required")
	}

	offset, _ := strconv.Atoi(os.Getenv("INPUT_TIMEZONE_OFFSET"))
	message := os.Getenv("INPUT_MESSAGE")

	return Configuration{
		token,
		int8(offset),
		message,
	}
}

func getRepository(metadata string) *Repository {
	tokens := strings.Split(metadata, "/")

	if len(tokens) != 2 {
		return nil
	}

	return &Repository{
		author: tokens[0],
		name:   tokens[1],
	}
}

func main() {
	config := getConfiguration()

	ctx := context.Background()

	token_source := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.token},
	)
	oauth_client := oauth2.NewClient(ctx, token_source)
	client := github.NewClient(oauth_client)

	repository := getRepository(os.Getenv("GITHUB_REPOSITORY"))

	if repository == nil {
		log.Fatalln("Failed to read repository metadata")
	}
}
