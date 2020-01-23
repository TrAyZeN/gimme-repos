package models

type Owner struct {
	Login string      `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}

func (o Owner) ToMap() map[string]interface{} {
	return map[string]interface{} {
		"login": o.Login,
		"avatar_url": o.AvatarUrl,
	};
}
