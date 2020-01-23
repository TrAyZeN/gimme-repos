package models

type Repository struct {
	Name             string
	Owner            Owner
	Html_url         string
	Description      string
	Size             int
	Stargazers_count int
	Language         string
	Forks            int
	Open_issues      int
}

func (r Repository) ToMap() map[string]interface{} {
	return map[string]interface{} {
		"name": r.Name,
		"owner": r.Owner.ToMap(),
		"html_url": r.Html_url,
		"description": r.Description,
		"size": r.Size,
		"stargazers_count": r.Stargazers_count,
		"language": r.Language,
		"forks": r.Forks,
		"open_issues": r.Open_issues,
	};
}
