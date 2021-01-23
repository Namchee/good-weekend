package main

import (
	"os"
	"strconv"
)

const (
	closed = "closed"
)

type Configuration struct {
	token   string
	offset  int8
	message string
}

func getConfiguration() Configuration {
	token := os.Getenv("INPUT_ACCESS_TOKEN")
	offset, _ := strconv.Atoi(os.Getenv("INPUT_TIMEZONE_OFFSET"))
	message := os.Getenv("INPUT_MESSAGE")

	return Configuration{
		token,
		int8(offset),
		message,
	}
}

func main() {
	config := getConfiguration()
}
