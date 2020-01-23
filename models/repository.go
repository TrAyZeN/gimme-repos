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

func (r Repository) ToMap() map[string]interface{} {
	return map[string]interface{} {
		"name": r.Name,
		"owner": r.Owner.ToMap(),
		"html_url": r.HtmlUrl,
		"description": r.Description,
		"size": r.Size,
		"stargazers_count": r.Stargazers_count,
		"language": r.Language,
		"forks": r.Forks,
		"open_issues": r.Open_issues,
	};
}
