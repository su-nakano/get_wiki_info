package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type BacklogItem struct {
	ID          int `json:"id"`
	ProjectID   int `json:"projectId"`
	MilestoneID int `json:"milestoneId"`
}

// url: /api/v2/projects/:projectIdOrKey/versions/:id
// method: Patch

func patchIssueMilestone(apiKey string, projectID int, issueID int, milestoneID int) {
	url := fmt.Sprintf("https://api.backlog.com/api/v2/projects/%d/versions/%d?apiKey=%s", projectID, issueID, apiKey)

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
	fmt.Println("Hello, world.")
}
