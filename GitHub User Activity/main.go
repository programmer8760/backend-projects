package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GitHubActivity struct {
	Type      string `json:"type"`
	Repo      Repo   `json:"repo"`
	CreatedAt string `json:"created_at"`
	Payload   struct {
		Action  string `json:"action"`
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: github-activity [username]")
		return
	}
	username := os.Args[1]

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))

	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode == 404 {
		fmt.Printf("User \"%s\" not found\n", username)
	} else if resp.StatusCode == 403 {
		fmt.Println("Access for activity info denied")
	} else if resp.StatusCode == 503 {
		fmt.Println("Github is unavailable at the moment")
	} else if resp.StatusCode == 200 {
		defer resp.Body.Close()

		var events []GitHubActivity
		err = json.NewDecoder(resp.Body).Decode(&events)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, event := range events {
			fmt.Println(event)
		}
	}
}
