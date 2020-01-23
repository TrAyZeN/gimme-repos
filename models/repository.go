package models

type Owner struct {
	Login string
	Avatar_url string
}

func (o Owner) ToMap() map[string]interface{} {
	return map[string]interface{} {
		"login": o.Login,
		"avatar_url": o.Avatar_url,
	};
}

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

type Response struct {
	Items []Repository `json:"items"`
}

func (r Response) ToMap() map[string]interface{} {
	var repositories []map[string]interface{}

	for _, repo := range r.Items {
		repositories = append(repositories, repo.ToMap())
	}

	return map[string]interface{} {
		"repositories": repositories,
	};
}
