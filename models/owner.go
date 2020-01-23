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
