package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type BacklogItem struct {
	ID          int `json:"id"`
	ProjectID   int `json:"projectId"`
	MilestoneID int `json:"milestoneId"`
}

// url: /api/v2/projects/:projectIdOrKey/versions/:id
// method: Patch

func patchIssueMilestone(apiKey string, issueKey string, milestoneID any) {
	baseURL := "https://gmo-office.backlog.com/api/v2/issues"

	// Prepare URL with query parameters
	url := fmt.Sprintf("%s/%d?apiKey=%s", baseURL, issueKey, apiKey)

	// Prepare URL with query parameters
	req, err := http.NewRequest("PATCH", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	issueKey := "197969"
	milestone := "Review/2024/Sprint-081"

	// update milestones in selected issues
	wikiPages := patchIssueMilestone(apiKey, issueKey, milestone)

	// Unmarshal the JSON data into an array of Items
	var items []Item
	if err := json.Unmarshal([]byte(wikiPages), &items); err != nil {
		panic(err) // Handle error appropriately in production code
	}

}
