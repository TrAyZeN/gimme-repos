package models

type Repository struct {
	Name             string `json:"name"`
	Owner            Owner  `json:"owner"`
	HtmlUrl         string `json:"html_url"`
	Description      string `json:"description"`
	Size             int    `json:"size"`
	Stargazers_count int    `json:"stargazers_count"`
	Language         string `json:"language"`
	Forks            int    `json:"forks"`
	Open_issues      int    `json:"open_issues"`
}
