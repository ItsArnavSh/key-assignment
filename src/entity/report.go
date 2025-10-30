package entity

type BasicReport struct {
	Context string //The text around it
	Source  string
	Secret  InventorySecret
}

type RepoInfo struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Language    string `json:"language"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	OpenIssues  int    `json:"open_issues_count"`
}

type Contributor struct {
	Login       string `json:"login"`
	HTMLURL     string `json:"html_url"`
	Avatar      string `json:"avatar_url"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Company     string `json:"company"`
	Blog        string `json:"blog"`
	Location    string `json:"location"`
	Email       string `json:"email"`
	Twitter     string `json:"twitter_username"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	PublicRepos int    `json:"public_repos"`
}
