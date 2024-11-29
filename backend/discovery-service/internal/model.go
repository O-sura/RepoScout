package internal

import "time"

// Holds the data belong to a repository after making a query
type Repository struct {
	Name              string    `json:"name"`
	Language          string    `json:"language"`
	CreatedAt         time.Time `json:"created_at"`
	PushedAt          time.Time `json:"pushed_at"`
	OpenIssuesCount   int       `json:"open_issues_count"`
	StargazersCount   int       `json:"stargazers_count"`
	ContributorsURL   string    `json:"contributors_url"`
	Tags              []string  `json:"topics"`
	DefaultBranch     string    `json:"default_branch"`
	IssueLabels       []string  `json:"issue_labels"`
	ContributorsCount int       `json:"contributors_count"`
	CommitsCount      int       `json:"commits_count"`
}

// ResultRepository holds the data that belongs to a repository which will be sent and displayed to a user
type ResultRepository struct {}