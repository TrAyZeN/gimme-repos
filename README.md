<h1 align="center">
    gimme-repos
</h1>

> An API that lets you discover unknown GitHub repositories
<!-- <div align="center">
    <img src="assets/image.png"/>
</div> -->

gimme-repos is a wrapper of the GitHub API. It provides you routes with certain
parameters so that you discover unknown repositories.

## Routes
Parameters available for all routes:

| Name       | Type      | Description                               |
|------------|-----------|-------------------------------------------|
| `language` | `string`  | The main language used in this repository |

### Really unknown repositories
Repositories with less than 20 stars.
```
GET /reallyunknown
```

### Unknown repositories
Repositories between 20 and 100 stars.
```
GET /unknown
```

### Maybe known repositories
Repositories with more than 100 stars.
```
GET /maybeknown
```

## Sample response
```
{
    "items":
        [
            {
                "name": "gimme-repos",
                "owner": {
                    "login": "TrAyZeN",
                    "avatar_url": "https://avatars0.githubusercontent.com/u/31016188?v=4"
                },
                "html_url": "https://github.com/TrAyZeN/gimme-repos",
                "description": "An API that lets you discover unknown GitHub repositories",
                "size": 360,
                "stargazers_count": 1337,
                "language": "Go",
                "forks": 42,
                "open_issues": 314
            }
        ]
}
```

## Requirements
- [Go](https://golang.org/)

## Install
```
git clone https://github.com/TrAyZeN/gimme-repos.git
cd gimme-repos
go install
```

## License
MIT TrAyZeN
