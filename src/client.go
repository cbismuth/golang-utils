package main

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cbismuth/golang-github-utils/src/domain"
)

type GitHubClient struct {
	PersonalAccessToken string
}

func (gitHubClient GitHubClient) GetGitHubRepositories() []domain.GitHubRepository {
	client := &http.Client{}

	payload, err := json.Marshal(struct{}{})
	if err != nil {
		panic(err)
	}

	url := "https://api.github.com/user/repos?per_page=100"
	req, err := gitHubClient.createGitHubRequest("GET", url, payload)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 400) {
		panic(string(bytes))
	}

	var repositories []domain.GitHubRepository
	if err := json.Unmarshal(bytes, &repositories); err != nil {
		panic(err)
	}

	return repositories
}

func (gitHubClient GitHubClient) UpdateGitHubRepository(owner, repo string) {
	client := &http.Client{}

	payload, err := json.Marshal(domain.NewDefaultGitHubRepositorySettings())
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := gitHubClient.createGitHubRequest("PATCH", url, payload)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 400) {
		panic(string(bytes))
	}
}

func (gitHubClient GitHubClient) createGitHubRequest(method string, url string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes2.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", gitHubClient.PersonalAccessToken))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	return req, err
}
