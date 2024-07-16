package dto

type RepoDetail struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"html_url"`
	ForkCount      int    `json:"forks_count"`
	StarCount      int    `json:"stargazers_count"`
	OpenIssueCount int    `json:"open_issues_count"`
	WatcherCount   int    `json:"watchers_count"`
	DateCreated    string `json:"created_at"`
	FullName       string
	DateUpdated    string `json:"updated_at"`
	Owner          struct {
		Login string `json:"login"`
	} `json:"owner"`
}
