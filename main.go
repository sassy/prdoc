package main

import (
  "fmt"
  "os"
  "strconv"
  "golang.org/x/oauth2"
  "github.com/google/go-github/github"
)

func main() {
    if len(os.Args) < 3 {
      fmt.Println("few arguments.")
      return
    }

    token := os.Getenv("GITHUB_TOKEN")
    var client *github.Client
    if token != "" {
      ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
      )
      tc := oauth2.NewClient(oauth2.NoContext, ts)
      client = github.NewClient(tc)
    } else {
      client = github.NewClient(nil)
    }

    prs, _, err := client.PullRequests.List(
        os.Args[1],
        os.Args[2],
        &github.PullRequestListOptions{State: "all"},
      )
    if err != nil {
      fmt.Println(err)
      return
    }

    docFile := "doc.md"
    fout, err := os.Create(docFile)
    if err != nil {
      fmt.Println(err)
      return
    }
    defer fout.Close()
    for _, pr := range prs {
      title := strconv.Itoa(*pr.Number) + " " + *pr.Title + "\n"
      fout.WriteString(title)
      fout.WriteString(*pr.Body)
      fout.WriteString("\n")
    }
    fmt.Println("done.")
}
