# go-prdoc

[![Go Report Card](https://goreportcard.com/badge/github.com/sassy/prdoc)](https://goreportcard.com/report/github.com/sassy/prdoc)

generate document using PullRequest body.

Collect PullRequest body and write file.

## Why

I think that you should write description of design and coding in Pull Request's body.
Because it is useful for review.
So, this tool get Pull Request information for documentation.

## Usage

```
$ go-prdoc "owner" "repos"
```

if you need to handle authentication, you set OAuth2 access token to GITHUB_TOKEN.

```
$ export GITHUB_TOKEN=....
```
