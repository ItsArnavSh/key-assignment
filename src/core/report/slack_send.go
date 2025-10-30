package sourceidentifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"stack/src/entity"
)

type AlertMessage struct {
	Text string `json:"text"`
}

func SendSlackAlert(webhookURL string, core_alert entity.BasicReport, repo entity.RepoInfo, contributors []entity.Contributor) error {
	// Build main alert message
	message := fmt.Sprintf(
		"ALERT: A potential secret/token has been found.\n\n"+
			"--- Leak Details ---\n"+
			"Source: %s\n"+
			"Context: %s\n"+
			"Provider: %s\n"+
			"Token Type: %s\n"+
			"Owner: %s\n"+
			"Description: %s\n\n"+
			"--- Repository Information ---\n"+
			"Repository: %s\n"+
			"Description: %s\n"+
			"Language: %s\n"+
			"Stars: %d | Forks: %d | Open Issues: %d\n"+
			"URL: %s\n\n"+
			"--- Major Contributors ---\n",
		core_alert.Source,
		core_alert.Context,
		core_alert.Secret.Provider,
		core_alert.Secret.TokenType,
		core_alert.Secret.Owner,
		core_alert.Secret.Description,
		repo.FullName,
		repo.Description,
		repo.Language,
		repo.Stars,
		repo.Forks,
		repo.OpenIssues,
		repo.HTMLURL,
	)

	// Append contributors
	for _, c := range contributors {
		message += fmt.Sprintf(
			"Name: %s (%s)\nCompany: %s\nLocation: %s\nFollowers: %d | Following: %d | Public Repos: %d\nEmail: %s\nProfile: %s\n\n",
			c.Name, c.Login,
			c.Company,
			c.Location,
			c.Followers, c.Following, c.PublicRepos,
			c.Email,
			c.HTMLURL,
		)
	}

	// Prepare and send payload
	payload := map[string]string{"text": message}
	data, _ := json.Marshal(payload)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
