package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	pat := os.Getenv("GITHUB_PAT")
	if len(pat) == 0 {
		log.Println("GITHUB_PAT environment variable not set")
		log.Println("Repository permissions:")
		log.Println("\t- Read access to metadata")
		log.Println("\t- Read and Write access to administration")

		return
	}

	client := GitHubClient{
		PersonalAccessToken: pat,
	}

	repositories := client.GetGitHubRepositories()
	for i, repository := range repositories {
		login := repository.Owner.Login
		name := repository.Name

		log.Println(strconv.Itoa(i+1) + ": " + login + "/" + name + " ...")

		client.UpdateGitHubRepository(login, name)
	}

	(&http.Client{}).CloseIdleConnections()
}
