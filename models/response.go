package models

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
