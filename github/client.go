package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	HttpAccept = "application/vnd.github+json"
	ApiVersion = "2022-11-28"
)

type GitHubClient struct {
	PersonalAccessToken string
}

func (c *GitHubClient) GetGitHubRepositories() []GitHubRepository {
	httpClient := &http.Client{}

	payload, err := json.Marshal(struct{}{})
	if err != nil {
		panic(err)
	}

	url := "https://api.github.com/user/repos?per_page=100"
	req, err := c.createGitHubRequest("GET", url, payload)
	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
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
		panic(string(b))
	}

	var repositories []GitHubRepository
	if err := json.Unmarshal(b, &repositories); err != nil {
		panic(err)
	}

	return repositories
}

func (c *GitHubClient) UpdateGitHubRepository(owner, repo string) {
	httpClient := &http.Client{}

	payload, err := json.Marshal(NewGitHubRepositorySettings())
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := c.createGitHubRequest("PATCH", url, payload)
	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
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
		panic(string(b))
	}
}

func (c *GitHubClient) createGitHubRequest(method string, url string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", HttpAccept)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.PersonalAccessToken))
	req.Header.Add("X-GitHub-Api-Version", ApiVersion)

	return req, err
}
