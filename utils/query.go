package utils

import (
	"net/url"
)

type Query map[string]string

func (q Query) String() string {
	s := ""

	for k, v := range q {
		s += k + "=" + v + "&"
	}

	return s
}

func BuildRequestString(q Query) string {
	u := url.URL{
		Scheme: "https",
		Host: "api.github.com",
		Path: "/search/repositories",
		RawQuery: q.String(),
	}

	return u.String();
}
