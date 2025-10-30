package sourceidentifier

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"stack/src/entity"
	"strings"
)

// GetInfo takes any GitHub blob/repo link, extracts repo info,
// gets top contributors, and enriches each with profile data.
func (r *Report) GetInfo(pageLink string) (entity.RepoInfo, []entity.Contributor, error) {
	u, err := url.Parse(pageLink)
	if err != nil {
		return entity.RepoInfo{}, nil, fmt.Errorf("invalid link: %v", err)
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 {
		return entity.RepoInfo{}, nil, fmt.Errorf("invalid GitHub URL format")
	}

	owner, repo := parts[0], parts[1]

	// Step 1: Get repo info
	repoURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	repoResp, err := http.Get(repoURL)
	if err != nil {
		return entity.RepoInfo{}, nil, fmt.Errorf("failed to fetch repo info: %v", err)
	}
	defer repoResp.Body.Close()

	if repoResp.StatusCode != 200 {
		return entity.RepoInfo{}, nil, fmt.Errorf("GitHub API error on repo: %s", repoResp.Status)
	}

	var repoInfo entity.RepoInfo
	if err := json.NewDecoder(repoResp.Body).Decode(&repoInfo); err != nil {
		return entity.RepoInfo{}, nil, fmt.Errorf("failed to parse repo info: %v", err)
	}

	// Step 2: Get contributors
	contribURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contributors?per_page=10", owner, repo)
	contribResp, err := http.Get(contribURL)
	if err != nil {
		return repoInfo, nil, fmt.Errorf("failed to fetch contributors: %v", err)
	}
	defer contribResp.Body.Close()

	if contribResp.StatusCode != 200 {
		return repoInfo, nil, fmt.Errorf("GitHub API error on contributors: %s", contribResp.Status)
	}

	var contributors []entity.Contributor
	if err := json.NewDecoder(contribResp.Body).Decode(&contributors); err != nil {
		return repoInfo, nil, fmt.Errorf("failed to parse contributors: %v", err)
	}

	// Step 3: Enrich each contributor with user profile info
	for i, c := range contributors {
		userURL := fmt.Sprintf("https://api.github.com/users/%s", c.Login)
		userResp, err := http.Get(userURL)
		if err != nil {
			continue
		}
		defer userResp.Body.Close()

		if userResp.StatusCode != 200 {
			continue
		}

		var user entity.Contributor // same fields overlap; safe decode
		if err := json.NewDecoder(userResp.Body).Decode(&user); err != nil {
			continue
		}

		// Merge profile info
		contributors[i].Name = user.Name
		contributors[i].Bio = user.Bio
		contributors[i].Company = user.Company
		contributors[i].Blog = user.Blog
		contributors[i].Location = user.Location
		contributors[i].Email = user.Email
		contributors[i].Twitter = user.Twitter
		contributors[i].Followers = user.Followers
		contributors[i].Following = user.Following
		contributors[i].PublicRepos = user.PublicRepos
	}
	return repoInfo, contributors, nil
}
